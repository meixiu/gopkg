package sign

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func RsaEncrypt(plainText string, key string) (string, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	data, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func RsaDecrypt(cipherText string, key string) (string, error) {
	byteText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("key error")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, pri, byteText)
	if err != nil {
		return "", nil
	}
	return string(data), nil
}
