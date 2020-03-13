package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _pokes_gengar_cow = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x95\xb9\x52\xf3\x30\x10\xc7\x7b\x3d\xc5\x4e\x26\xdd\xf7\x15\xbe\x72\x08\x87\x2a\xa4\x25\x05\x49\x95\xcc\x30\x03\x18\x42\x21\x52\xe0\x0c\x05\x93\x77\x67\x42\x6c\x1d\xd6\xae\xbc\x32\xa9\x8c\x58\xed\xf1\xdb\xbf\x76\x9f\xde\x3f\xd4\xf1\xa5\x82\x87\xcd\xdd\x7a\xbb\xf9\x0f\xa3\x9b\x53\xfd\x3a\x1f\x95\x62\x5c\x1f\xaa\xc7\xe7\xe3\x17\xdc\x2e\x16\xab\xf5\xb2\x14\x60\x7e\xe3\xfa\x70\x3c\xbd\x1d\xea\x4f\xfb\x90\x38\xa5\x8e\xad\xf3\x7d\xb5\x2b\xa4\x02\xef\xb7\xaf\x76\xf9\xbc\x9c\x94\x59\x3e\x55\xfb\xfb\xef\xed\xbf\x6c\x32\x2f\xce\x17\x6b\x7d\xda\x58\xc8\x99\x65\x00\x57\x7f\xc2\x36\xd4\xbe\xd2\xa4\xb0\x7d\x39\x5e\xa5\x22\x22\x1a\xd7\x00\xe8\x31\x27\x54\x20\x00\x44\x54\xe7\x9b\xca\x99\x0a\xa0\x92\xfa\xef\xe4\x2c\xbc\x44\x6d\x07\x85\x4e\x9a\xa4\xda\xcd\x6c\x10\x4d\x2c\x85\x8b\x23\x46\x04\xd7\xde\xbd\x80\x71\x31\x19\x48\xe5\x76\xb1\x9b\x84\x54\x02\x39\x6d\xe9\xf4\xe3\x41\x41\xc9\x60\xdf\xae\x35\xb4\x4a\x37\x57\x48\xf9\x78\x09\xf7\xc9\x15\x03\x3a\x4d\x22\x51\x41\x5c\xf5\xee\x13\x41\x14\xe6\x74\x37\xcf\x98\x4f\xc4\x7c\x09\x30\x82\x6e\x7d\x3a\x45\x01\xc4\x28\xe0\x37\x05\xb4\xa6\xf6\x66\x22\x09\x62\x9e\xa9\x5b\x0e\x63\x86\xa5\xb3\x5c\x39\xcd\x07\xab\x22\x94\x98\xd7\x3e\x0a\x98\x18\x20\x83\x26\x28\x02\x0c\x13\x68\x17\x81\x79\x2d\x80\x46\x61\x21\x2f\x34\x72\x52\x22\x2d\xb8\x70\x83\x88\xa1\x93\xb0\xf2\x8f\xa0\x80\xc7\x71\xf1\xc6\x0e\x64\xdb\xff\x80\xf7\xca\x71\xd3\xdc\x4d\x63\x9c\x0f\xf8\xc2\x80\x15\x92\x9d\x4a\x16\xdc\x22\x2e\x49\x72\x7a\x47\x14\xd3\xab\x72\x5c\x2f\x78\xd7\xb5\xb3\x34\xdc\x4c\x0f\x88\xbe\x06\xb6\x09\x74\xff\x83\xd3\xf4\xb8\x51\x32\x67\x72\x1d\x3e\x46\xfc\xa1\xd6\x0c\x07\x5c\xad\x17\x57\x6c\xe6\xbd\x3d\xe5\xbe\xa4\xde\x7d\x9a\x9c\x7b\xd6\x0d\xb1\xee\xbb\xf3\xf0\x2f\x35\x12\xdb\x9d\x6c\x5a\x38\x25\xf6\x6c\xb2\x77\x79\x42\x6c\x3e\x5e\x36\xb9\x54\x42\xac\xd6\x4b\xf1\x13\x00\x00\xff\xff\x77\xa7\x76\xfe\x6f\x0c\x00\x00")

func pokes_gengar_cow() ([]byte, error) {
	return bindata_read(
		_pokes_gengar_cow,
		"pokes/Gengar.cow",
	)
}

