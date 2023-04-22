package gcrypt

import "golang.org/x/crypto/bcrypt"

/*
   @File: pw.go
   @Author: khaosles
   @Time: 2023/2/27 23:26
   @Desc:
*/

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password, salt string) string {
	// 第一次加密
	encrypt1, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// aes加密
	encrypt2 := AesCbcEncrypt(string(encrypt1), salt)
	return encrypt2
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash, salt string) bool {
	// aes解密
	password = AesCbcDecrypt(password, salt)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
