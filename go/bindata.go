// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../definitions/fields.toml (630B)
// ../definitions/info_object_meta.toml (1.497kB)
// ../definitions/info_storage_meta.toml (107B)
// ../definitions/operations.toml (4.658kB)
// ../definitions/pairs.toml (1.346kB)

package specs

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

var _fieldsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xb1\x4e\xc3\x40\x10\x44\x7b\x7f\x45\x94\x6f\x40\x74\x34\x74\x14\x08\x44\x0a\x8a\x08\xa1\xb5\xbd\x09\x0b\x89\xcf\x9a\x1d\x04\xe1\xeb\xd1\x9d\x13\xdd\x39\xbe\xd2\x6f\x9e\x77\x6e\x77\xdb\xda\x5b\xc3\xd3\xa8\xab\xbb\xd5\xfa\xfe\x10\xba\xaf\x07\x2a\x84\x01\xeb\xa6\xd9\xb6\xd6\xe7\xd4\x09\x1b\xf6\x67\xec\xd7\xfc\x5d\x00\x39\xc5\xb4\x77\xd6\x7e\x52\x20\x63\x05\xa6\x06\x1b\x7a\xfd\xcd\xdc\x06\x46\x7a\x54\x4a\x86\x1b\x06\xc8\x5e\x1f\x95\x12\xc3\x61\xa6\xdf\xde\x24\x26\x47\xad\x95\x86\x0c\x9f\xda\x4f\xed\xd2\xf4\xb0\xdb\xb9\xb2\x32\x25\xd8\xb5\x5e\x1e\x63\x14\x43\xb1\xf6\x73\xfc\x9c\x38\x58\x62\xf0\x42\x7d\x8e\xcf\x36\x3f\x6a\x2f\x1d\x6d\x2e\x97\xcd\xc5\xe1\x5e\x54\x7a\x4d\xd0\xed\x4f\x2b\x3b\x38\xba\xda\x78\xa7\xd0\x9c\xd6\x2d\xce\xba\xb9\x24\x93\x66\x0b\x01\xe5\x53\x9c\x01\xba\x54\x62\xf4\x8d\x43\xad\xf9\x27\xc3\x57\x18\x93\xfb\x1f\x00\x00\xff\xff\xab\x67\x1f\xe1\x76\x02\x00\x00")

func fieldsTomlBytes() ([]byte, error) {
	return bindataRead(
		_fieldsToml,
		"fields.toml",
	)
}

func fieldsToml() (*asset, error) {
	bytes, err := fieldsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fields.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0xd, 0x48, 0xb2, 0x4b, 0xc9, 0xdb, 0x61, 0xf5, 0x19, 0xdf, 0x4f, 0xfa, 0xb8, 0x12, 0x4a, 0xc0, 0x58, 0xb, 0xc3, 0x1a, 0x30, 0x9d, 0x11, 0xf5, 0x8f, 0xbe, 0x73, 0xdb, 0x3c, 0x20, 0xf7}}
	return a, nil
}

