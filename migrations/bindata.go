package migrations

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

var __000001_todo_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xb1\x6e\x83\x30\x10\x87\xf1\xdd\x4f\xf1\x17\x53\x2b\xb5\x4f\xc0\x64\xca\xb5\x42\x75\xdd\xc8\x31\x83\xc7\x24\x5c\xd0\x0d\xd8\x11\x3e\x24\x1e\x3f\x62\xcd\xfe\xe9\xd3\xef\xca\xb3\xe4\xd6\x98\xdb\xca\x17\x65\xf0\xae\x9c\xab\x94\x0c\xb9\x23\x17\x05\xef\x52\xb5\xa2\xd9\x36\x99\x3e\x4b\xad\x8f\xa6\x35\xe6\x2b\x90\x8d\x84\x68\x3b\x47\x2f\xa1\x96\xa9\x54\xbc\x19\x99\x00\x40\xb2\xf2\xcc\x2b\x4e\x61\xf8\xb3\x21\xe1\x97\x12\x7e\xc8\x53\xb0\x91\x7a\x74\x09\x3d\x7d\xdb\xd1\x45\xd8\x33\x86\x9e\x7c\x1c\x62\xfa\x30\xc7\x03\x50\xde\x15\xfe\x3f\xc2\x8f\xce\x99\xf7\xc3\x58\x96\x45\xb4\x7d\x06\x00\x00\xff\xff\x26\x89\x8c\x81\xb3\x00\x00\x00")

func _000001_todo_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_todo_up_sql,
		"000001_todo.up.sql",
	)
}

var __000002_todo_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4a\x4d\xcf\xcc\xb3\xe6\xe5\xe2\xe5\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc9\x4f\xc9\x2f\xe6\xe5\x72\x74\x71\x51\x70\xf6\xf7\x0b\x0e\x09\x72\xf4\xf4\x0b\x01\x8b\x2a\x84\xfa\x79\x06\x86\xba\x2a\x68\x80\x38\x9a\x60\x8d\xc9\xf9\xb9\xb9\x99\x25\xd6\x80\x00\x00\x00\xff\xff\x49\x95\x8d\x26\x4a\x00\x00\x00")

func _000002_todo_up_sql() ([]byte, error) {
	return bindata_read(
		__000002_todo_up_sql,
		"000002_todo.up.sql",
	)
}

var __000003_todo_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x41\x6b\x83\x40\x10\x85\xef\x82\xff\xe1\x1d\x15\x72\x2a\xcd\x29\xa7\xb4\x78\x28\x04\x4b\x8b\x9e\x65\xe2\x8e\xe9\x50\x77\x57\x66\xc7\x90\xfc\xfb\xa2\x5e\xec\x65\x61\x3f\xde\x7c\xef\x5d\xf9\x26\xe1\x94\x67\x79\xf6\xfe\x5d\x9d\x9b\x0a\xcd\xf9\xed\x52\x41\x06\x84\x68\xe0\x87\x24\x4b\xb0\xe8\x62\xea\xcc\x4f\x28\xf2\x0c\x00\xc4\x2d\xef\x9d\xb4\xff\x21\x2d\x5e\x8e\xc7\x12\x93\x8a\x27\x7d\xe2\x97\x9f\x70\x3c\xd0\x3c\x1a\xe6\x59\x5c\x77\xe3\xc0\x4a\xc6\xdd\xfd\xb5\x28\x0f\xdb\xfd\x22\x04\x8c\x1f\x86\xfa\xb3\x41\xdd\x5e\x2e\x68\xeb\x8f\xaf\xb6\xca\xb3\x72\x5d\x23\x21\xb1\x1a\x24\x58\xdc\xd7\x8b\x3b\xac\xdf\x32\xcf\x12\x8f\xdc\x1b\x7a\x4a\x56\x88\x03\xa5\x7f\x7b\xca\x2d\x87\x41\xa3\xdf\x04\xab\xd6\x69\x9c\x60\x74\x1d\x79\x07\x69\x34\xd6\x3d\x5d\xbb\x94\x03\xf9\x05\xec\x92\x7d\xf4\x5e\xec\xf4\x17\x00\x00\xff\xff\xd5\x07\x1a\x1c\x35\x01\x00\x00")

func _000003_todo_up_sql() ([]byte, error) {
	return bindata_read(
		__000003_todo_up_sql,
		"000003_todo.up.sql",
	)
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

var _generate_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\x2b\xe6\xe2\xd2\xd7\x4f\xcf\xb7\x4a\x4f\xcd\x4b\x2d\x4a\x2c\x49\x55\x48\xcf\xd7\x4d\xca\xcc\x4b\x49\x2c\x49\x54\xd0\x2d\xc8\x4e\x47\x52\xa9\xa0\x9b\xaf\x00\x95\xd2\x4b\xcf\xe7\xd4\xe3\x02\x04\x00\x00\xff\xff\xcc\xf6\x80\x91\x4d\x00\x00\x00")

func generate_go() ([]byte, error) {
	return bindata_read(
		_generate_go,
		"generate.go",
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
	"000001_todo.up.sql": _000001_todo_up_sql,
	"000002_todo.up.sql": _000002_todo_up_sql,
	"000003_todo.up.sql": _000003_todo_up_sql,
	"bindata.go":         bindata_go,
	"generate.go":        generate_go,
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_todo.up.sql": &_bintree_t{_000001_todo_up_sql, map[string]*_bintree_t{}},
	"000002_todo.up.sql": &_bintree_t{_000002_todo_up_sql, map[string]*_bintree_t{}},
	"000003_todo.up.sql": &_bintree_t{_000003_todo_up_sql, map[string]*_bintree_t{}},
	"bindata.go":         &_bintree_t{bindata_go, map[string]*_bintree_t{}},
	"generate.go":        &_bintree_t{generate_go, map[string]*_bintree_t{}},
}}
