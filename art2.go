package main

import (
	"fmt"
	"time"
)

func face() {
	var count = 1

	for {
		fmt.Println("∩――――∩")
		fmt.Println("|| ∧ ﾍ　 ||")
		fmt.Println("||(* ´ ｰ`) は？ねるわ")
		fmt.Println("|ﾉ^⌒⌒づ`￣＼")
		fmt.Println("(　ノ　　⌒ ヽ ＼")
		fmt.Println("＼　　||￣￣￣￣￣||")
		fmt.Println("　 ＼,ﾉ||―――――||")
		time.Sleep(1 * time.Second)
		count++
		if count > 1 {
			break
		}
	}
}

func main() {
	face()
}
