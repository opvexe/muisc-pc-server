package util

import (
	"errors"
	"regexp"
)

/*
	判断用户名的合法性
*/
func ValidateUserName(n string) error {
	if len(n) == 0 {
		return errors.New("请输入您的用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", n)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", n)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

/*
	校验邮箱的合法性
*/
func ValidateEmail(e string) error {
	if len(e) == 0 {
		return errors.New("请输入您的邮箱")
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, e)
	if !matched {
		return errors.New("邮箱格式不符合规范")
	}
	return nil
}

/*
	校验密码的合法性
*/
func ValidatePassWord(s string) error {
	if len(s) == 0 {
		return errors.New("请输入您的密码")
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", s)
	if !matched {
		return errors.New("密码格式不符合规范")
	}
	return nil
}
