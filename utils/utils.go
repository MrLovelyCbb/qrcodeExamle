package utils

import (
	"crypto/rand"
	"encoding/base64"
)

const (
	INT_MAX = int64(^uint32(0) >> 1)
	INT_MIN = ^INT_MAX
)

//sessionId函数用来生成一个session ID，即session的唯一标识符
func GetSessionId(n int) string {

	//id := uuid.NewV4()
	//uuidHash := int(crc32.ChecksumIEEE([]byte(id.String())))
	//fmt.Println("uuidHash=> ", uuidHash)
	//result, _ := rand.Int(rand.Reader, big.NewInt(utils.INT_MAX))
	//fmt.Println("result => ", result)

	b := make([]byte, n)
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	rand.Read(b)
	//fmt.Println(b[:n])
	return base64.URLEncoding.EncodeToString(b) //将生成的随机数b编码后返回字符串,该值则作为session ID
}
