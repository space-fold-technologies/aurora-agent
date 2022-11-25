// Code generated for package server by go-bindata DO NOT EDIT. (@generated)
// sources:
// resources/application.yml
// resources/boot.txt
// resources/settings.yml
package server

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

var _resourcesApplicationYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xca\x31\x6f\x85\x20\x10\x00\xe0\xdd\xc4\xff\x70\x1b\x53\x5f\x00\x51\x9f\xcc\x5d\x9b\x0e\xed\x6e\xae\x70\x57\x49\x9e\x9c\x01\x34\xe9\xbf\x6f\x3a\x74\x79\xe3\x97\x7c\x19\x77\xf2\x80\x67\x91\x82\x2f\xf8\x4d\xb9\xf5\xdd\x45\xa5\x26\xc9\x1e\xcc\x4d\xdf\x74\xdf\x15\x3a\xa4\xa6\x26\xe5\xe7\x79\x52\xbe\x52\x91\xbc\x53\x6e\x1e\x54\xa4\x8b\x1e\x72\xfc\x49\xf5\x5d\x90\x7d\x4f\x6d\xdd\xb0\x6e\x1e\xa2\x89\xc8\x66\x5a\x58\xdf\x27\x37\x87\x61\x64\x0e\x66\xa0\x25\xe2\x64\x2c\x8f\x73\x60\x62\x1c\x1d\xf7\xdd\xd7\x99\x1e\x71\x8d\xd8\xc8\x83\xfa\x38\x33\xbc\x87\x06\x60\x16\x30\xce\x5b\xe7\xdd\x00\x6f\xaf\x9f\x60\xb5\xb5\xea\x3f\xd3\x21\x61\x5b\x2b\x05\x0f\xca\x8c\xda\xb9\x59\xdb\xfb\xa0\x7e\x03\x00\x00\xff\xff\xac\x09\x25\xbf\xdb\x00\x00\x00")

func resourcesApplicationYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesApplicationYml,
		"resources/application.yml",
	)
}

func resourcesApplicationYml() (*asset, error) {
	bytes, err := resourcesApplicationYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/application.yml", size: 219, mode: os.FileMode(511), modTime: time.Unix(1668535923, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesBootTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x90\xe1\x09\xc4\x20\x0c\x85\xff\x0b\xee\xf0\x6d\xd0\x85\x02\x2e\xd2\x2d\x0e\x6e\xc0\x9b\xe4\x50\x5b\x94\x44\x1a\x2f\x7f\x4e\x82\x0d\xcf\x2f\xbe\xfa\xa0\xb4\x55\xbf\xd0\xf6\x72\x2b\xd3\x91\x56\x72\xe2\xa0\x2d\x39\x41\xa0\xef\x5d\x99\x8f\x94\x42\x4e\x27\x7c\xde\x2f\x55\x72\x69\xf5\xa2\xab\xb3\x50\xe3\x1e\x86\xb7\xd8\xfa\x03\xbd\x2d\x45\x79\x19\xf3\x00\x31\x84\x61\x54\x5f\xfe\x7c\x4d\x88\xc0\x18\xd9\xda\x8e\xf5\x37\xce\xd6\xfc\x5a\x13\xc9\xe1\xa6\xe8\x11\x43\x58\x19\x89\x9b\xa2\x4f\x2c\x8c\x18\x53\x77\x27\x8b\x88\x83\x94\x11\x72\x22\xb6\xfe\x31\xf8\x0d\x00\x00\xff\xff\x61\x15\x01\x0e\x38\x04\x00\x00")

func resourcesBootTxtBytes() ([]byte, error) {
	return bindataRead(
		_resourcesBootTxt,
		"resources/boot.txt",
	)
}

func resourcesBootTxt() (*asset, error) {
	bytes, err := resourcesBootTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/boot.txt", size: 1080, mode: os.FileMode(511), modTime: time.Unix(1668535630, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesSettingsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\x31\xcb\xc2\x30\x10\x80\xe1\x3d\x90\xff\x70\xf0\xcd\x4d\xf3\x39\xb4\x98\x4d\xd4\x49\x44\xa9\x83\x73\x68\xae\x36\x50\x7b\xe5\xee\x5a\xe8\xbf\x17\x5d\x94\x77\x7c\x9f\x89\x69\xc9\x09\x39\xc0\xe1\xb2\x3f\x1d\x9b\xe2\x76\xdf\x35\x67\x6b\x26\xa6\x2e\x0f\x58\xa4\xcc\xd8\x2a\xf1\x1a\xa0\x44\x6d\xcb\x38\x33\x71\x84\x3f\xb8\x46\xed\x81\x46\x78\x2b\x90\x55\x14\x9f\xa0\x04\x0f\xa2\x04\xa2\x73\xd7\x59\x33\x64\x51\x1c\x8b\x98\x12\xa3\x48\x00\xef\x3e\x59\x13\xd3\x82\xac\x59\xf0\xfb\xfe\xbd\xab\xdd\xb6\x72\x95\x35\x3d\x89\xfe\xe0\x89\x58\x03\x6c\x6a\xef\x5f\x01\x00\x00\xff\xff\xc7\x10\x37\x4d\xac\x00\x00\x00")

func resourcesSettingsYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesSettingsYml,
		"resources/settings.yml",
	)
}

func resourcesSettingsYml() (*asset, error) {
	bytes, err := resourcesSettingsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/settings.yml", size: 172, mode: os.FileMode(511), modTime: time.Unix(1669070313, 0)}
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
	"resources/application.yml": resourcesApplicationYml,
	"resources/boot.txt":        resourcesBootTxt,
	"resources/settings.yml":    resourcesSettingsYml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"application.yml": &bintree{resourcesApplicationYml, map[string]*bintree{}},
		"boot.txt":        &bintree{resourcesBootTxt, map[string]*bintree{}},
		"settings.yml":    &bintree{resourcesSettingsYml, map[string]*bintree{}},
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
