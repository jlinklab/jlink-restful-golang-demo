package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"strings"
)

// 反转字符串

func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func AesDecrypt(encodeData, key string) ([]byte, error) {
	crypted, err := hex.DecodeString(encodeData)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("err is:", err)
		return nil, err
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)

	return origData, nil
}

func AesEncrypt(src, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if src == "" {
		return "", fmt.Errorf("plain content empty")
	}
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	// 普通base64编码加密 区别于urlsafe base64
	// fmt.Println("base64 result:", base64.StdEncoding.EncodeToString(crypted))

	encodeData := hex.EncodeToString(crypted)
	encodeData = strings.ToUpper(encodeData)

	return encodeData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	if origData == nil || len(origData) == 0 {
		return nil
	}

	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unPadding := int(origData[length-1])
	if unPadding > length {
		return origData
	}

	paddingLen := length - unPadding
	if paddingLen > length {
		paddingLen = length
	}

	return origData[:paddingLen]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		// xmlog.Error("crypto/cipher: input not full blocks")
		return
	}
	if len(dst) < len(src) {
		// xmlog.Error("crypto/cipher: output smaller than input")
		return
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		// xmlog.Error("crypto/cipher: input not full blocks")

		return
	}
	if len(dst) < len(src) {
		// xmlog.Error("crypto/cipher: output smaller than input")
		return
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func KeyFilter(timeMillis string, appSecret string) string {
	var key strings.Builder

	timeLength := len(timeMillis)
	if timeLength/2 > 0 {
		key.WriteString(timeMillis[timeLength/2:])
	}
	key.WriteString(appSecret)

	return keyFilterFunc(key.String())
}

func keyFilterFunc(key string) string {
	length := len(key)
	if length > 16 {
		key = key[:16]
	} else {
		for i := 0; i < 16-length; i++ {
			key += string(i)
		}
	}
	return key
}
