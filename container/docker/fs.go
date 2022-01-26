// Copyright 2022 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/google/cadvisor/container"
	"github.com/google/cadvisor/container/common"
	"github.com/google/cadvisor/fs"
	info "github.com/google/cadvisor/info/v1"
)

func GetFsStats(
	stats *info.ContainerStats,
	machineInfoFactory info.MachineInfoFactory,
	metrics container.MetricSet,
	storageDriver StorageDriver,
	fsHandler common.FsHandler,
	globalFsInfo fs.FsInfo,
	poolName string,
	rootfsStorageDir string,
	zfsParent string,
) error {
	mi, err := machineInfoFactory.GetMachineInfo()
	if err != nil {
		return err
	}

	if metrics.Has(container.DiskIOMetrics) {
		common.AssignDeviceNamesToDiskStats((*common.MachineInfoNamer)(mi), &stats.DiskIo)
	}

	if metrics.Has(container.DiskUsageMetrics) {
		var device string
		switch storageDriver {
		case DevicemapperStorageDriver:
			device = poolName
		case AufsStorageDriver, OverlayStorageDriver, Overlay2StorageDriver, VfsStorageDriver:
			deviceInfo, err := globalFsInfo.GetDirFsDevice(rootfsStorageDir)
			if err != nil {
				return fmt.Errorf("unable to determine device info for dir: %v: %v", rootfsStorageDir, err)
			}
			device = deviceInfo.Device
		case ZfsStorageDriver:
			device = zfsParent
		default:
			return nil
		}

		var (
			limit  uint64
			fsType string
		)

		var fsInfo *info.FsInfo

		for _, fs := range mi.Filesystems {
			if fs.Device == device {
				limit = fs.Capacity
				fsType = fs.Type
				fsInfo = &fs
				break
			}
		}

		fsStat := info.FsStats{Device: device, Type: fsType, Limit: limit}
		usage := fsHandler.Usage()
		fsStat.BaseUsage = usage.BaseUsageBytes
		fsStat.Usage = usage.TotalUsageBytes
		fsStat.Inodes = usage.InodeUsage

		if fsInfo != nil {
			fileSystems, err := globalFsInfo.GetGlobalFsInfo()

			if err == nil {
				addDiskStats(fileSystems, fsInfo, &fsStat)
			} else {
				klog.Errorf("Unable to obtain diskstats for filesystem %s: %v", fsStat.Device, err)
			}
		}

		stats.Filesystem = append(stats.Filesystem, fsStat)
	}

	return nil
}

func addDiskStats(fileSystems []fs.Fs, fsInfo *info.FsInfo, fsStats *info.FsStats) {
	if fsInfo == nil {
		return
	}

	for _, fileSys := range fileSystems {
		if fsInfo.DeviceMajor == fileSys.DiskStats.Major &&
			fsInfo.DeviceMinor == fileSys.DiskStats.Minor {
			fsStats.ReadsCompleted = fileSys.DiskStats.ReadsCompleted
			fsStats.ReadsMerged = fileSys.DiskStats.ReadsMerged
			fsStats.SectorsRead = fileSys.DiskStats.SectorsRead
			fsStats.ReadTime = fileSys.DiskStats.ReadTime
			fsStats.WritesCompleted = fileSys.DiskStats.WritesCompleted
			fsStats.WritesMerged = fileSys.DiskStats.WritesMerged
			fsStats.SectorsWritten = fileSys.DiskStats.SectorsWritten
			fsStats.WriteTime = fileSys.DiskStats.WriteTime
			fsStats.IoInProgress = fileSys.DiskStats.IoInProgress
			fsStats.IoTime = fileSys.DiskStats.IoTime
			fsStats.WeightedIoTime = fileSys.DiskStats.WeightedIoTime
			break
		}
	}
}
