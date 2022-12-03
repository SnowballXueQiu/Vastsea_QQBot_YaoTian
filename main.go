package main

import (
	_ "YaoTian/Commands"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func main() {
	zero.RunAndBlock(&zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{123456},
		Driver: []zero.Driver{
			// 正向 WS
			driver.NewWebSocketClient("ws://127.0.0.1:6700", ""),
		},
	}, nil)
}
