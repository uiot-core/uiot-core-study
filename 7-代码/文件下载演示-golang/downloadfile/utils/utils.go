package utils

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

//生成随机字符串
func GetRandomString(lenght int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	bytesLen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

// 写文件

func WriteFile(filePath string, reader io.Reader) error {
	b:= make([]byte, 1024)

	for {
		m, errfile:= reader.Read(b)
		fl, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer fl.Close()
		n, err := fl.Write(b[:m])
		if err == nil && n < m {
			err = io.ErrShortWrite
			fmt.Println(err)
		}

		if errfile == io.EOF{
			break
		}
	}
	return nil
}
