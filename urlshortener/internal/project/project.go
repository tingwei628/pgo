package project

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	WorkingDir = filepath.Join(filepath.Dir(b), "../..")
)
