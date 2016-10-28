package main

import (
	//"fmt"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
)

func main() {
	//test()
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}

// func test() {
// 	var ss []int
// 	for i := 0; i < 10; i++ {
// 		ss = append(ss, i)
// 	}
// 	fmt.Print(ss, "\n")
// 	ss = ss[:5]
// 	fmt.Print(ss, "\n")
// 	index := 5
// 	ss = append(ss[:index], ss[index+1:]...)
// 	fmt.Print(ss, "\n")
// 	index = 0
// 	rear := append([]int{}, ss[index:]...)
// 	ss = append(ss[0:index], 11)
// 	ss = append(ss, rear...)
// 	fmt.Print(ss, "\n")
// }
