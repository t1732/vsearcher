package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	_rootPath  = filepath.Dir(filepath.Join(b, "../../"))
	App        = &Application{RootPath: &rootPath{}}
)

type Application struct {
	RootPath *rootPath
}

type rootPath struct{}

func (r *rootPath) ToString() string {
	return _rootPath
}

func (r *rootPath) Join(path string) string {
	return filepath.Join(r.ToString(), path)
}
