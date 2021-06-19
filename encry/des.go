package encry

import (
	"bytes"
	"crypto/des"
	"errors"
	"fmt"
)

//des 加密 用户加密请求串
func DesECBEncrypt(data, key []byte) ([]byte, error) {
	if len(key) > 8 {
		key = key[:8]
	}
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bs := block.BlockSize()
	data = Pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("need a multiple of the blocksize")
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

//des 填充
func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}