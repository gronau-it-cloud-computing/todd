// Code generated by go-bindata.
// sources:
// agent/testing/bashtestlets/http
// agent/testing/bashtestlets/iperf
// agent/facts/collectors/get_addresses
// agent/facts/collectors/get_hostname
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
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

var _testingBashtestletsHttp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x97\xcf\x72\xdb\x36\x10\xc6\xef\x7c\x8a\x2d\xad\x46\x71\xc6\x16\x23\x1f\x7a\x70\xc6\x99\x3a\xb2\x1c\x3b\x91\xff\x44\xb2\xa7\xed\x44\x19\x0d\x45\x42\x12\x1b\x0a\x60\x09\x50\xae\x6b\xfb\xdd\xbb\x00\x48\x10\xa0\x1c\x59\xd3\xf6\x52\x62\xf9\xdb\x8f\xc0\x62\xf7\x93\xb3\xf3\x53\x50\xf0\x3c\x98\x26\x34\x20\x74\x05\xd3\x90\x2f\x3c\x6f\x07\x7a\x2c\xbb\xcf\x93\xf9\x42\xc0\xc1\xdb\xee\x2f\x70\x11\x0a\x01\x57\xfc\x2e\x4c\x45\x07\x6e\x39\x01\x96\xc3\x92\xc5\xc9\x2c\x89\x42\x91\x30\x0a\x6c\x06\x62\x91\x70\xcc\xe4\xac\xc8\x23\x02\x11\x8b\x09\x24\x1c\xe6\x6c\x45\x72\x4a\x62\x98\xde\x23\x41\x20\x4d\x22\x42\x51\x20\xcb\xd9\x2a\x89\x31\xbe\x20\x39\x39\xc4\xbc\x85\x10\x19\x3f\x0c\x82\x79\x22\x16\xc5\xb4\x13\xb1\x65\x20\x58\x1c\x23\xf7\x27\x89\x84\x7a\x0e\xa6\x29\x9b\x06\xcb\x90\x0b\x92\x07\x83\xf3\x5e\xff\x72\xd4\xf7\xbc\x68\x41\xa2\xef\xaf\x77\xe1\xc1\x03\xfc\x8f\x44\x0b\x06\xfe\xb0\xa0\x34\xa1\x73\x48\x28\xa8\xd7\x72\xb3\xc4\x57\x00\x0a\x2f\x43\x1a\xc3\xfe\x0a\xa2\x22\x4f\xe1\x7d\x10\x93\x55\x40\x8b\x34\x85\x83\xf7\xaf\xba\xf0\xf8\x08\x0f\x5a\xe5\xfd\xab\x03\xf0\x15\x83\xe7\xa0\x4c\xa0\x1a\x17\x61\x9a\x92\xb8\x03\xfd\xbf\x13\x81\x1f\xe8\xf8\xef\x80\xe0\x23\x74\xdf\xc1\x93\xf5\xfd\x9e\xf9\x28\x5c\x1f\x8f\x46\xfd\x13\xfd\x6d\x85\xbe\xf5\x9e\x3c\x2f\x99\xc1\xd7\xaf\xd0\xea\xc2\xd1\x11\x7e\x43\xd2\x3e\x7c\xfb\xf6\x0e\x29\x2c\x12\xd5\x1b\x95\x51\x6f\x96\xc8\xeb\xb8\xc1\xca\xcd\x92\x9c\x0b\x08\xf3\x79\xb1\x24\x14\x1f\x66\x58\x05\x55\x52\x41\xb8\x48\x89\x00\x1a\x2e\x09\xdc\x25\x78\x90\xe3\xc1\x6f\xc7\x7f\x8c\x60\x4a\xf4\x7b\xcc\x21\x78\x6f\x23\xb6\x27\xbf\x88\x87\x09\xd3\xbb\xf0\x9e\x5b\x2f\xf7\x40\x96\x04\xcf\x66\xf4\x39\xcc\x72\xb6\x84\xd6\x81\x7a\x53\x64\xc0\x17\xac\x48\x63\xa9\x99\x85\x9c\xe3\xc5\x25\x54\x30\x25\x11\x66\x59\x55\x55\x0f\x9f\x27\x28\xc1\x8f\xfc\xd6\xaf\x7e\xbd\x6a\x3d\x54\x8f\x3b\xad\xb7\xf0\x04\x3b\x30\x24\x4b\xec\x0c\xb8\xb1\xf6\xfe\x3c\xde\x75\x70\xb5\x5b\x2c\x48\x28\x54\xbf\x41\xc6\x70\x1b\x7b\x50\xe1\xd6\x2e\xe5\x61\xc4\x22\x14\x6d\x0e\x29\x99\x09\xcf\x63\x85\xc8\x0a\x71\xd4\x7a\xad\xae\x74\x9f\xc3\x3e\x83\xfa\xee\xf7\xef\x60\xac\xca\xde\x1e\xeb\xf2\x8f\xe9\xd9\xcd\xcd\x35\x0e\x42\x4c\x0e\xe1\xe7\x07\xd9\x9e\x13\xd9\xd4\x4f\x30\xe9\x5f\x9e\x1c\x1b\x6a\xc0\xd8\x77\x2c\x8f\x48\x96\x8a\x93\xff\x9f\xc8\xd3\xa4\x2a\xae\xe9\x0f\x86\xee\x31\x4a\xb1\x9d\x5d\x3c\xd2\x41\xcd\xf6\x0c\x7b\x9d\x93\xdf\x67\xf2\x8e\x6d\x36\xcb\x89\xc8\x43\xca\xf1\x85\xe6\x4f\x0c\x3f\xc2\xbb\x14\xeb\x19\x5c\x86\xdd\x9c\xbe\xc9\xb9\x61\xd8\xd1\x2e\x2f\x64\x48\x73\xa7\xf6\xbe\x85\xec\x3a\x71\x9f\x29\x32\xd2\xeb\x89\x5c\x6b\xf6\x63\xb3\x6e\xea\x4c\x56\xe9\xac\x33\x9e\x19\xf6\xb2\x58\x4e\x71\xc3\x68\x1f\x65\x02\x9a\x09\x97\x49\xb4\x58\x56\x39\x5c\x27\x9d\x3f\x93\x34\x24\x71\x92\x37\xb2\xf2\x32\x56\xa6\x7d\xaa\xeb\x93\xfc\x43\x64\x52\xcc\xee\x68\xca\xc2\x58\x26\x70\x8c\x4d\xaa\x80\x4e\xf8\xbc\x96\xb0\x20\x61\x4c\x72\x6e\x78\xbd\xd6\xf4\x60\x8d\xce\xc9\x5f\x05\x76\xb5\xa1\xcb\xb5\xc6\x2f\xd6\xf0\x22\x73\xf6\xa2\x97\x1a\xbe\x34\xf0\xf1\x6a\x0e\x27\xe5\x2e\x61\x94\x11\xa2\x13\xe4\x43\x63\xf7\x57\x4e\xce\x6d\xf6\x5c\x86\xfd\x8d\x6b\xc3\x57\xb5\x74\xbb\xa1\xaa\xa6\xa6\xbf\xd4\x8d\x83\x2f\x01\x0d\x40\x75\x17\x54\xed\xb5\xa9\xe9\x86\xf5\xc8\xa0\x83\xc3\xed\x70\x00\xa7\x44\xa0\xcb\xa9\x9d\xe1\x4c\x4e\xc8\x6c\x26\xef\x72\x55\x76\xd4\xc8\x24\xb4\x71\x36\x5b\xdd\x5d\xcf\x33\x53\x88\x63\xac\x9c\xb6\xa5\xa7\x1a\x1e\x41\x5a\xd2\x3e\x85\x36\x0f\x3a\x6f\xac\xc9\x1d\xbf\xee\xbc\x19\xef\xea\x99\xed\xbc\x09\xc6\xdd\x20\x6b\xef\x7a\x8d\x29\xdd\xac\xe6\x4c\xb8\xa5\xf7\xa1\xa1\x57\xb6\xeb\x66\x31\xd7\x00\x2c\xb5\x5e\x43\xcd\x1a\xf4\xcd\x8a\xae\x4d\x58\x8a\x27\x0d\x45\xe7\x4e\x36\x6b\x36\xad\xc4\x52\xed\x37\x54\x95\x5d\x6c\x56\xb3\x4d\xc6\x52\x3a\xad\x95\x6c\x3b\x79\xb1\x7e\x96\x11\x59\x6a\x1f\x6b\x35\xdb\x70\xb6\x6a\x94\xd2\xaa\x2c\xb5\xb3\x5a\xcd\x76\xa2\xcd\x6a\x3f\x30\x33\x4b\xf6\xdc\x95\x35\x56\xb5\xad\xae\xeb\x77\x96\xf0\xa7\x5a\xd8\xb1\xb4\x17\xae\xb9\xe9\x88\x8e\xe6\xe7\x86\xa6\xb6\xbd\xed\x14\x8d\x65\x5a\x7a\x83\x86\x5e\x69\x8c\xdb\x09\x1a\x57\xb5\x04\x2f\x1a\x82\xda\xd7\xb6\xd3\xab\x6c\xd7\x92\xbb\xb4\xe4\x1c\x67\xdd\xac\xf8\x9c\x37\x5b\xaa\x57\x4d\xd5\x6d\x76\xb9\xee\xdd\x96\xe2\x75\x63\x02\xab\x2e\xda\x2c\xd9\xb0\x77\x4b\xef\xcb\x7f\xf7\x89\x1f\xfd\x0a\x58\xea\xc3\x5a\xdd\xf1\xf8\x17\x3c\x77\xed\x27\xc2\x92\x1c\xd5\x92\x1e\xfe\xf1\x8b\x7f\x3f\x4e\xe2\x50\x84\x47\x6d\xfd\x4f\x00\xdf\xfc\x4a\xf8\x87\x7e\xbb\x65\x56\x6d\x7f\x4f\xbf\x6f\xb8\xbf\xa2\x1a\x31\x97\x2d\xc7\xbf\x06\xcb\x80\x4b\x59\x8e\x5d\x93\x56\xd0\xa5\x9d\x2a\xd7\xbc\x13\x76\x33\x94\xd3\xd6\xa4\x5a\x1a\xc2\x76\x50\xc5\xd8\x01\x43\xd9\xce\x68\x17\xc7\x3d\x8d\xed\x78\x8a\xb2\x03\x0e\x65\x0c\xcc\x60\x26\x62\x38\xc7\x8f\x14\xe7\x44\x5c\x4e\xfb\x46\x4d\xe9\xb5\xcb\x94\x56\x50\x43\x65\xc0\xa5\xf4\x98\xd5\x90\x5e\xd7\x8c\x33\xe1\x1a\x73\x42\x0d\xd2\x96\xb3\x02\xee\x0d\x55\x87\xaf\x2f\xa9\x8a\xfc\x8f\xbb\x77\xa6\x46\xc1\x4e\xa4\xed\x7b\x4f\x6d\x4f\x8f\x92\x35\x0c\xde\xbf\x01\x00\x00\xff\xff\x7f\x9e\x7c\xb1\xd9\x0f\x00\x00")

