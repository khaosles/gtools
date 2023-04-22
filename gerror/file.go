package gerror

/*
   @File: file.go
   @Author: khaosles
   @Time: 2023/2/23 23:36
   @Desc:
*/

var ()

var (
	IOException            = Exception("IO")
	FileNotFoundException  = Exception("FileNotFound")
	FileNotAccessException = Exception("FileNotAccess")
	FileExistException     = Exception("FileExist")
)
