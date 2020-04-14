// Code generated by go-bindata.
// sources:
// static/javascript/catalog.js
// static/javascript/sfomuseum.geotag.init.js
// static/javascript/sfomuseum.maps.js
// DO NOT EDIT!

package sfomuseum

import (
	"github.com/whosonfirst/go-bindata-assetfs"
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

var _staticJavascriptCatalogJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x97\xe1\x6a\x83\x30\x14\x85\x5f\x45\xf2\xbb\xab\xb9\x51\x63\xd2\x57\x19\x63\xd8\xce\x76\x82\xd6\x62\x2a\x74\x15\xdf\x7d\xb8\xc1\xa0\xb7\x83\xc8\x21\xf9\x69\xa1\xf1\xe3\xe4\xc4\xef\xe6\x75\x12\xe3\xd0\x8a\x5d\x22\x3e\xaf\xd7\x8b\xdb\xa5\x69\xd7\xb4\xad\x3b\x36\x75\xfb\xb1\x75\xc7\xbe\x1b\x5d\x3d\x76\xdb\x7e\x38\xa5\x55\x3d\x34\x55\x9b\x92\xcd\xca\x74\xba\xcf\xe9\x74\x9b\xd3\xe9\xe5\x6b\xde\x5e\xce\x27\xb1\x49\x84\xeb\xc7\xe1\x50\x2f\x4b\xb9\x63\x7f\x6a\xdc\xf2\x5b\xd7\x9c\xdf\xef\x7d\xdf\x89\x5d\x42\x6a\x79\xae\x6e\x7f\xcf\xe5\x26\x11\x6d\xb5\xaf\x7f\xde\xbe\xac\x2a\xe6\x4d\x02\xe0\xe4\xe4\xc1\x19\x0f\x6e\xdf\x36\x7b\xce\x93\x31\x1e\xf3\xc8\x93\x13\xca\x93\x05\x89\xa7\x60\x38\x19\x8a\xa3\xb1\x78\x88\xf1\x68\xc6\xa3\x51\x1e\xac\x3d\x1c\x87\xc7\x03\xb7\xc7\xfa\x71\x7e\xff\xe7\x23\xca\x19\x91\x05\x89\x0a\x19\xa3\x3f\x85\x44\x71\xe2\xf4\xa7\x40\xfb\xa3\xc3\xc4\xc3\x70\x34\x1a\x8f\x2e\xb0\x78\x3c\x5f\x43\x5d\x80\x3c\xe5\x8a\x78\xd6\xf5\x99\x15\xa8\x44\x13\x2a\x55\x2c\x22\x85\x12\x99\x18\x02\x2b\x0d\x88\x63\xc2\x34\x9a\xe1\x18\x74\xbf\x8c\xcf\xa7\x18\x0e\xaa\x53\xe3\x3b\x60\xff\xe3\x78\xec\x6e\xd0\xf3\x65\xa2\x74\xc7\xc0\xdd\x59\xa1\x2f\x00\x07\x75\x97\x8d\x32\x1a\x5a\x54\xee\x36\xca\x66\x59\x74\xb3\x2c\x3c\x6b\x78\xe4\x6e\xb1\xfd\x52\x52\x46\xf8\xf4\x2c\xab\xa2\x38\x2b\x54\xb1\x62\x36\xd4\x0c\x07\xf3\x84\x92\x32\x8f\x92\x4e\x8e\xe2\xf8\x26\xb1\xd5\x22\xe5\x01\x61\xb3\x98\x92\x14\xa6\x3e\xf6\x11\x87\xd0\xfa\x10\x56\x1f\x1f\x0e\x5a\x1f\x0a\x73\x11\x64\xf5\x21\xec\x22\xa8\x24\x45\x69\x33\xa1\x6d\x26\xcc\xeb\x0c\x47\x49\x86\x83\x79\x5d\x49\x5a\x71\xb8\x00\x1c\xf8\x64\x05\x11\xe9\x13\x0e\x26\x52\x25\x29\x88\x48\x9f\x70\x30\x91\x2a\x49\x41\xa6\x9e\x27\x1c\x2b\xe6\xb7\xef\x00\x00\x00\xff\xff\x60\x62\xc6\x68\x11\x13\x00\x00")

func staticJavascriptCatalogJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptCatalogJs,
		"static/javascript/catalog.js",
	)
}

