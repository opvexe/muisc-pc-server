package plus

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrap400Response(t *testing.T) {
	var er = errors.New("请求错误-------")
	err := Wrap400Response(er, "我是不是请求错误了")
	fmt.Println(err)
}
