package g

import (
	"runtime"
)

/*
   @File: system.go
   @Author: khaosles
   @Time: 2023/5/25 15:05
   @Desc:
*/

func IsWindows() bool {
	// darwin linux
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
