package g

import (
	"strings"

	"github.com/google/uuid"
)

/*
   @File: ID.go
   @Author: khaosles
   @Time: 2023/4/30 16:39
   @Desc:
*/

func UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", 4)
}
