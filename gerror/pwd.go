package gerror

/*
   @File: pwd.go
   @Author: khaosles
   @Time: 2023/3/2 23:46
   @Desc:
*/

var (
	KeyLengthException       = Exception("KeyLength")
	PasswordDecryptException = Exception("PasswordDecrypt")
	PasswordEncryptException = Exception("PasswordEncrypt")
)
