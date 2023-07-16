package cbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func pkcs7Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(plainText, padtext...)
}

func pkcs7Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])

	return origData[:(length - unpadding)]
}

// aesCBCEncrypt aes加密，填充秘钥key的16位，24，32分别对应AES-128, AES-192, or AES-256
func aesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	//fmt.Printf("encrypt key: %v\n", string(key))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充原文
	blockSize := block.BlockSize()
	rawData = pkcs7Padding(rawData, blockSize)
	// 初始向量IV必须是唯一，但不需要保密
	//cipherText := make([]byte, blockSize+len(rawData))
	cipherText := make([]byte, len(rawData))
	// block大小 16
	//iv := cipherText[:blockSize]
	iv := key[:blockSize]
	//if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//	return nil, err
	//}

	// block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	//mode.CryptBlocks(cipherText[blockSize:], rawData)
	mode.CryptBlocks(cipherText, rawData)

	return cipherText, nil
}

func aesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	//iv := encryptData[:blockSize]
	iv := key[:blockSize]
	//encryptData = encryptData[blockSize:]

	// CBC mode always worksin whole blocks
	if len(encryptData)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = pkcs7Unpadding(encryptData)
	return encryptData, nil
}

// 采用sha1进行2次hash，取32位
// func sha1Hash2(key []byte) []byte {
// 	h := sha1.New()
// 	h.Write(key)
// 	hashData := h.Sum(nil)
// 	//fmt.Printf("hash data 1: %v\n", hashData)
// 	keyBuffer := bytes.NewBuffer(hashData)

// 	h.Reset()
// 	h.Write(hashData)
// 	//fmt.Printf("hash data 2: %v\n", h.Sum(nil))
// 	keyBuffer.Write(h.Sum(nil))
// 	//fmt.Println(keyBuffer.Bytes())

// 	return keyBuffer.Bytes()[:32]
// }

// Encrypt aes cbc 加密
func Encrypt(data, key []byte) ([]byte, error) {
	//fmt.Printf("encrypt key: %v\n", string(key))
	//return aesCBCEncrypt(data, sha1Hash2(key))
	return aesCBCEncrypt(data, key)
}

// Decrypt aes cbc 解密
func Decrypt(data, key []byte) ([]byte, error) {
	//fmt.Printf("encrypt key: %v\n", key)
	//return aesCBCDecrypt(data, sha1Hash2(key))
	return aesCBCDecrypt(data, key)
}
