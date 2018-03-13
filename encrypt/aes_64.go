package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"encoding/base64"
)

type Base64Aes struct {

}

func (b Base64Aes) AesEncryptStr(origData string) string {
	dataBytes,err := b.AesEncrypt([]byte(origData),[]byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = base64.StdEncoding.EncodeToString(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

func (b Base64Aes) AesDecryptStr(encrypted string) string {
	// 验证输入参数
	// 必须为aes.BlockSize的倍数 --- 16
	encBytes,err := base64.StdEncoding.DecodeString(encrypted)
	if len(encBytes)%aes.BlockSize != 0 || err != nil{
		return "invalid"
	}
	dataBytes,err := b.AesDecrypt(encBytes,[]byte(DefaultAESKey))
	retStr := ""
	if dataBytes != nil {
		retStr = string(dataBytes)
	}
	if err != nil {
		retStr = err.Error()
	}
	return retStr
}

func (b Base64Aes) AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (b Base64Aes) AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}