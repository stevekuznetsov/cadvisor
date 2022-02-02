// Copyright 2016 Google Inc. All Rights Reserved.
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

package utils

import (
	"fmt"
	"os"
	"strings"

	dockertypes "github.com/docker/docker/api/types"
	v1 "github.com/google/cadvisor/info/v1"
)

const (
	DriverStatusPoolName      = "Pool Name"
	DriverStatusMetadataFile  = "Metadata file"
	DriverStatusParentDataset = "Parent Dataset"
)

func DriverStatusValue(status [][2]string, target string) string {
	for _, v := range status {
		if strings.EqualFold(v[0], target) {
			return v[1]
		}
	}

	return ""
}

func DockerThinPoolName(info dockertypes.Info) (string, error) {
	poolName := DriverStatusValue(info.DriverStatus, DriverStatusPoolName)
	if len(poolName) == 0 {
		return "", fmt.Errorf("Could not get devicemapper pool name")
	}

	return poolName, nil
}

func DockerMetadataDevice(info dockertypes.Info) (string, error) {
	metadataDevice := DriverStatusValue(info.DriverStatus, DriverStatusMetadataFile)
	if len(metadataDevice) != 0 {
		return metadataDevice, nil
	}

	poolName, err := DockerThinPoolName(info)
	if err != nil {
		return "", err
	}

	metadataDevice = fmt.Sprintf("/dev/mapper/%s_tmeta", poolName)

	if _, err := os.Stat(metadataDevice); err != nil {
		return "", err
	}

	return metadataDevice, nil
}

func DockerZfsFilesystem(info dockertypes.Info) (string, error) {
	filesystem := DriverStatusValue(info.DriverStatus, DriverStatusParentDataset)
	if len(filesystem) == 0 {
		return "", fmt.Errorf("Could not get zfs filesystem")
	}

	return filesystem, nil
}

func SummariesToImages(summaries []dockertypes.ImageSummary) ([]v1.DockerImage, error) {
	var out []v1.DockerImage
	const unknownTag = "<none>:<none>"
	for _, summary := range summaries {
		if len(summary.RepoTags) == 1 && summary.RepoTags[0] == unknownTag {
			// images with repo or tags are uninteresting.
			continue
		}
		di := v1.DockerImage{
			ID:          summary.ID,
			RepoTags:    summary.RepoTags,
			Created:     summary.Created,
			VirtualSize: summary.VirtualSize,
			Size:        summary.Size,
		}
		out = append(out, di)
	}
	return out, nil
}
