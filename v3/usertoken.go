package v3

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jlink-restful-golang-demo/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	kUserTokenUrl = "https://rs.xmeye.net/login/va1/%s/%s.rs"
)

type usertokenResponse struct {
	Msg  string   `json:"msg"`
	Code int      `json:"code"`
	Data userdata `json:"data"`
}
type userdata struct {
	AccessToken string `json:"accessToken"`
}

// GetUserToken
func (jUser *JUser) GetUserToken() (string, error) {
	timeMillis := utils.GetTimeMillisUtil()
	secret := jUser.JClient.Uuid + jUser.JClient.AppKey + jUser.JClient.AppSecret + timeMillis
	change := utils.Change(secret, jUser.JClient.Movecard)
	merge := utils.MergeByte(secret, change)
	secret = fmt.Sprintf("%x", md5.Sum(merge))

	tokenUrl := fmt.Sprintf(kUserTokenUrl, timeMillis, secret)

	key := KeyFilter(timeMillis, jUser.JClient.AppSecret)
	account, _ := AesEncryption(jUser.Username, key)
	pass, _ := AesEncryption(jUser.Password, key)

	value := url.Values{}
	value.Add("account", account)
	value.Add("pass", pass)

	payload := strings.NewReader(value.Encode())
	req, _ := http.NewRequest("POST", tokenUrl, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("uuid", jUser.JClient.Uuid)
	req.Header.Add("appKey", jUser.JClient.AppKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var body, _ = ioutil.ReadAll(res.Body)
	msg, _ := url.PathUnescape(string(body))
	resp := &usertokenResponse{}
	err = json.Unmarshal([]byte(msg), resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 2000 {
		fmt.Errorf(resp.Msg)
		return "", errors.New(resp.Msg)
	}
	jUser.Jutoken = resp.Data.AccessToken

	// fmt.Println(resp.Data.AccessToken)
	return resp.Data.AccessToken, nil
}

// generate key
func KeyFilter(timeMillis, appSecret string) string {
	var key string
	if len(timeMillis)/2 > 0 {
		key = timeMillis[len(timeMillis)/2:]
	}
	key += appSecret
	length := len(key)
	if length > 16 {
		key = key[:16]
	} else {
		for i := 0; i < 16-length; i++ {
			key += strconv.Itoa(i)
		}
	}

	return key
}

func AesEncryption(src, key string) (string, error) {
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

	encodeData := hex.EncodeToString(crypted)
	encodeData = strings.ToUpper(encodeData)
	return encodeData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
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
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
