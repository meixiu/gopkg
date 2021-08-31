package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
)

// Md5 生成md5字符串
func Md5(str string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// sha1()
func Sha1(str string) string {
	hash := sha1.New()
	_, _ = hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Crc32 crc32()
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// base64_encode()
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64_decode()
func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Base64WithSha256
func Base64WithSha256(str string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	_, _ = hash.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// Base64WithSha1
func Base64WithSha1(str string, secret string) string {
	hash := hmac.New(sha1.New, []byte(secret))
	_, _ = hash.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// HmacSha256 生成hmacSha256
func HmacSha256(message string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	_, _ = hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSha1 生成hmacSha1
func HmacSha1(s string, secret string) string {
	hash := hmac.New(sha1.New, []byte(secret))
	_, _ = hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