var _info_object_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x3b\x8f\xd3\x40\x10\xc7\xfb\x7c\x8a\x51\x9a\xab\x4c\x05\x74\x57\x20\x5d\x83\x44\x74\x48\x80\x28\x4e\x08\x6d\xb2\xe3\x64\x88\xf7\xc1\xec\xf8\x88\xf9\xf4\x68\x76\xe3\xc7\xd9\x16\xdc\x51\x25\xbb\xf3\x9f\xdf\xdf\xf3\xb0\x1f\x4c\x8c\xe8\x6d\x15\xea\x3a\xa1\x7c\xdb\x48\x17\x11\x6e\x61\x4b\x5e\xde\xbe\xde\x6e\x2c\xa6\x03\x53\x14\x0a\x5e\x6f\xdf\x65\xf1\x7d\xd6\x02\x25\x90\x13\x42\xc9\x84\x50\xe7\x53\xc1\x41\xd8\xff\xc0\x83\x6c\x37\x9b\x9e\xef\x5b\xb7\x47\xae\x9c\xb9\x90\x6b\xdd\x13\x9f\x85\xcb\xce\x5c\x7a\x4e\x49\x4b\x40\x7e\x20\x47\x64\xa3\xca\x09\x3c\xd1\x6f\x5c\x45\xaf\x94\x30\x81\x6b\x9a\x92\x23\xf2\xdf\xe8\x12\xc4\x34\xff\xeb\x91\x93\x07\xa7\x35\x97\x43\xf0\x82\x5e\xaa\x06\xfd\x51\x4e\x0b\xf4\x28\x70\xf6\xcd\x18\x4d\xc2\xe4\x8f\xd3\xb0\x46\xd6\xe2\x28\xe6\xb8\x76\x4f\x76\x79\x8b\x97\x18\x58\xe0\x16\x84\x5b\x9c\x57\xf5\xfe\xae\x1f\x79\xeb\xe9\x67\x8b\x70\xc6\x4e\x8b\x4a\x12\xd8\x1c\xf1\x95\x42\x1b\x93\xa4\x72\xc1\x52\x4d\x38\xe1\x0b\x39\xcc\x61\xf2\xe7\x4a\x0c\x1f\xa7\xab\xd6\x9b\xcf\xec\x3e\x90\x3f\x7f\xce\xd2\xde\x36\x75\x4e\x01\x50\x00\x50\x07\x86\x7c\x2e\xcb\x96\xfd\x5d\xb0\x93\x26\xdc\xe7\xc0\x2e\x58\x9c\x97\xb6\x79\x70\x6d\x23\x14\x0d\x4b\xb5\xd6\x88\xf9\x40\x7b\xf1\xd8\x03\x3d\x01\x59\xdd\xfb\xfc\x77\xfa\x10\x03\xfa\x85\x5b\xaf\xa2\x42\x2b\x89\xda\xdd\x01\xf6\x74\x6b\x46\x8f\x17\x2e\xe6\xe8\x91\x97\xd2\x62\x4d\x1e\x2d\xec\xbb\x7e\x8e\xbc\x86\x27\xff\x2c\x7c\x91\xfd\x1b\x1f\xcd\x74\xd3\x9f\xb5\x7d\x1f\x8d\x9c\xb4\xf7\x48\x72\x42\x2e\xdf\x9a\x7d\x0a\x4d\x2b\x3a\x0b\x39\x41\x28\x97\x8c\x8d\x11\x7a\xbc\x5e\x4a\xf8\x65\xd8\xa6\xde\xfc\x26\xc1\xd7\xc0\xe7\x3b\x62\xb0\xa8\xaf\x62\x82\xe0\xa1\x4d\xc8\x37\xfa\x8d\x89\x6d\x19\x60\x42\x7e\xa4\x03\x56\x0e\xc5\x58\x23\x66\x7c\x54\xe3\xbb\x45\xd5\x9f\x8a\x7a\x77\x15\x67\x2f\x4c\x70\x85\x0c\x2d\xe8\x61\x6a\xa0\x8e\x2b\xf4\xd2\x88\xef\xd7\x1f\x67\xe2\xc2\xeb\x4b\x42\x9e\x1b\x29\x6c\xe9\xf2\x27\x00\x00\xff\xff\x8c\x8e\xe3\x56\xd9\x05\x00\x00")

func info_object_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_info_object_metaToml,
		"info_object_meta.toml",
	)
}

func info_object_metaToml() (*asset, error) {
	bytes, err := info_object_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "info_object_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0x22, 0x3, 0x91, 0xf9, 0xc2, 0xc4, 0x6e, 0x2b, 0x1e, 0xc6, 0xe, 0xa4, 0x62, 0x72, 0x3f, 0x13, 0xd2, 0x51, 0x84, 0x12, 0x80, 0x34, 0x83, 0xf3, 0x7d, 0xbe, 0xea, 0x96, 0xa5, 0x8d, 0x50}}
	return a, nil
}

var _info_storage_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xce\xc9\x4f\x4e\x2c\xc9\xcc\xcf\x8b\xe5\x2a\xa9\x2c\x48\x55\xb0\x55\x50\x2a\x2e\x29\xca\xcc\x4b\x57\xe2\xe2\x8a\xce\x4b\xcc\x4d\xc5\x14\x4f\xad\x28\xc8\x2f\x2a\x51\xb0\x55\x28\x29\x2a\x4d\xe5\xe2\x8a\x2e\xcf\x2f\xca\xd6\x4d\xc9\x2c\x22\xa4\x12\x10\x00\x00\xff\xff\x5f\xe0\xd4\x5a\x6b\x00\x00\x00")