func staticJavascriptCatalogJs() (*asset, error) {
	bytes, err := staticJavascriptCatalogJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/catalog.js", size: 4881, mode: os.FileMode(420), modTime: time.Unix(1586891415, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptSfomuseumGeotagInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xc1\x6a\xc3\x30\x0c\x86\xcf\xcd\x53\x68\x3e\xd9\x50\xfc\x00\x2b\xbd\x6c\x2b\x6c\xe0\x51\xd8\x65\xc7\x21\x62\xc5\x04\x6c\xcb\xd8\x4e\x43\x09\x7b\xf7\x91\xb8\x97\xd2\x5d\x04\x92\x7e\xfd\xff\x87\xe6\x31\x5a\x9e\x35\x5a\x7b\xba\x50\xac\x66\x2c\x95\x22\x65\x29\x3c\xa3\x15\x7b\x18\xa6\xd8\xd7\x91\x23\xac\xbd\xa4\x55\xa3\x96\xae\x03\x00\xe8\x39\x16\xf6\xa4\x3d\x3b\x29\xde\x4f\xc6\x9c\x85\x3a\x6c\x9b\xad\x8c\x03\xc8\x27\x30\xfa\x95\x63\xcd\xec\xb5\xc1\x2b\xe5\xa2\x96\x6e\x97\xa9\x4e\x39\x36\xe9\xef\x3f\x5e\xdf\xe7\x2f\xf3\x76\xe7\x75\xc1\x0c\x01\x13\x1c\xc1\x11\x57\x74\x3a\x60\x2a\xda\x51\xfd\xc4\xf4\x72\xfd\xb0\x52\x04\x4c\x0f\x17\x7e\x4b\xfc\xe9\x1b\x00\x1c\x21\xd2\xfc\x00\x24\x97\x6e\xd7\x63\x45\xcf\xee\x19\xca\xc0\x61\x2a\x34\x85\x16\x70\x9b\xef\x1b\xa9\x3a\x34\xd6\x80\x69\xfd\xd7\xcd\x46\xde\xa7\xa8\x43\xb7\x0a\xff\x02\x00\x00\xff\xff\x8d\xf3\xb8\xe8\x57\x01\x00\x00")

func staticJavascriptSfomuseumGeotagInitJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptSfomuseumGeotagInitJs,
		"static/javascript/sfomuseum.geotag.init.js",
	)
}