func testingBashtestletsHttpBytes() ([]byte, error) {
	return bindataRead(
		_testingBashtestletsHttp,
		"testing/bashtestlets/http",
	)
}

func testingBashtestletsHttp() (*asset, error) {
	bytes, err := testingBashtestletsHttpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testing/bashtestlets/http", size: 4057, mode: os.FileMode(493), modTime: time.Unix(1477387277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testingBashtestletsIperf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\xcc\x2a\x46\x95\x04\x8e\x15\xe7\xb0\x87\x04\x0e\x36\x9b\xcd\x21\x40\xdb\x0d\xea\x2c\x8a\x45\x53\x18\x94\x34\xb2\xd8\x48\xa4\x40\x8e\xac\xf5\x3a\xfe\xef\x1d\x92\xfe\x6a\xe1\x4b\x93\x0b\xcd\x99\x79\x7c\xef\xcd\x8c\x4e\x7e\x4b\x3b\x6b\xd2\x4c\xaa\x14\xd5\x02\x32\x61\xab\x28\x3a\x81\x7b\xdd\x2e\x8d\x9c\x57\x04\x57\x97\xe3\xdf\xe1\x83\x20\x82\xbf\x6d\x2f\x6a\x1a\xc1\x3f\x16\x41\x1b\x68\x74\x21\x4b\x99\x0b\x92\x5a\x81\x2e\x81\x2a\x69\xb9\xd2\xea\xce\xe4\x08\xb9\x2e\x10\xa4\x85\xb9\x5e\xa0\x51\x58\x40\xb6\xe4\x0c\x84\x5a\xe6\xa8\x18\xa0\x35\x7a\x21\x0b\xbe\xaf\xd0\xe0\x35\xd7\x55\x44\xad\xbd\x4e\xd3\xb9\xa4\xaa\xcb\x46\xb9\x6e\x52\xd2\x45\xc1\x79\xdf\x30\x27\x7f\x4e\xb3\x5a\x67\x69\x23\x2c\xa1\x49\xdf\x3f\xde\x3f\x7c\x9c\x3e\x44\x51\x5e\x61\xfe\x7a\x7a\x06\xab\x08\xf8\x0f\xf3\x4a\x43\xfc\xa9\x53\x4a\xaa\x39\x48\x05\x3e\xec\xc8\x62\xec\x13\x18\xb8\x11\xaa\x80\x8b\x05\xc8\x16\x4d\x09\xb7\x69\x81\x8b\x54\x75\x75\x0d\x57\xb7\xef\xc6\xf0\xf6\x06\xab\x00\x73\xfb\xee\x0a\xe2\x90\xc4\x4a\x94\x26\xc6\xb3\x24\xea\x1a\x8b\x11\x3c\xfc\x27\x89\x9f\x18\xc5\x37\x80\x7c\x84\xf1\x0d\xac\x0f\x18\xdc\xef\x9e\x85\xa7\xbb\xe9\xf4\xe1\xaf\xf0\xba\x4f\xbd\x8c\xd6\x51\x24\x4b\xf8\xf2\x05\x06\x63\x98\x4c\x20\xf6\x24\x63\xf8\xfa\xf5\x86\xb3\xd8\x26\x15\xa8\xba\xdb\xa8\x94\xae\x21\xcf\xec\x5d\x29\x8d\x25\x10\x66\xde\x35\xa8\xf8\x50\xb2\x0f\xde\x54\x42\x4b\x35\x12\x28\xd1\x20\xf4\x92\x95\xdc\xbd\xff\x7c\xf7\xef\x14\x32\x0c\x71\xae\x41\xee\xdc\x54\x0f\xdd\x8b\x2c\x46\xd4\xbd\x58\xda\x83\xe0\x10\x9c\x29\xac\x6d\x87\x6f\xa1\x34\xba\x81\xc1\x95\x8f\x74\x2d\xd8\x4a\x77\x75\xe1\x30\x5b\x61\x2d\xb7\x4e\x2a\xd2\x1e\x42\xb4\xed\xd6\xd7\x88\xcf\x33\x86\xb0\x93\x78\xf0\x47\xbc\xff\x35\x58\x6d\x8f\x27\x83\x4b\x58\xc3\x09\x7c\xc2\x86\x67\x03\x9e\x0f\xb8\x1f\x4f\x1f\xff\x90\xee\xd9\xb2\x21\x82\xfc\xc4\x41\xab\x99\xc6\x10\xb6\xe9\x07\x2c\x9d\x18\xaa\x04\x25\x16\x6a\x2c\x29\xb8\xc8\x15\x5b\xb7\x3a\x8b\x76\xa3\x1e\x6c\x97\x59\x92\xd4\xb9\x61\x1e\xf2\x0c\x73\x85\xab\x6b\xc4\x2b\x72\xcc\xa0\x07\x02\xee\x19\xb7\x8f\x7b\x68\xd9\x1d\x9e\x2d\xaf\x9d\x1f\x1d\x42\x8f\x60\xb0\xad\x05\x4f\x3e\x67\xf4\x3c\xc3\x21\x98\x53\x27\xea\xed\x23\x8f\x4f\xa3\x03\x85\xa7\x7e\x50\x06\x3b\xe2\x6f\xe0\x3c\x4d\x6c\xba\x5a\x6d\x0b\xd6\xeb\x34\x89\x07\xe3\x38\x49\x93\x33\x47\xff\xa3\xa6\x0d\x93\x1e\x13\x26\x55\x6a\x93\xbb\x31\x6f\x70\x2e\x32\xc9\xba\xfc\x36\xb1\x30\x77\x19\x5f\x94\xd0\xc4\x4e\x0b\x93\xcb\x85\x82\x4a\x2c\xdc\xde\x61\x21\x73\x12\x59\xed\xda\x68\x5c\xe6\x08\x3e\xf3\xb9\x73\x7e\x62\x13\x7c\x45\x40\xee\xb9\xd5\xe1\x31\xa1\x96\xdc\x5e\x55\xf2\xee\xba\x91\x77\xc6\x99\xfd\x02\x7b\xf2\xbd\x56\x09\xed\x92\xb6\x06\x48\x1b\x04\x9b\x4e\xcd\x36\xf3\x31\xd9\x2c\xd4\x5e\xb7\xa3\x09\x71\x14\xe9\x8e\x98\x04\xfb\x32\xf8\xa9\x84\xa5\x53\xde\xce\x7a\xa9\x0a\xdd\xcf\xac\xfc\x1f\x67\x4d\xb6\x24\xdc\x59\x18\x2a\xd9\x40\x32\x90\xbc\xa8\x04\xdc\x7f\xb0\xf3\x42\x39\x47\x47\xe7\xcf\xf7\x4f\x10\x00\xc0\x01\x5c\xc3\xcb\xe9\xe8\xfc\xe5\x6c\x74\x9e\xbe\x8c\xd3\x36\xd9\x9b\x0f\x1f\xfe\x64\x68\xbe\x67\xcb\x01\xdc\xc4\xb0\x01\x35\x7f\x72\x7c\x02\x6f\x3f\xa1\x5f\xbe\x30\x66\x4e\x34\x4f\x9a\xc2\x1c\xad\x15\x66\x09\x19\xf3\x78\x64\x1f\x36\x31\x8b\xec\x28\x2f\x89\x6b\x26\xed\x36\x58\x2b\x74\x97\xbd\x36\xaf\xf0\xad\xe3\x0b\xe7\x08\x7f\xe4\xea\x65\x44\x46\x28\x5b\xa2\x09\x0a\xed\xaf\x48\xb4\x98\xc3\x79\x90\x15\x44\xd8\xad\xba\xb3\x28\x63\x1b\x7b\x59\x50\xc5\xb8\x92\xec\x8c\x73\x7f\x05\x3a\xc0\xed\xd1\x1d\xc6\x1e\x3c\xe2\x18\xaf\xcb\xac\x10\x24\x26\xc9\x2a\x3e\xda\xac\xf8\x3a\x4e\x06\x47\x23\x49\x3c\x84\xf8\x27\xdd\x21\xfb\xc7\x3b\x9f\x77\x44\x88\xcf\x3d\x72\x9f\xc4\xeb\x24\x0a\x12\x0f\x08\x46\xdf\x03\x00\x00\xff\xff\x16\x5e\xc0\xe2\xf1\x06\x00\x00")

