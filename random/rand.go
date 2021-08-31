package random

import (
	"math/rand"
	"time"
)

func Int(min, max int64) int64 {
	if max == min {
		return min
	}
	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(max-min)
}

func String(c int) string {
	ascii := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")
	return makeString(c, ascii)
}

func StringWord(c int) string {
	ascii := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	return makeString(c, ascii)
}

func makeString(c int, ascii []byte) string {
	b := make([]byte, c)
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	_, err := rd.Read(b)
	if err != nil {
		return ""
	}
	r := make([]byte, 0, c)
	l := len(ascii)
	for _, v := range b {
		r = append(r, ascii[int(v)%l])
	}
	return string(r)
}
