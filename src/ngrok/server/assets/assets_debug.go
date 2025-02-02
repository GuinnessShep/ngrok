// +build debug

package assets

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// assets_server_tls_snakeoil_crt reads file data from disk. It returns an error on failure.
func assets_server_tls_snakeoil_crt() ([]byte, error) {
	return bindata_read(
		"/Users/rphlo/ngrok/assets/server/tls/snakeoil.crt",
		"assets/server/tls/snakeoil.crt",
	)
}

// assets_server_tls_snakeoil_key reads file data from disk. It returns an error on failure.
func assets_server_tls_snakeoil_key() ([]byte, error) {
	return bindata_read(
		"/Users/rphlo/ngrok/assets/server/tls/snakeoil.key",
		"assets/server/tls/snakeoil.key",
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
	"assets/server/tls/snakeoil.crt": assets_server_tls_snakeoil_crt,
	"assets/server/tls/snakeoil.key": assets_server_tls_snakeoil_key,
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
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"server": &_bintree_t{nil, map[string]*_bintree_t{
			"tls": &_bintree_t{nil, map[string]*_bintree_t{
				"snakeoil.crt": &_bintree_t{assets_server_tls_snakeoil_crt, map[string]*_bintree_t{
				}},
				"snakeoil.key": &_bintree_t{assets_server_tls_snakeoil_key, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}