func testingBashtestletsIperfBytes() ([]byte, error) {
	return bindataRead(
		_testingBashtestletsIperf,
		"testing/bashtestlets/iperf",
	)
}

func testingBashtestletsIperf() (*asset, error) {
	bytes, err := testingBashtestletsIperfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testing/bashtestlets/iperf", size: 1777, mode: os.FileMode(493), modTime: time.Unix(1477387277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _factsCollectorsGet_addresses = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\xcd\x6a\x22\x41\x10\xc7\xef\xf3\x14\xb5\xe3\xc0\xe8\xb2\x6b\x27\x1e\x72\x98\x20\x24\x88\x87\x40\x3e\x0e\x21\x27\x23\xa1\xa7\xbb\x74\x2a\xd1\xee\xa1\xab\x54\xc4\xf8\xee\x29\x05\x35\x7d\xaa\xae\xff\x07\xbf\xea\xfc\x31\x2b\x4e\xa6\xa6\x60\x30\xac\xa1\xb6\xdc\x64\x59\x07\x46\xb1\xdd\x26\x9a\x37\x02\x83\xab\xeb\x1b\x78\xb2\x22\xf0\xc2\x1b\xbb\x90\x3e\xbc\x31\x42\x4c\xb0\x8c\x9e\x66\xe4\xac\x50\x0c\x10\x67\x20\x0d\xb1\x26\x39\xae\x92\x43\x70\xd1\x23\x10\xc3\x3c\xae\x31\x05\xf4\x50\x6f\xd5\x81\xb0\x20\x87\x41\x0b\xda\x14\xd7\xe4\x75\xdf\x60\xc2\x4a\x73\x8d\x48\xcb\x95\x31\x73\x92\x66\x55\xf7\x5d\x5c\x1a\x89\xde\xab\xef\x13\x9d\x1c\x67\x53\x2f\x62\x6d\x96\x96\x05\x93\x79\x7c\x18\x8d\x9f\x5f\xc7\x59\x66\x53\xb2\xdb\x0f\x96\x34\x2c\x77\xf9\xbd\xf7\x09\x99\x91\xf3\x6a\x92\x97\x2a\x9e\xfe\xc3\x6e\xd1\xa5\x16\x0e\x7f\xf8\x06\xb0\x9b\x2f\x28\x0d\x05\x14\xb3\x6b\x13\x05\x81\x62\xb0\x2f\x55\x61\x45\x2a\xd9\xbc\x9b\xfe\x5f\x63\xca\x5e\x2f\xcb\x66\x7a\x2b\x01\x05\xc8\x8b\xdd\xb9\x6e\x72\x37\xdd\xe7\xb7\xe0\x63\x06\xfa\x3a\x87\x73\x9c\x0a\xea\xa1\xfc\xb8\xba\x50\x15\xe7\xb1\xa0\x32\xff\xa7\x54\x3e\x06\xfc\xcd\xad\xbd\xa7\xb9\xaa\xfe\x2b\xc7\x74\xaf\xe8\xe8\x9a\x08\x97\xf0\x4f\x00\x00\x00\xff\xff\xd8\xb5\x56\x13\xa9\x01\x00\x00")