func info_storage_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_info_storage_metaToml,
		"info_storage_meta.toml",
	)
}

func info_storage_metaToml() (*asset, error) {
	bytes, err := info_storage_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "info_storage_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0xd1, 0x40, 0x2c, 0xba, 0x2b, 0x76, 0x2b, 0x8d, 0x58, 0xb2, 0xa, 0x71, 0xa7, 0xb1, 0xb5, 0x96, 0x12, 0x28, 0xb9, 0x77, 0x6d, 0x13, 0xc5, 0x3a, 0x98, 0x62, 0x3e, 0x73, 0xf2, 0x73, 0x74}}
	return a, nil
}

var _operationsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x58\x4f\x93\xdb\xb6\x0f\xbd\xeb\x53\x60\x94\x43\x2e\x1b\xdf\x7e\xc7\x1c\x7e\xe9\xa6\xd3\x43\x9b\xed\x34\xcd\x29\xb3\xb3\xa1\x29\xc8\x46\x97\x22\x38\x24\x65\xc7\xfd\xf4\x1d\x90\xb2\xad\xbf\xb6\xf6\xb2\x63\x51\x00\x1e\x88\xf7\x08\x50\xfb\x5d\x39\x87\xb6\x42\xff\x5c\x54\x18\xb4\x27\x17\x89\x2d\x7c\x84\x92\x02\xc4\x3d\x02\xd9\x88\xbe\x56\x1a\xa1\x66\x0f\xff\x4f\xd6\xe0\xd1\xa8\x88\x15\xb0\x43\xaf\xc4\x21\x6c\xca\xa2\xb8\xc4\xda\xb0\xdb\x68\x8f\x2a\xe2\x4b\x5e\x9a\x04\x3f\x92\x31\x90\x4d\x40\x59\xc8\x56\xc0\xdb\x7f\x50\xc7\x4d\x59\x38\xe5\x55\x13\xe0\x23\x7c\x2f\x9d\x8a\xfb\xf2\xb9\xf0\x18\x5a\x13\xf3\x12\x97\xcf\x23\xb0\xa3\xa7\x3b\x58\x1d\x82\x66\x1b\xd1\x46\x88\x7c\x07\x96\xcb\x07\x28\xbd\xfc\x09\xf4\x2f\x8e\x12\xb0\xe5\x73\x91\x9d\x5e\x1a\xae\x50\x40\x72\xa8\x49\x11\xb8\x69\x28\xde\x2e\x42\x32\x01\x65\x2b\xa8\xc9\x52\xd8\xf7\x12\x73\x9e\x35\x86\x30\xce\xec\x06\xf8\xd6\xb0\x7e\x5d\x4b\xe6\x27\x31\x5e\xe2\xb2\x8b\xd4\xa3\x32\xad\xdc\x66\x12\x2c\x1e\x21\xd9\xbd\x8d\xcb\x1e\x58\xa6\xf2\x06\x56\x32\x18\x10\x99\x11\x6f\xf2\xf7\x00\xe5\x96\xaa\xfb\x34\xa6\x48\xe3\xed\x73\xb3\x25\x7b\x7b\xff\xd9\x24\x27\x12\xa4\xcc\x49\x5f\x8b\xc2\xda\x52\x15\xd6\x81\x1b\x0a\xf1\x16\xb2\xbc\x3f\xc3\x6e\xd1\xb0\xdd\x49\x49\xe2\x9e\xc2\x02\xfa\xb0\x04\x5b\xba\x91\x86\x66\x47\x6b\xb5\xf4\x0b\xbb\xd3\xe6\xea\x94\xeb\xe6\x4e\x4b\xe5\x72\x27\xa9\xcf\x53\x02\x06\xf6\xd0\xb4\x26\x92\x33\xd8\x25\x0d\x64\x13\x46\x40\x7f\x20\x8d\xc3\x3d\x04\xaf\xa5\x86\x55\x88\x49\x3b\x15\xf9\xb5\x49\x3e\x92\x47\x1d\xd9\xe7\x4c\x93\x63\x4f\xdf\x15\x4d\xc3\x4c\xd5\x5d\x91\x7f\x9b\xb6\x6b\x8c\x7a\xbf\x36\xc3\x5f\xc5\x38\x65\xd7\xb9\x49\x7e\xe9\xe7\x7c\x6a\xe9\x15\xd4\x9e\x1b\x50\xb0\xa3\x03\x5a\x68\xbd\x11\x09\x48\x3e\x33\x19\x3e\x40\xd9\x7a\x93\x12\x6b\xf8\xb0\x36\xad\x3f\xf8\x80\x29\xab\xe4\x23\x39\xc9\x8f\xf9\x94\xe4\xcd\x55\xfc\x6f\x61\x32\x8b\x40\xf9\xb8\x3a\xad\xb3\xc3\x52\x13\xeb\x45\xec\x11\x7d\x59\x5d\x41\xf7\xc5\x76\x25\xdb\x23\xc4\xdc\xcd\xee\x00\xce\x74\xb4\x05\xd8\x69\x57\x23\x5b\xe1\xcf\x49\x5f\x7b\x80\x52\x9c\xa7\x67\x3b\xad\xce\x14\x86\x1b\x67\x70\x45\x69\x3a\xbb\x7e\x86\xd0\x3a\xc3\xaa\x4a\x13\x4c\xb3\x0d\xd1\xb7\x3a\x5e\x4f\xf7\xcc\x06\xc4\x6b\xa6\xfd\x2d\xe4\x96\x3a\xe0\x9d\xbc\x52\x17\x4c\x61\xc7\x4d\x70\xb1\x90\xc3\x92\xb9\x99\x3e\x78\xce\xc7\xa9\xdd\x5a\x41\xfe\xa9\x76\x38\xa3\x45\x38\xee\x49\xef\x21\xb4\xce\xb1\x88\x55\xd9\x8a\x9b\x4c\xfb\xe6\x82\xd0\x13\xa8\x3c\xaf\xd0\xa6\x98\xbd\xad\x17\x5d\x80\xb2\x2e\x97\x71\x26\x92\x0c\x0e\x35\xd5\xa4\x81\xeb\x3a\xe0\x5d\x59\x66\xab\xfb\xf3\x56\x32\x90\x0a\x78\x54\xeb\x9b\xe4\x5f\x62\xbc\xe9\xb9\xc9\x8e\xd2\xcf\xf9\xcd\x38\xcf\x07\xaa\xa4\x6a\x47\x75\x7a\xe8\xc8\xd0\xca\x42\xf2\x49\x00\x2b\xab\x78\x6e\x9d\x5d\x37\x9b\xe6\x2b\x51\x1b\x45\x36\x2a\xb2\xbd\x23\x12\x22\x7b\x21\xab\xf3\xcb\xcd\xe9\x1c\xe4\x4a\xfc\x0a\xce\xbb\x48\x1e\xc8\x86\xa8\xec\xb8\xa3\x5a\xd5\x8c\x2f\xab\xe2\x81\x83\xac\x05\xb0\x42\x39\xc7\xf3\x80\xf9\x1d\xa8\x95\x60\xc3\xc0\x3b\x5c\x38\xa3\x3b\x8c\xa0\xe0\xa0\x0c\x55\xd3\xc0\x89\xd7\xd9\x19\xb1\x7a\x47\xd2\x03\x6e\x74\x07\x65\xcc\x14\x36\x40\x2b\x57\xf5\xdc\x2a\xae\xf0\x43\x30\xca\x50\x9d\xef\x3a\x89\x8e\x08\xcf\x7c\x77\x11\xfa\xe5\x9f\xca\x6d\x14\xbd\x2c\x8b\x01\x27\x97\xb9\x9a\x66\xfe\x39\x7c\x51\xbc\x7b\x07\x9f\x70\xaf\x0e\xc4\xbe\x28\x3e\xc0\x63\x36\x67\x6b\x4e\x67\x57\xb6\x98\x7a\x74\x5a\x93\x87\x4e\xf3\x05\xc0\x07\xf8\x9a\x03\xc1\xe3\xd3\x97\xf7\x7f\xc3\x97\xcf\x9f\x1f\xd3\xa1\x3f\x37\x2c\xcc\x43\xdd\x98\x6c\xfd\x2d\xa0\xbf\x18\x91\x8c\x84\x46\xda\x44\x36\x7b\x91\x4a\x6f\x4f\x52\x94\x26\xa0\xa9\x37\xd7\x74\x28\x00\x55\xd8\x38\x96\xae\xd2\x01\xb7\x5a\x3e\x71\xea\xf6\xba\x45\x73\x54\xa7\x00\x1e\x63\xeb\x2d\x58\x32\x80\xde\xb3\xcf\xe6\x5d\xa0\xaf\xbf\x3d\x7d\xfb\xfd\x11\x2c\x1e\xd0\x9f\x2d\x7f\xe4\x79\xf3\x85\xe3\xe7\x9f\x14\xe2\x8f\xbe\xfd\x70\x5b\x7a\x8f\xfa\xb5\x77\xec\x01\xc5\x41\xae\xa1\x96\xe3\xa6\x90\x9a\x0f\xd9\x6a\x30\xaa\x4a\x45\x35\x2f\xaf\x0e\x5f\xb7\xde\x4b\x15\x2e\x22\x3b\x7b\x8d\x04\x25\xcb\xe5\x73\x61\x58\x2b\x03\x1f\x21\xfa\x16\x47\x70\xcb\x4a\xee\xa0\xb2\xa0\xaf\x3d\x79\xe1\xae\x27\xf2\x22\x9f\x17\xf2\x14\xe5\x6a\x7c\x96\x78\xa8\xee\xae\x93\x2e\x7c\xab\xca\x9b\x54\xb7\x9a\x0c\xbe\x0f\xd0\xed\x6e\xe6\x8e\x79\x1c\x80\x8f\xa6\x82\x5c\x5b\xf8\x45\x2b\x63\xb6\x4a\xbf\xce\x0c\x89\x61\x3e\xab\x7b\x63\xc7\xe6\x91\xe2\x9e\x5b\xb9\x82\x9c\x40\x39\x02\x01\x5a\x35\x1f\x6f\x50\x12\xa2\x5a\xa0\x44\xde\x80\x4a\x14\x88\xb6\xa4\xcb\x91\xad\x19\xb8\x5e\xfa\x02\x5c\x9c\xce\x7d\xc0\x34\x80\x6f\xcd\x66\xa9\x7d\xf7\xa1\x99\xe8\x98\xe7\xa1\xff\x1f\x8c\x2b\x1f\xdd\x5c\x7f\x89\x27\x97\x78\x39\x3f\x37\xd5\xff\xee\x93\xf3\x5f\x00\x00\x00\xff\xff\xbf\xc0\x48\x68\x32\x12\x00\x00")

