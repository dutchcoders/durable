// +build !windows

package durable

// Original: https://github.com/nsqio/nsq/blob/master/nsqd/rename.go

import (
	"os"
)

func atomicRename(sourceFile, targetFile string) error {
	return os.Rename(sourceFile, targetFile)
}
