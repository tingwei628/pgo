package utils

import (
	"path/filepath"
	"runtime"
)

var Basepath = func() string {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../")
	// blog root path
	return root
}()
