package main

import (
	"fmt"
	"strconv"
	"test222/app"
	"time"
)

func main() {
	i := 0
	for {
		if i < 10 {
			val := i
			if val%2 == 0 {
				val++
			}
			app.HandleRequest(strconv.Itoa(val))
		}
		time.Sleep(time.Second)
		cnt := app.GetUniqueCount()
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "CNT: ", cnt)
		if cnt == 0 {
			return
		}
		i++
	}
}
