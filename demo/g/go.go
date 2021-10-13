package g

import "fmt"

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic in go:", err)
				//TODO: write to log
			}
		}()

		x()
	}()
}
