package encry

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

//sha1 加密用于签名
func HmacSha1(key string, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	var sign string = hex.EncodeToString(mac.Sum(nil))
	return strings.ToUpper(sign)
}

//随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