func factsCollectorsGet_addressesBytes() ([]byte, error) {
	return bindataRead(
		_factsCollectorsGet_addresses,
		"facts/collectors/get_addresses",
	)
}

func factsCollectorsGet_addresses() (*asset, error) {
	bytes, err := factsCollectorsGet_addressesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "facts/collectors/get_addresses", size: 425, mode: os.FileMode(493), modTime: time.Unix(1477387277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _factsCollectorsGet_hostname = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\x8c\xbd\x4a\xc4\x50\x10\x85\xfb\x3c\xc5\x98\x15\xa2\xcd\x8e\x5a\x58\xa4\x5d\x16\x14\xfc\x29\xc4\x4a\x2c\xee\xcf\x6c\xee\x48\x72\x27\xdc\x99\x44\x82\xf8\xee\x5e\x64\xbb\x73\x0e\xdf\x77\x76\x17\xb8\x68\x41\xcf\x19\x29\xaf\xe0\x9d\xa6\xa6\xd9\xc1\x41\xe6\xad\xf0\x90\x0c\xee\x6e\x6e\xef\xe1\xd9\x99\xc1\xab\x7e\xbb\xd1\xf6\xf0\xae\x04\x52\x60\x92\xc8\x27\x0e\xce\x58\x32\xc8\x09\x2c\xb1\x56\x53\x65\x29\x81\x20\x48\x24\x60\x85\x41\x56\x2a\x99\x22\xf8\xad\x12\x04\x23\x07\xca\xf5\x60\x2e\xb2\x72\xac\x7b\xa2\x42\x7d\xf5\x92\xd9\xac\x3d\xe2\xc0\x96\x16\xbf\x0f\x32\xa1\x49\x8c\x95\xfb\xa2\x60\xff\x19\xfd\x28\x1e\x27\xa7\x46\x05\x9f\x1e\x0f\xc7\x97\xb7\x63\xd3\x50\x48\x02\xdd\x4f\xfb\x20\x6a\xd9\x4d\xd4\xf6\x1f\x6d\x77\x79\x95\xce\xf5\xba\x6b\x3f\x7f\xbb\xbf\x00\x00\x00\xff\xff\x92\x35\x2d\x49\xe7\x00\x00\x00")

func factsCollectorsGet_hostnameBytes() ([]byte, error) {
	return bindataRead(
		_factsCollectorsGet_hostname,
		"facts/collectors/get_hostname",
	)
}

func factsCollectorsGet_hostname() (*asset, error) {
	bytes, err := factsCollectorsGet_hostnameBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "facts/collectors/get_hostname", size: 231, mode: os.FileMode(493), modTime: time.Unix(1477387277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"testing/bashtestlets/http": testingBashtestletsHttp,
	"testing/bashtestlets/iperf": testingBashtestletsIperf,
	"facts/collectors/get_addresses": factsCollectorsGet_addresses,
	"facts/collectors/get_hostname": factsCollectorsGet_hostname,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"facts": &bintree{nil, map[string]*bintree{
		"collectors": &bintree{nil, map[string]*bintree{
			"get_addresses": &bintree{factsCollectorsGet_addresses, map[string]*bintree{}},
			"get_hostname": &bintree{factsCollectorsGet_hostname, map[string]*bintree{}},
		}},
	}},
	"testing": &bintree{nil, map[string]*bintree{
		"bashtestlets": &bintree{nil, map[string]*bintree{
			"http": &bintree{testingBashtestletsHttp, map[string]*bintree{}},
			"iperf": &bintree{testingBashtestletsIperf, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

