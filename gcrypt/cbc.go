package gcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/khaosles/gtools/gerror"
)

/*
   @File: gaes.go
   @Author: khaosles
   @Time: 2023/3/2 23:12
   @Desc:
*/

func AesCbcEncrypt(text, salt string) string {

	if len(salt) != 16 && len(salt) != 24 && len(salt) != 32 {
		panic(gerror.KeyLengthException.New(fmt.Sprintf("salt length error -- %v", len(salt))))
	}

	// 将密钥转换为字节数组
	keyBytes, _ := hex.DecodeString(salt)
	// 创建一个AES加密块
	block, _ := aes.NewCipher(keyBytes)
	iv := make([]byte, aes.BlockSize)
	// 创建一个CBC加密模式
	mode := cipher.NewCBCEncrypter(block, iv)
	// 填充原文，使其长度为16的倍数
	plaintextBytes := []byte(text)
	plaintextBytes = PKCS5Padding(plaintextBytes, aes.BlockSize)
	// 加密
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	return hex.EncodeToString(ciphertext)
}

func AesCbcDecrypt(ciphertext_, salt string) string {
	if len(salt) != 16 && len(salt) != 24 && len(salt) != 32 {
		panic(gerror.KeyLengthException.New(fmt.Sprintf("salt length error -- %v", len(salt))))
	}
	ciphertext, _ := hex.DecodeString(ciphertext_)
	keyBytes, _ := hex.DecodeString(salt)
	// 创建一个AES加密块
	block, _ := aes.NewCipher(keyBytes)
	iv := make([]byte, aes.BlockSize)
	// 创建一个CBC解密模式
	mode := cipher.NewCBCDecrypter(block, iv)
	// 解密
	plaintextBytes := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintextBytes, ciphertext)
	// 去除填充
	plaintextBytes = PKCS5UnPadding(plaintextBytes)

	return string(plaintextBytes)
}

// PKCS5Padding 填充函数
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 去除填充函数
func PKCS5UnPadding(plaintext []byte) []byte {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	return plaintext[:(length - unpadding)]
}
