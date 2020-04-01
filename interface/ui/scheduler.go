package ui

import (
	"fmt"

	"github.com/spf13/viper"
)

func Schedule(openID, msg string) string {
	return ""
}

//Realese 释放取消关注的用户的资源
func Realese(userID string) {
}

//Prologue 开场白
func Prologue(userID string) string {
	return fmt.Sprintf(viper.GetString("prologue"))
}

//NotSuport 不支持
func NotSuport() string {
	return _NOT_SUPORT
}
