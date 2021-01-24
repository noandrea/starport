// Code generated for package cosmosfaucet by go-bindata DO NOT EDIT. (@generated)
// sources:
// starport/pkg/cosmosfaucet/openapi/index.html
// starport/pkg/cosmosfaucet/openapi/openapi.yml.tmpl
package cosmosfaucet

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _openapiIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x4f\x6f\xd4\x30\x10\xc5\xef\xfb\x29\x06\x73\xa0\x95\x48\xdc\x0a\x54\xa1\xc5\x89\x50\xf9\x23\x21\x55\x02\x09\x38\x70\x42\x6e\x3c\x49\x86\x75\xec\xc8\x1e\xef\x36\x42\xfd\xee\x68\x13\x4a\xb2\xdd\x5e\x36\x97\x64\x26\xbf\x79\x2f\x79\x23\xab\x67\x1f\xbe\xbc\xff\xfe\xf3\xeb\x47\x68\xb9\xb3\xe5\x4a\xed\x6f\x60\xb5\x6b\x0a\x81\x4e\x94\x2b\x00\x00\xd5\xa2\x36\xd3\xe3\x58\x76\xc8\x1a\xaa\x56\x87\x88\x5c\x88\xc4\x75\xf6\x46\x80\x5c\x00\x4c\x6c\xb1\xfc\xa4\x53\x85\xac\xe4\x54\xcd\x6f\x2d\xb9\x0d\x04\xb4\x85\x88\x3c\x58\x8c\x2d\x22\x0b\xe0\xa1\xc7\x42\x30\xde\xb1\xac\x62\x14\xd0\x06\xac\x0b\x21\x65\x72\xfd\xa6\xc9\x2b\xdf\xc9\xb8\xd3\x4d\x83\x21\x4b\x94\x19\x8a\xfc\xee\x55\xfe\xfa\x22\xbf\x58\xb4\xf3\x71\x50\x3e\x69\x45\x95\x77\x0f\x26\xd4\xe9\x06\x65\xef\x9a\x13\x5c\x6a\xbd\xdd\x4b\x64\x97\x57\x77\x97\x57\xf9\x38\xfb\xcf\x48\xc9\x39\x1e\x75\xeb\xcd\xb0\xf0\x37\xb4\x05\x32\x85\x98\x45\x45\xa9\xa4\xa1\x6d\xb9\x9a\xa1\x58\x05\xea\x19\x62\xa8\x4e\xfc\xdf\xec\x36\x39\x63\x31\xff\x1d\xf7\xaa\x93\x4c\xf9\x58\x77\x6e\xec\x2f\x29\x81\x1c\x31\x7c\x9b\x44\xa0\xf6\x01\xea\x71\x4f\x2f\x22\xf8\x1e\x9d\xee\x29\x1f\x3a\x9b\x1f\x4c\xed\xc8\x19\xbf\xcb\xbd\xb3\x5e\x1b\x28\xa0\x4e\xae\x62\xf2\xee\xec\x1c\xfe\x1c\x80\xff\xd1\x44\x50\x3c\x98\xfc\xf8\x7c\x3d\x7e\xe7\xd9\x63\x16\x20\x05\xbb\x06\xb1\xf0\x15\x2f\x8f\x18\xe3\xbb\x5f\x64\xd6\x20\x9e\x2f\x52\x7c\x02\x43\xec\x6f\xc8\x6d\xc8\x35\x6b\xe0\x90\xf0\x18\xb1\x7a\xf0\x89\xd7\x20\xae\x75\xc4\x9b\xb1\x38\x12\xba\x3f\x7f\x7b\xd0\xb9\x9f\xf3\x3c\x48\x58\xc9\x69\xd5\x4a\x4e\x07\xe7\x6f\x00\x00\x00\xff\xff\x1d\x27\x6d\x15\x49\x03\x00\x00")

func openapiIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_openapiIndexHtml,
		"openapi/index.html",
	)
}

func openapiIndexHtml() (*asset, error) {
	bytes, err := openapiIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi/index.html", size: 841, mode: os.FileMode(436), modTime: time.Unix(1611427875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _openapiOpenapiYmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x95\x4d\x73\xdb\x38\x0f\xc7\xef\xfa\x14\x18\x3d\xcf\x4c\x2e\x8d\xe5\xa4\xdb\x43\x74\xf3\x6c\x67\x67\x72\xeb\x6c\x7b\xec\xa1\x34\x05\xd9\x4c\x44\x40\x01\x40\xd7\x4e\x27\xdf\x7d\x87\x92\xfc\xd2\x75\x5e\xba\x7b\xdb\x8b\xc7\x24\x01\xf0\x8f\xff\x8f\xa4\xf4\xbb\x5b\xad\x50\x6a\x28\xaf\x67\xf3\xb2\x28\x02\xb5\x5c\x17\x00\x0d\xaa\x97\xd0\x5b\x60\xaa\xa1\xfc\xc3\x25\x8f\x06\x8b\x4f\xb7\xd0\xb0\x07\x47\x0d\xe0\xb6\xef\x58\x50\x66\x5f\xe9\x2b\x7d\x46\x6a\xc0\x73\x20\x85\x56\x38\x82\xad\x11\xda\x31\xc7\x79\xcf\x89\x0c\x3c\x53\x1b\x56\x49\xb0\x81\x40\xf0\x6d\x1c\xce\x76\xb1\xfb\x06\xc6\x43\x82\xa0\xc7\xb0\x41\xd9\xa7\xcc\xca\x02\x60\x83\xa2\xa3\x86\xab\xd9\x3c\x2b\x04\xb0\x60\x1d\x1e\x45\xb5\x2c\xf0\xe3\x07\xcc\x7e\x5f\xbb\x40\xb7\x1f\xe1\xe9\xa9\x2c\x0a\x45\xc9\x99\xb9\x93\x4b\x48\xd2\xd5\x50\x41\x51\xf4\xce\xd6\xc3\x5c\x95\x7f\x00\x7a\x56\x1b\xff\x01\x68\x8a\xd1\xc9\xae\x86\x72\x68\xc6\xf8\x1e\x49\xb3\xb4\xbf\xcb\x2a\xa7\x04\xcf\xa4\x29\xa2\xee\x0b\x5c\x42\xe9\xfa\xbe\x0b\xde\x65\xd3\xaa\x3b\x65\xda\x87\xf6\xc2\x4d\xf2\xbf\x16\xea\xc4\x45\xb4\x49\xfb\x18\x1c\x72\xfb\x4b\x6e\x76\xfb\x20\x00\x72\x11\xcf\x26\x7f\x66\x76\xc2\x44\xf0\x21\xa1\x1a\xf0\xf2\x0e\xbd\x65\x60\x8b\xd6\x50\x20\xba\xfb\x40\x2b\x70\xa0\x2e\xf6\x1d\x02\x6e\xd1\xa7\x9c\x0e\xcb\xdd\x80\xe4\xe2\x8b\xec\x20\x18\x70\xb2\x0b\x58\x26\x33\xa6\x4c\x6f\xa0\x15\x56\xeb\x4c\x55\x08\xe5\x1d\x6c\x82\x06\x1b\xb1\x73\xd7\xf1\xf7\x5c\xb6\x0b\x74\x9f\x0d\x54\xc4\x61\xa5\x09\x6d\x8b\x82\xe4\x31\xd7\x98\xb6\x9c\x4c\xbd\x50\x58\xba\xce\x91\xc7\x7a\x80\xb9\xf8\x74\xbb\x68\x1a\x41\x55\x78\x7a\xaa\x96\x8e\xee\xab\x69\x5d\x2b\xcf\x1a\x59\xaf\xd2\xe3\xe6\xb7\xcd\xcd\xea\x66\xdb\xd1\xf5\xc3\xf6\x7a\x63\x0f\xeb\xc7\x9b\x9b\xdd\x36\xc5\xf7\xef\xbd\xeb\xee\xdc\x87\x8d\xa4\xc7\xa3\x3b\xd9\x84\x20\xd8\xd4\x60\x92\xf0\x30\xad\x7e\x8d\xd1\xd5\x87\x31\xc0\xff\x05\xdb\x1a\xca\xff\x55\x0d\xb6\x81\x42\xf6\x43\xab\xec\xe6\x9f\xa3\x8d\xfb\x92\x82\xda\x33\xe9\x11\x2b\x40\xf9\x61\x3e\x2f\x4f\x6b\xfd\x8c\xe4\x96\x0c\x85\x5c\x07\x28\xc2\x72\x94\x56\x5e\xbf\x9a\xb6\xe8\xba\x77\xa0\x1c\x11\x58\x80\x98\x26\xaa\x4e\x10\x14\xe9\x57\x78\xfe\x27\x00\x3d\xc7\xe2\x55\x1a\xa3\xfd\x65\x51\x9c\x2c\xe5\xe4\x13\x56\x63\x2d\xdb\xf5\xf9\xb6\x8c\xc7\x7f\xdc\xf0\x70\x1a\x0e\xb7\xcc\x8d\xed\x14\xd3\x85\xed\x51\x2c\x1c\xd9\x4e\xab\x47\x6d\x53\x4d\x35\x09\xb4\x3a\xbd\x83\xad\x4b\x9d\xd5\x50\xfe\x2b\x17\x06\xb2\x67\x9b\x38\x11\xb7\x3b\xdf\xe3\xc4\xa7\x4b\xb8\x9a\x0f\x4f\xd6\x61\x2e\x18\x46\x3d\x0d\x39\x13\x7c\x70\x6a\xf4\xf1\x45\xab\xce\xcd\x18\x8e\xef\x1b\x56\x98\x38\xd2\xf6\xe4\x1d\x7b\xa9\x9b\x33\xa1\xcf\x12\xff\x32\x95\x9b\x84\xef\x87\xff\x80\x6f\xb6\xf6\x30\x50\x73\x96\x5e\x62\x9d\x23\xdf\xe8\x6e\xcc\x7f\x23\xe8\x75\x9b\x0a\xdc\x8e\xaf\xc1\x47\xf6\xfa\xcc\x27\x37\x50\x93\xdf\x5d\x88\x2c\x08\x6e\x99\xff\x7e\x36\x27\x3d\xcb\xd0\xe2\xf0\x51\x2b\xd7\x66\xbd\xd6\x55\xb5\x0a\xb6\x4e\xcb\x99\xe7\x58\x19\x52\x83\x12\x03\x59\xa5\x53\x7c\x65\x82\x58\x35\xb8\xc1\x8e\xfb\xaa\x61\xaf\x65\xf1\x57\x00\x00\x00\xff\xff\xa0\x74\xcc\x9a\xf9\x07\x00\x00")

func openapiOpenapiYmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_openapiOpenapiYmlTmpl,
		"openapi/openapi.yml.tmpl",
	)
}

func openapiOpenapiYmlTmpl() (*asset, error) {
	bytes, err := openapiOpenapiYmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi/openapi.yml.tmpl", size: 2041, mode: os.FileMode(436), modTime: time.Unix(1611431517, 0)}
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
	"openapi/index.html":       openapiIndexHtml,
	"openapi/openapi.yml.tmpl": openapiOpenapiYmlTmpl,
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
	"openapi": &bintree{nil, map[string]*bintree{
		"index.html":       &bintree{openapiIndexHtml, map[string]*bintree{}},
		"openapi.yml.tmpl": &bintree{openapiOpenapiYmlTmpl, map[string]*bintree{}},
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