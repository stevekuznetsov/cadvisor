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
// generated by build/assets.sh; DO NOT EDIT

// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cmd/internal/pages/assets/html/containers.html (9.64kB)

package pages

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cmdInternalPagesAssetsHtmlContainersHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x5f\x73\xe2\x38\x12\x7f\x86\x4f\xd1\xeb\xba\x87\xd9\xaa\xb1\x49\x26\xb9\x87\x9b\x23\x54\xb1\xcc\xcc\x2d\xb7\x99\x24\x15\x92\xdd\xda\x47\xd9\x6e\x6c\x4d\x84\xa5\x95\x64\x08\x97\xca\x77\xbf\x92\x64\x83\xff\x01\xf9\x57\x33\xcb\x0b\x20\xa9\xbb\x7f\xfd\x53\xb7\xd4\xb2\x3c\xfc\xc9\xf7\xfb\x00\x13\x2e\xd6\x92\x26\xa9\x86\x0f\x47\xc7\xa7\xf0\x1f\xce\x13\x86\x30\xcd\xa2\x00\xc6\x8c\xc1\xb5\xe9\x52\x70\x8d\x0a\xe5\x12\xe3\xa0\xdf\x07\x38\xa7\x11\x66\x0a\x63\xc8\xb3\x18\x25\xe8\x14\x61\x2c\x48\x94\x62\xd9\xf3\x1e\x7e\x47\xa9\x28\xcf\xe0\x43\x70\x04\xef\xcc\x00\xaf\xe8\xf2\x7e\xfe\x77\x1f\x60\xcd\x73\x58\x90\x35\x64\x5c\x43\xae\x10\x74\x4a\x15\xcc\x29\x43\xc0\xfb\x08\x85\x06\x9a\x41\xc4\x17\x82\x51\x92\x45\x08\x2b\xaa\x53\x6b\xa6\x50\x12\xf4\x01\xfe\x2c\x54\xf0\x50\x13\x9a\x01\x81\x88\x8b\x35\xf0\x79\x75\x1c\x10\x6d\xf0\x9a\x4f\xaa\xb5\xf8\x38\x18\xac\x56\xab\x80\x58\xac\x01\x97\xc9\x80\xb9\x71\x6a\x70\x3e\x9d\x7c\xbe\x98\x7d\xf6\x3f\x04\x47\x46\xe2\x36\x63\xa8\x14\x48\xfc\x2b\xa7\x12\x63\x08\xd7\x40\x84\x60\x34\x22\x21\x43\x60\x64\x05\x5c\x02\x49\x24\x62\x0c\x9a\x1b\xb4\x2b\x49\x35\xcd\x92\xf7\xa0\xf8\x5c\xaf\x88\xc4\x3e\x40\x4c\x95\x96\x34\xcc\x75\x8d\xaa\x12\x1b\x55\xb5\x01\x3c\x03\x92\x81\x37\x9e\xc1\x74\xe6\xc1\x2f\xe3\xd9\x74\xf6\xbe\x0f\xf0\xc7\xf4\xe6\xd7\xcb\xdb\x1b\xf8\x63\x7c\x7d\x3d\xbe\xb8\x99\x7e\x9e\xc1\xe5\x35\x4c\x2e\x2f\x3e\x4d\x6f\xa6\x97\x17\x33\xb8\xfc\x02\xe3\x8b\x3f\xe1\xb7\xe9\xc5\xa7\xf7\x80\x54\xa7\x28\x01\xef\x85\x34\xf8\xb9\x04\x6a\x48\x34\xf3\x06\x30\x43\xac\x01\x98\x73\x07\x48\x09\x8c\xe8\x9c\x46\xc0\x48\x96\xe4\x24\x41\x48\xf8\x12\x65\x46\xb3\x04\x04\xca\x05\x55\x66\x2a\x15\x90\x2c\xee\x03\x30\xba\xa0\x9a\x68\xdb\xd2\x72\x2a\xe8\xfb\xfe\xa8\xdf\x1f\xa6\x7a\xc1\x46\x7d\x80\x61\x8a\x24\x1e\xd9\x29\x18\x6a\xaa\x19\x8e\xa2\x71\xbc\xa4\x8a\x4b\xf0\xe1\xe1\x21\xf8\x44\x95\x60\x64\x7d\x41\x16\xf8\xf8\x38\x1c\xb8\x21\x6e\xb8\x8a\x24\x15\x1a\x94\x8c\xce\xbc\x87\x87\xe0\x9a\x73\xfd\xf8\xa8\x8c\xe5\x68\x20\xb8\x10\x28\x83\x05\xcd\x82\x6f\xca\x1b\x0d\x07\x6e\x70\x21\xf9\x93\xef\xc3\x39\xd1\xa8\xb4\x8d\x21\xca\x30\x36\xd8\x61\x41\x33\x3a\xa7\x18\xc3\x64\x36\x03\x83\xd3\x8e\x66\x34\xbb\x03\x89\xec\xcc\x53\x7a\xcd\x50\xa5\x88\xda\x83\x54\xe2\xbc\x6d\x37\xe4\x5c\x2b\x2d\x89\xf0\x4f\x83\xa3\xe0\xc8\x0f\x51\x93\xe0\x83\xc5\x11\x29\xe5\x8d\xfa\x5b\x00\x97\xc2\x50\x44\x98\x61\x67\x81\xaf\x35\x67\x95\xf8\x27\xc1\x71\x70\xdc\xb2\xf6\x1c\x8d\x11\xcf\x4c\xb6\xa0\x54\x2d\xc0\x7b\x19\xfb\x2f\x59\x92\x99\x9b\x90\x8d\x27\xfb\x26\xe8\xdb\x5f\x39\xca\xb5\x7f\x12\xfc\xb3\x00\xdc\x31\x4d\xfb\xe4\xf7\x10\xdd\xd6\xb4\xd5\xa5\xd7\x02\xcf\x3c\x8d\xf7\x7a\xf0\x8d\x2c\x89\x6b\xf5\xba\x4d\x30\x4e\x62\x94\x7b\x80\x3d\x47\x59\x85\xd7\xa6\xc2\xe1\xa0\xcc\x81\x61\xc8\xe3\x75\x61\x23\xa6\x4b\x88\x18\x51\xea\xcc\xdb\xc8\xba\x50\xf1\x55\xca\x57\x11\x51\xe8\xc1\xc6\x3d\xd2\x9c\x4e\x6f\x2b\xcc\x7c\xb5\xf0\x8f\x3f\x78\x40\xe3\x33\x8f\xf1\x84\x7b\x1b\xb1\x01\xd9\xfc\xac\xd9\x2b\x45\x46\xfd\x5e\xb5\x43\x90\x04\x7d\x03\x16\xa5\xe9\x32\xd9\x7b\x3c\x6a\x27\x69\x7a\x6c\xe4\x06\x31\x5d\x9a\x6f\xce\x4a\xf1\x50\x22\x89\x23\x99\x2f\x42\x27\xfd\xf0\x20\x49\x96\x20\xfc\x43\x10\x89\x99\x9e\x6c\xdc\xfc\x78\x06\xc1\x55\xbd\x4d\x3d\x3e\x5a\x83\x8c\x8e\x2a\xce\x36\x25\x83\x73\x9a\xdd\x3d\x3e\x7a\xa3\x8e\xae\x1b\xbc\xd7\x06\x1d\x19\x0d\x07\x8c\x16\x00\x30\x8b\x8d\xe2\xe1\x80\xb3\x2d\x29\x16\xb8\xfb\xf3\xf0\x40\xe7\x10\x4c\x95\x23\xf5\x00\x57\x50\x7c\x86\xe9\xe9\x16\x64\x10\x0c\x62\x1e\xdd\x19\xc6\x3e\xd9\x6f\xd8\xfa\xe4\xc0\xa4\xa7\x9d\xa6\x0f\x59\x69\xdb\x11\x3c\x5e\x90\xcc\x1b\x5d\xd9\xef\xa7\xda\x29\x49\xa8\x3a\x3c\xcb\xc3\xa8\xca\xfc\xeb\x62\xe4\x64\x54\xd3\x37\x1c\xa4\x27\xd5\x00\xa9\x08\x33\xaa\xb4\x9f\x48\x9e\x8b\x46\x84\xa8\x8a\x02\x1b\x1e\x4d\x84\xbd\x5a\x12\xd4\xc6\x97\x41\xd1\x36\xe2\x53\x8d\x0b\x1b\x2c\xb5\xf1\xdb\x48\x69\x04\x49\x75\x76\x76\x52\xe8\x18\x74\x73\x3d\xd3\x44\xe7\x6f\x41\xe0\x27\x49\x97\x28\xc1\xe9\x6b\x12\x98\xb3\x83\xfc\xb9\x10\x54\x56\xdc\xf2\xd7\xc0\xe7\x52\xcb\xa9\x81\x0e\x8a\x86\x4a\x90\xac\xb4\x62\xd4\xf8\x8c\x84\xc8\x2c\x77\x55\xdd\xc1\x6f\xb8\x36\xd4\x99\xe1\x23\x68\x76\xfe\x4e\x58\x6e\x57\x88\x66\xfe\xd5\x59\x73\xce\x6e\xb1\xf5\x5e\x06\x6d\xa6\xb9\x24\x09\x0e\x43\x39\x2a\x00\x19\x55\xbb\xc8\xea\x6d\xb9\xb2\xe6\x5b\x5c\xed\x46\xf5\x5c\xbe\x2a\xfa\xdb\x7c\x55\x3b\xeb\x7c\xf5\x36\x74\xf5\x86\x83\x9c\x59\x6f\x4a\x26\x8b\x86\x5d\xd1\xda\x95\xe3\xce\xab\xe9\x82\x24\x78\x38\x42\x2b\x8b\xce\xce\x50\x05\xa8\x2e\x4d\x27\x23\xa7\xda\x05\x6b\xa5\xa7\x8a\xcb\x69\x33\xfb\x92\x8b\x13\x9f\x5a\x19\xb3\x3f\xd6\x46\x99\x29\x0c\xe5\xf6\xff\x21\xdf\xae\x51\xf1\x5c\x46\xa8\xc6\x4b\x42\x99\x29\xc9\xdf\x20\x07\xa7\x8a\x33\x5b\xd6\x36\xf2\xcf\x99\x9c\x88\xbc\x6a\x6c\x67\xa0\x55\x98\xd8\x19\x3f\x40\x22\x4d\x97\xe6\x00\x50\x58\xf4\x6d\xdd\x0b\x82\x64\xc8\xdc\x6f\x6f\x34\xb9\xba\x75\xd3\xbf\xd5\x58\x2c\xde\x02\x23\x03\x27\x38\x37\x85\xf8\xc6\xf1\xfd\x26\xf7\xe5\x51\x4a\xa4\x99\xc7\x32\x46\x85\xa4\x99\x76\x8d\x6d\x63\x50\x53\x93\x67\x74\xa3\x46\x55\xd5\xb4\x91\x57\x27\xb1\xc3\x97\xaf\xe4\xfe\x8d\xdc\xf9\x4a\xee\xc1\xaa\x6a\x78\x34\xe1\x75\x87\xb6\x16\x77\xfb\x14\xf1\x57\xb9\xa4\xee\x5e\xef\xce\x98\x31\xbe\x32\x47\x16\xde\x9e\x24\x63\xa1\x61\x10\x82\xaf\x24\x4a\x69\x86\xd3\x6c\xce\x83\x8b\x7c\x61\xe5\xca\x35\xa6\x8d\xbe\x5c\x6a\x36\xff\x9d\x13\x5f\x71\xc1\xe5\xfa\xfb\x06\xbc\xb3\xb9\x27\xe6\xdd\x80\xc0\x3d\x89\xb0\x6a\x5e\x4f\x6f\x45\x59\x33\x03\xe8\xff\x70\x8f\xe1\xdd\x41\x53\xc8\xdf\x66\x54\xef\x91\x7f\x49\x54\x15\x7a\xde\x28\x51\xba\x92\xa4\xed\xf4\xc1\x1c\xd9\xe9\x6e\x21\xf9\x0a\x47\x67\x2b\x22\xde\x6a\x91\x5b\x11\xd1\xb9\x2c\xb4\x3d\xae\x58\x7d\x81\xd7\x15\xe9\x03\x9e\x37\x53\xaf\xf0\xee\x29\x67\x84\xc3\x9b\xd9\xad\x32\xa5\xd1\xee\x4a\xdc\x66\x5e\x91\x7f\x42\xd2\x05\x91\xeb\x3d\x65\x80\x19\x65\x2c\xd0\x2c\x69\x17\x02\xf5\x61\x45\x32\x5f\x2e\x51\x2e\x29\xae\xf6\x97\x07\xd5\x0a\x21\x37\x88\xfd\x84\xe4\x09\x7a\x75\x95\xe6\xd4\xbc\x29\x19\x7e\x88\x37\x57\x92\x47\xa8\xd4\xa1\x6a\xa7\xea\x8e\x28\x45\x7c\xcd\xc5\x93\x1c\xda\x51\x67\x7c\x47\x37\x6d\xc9\xf1\x14\x07\x3b\xbc\x69\x18\x38\x1d\xdd\x70\x4d\x18\x94\x71\x78\x6a\x23\xb3\xc2\x4f\x24\x72\x5f\x9b\x21\xbe\x9b\xf8\x28\x25\x52\x6f\x49\x81\xf2\xa9\x94\x51\x35\xb9\xba\x85\x73\x4e\x62\x18\x2f\x51\xee\xd1\xc7\x38\x89\xeb\x8a\x36\x0f\xab\xaa\xc8\x2c\x26\x10\xf6\xa8\x2e\x77\x2a\x13\x28\x7d\xb3\xff\x77\xe2\xeb\x56\xf9\x8b\x44\x72\x17\xf3\x55\xb6\x4b\xa7\x53\x15\x96\xc3\x76\x2a\x6d\x87\xc6\xc1\xdd\xf9\x3b\x86\x49\xb9\x51\x7f\xa7\x48\x59\x58\x73\x87\xa7\x21\x94\x83\x46\x4b\x05\x80\xe4\x2b\xe8\x3e\xf0\xec\x9d\xc2\xc6\xb0\xf6\x72\xfc\x2f\x7b\xb6\xac\xb9\x2a\x79\x22\xd1\x3e\x5b\x85\xd6\xa7\x6b\xa0\x1f\x12\x09\xd5\x3f\x7e\x6c\x0e\xaa\xd2\x2b\xd7\x11\xd7\x91\x72\xed\x3b\x2a\x3a\x35\x43\x7d\xaf\x52\xd2\xe7\x19\x5b\x7b\xa3\x5f\xb9\x86\x72\xc2\xdc\x21\xb9\x43\xb2\xcd\xe6\x73\xe0\xd2\x6c\xce\x1b\x60\x23\xce\xe2\x97\xa0\x9d\x70\x16\x3f\x15\x6e\xaf\xd7\x89\xbb\xbb\xb1\x3d\x73\x27\x5e\x35\xba\x34\xde\x37\x57\x9f\x67\x26\xe5\x05\xea\x15\x97\x77\xcf\xcc\xca\xde\xeb\xd3\xb1\x30\x5c\x6c\xf6\xcf\x49\xc4\x5e\xb3\x37\x96\x5c\x98\xe0\x6f\x27\x48\x98\x6b\xcd\x37\xf3\x15\xea\x0c\x42\x9d\xf9\x31\xce\x49\xce\x34\x94\x72\xbe\xe6\x49\xc2\xd0\x2b\x1e\x9d\x3b\x21\xc7\x73\xe6\x50\xfa\x0a\x19\x46\xf6\x08\xb0\x31\x06\x31\xd1\xa4\x10\xad\x60\x00\x22\x29\xf1\x53\xa2\x04\x17\xb9\x38\xf3\xb4\xcc\xb1\x68\xc4\x7b\x41\xb2\x18\xe3\x33\x6f\x4e\x98\xc2\x8e\x10\x73\xe1\xd5\x6d\xb8\x9c\xeb\xee\xf8\xaa\x05\x66\x44\x24\x56\xc6\xf6\xca\x48\x70\x9e\xb5\x58\xca\x59\xb7\x49\xaf\x49\xb0\xbf\xc0\x2c\xf7\x40\x72\xe3\xb1\xfb\x6d\x1d\xb3\xd5\x25\xc3\x38\x5c\xef\x65\xac\x1d\xf3\xc5\xe3\xa1\x3d\x61\xfb\x9c\x05\x39\x95\x3c\x4f\x52\x91\xeb\xf6\x2a\xb8\x59\x96\x4b\x78\xe1\x5a\xa3\x6a\x6f\xdf\x2f\x30\xfb\x59\x4a\x6e\x1f\x1f\xb7\xb6\x80\xd2\x16\xda\x11\xbb\x8d\x35\x9c\x6f\x64\xe8\x17\xf5\xc3\xb6\xcc\x2f\x94\xa1\x5a\x2b\x8d\x8b\xa7\x57\x90\xf3\x8d\x8c\xdb\xfb\x3a\x8b\xc8\xdd\x9a\x76\x2c\x53\x93\x5c\x69\xbe\xf8\x8a\x5a\xd2\xe8\xb9\x7c\x1c\x58\xac\x7a\xfb\x18\x18\xbb\xdb\x73\x13\xc7\x50\x58\x6f\xae\x58\xfb\x62\xa5\x51\x4b\x59\x27\xfc\x85\xd3\x73\x30\x1e\x7a\xcd\xc3\x66\xc7\x2d\xc8\x0f\x0b\x8d\x8e\xbb\x93\x43\xd1\xf1\xb4\xa2\x4a\x80\xa9\x9b\x6d\x59\xf3\xb1\xb9\x5e\xd0\x4c\xe4\xba\x56\xea\x56\x6f\x48\xfc\xd8\x5d\xf8\xf9\x11\xcf\x33\xed\x75\xee\xdf\x9b\xad\xbb\x4b\xce\xaa\xdf\x21\xb7\x24\x2c\xc7\xb3\xe3\xa3\x06\xe4\xdd\x0b\x4d\x27\xc2\x5a\x35\xd8\xd0\xd4\xbd\x00\xbe\x90\x43\x57\x8c\x1c\xa4\xb1\x28\x23\xfe\x9e\x4c\xd6\x4a\x2d\x67\x45\x72\xc6\x2a\x66\x42\xc6\xa3\x3b\xaf\xab\x7e\xde\xe7\xdc\xcb\x27\x61\xc7\x42\xdd\xd1\x59\xed\xaa\x74\xec\xbf\xa3\x2f\x85\x13\xfb\x2e\x53\x60\x11\xaa\x40\xa1\xbe\xcc\xcc\x39\x72\x42\x18\x0b\x49\x74\xf7\x4e\x69\x22\xf5\x15\x49\xf0\xdd\xc3\x43\xb0\xb9\x4f\x75\xf7\xdc\xef\xc1\xb4\xd5\x4e\xe3\xb6\xa9\x75\xf8\xb2\xad\xee\x02\xd9\xfe\x2c\x6f\x93\x7f\xb6\x2f\x3a\x99\x4f\x2c\xc9\xca\xdd\x96\x18\x3b\xf5\x8b\x99\x62\x50\xfd\x85\x01\xf7\x9e\xc0\x70\xe0\xde\xa2\xf9\x7f\x00\x00\x00\xff\xff\x19\x81\xf2\xde\xa8\x25\x00\x00")

func cmdInternalPagesAssetsHtmlContainersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_cmdInternalPagesAssetsHtmlContainersHtml,
		"cmd/internal/pages/assets/html/containers.html",
	)
}

func cmdInternalPagesAssetsHtmlContainersHtml() (*asset, error) {
	bytes, err := cmdInternalPagesAssetsHtmlContainersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/internal/pages/assets/html/containers.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3b, 0xb9, 0x8c, 0xe6, 0xbf, 0xf1, 0x9d, 0x7, 0xe7, 0xbc, 0xd5, 0xf1, 0xb1, 0xc2, 0x47, 0xe5, 0x2e, 0x76, 0xdf, 0x59, 0x12, 0x64, 0x50, 0x9d, 0x91, 0xb3, 0xaf, 0xb3, 0x1, 0x60, 0xd, 0x6c}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmd/internal/pages/assets/html/containers.html": cmdInternalPagesAssetsHtmlContainersHtml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cmd": {nil, map[string]*bintree{
		"internal": {nil, map[string]*bintree{
			"pages": {nil, map[string]*bintree{
				"assets": {nil, map[string]*bintree{
					"html": {nil, map[string]*bintree{
						"containers.html": {cmdInternalPagesAssetsHtmlContainersHtml, map[string]*bintree{}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