func operationsTomlBytes() ([]byte, error) {
	return bindataRead(
		_operationsToml,
		"operations.toml",
	)
}

func operationsToml() (*asset, error) {
	bytes, err := operationsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "operations.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x91, 0xf0, 0xc6, 0xf2, 0x6b, 0x8a, 0x67, 0x58, 0xe8, 0xe2, 0x34, 0x34, 0x59, 0xd0, 0xda, 0x8d, 0xf6, 0x13, 0xd3, 0x28, 0x71, 0x8e, 0xad, 0xc, 0xc0, 0x82, 0xb9, 0x59, 0xf4, 0xa4, 0xde, 0x13}}
	return a, nil
}

var _pairsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x41\x8f\xd4\x30\x0c\x85\xef\xfd\x15\x56\xcf\xab\xd9\x0b\x70\xe3\x84\x84\x84\xc4\x9e\x80\x13\x42\x55\xa6\x76\xb6\xd6\xa4\x71\x71\xdc\x29\xc3\xaf\x47\x49\xdb\x99\x02\x3b\xb3\x0b\xb7\x2a\xf1\xfb\x9e\xfd\x9c\x7e\x6d\x25\x1a\x45\x6b\x7a\x7c\xfd\xad\xb2\xd3\x40\xf0\x16\xea\x64\xca\xf1\xb1\xae\xaa\xf3\x75\xbe\xb9\x76\xcf\x71\x74\xc6\x12\x1b\x93\x03\xc5\xbf\xab\x90\x52\xab\x3c\xe4\x92\x72\x3c\x50\xcb\xfe\x04\xd6\x11\x6c\xe5\x50\xe4\xe0\x45\x21\x70\xb2\x42\x57\x42\x8a\xc6\x2e\xbc\x98\xda\xc9\x04\x26\x30\xa8\x1c\x19\x09\x2e\x84\x02\x4e\xa4\x47\x6e\x09\xf2\xa7\x89\xba\x47\xca\x36\x14\x71\x10\x8e\xf6\xbf\x26\xab\xfe\x96\xc5\x8f\x81\x75\x93\x20\x47\xbb\x4a\x9f\x3a\x8a\x25\x9d\x51\x03\x28\xd9\xa8\x91\x10\xf6\x27\x50\x72\x6d\x07\x13\x87\x00\x33\x2f\x93\x39\x1a\x69\x4b\x83\x89\x5e\xf0\x1f\x2e\x87\xb9\x26\xe7\xd9\xf4\x82\x9b\x06\x3e\x72\xb2\x07\xc1\x82\x08\xd2\x96\x0d\xfc\xd3\xea\x56\xd1\x8d\xa1\xa3\xeb\x9f\x78\x34\x37\x98\x8b\x16\xb2\x30\x03\xc4\xfb\x44\xf6\x5b\x6a\x6f\x5e\x5d\x25\xcc\xd5\xa5\x1f\xeb\x38\x81\xd2\xf7\x91\x92\xdd\x9d\xb1\x25\xb9\x44\x74\xc8\xcb\x2b\x25\x8b\x64\x4f\x5e\x94\x72\xbe\x98\x6d\xfb\x31\x18\x0f\x4e\xad\x61\x7c\xea\xd1\xb3\x34\xad\x0b\x61\xef\xda\xc3\x26\x72\x79\xb7\x9c\xdd\x58\xac\x33\x30\x41\x01\x3a\x92\x9e\xc0\xb8\x27\x98\x66\x5f\x40\x67\x0e\xbc\x4a\x0f\x49\x46\x6d\xcb\xfc\x89\x7f\xd2\x4b\xa7\xcf\xb5\xcf\xcd\x2e\x31\x9c\x66\xb7\xc0\x3d\x1b\x21\x2c\xff\x77\x71\xcf\x8e\x93\xe8\xa1\x41\xd6\x67\xd7\x56\xd7\xd5\x76\x73\x59\x07\xc8\x7a\xe5\x35\xdc\x2d\x13\xcb\x40\x3a\xbf\x9a\xd2\xcf\x3e\xcf\x1e\x9c\xf1\x91\xce\x2b\x41\xd6\x5d\xb5\xb6\x01\x0f\x5f\x3e\x7d\x86\x64\x4e\x0d\x26\xb6\x0e\xee\x8b\xc1\x0c\x5b\x47\x5b\xec\xd2\x46\xb6\xd2\x91\xbc\x1b\x43\x0e\x1d\xee\x81\x3d\x44\x31\x48\x64\xbb\xea\xbd\x28\xf8\xf4\x27\x02\x4a\x63\x11\x65\x4a\x30\x04\x67\x5e\xb4\xbf\x2b\xf3\xed\xa9\x73\x47\x16\x05\x4e\x30\x46\x24\xcf\x91\x70\x57\xd5\x75\xfd\x2b\x00\x00\xff\xff\x63\x67\x32\x3e\x42\x05\x00\x00")

func pairsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pairsToml,
		"pairs.toml",
	)
}

func pairsToml() (*asset, error) {
	bytes, err := pairsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pairs.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4d, 0x82, 0x39, 0xfa, 0x3f, 0x1, 0x7, 0x7d, 0xaf, 0xfb, 0xd1, 0x87, 0xa5, 0xf8, 0x43, 0xf6, 0x69, 0xa3, 0x71, 0x7f, 0x12, 0xa3, 0xb1, 0xd3, 0xba, 0xdd, 0x72, 0x5d, 0x95, 0xbf, 0xf3, 0xcb}}
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
	"fields.toml":            fieldsToml,
	"info_object_meta.toml":  info_object_metaToml,
	"info_storage_meta.toml": info_storage_metaToml,
	"operations.toml":        operationsToml,
	"pairs.toml":             pairsToml,
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
	"fields.toml":            {fieldsToml, map[string]*bintree{}},
	"info_object_meta.toml":  {info_object_metaToml, map[string]*bintree{}},
	"info_storage_meta.toml": {info_storage_metaToml, map[string]*bintree{}},
	"operations.toml":        {operationsToml, map[string]*bintree{}},
	"pairs.toml":             {pairsToml, map[string]*bintree{}},
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