var _pokes_pikachu_cow = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xbf\x6e\xf2\x30\x10\xdf\xfd\x14\x27\xc4\xf6\x7d\x83\xed\x04\x08\x0d\x9d\x28\x6b\x19\x0a\x13\x48\x95\x68\xd3\xd2\x21\x65\x68\x50\x2b\x55\x7d\xf7\x2a\x11\xc4\xe7\xe4\xce\x3e\x28\x53\xb8\x9c\x7d\xf7\xfb\xe3\x8b\x77\x6f\xef\xe5\xe1\xb9\x80\x87\xd5\xdd\x72\xbd\xfa\x0f\x83\x9b\x63\xf5\x92\x0d\x72\x35\xac\xf6\xc5\xe3\xd3\xe1\x13\x6e\x67\xb3\xc5\x72\x9e\x2b\x70\xbf\x61\xb5\x3f\x1c\x5f\xf7\xd5\x07\x0e\x32\x51\x2e\x8c\xe2\xdb\x2f\xb3\xdb\xa4\xd3\x12\xa0\x79\x4a\xb2\x7c\x94\xdb\x64\x5c\x6e\xef\xbf\xd7\xff\xec\x28\x4b\x7f\xdc\x53\xbb\x7c\x5b\x6c\xd2\x73\x62\xfd\xe7\xb4\x6a\x8a\x57\x15\xf5\xae\xee\x1d\xde\x51\xf9\x69\x6d\xc2\x29\xd9\x4c\xb2\xb2\x9f\x51\x47\x5d\x61\x6b\x88\x14\x9b\x4c\x5c\x45\xab\x85\xdd\x84\x9f\x28\xb8\x69\x0b\xf7\xdc\x71\x6a\xb9\x76\x84\x84\xb4\x61\xb7\x44\x3b\xe8\xe4\xf6\x0d\x4d\x64\x03\xe0\xf3\x14\xa4\x95\x87\xe1\xe3\x76\xbc\xb2\xcc\xb3\x50\xf1\x3b\x20\xe0\x69\xc3\x49\x83\x09\x88\xc0\x38\x49\xc3\xd5\x55\x40\x55\xbe\x9a\x58\x49\x47\x51\x7f\x58\x8c\xbb\x73\x78\x08\x13\x93\x7a\x98\x2e\xf3\x8c\xfb\xfb\xda\x21\x30\x08\xf0\x35\xfa\x8e\x7d\x9e\xe1\x2f\xc7\x3b\xc0\x17\xe6\x5c\xda\x90\xee\x3b\xa4\x65\x2d\x8c\xd8\x1a\x34\xda\xba\x74\x86\xbc\x8a\xc9\x4c\xe8\xad\x75\xcc\x3b\x75\x71\xbe\x3a\x31\xec\xba\x75\x82\xf2\x5c\x30\x6c\x80\x2a\xd7\x97\x87\x9d\xad\xe4\x04\xe1\xbc\x89\xe7\x96\x16\xe8\xd3\x66\xe8\x84\x3b\x80\x91\x4e\x2e\x55\xbe\xaf\xb6\xc8\x48\x3c\x39\xf4\xb8\xd2\x8d\x4e\x20\xe7\xd4\x90\x46\x6b\xa2\xec\xd7\xb9\x89\xa6\x3a\xa2\x8b\x11\xf1\xc5\x0c\x2c\xa7\x67\xad\x50\x64\xd4\x71\x47\x45\xf6\x31\x13\xcf\x0d\x9a\x4a\x8a\x1c\x0f\x3a\x35\x6f\x7c\xb5\x48\xb9\xf4\x25\x35\x78\x3b\x42\xc4\xd4\x58\x82\x34\xec\x58\xfc\x5d\x62\x24\xc7\x37\xad\x9a\x3c\xfa\x22\x40\x30\xa2\x40\x4a\x04\xbb\x89\x83\x1c\xf3\x7f\x4f\x6c\xb1\x50\x40\xdd\x05\xc2\xe7\x54\x82\xde\xdb\x33\x99\x96\x4a\x2d\x96\x73\xf5\x1b\x00\x00\xff\xff\x59\xff\x9b\x8c\xe5\x0b\x00\x00")

func pokes_pikachu_cow() ([]byte, error) {
	return bindata_read(
		_pokes_pikachu_cow,
		"pokes/Pikachu.cow",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"pokes/Gengar.cow": pokes_gengar_cow,
	"pokes/Pikachu.cow": pokes_pikachu_cow,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"pokes": &_bintree_t{nil, map[string]*_bintree_t{
		"Gengar.cow": &_bintree_t{pokes_gengar_cow, map[string]*_bintree_t{
		}},
		"Pikachu.cow": &_bintree_t{pokes_pikachu_cow, map[string]*_bintree_t{
		}},
	}},
}}