func staticJavascriptSfomuseumGeotagInitJs() (*asset, error) {
	bytes, err := staticJavascriptSfomuseumGeotagInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/sfomuseum.geotag.init.js", size: 343, mode: os.FileMode(420), modTime: time.Unix(1586901931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptSfomuseumMapsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x98\xef\x8a\xa3\x30\x14\xc5\x3f\xaf\x4f\x11\xf2\xa9\x85\x6e\x4d\x52\xff\x24\x2d\x7d\x92\x65\x59\x6c\x57\xbb\x42\xd4\x62\x74\x99\x1d\xc7\x77\x5f\xb4\x30\xbb\x93\x0e\x44\x0e\xe6\x53\x6b\xa0\xb7\x3f\xce\x3d\xd7\x73\xf5\x77\xd6\x12\x53\x34\x55\x6f\xf2\xbe\x22\xe7\xff\xbe\xbf\xbd\x91\x61\x3c\x05\xc1\xfb\xc9\xbe\xca\xee\x86\x9c\xc9\xa6\xe8\xeb\x6b\x57\x36\xf5\x66\x3b\x04\x01\x21\x84\xcc\x45\x72\x5d\x90\x33\x19\x82\xe0\xcb\x35\xeb\x32\xdd\xdc\x8e\xe4\xdb\x40\xfb\x56\xd3\x23\xa1\xbf\xba\xee\x6e\x8e\x61\x58\x95\x5a\x9b\xa2\xcc\xf5\xcf\xfd\xbf\xba\x4d\x7b\x0b\xb3\xbc\x2d\x33\x1d\x72\x75\x48\xc3\xe1\x75\x0c\x87\x97\x31\x1c\xbe\xfe\x19\xf7\xf7\xfa\x46\x77\x84\x9a\xa6\x6f\xaf\xf9\x54\xca\x14\xcd\xad\x34\xd3\x59\x55\xd6\x3f\x5e\x9b\xa6\xa2\x47\xc2\xc5\x74\x9d\xbd\xbc\x5f\xa7\x3b\x42\x75\x76\xc9\xe7\x7f\x9f\xaa\xd2\x71\x47\x00\x9c\x88\x3b\x70\xfa\xab\xb9\xe8\xf2\x62\xf3\x1c\x2c\x1e\xf9\x91\x27\xe2\x28\xcf\x61\x15\x79\x62\x0b\xe7\x80\xe2\x24\x98\x3c\xdc\xe2\x49\x2c\x9e\x04\xe5\xc1\xdc\x63\xe3\xd8\xf2\xc0\xee\x51\x6e\x9c\xc7\xef\x5c\x44\x91\x45\xa4\x40\xa2\x98\xf9\xf0\x4f\xcc\x50\x1c\x3f\xfe\x89\x51\xff\x24\xeb\xc8\x63\xe1\x24\xa8\x3c\x49\x8c\xc9\xe3\xb8\x1b\x26\x31\xc8\x93\x2e\x90\x67\x99\x9f\x2d\x03\xa5\xa8\x42\xa9\xf0\x45\x24\x50\x22\xe9\x23\xc0\x52\x09\xe2\xc8\x75\x1c\x6d\xe1\x48\xb4\x5f\xd2\x95\xa7\x18\x0e\x1a\xa7\xd2\x35\x60\x9f\xe3\x38\xd2\x5d\xa2\xf3\x25\xbd\x78\x47\xc2\xde\x59\x10\x5f\x00\x0e\x9a\x5d\xca\xcb\x6a\xa8\xd0\x70\x57\x5e\x9a\xa5\xd0\x66\x29\x78\xd7\x70\x84\xbb\xc2\xfa\x25\x18\xf3\x70\xeb\x99\xaa\xa2\x38\x0b\xa2\x62\xc1\x6e\x98\x58\x38\x58\x4e\x08\xc6\x22\x2f\xea\x44\x28\x8e\x6b\x13\x5b\x1c\xa4\xb6\x40\xd8\x2e\x26\x18\x5f\xc7\x3e\xea\x23\x0e\x47\xed\xc3\x31\xfb\xb8\x70\x50\xfb\xf0\x75\x1e\x04\x2d\xfb\x70\xec\x41\x50\x30\xee\xc5\xcd\x1c\x75\x33\xc7\x72\xdd\xc2\x11\xcc\xc2\xc1\x72\x5d\x30\xbe\x60\xb8\x00\x1c\x78\xb2\x56\x09\xd2\x27\x1c\x2c\x48\x05\xe3\xab\x04\xe9\x13\x0e\x16\xa4\x82\xf1\x55\xb6\x9e\x27\x1c\x45\xc7\xef\xbb\xc7\xfb\xb2\xf1\xf4\xf8\x6c\xf3\xae\x6f\xeb\xf9\xd5\xd9\x69\x3e\x08\xc6\xed\x66\x7b\x0a\xfe\x06\x00\x00\xff\xff\xc2\xcd\x22\xa5\x95\x13\x00\x00")

func staticJavascriptSfomuseumMapsJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptSfomuseumMapsJs,
		"static/javascript/sfomuseum.maps.js",
	)
}

func staticJavascriptSfomuseumMapsJs() (*asset, error) {
	bytes, err := staticJavascriptSfomuseumMapsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/sfomuseum.maps.js", size: 5013, mode: os.FileMode(420), modTime: time.Unix(1586900418, 0)}
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
	"static/javascript/catalog.js": staticJavascriptCatalogJs,
	"static/javascript/sfomuseum.geotag.init.js": staticJavascriptSfomuseumGeotagInitJs,
	"static/javascript/sfomuseum.maps.js": staticJavascriptSfomuseumMapsJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"javascript": &bintree{nil, map[string]*bintree{
			"catalog.js": &bintree{staticJavascriptCatalogJs, map[string]*bintree{}},
			"sfomuseum.geotag.init.js": &bintree{staticJavascriptSfomuseumGeotagInitJs, map[string]*bintree{}},
			"sfomuseum.maps.js": &bintree{staticJavascriptSfomuseumMapsJs, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
