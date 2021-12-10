package path

import (
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

var (
	Abs  = filepath.Abs
	Dir  = filepath.Dir
	Base = filepath.Base
	Ext  = filepath.Ext
)

func ResolvePaths(base string, paths []string) []string {
	resolved := make([]string, len(paths))
	for k, path := range paths {
		resolved[k] = Resolve(base, path)
	}
	return resolved
}

func Resolve(base, path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(base, path)
}

func Expand(path string) (string, error) {
	if filepath.IsAbs(path) || len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, path[1:]), nil
}

func SplitFilename(s string) (name, ext string) {
	parts := strings.Split(s, ".")
	switch {
	// if first part is empty then we probably dealing with dotfile, not ext
	case len(parts) < 2 || parts[0] == "":
		return s, ""
	default:
		return strings.Join(parts[:len(parts)-1], "."), parts[len(parts)-1]
	}
}

func Explode(s string) (dir, name, ext string) {
	var file string
	dir, file = path.Split(s)
	name, ext = SplitFilename(file)
	return
}
