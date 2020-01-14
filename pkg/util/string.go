package util

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type S string

/*
	计算文件大小
*/
func HumanSizeString(bytes int64) string {
	var thresh int64 = 1024
	if bytes < 0 {
		bytes = 0
	}
	if bytes < thresh {
		return fmt.Sprintf("%dB", bytes)
	}
	var units = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	var u = 0
	var tmp = float64(bytes)
	var standard = float64(thresh)
	for tmp >= standard && u < len(units)-1 {
		tmp /= float64(standard)
		u++
	}
	str := strconv.FormatFloat(tmp, 'f', 1, 64)
	return fmt.Sprintf("%s%s", str, units[u])
}

/*
	获取uuid
*/
func UUID() string {
	u := uuid.NewV1()
	str := u.String()
	return strings.ReplaceAll(str, "-", "")
}

/*
	获取随机字符串
*/
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijkmnpqrstuvwxyz23456789")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

/*
	获取随机数
*/
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/*
	获取4位随机数
*/
func RandomNumber() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31()%10000)
}

/*
	文件目录
*/
func FilePathJoin(elem ...string) string {
	split := "/"
	if runtime.GOOS == "windows" {
		split = `\`
	}
	return strings.Join(elem, split)
}
