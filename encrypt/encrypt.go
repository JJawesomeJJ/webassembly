package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesEncrypt struct {
}
//加密字符串
func (this *AesEncrypt) Encrypt(strMeg string,encryptKey string) (string, error) {
	key := []byte(encryptKey)
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMeg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMeg))
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

//解密字符串
func (this *AesEncrypt) Decrypt(encrypted string,encryptKey string) (strDesc string, err error) {
	src,err := base64.StdEncoding.DecodeString(string(encrypted))
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	key := []byte(encryptKey)
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
