package gcmd

import "testing"

/*
   @File: command_test.go
   @Author: khaosles
   @Time: 2023/5/25 20:14
   @Desc:
*/

func TestSync(t *testing.T) {
	Sync("echo", "你好")
}
