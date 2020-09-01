package main

import(
	"fmt"
	"time"
	"os"
	"strconv"
)

func main() {
	timeOut := 0
	if len(os.Args) > 1 {
		arg1 := os.Args[1] 
		timeOut, _ = strconv.Atoi(arg1)
	}
	i := 0
	infiniteProcess := func() {
		for {
			fmt.Printf("-------------line%v---------------", i)
			i++
			time.Sleep(time.Second)
		}
	}
	go infiniteProcess()
	 <-time.After(time.Duration(timeOut * 1e9))
	
	
	
}