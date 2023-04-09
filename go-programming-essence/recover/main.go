package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover: ")
			log.Printf("%#+v\n", err)
		}
	}()

	panicFunc()
}

func panicFunc() {
	panic("panic")
}
