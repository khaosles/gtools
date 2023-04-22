package gerror

import "fmt"

/*
   @File: base.go
   @Author: khaosles
   @Time: 2023/2/23 23:38
   @Desc:
*/

type GError interface {
	Error() string
	New(s any) *gerror
}

func Error(s string) GError {
	return &gerror{s + "Error", ""}
}

func Exception(s string) GError {
	return &gerror{s + "Exception", ""}
}

type gerror struct {
	prefix string
	err    string
}

func (e *gerror) Error() string {
	if e.err == "" {
		return e.prefix
	} else {
		return fmt.Sprintf("%v: %v", e.prefix, e.err)
	}
}

func (e *gerror) New(s any) *gerror {
	e.err = fmt.Sprintf("%v", s)
	return e
}
