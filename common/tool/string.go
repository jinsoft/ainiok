package tool

import (
	"math/rand"
	"time"
)

func RandomString(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := []byte(str)
	var res []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		res = append(res, b[r.Intn(len(b))])
	}
	return string(res)
}
