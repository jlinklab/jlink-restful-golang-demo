package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

func GetEncryptStr(uuid, appKey, secret, tm string, moved int) string {
	secret = uuid + appKey + secret + tm
	change := Change(secret, moved)
	merge := MergeByte(secret, change)
	secret = fmt.Sprintf("%x", md5.Sum(merge))
	return secret
}

var counter int64 //
func GetTimeMillisUtil() string {
	var value string
	counter++
	if counter < 10 {
		value = "000000"
	} else if counter < 100 {
		value = "00000"
	} else if counter < 1000 {
		value = "0000"
	} else if counter < 10000 {
		value = "000"
	} else if counter < 100000 {
		value = "00"
	} else if counter < 1000000 {
		value = "0"
	} else if counter >= 10000000 {
		counter = 1
		value = "000000"
	}

	return fmt.Sprintf("%s%d%d", value, counter, time.Now().UnixNano()/1000000)
}

func Change(data string, mc int) string {
	src := []byte(data)
	length := len(src)
	for i := 0; i < length; i++ {
		temp := src[length-(i+1)]
		if (i % mc) > ((length - i) % mc) {
			temp = src[i]
		}

		src[i] = src[length-(i+1)]
		src[length-(i+1)] = temp
	}

	return string(src)
}

func MergeByte(encrypt, change string) []byte {
	length := len(encrypt)
	length2 := length * 2
	temp := make([]byte, length2)
	for i := 0; i < length; i++ {
		temp[i] = encrypt[i]
		temp[length2-1-i] = change[i]
	}

	return temp
}
