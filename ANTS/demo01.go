package main

import (
	"fmt"
	// "time"

	"github.com/panjf2000/ants/v2"
)

func main()  {
	newantsPool,err := ants.NewPool(10)
	if err != nil {

		fmt.Println("new pool err.")

		return
	}
	defer newantsPool.Release()
	// myTicker := time.NewTicker(1*time.Second)
	var i int
	for {
		// <- myTicker.C
		myFun := func ()  {
			fmt.Println(i)
			i++
		}
		newantsPool.Submit(myFun)

	}


}