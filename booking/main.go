package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go checkRoomAvailabilityAPI(cancel)
	bookingProcessAPI(ctx)
}

func checkRoomAvailabilityAPI(cancel context.CancelFunc) {
	time.Sleep(time.Second * 3)
	cancel()
}

func bookingProcessAPI(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("timeout exceeded for room booking")
	case <-time.After(time.Second * 5):
		fmt.Println("room booked successfully")
	}
}
