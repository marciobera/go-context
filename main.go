package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)

	select {
		case <-ctx.Done():
			fmt.Println("Context timeout excedded")
		case <-time.After(time.Second * 10):
			fmt.Println("Invoked cancel() after a period of time")
			cancel()
	}
}
