package gcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	tsgutils "github.com/typa01/go-utils"

	"github.com/khaosles/dor/gtools/gerror"
)

/*
   @File: gcm.go
   @Author: khaosles
   @Time: 2023/3/4 17:23
   @Desc:
*/

// GetKeyByIp 根据ip生成token密钥 16位
func GetKeyByIp(ip string) string {

	length := len(ip)
	if length > 16 {
		return ip[:16]
	}
	for i := length; i < 16; i++ {
		ip += "a"
	}
	return ip
}

// AesGcmEncrypt  GCM加密 每次生成随机密文 后面24位为iv
func AesGcmEncrypt(data, key string) string {
	secretKey := []byte(key)
	plaintext := []byte(data)
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(gerror.KeyLengthException.New(err))
	}
	// 生成随机iv
	iv := tsgutils.GUID()[:24]
	nonce, _ := hex.DecodeString(iv)
	// gcm模式
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(gerror.PasswordEncryptException.New(err))
	}
	ciphertext := aesGcm.Seal(nil, nonce, plaintext, nil)
	return fmt.Sprintf("%x%v", ciphertext, iv)
}

// AesGcmDecrypt GCM解密
func AesGcmDecrypt(data, key string) (string, error) {
	// iv
	iv := data[len(data)-24:]
	// 密文
	data = data[:len(data)-24]
	secretKey := []byte(key)
	ciphertext, _ := hex.DecodeString(data)
	nonce, _ := hex.DecodeString(iv)
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}
	// gcm模式
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// 解密
	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
