package main

import (
	"fmt"
	"sync"

	"time"
)

type Order struct {
	Item string
}

func sequentiall_processing(orders []Order) {
	for _, order := range orders {
		fmt.Printf("%s received\n", order.Item)
		time.Sleep(time.Second * 1)
		fmt.Printf("%s prepared\n", order.Item)
		time.Sleep(time.Second * 1)
		fmt.Printf("%s delivered\n", order.Item)
	}
}

func Concurrently_Processing(orders []Order) {
	channel := make(chan Order, len(orders))
	deliverychan := make(chan Order)
	var wg sync.WaitGroup

	go func() {
		for order := range channel {
			fmt.Printf("%s received\n", order.Item)
			time.Sleep(time.Second * 1)
			fmt.Printf("%s prepared\n", order.Item)
			deliverychan <- order
		}
		close(deliverychan)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for order := range deliverychan {
			time.Sleep(time.Second * 1)
			fmt.Printf("%s delivered\n", order.Item)
		}
	}()

	for _, order := range orders {
		channel <- order
	}
	close(channel)

	wg.Wait()
}

func main() {
	orders := []Order{
		{Item: "Chicken"},
		{Item: "beef"},
		{Item: "shawerma"},
		{Item: "kosharid"},
		{Item: "Fries"},
	}

	fmt.Println("Sequential:")
	start := time.Now()
	sequentiall_processing(orders)
	fmt.Printf("Time: %v\n", time.Since(start))

	fmt.Println("___________________")

	fmt.Println("Concurrent (fast):")
	start = time.Now()
	Concurrently_Processing(orders)
	fmt.Printf("Time: %v\n", time.Since(start))
}
