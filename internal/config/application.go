package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Dir(filepath.Join(b, "../../"))
)
