package main

import (
	"fmt"
	"time"

	heavylock "github.com/viney-shih/go-lock"
)

func main() {
	myCAXMutex := heavylock.NewCASMutex()
	go func() {
		ticker := time.NewTicker(10 * time.Microsecond)
		defer ticker.Stop() 
		for {
			<-ticker.C
			myCAXMutex.Lock()
			fmt.Printf("111\n")
			myCAXMutex.Unlock()
		}
	}()

	go func() {
		ticker := time.NewTicker(100 * time.Microsecond)
		defer ticker.Stop() 
		for {
			<-ticker.C
			if !myCAXMutex.TryLockWithTimeout(50 * time.Microsecond) {
				fmt.Printf("Try Lock failed!\n")
				continue
			}
			fmt.Printf("222\n")
			myCAXMutex.Unlock()
		}
	}()
	time.Sleep(1 * time.Second)
}
