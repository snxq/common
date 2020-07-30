package sys

import (
	"os"
	"path/filepath"
	"time"

	"github.com/djherbis/atime"
)

// Du impl command du.
func Du(dir string) int64 {
	var size int64
	_ = filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}

// GetATime return file's ATime.
func GetATime(path string, defaultTime time.Time) time.Time {
	at, err := atime.Stat(path)
	if err != nil {
		return defaultTime
	}
	return at
}
