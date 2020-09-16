package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "worker1")
	go worker(ctx, "worker2")
	go worker(ctx, "worker3")
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	fmt.Println("-----------")
	ctxTimeout()
	time.Sleep(2 * time.Second)

	fmt.Println("-----------")
	ctxCancel()
	time.Sleep(2 * time.Second)

	fmt.Println("-----------")
	ctxValue()
	time.Sleep(2 * time.Second)

	fmt.Println("-----------")
	ctxExtend()
	time.Sleep(2 * time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "is over")
			return
		default:
			fmt.Println(name, "is running ")
			time.Sleep(1 * time.Second)
		}
	}
}

func ctxCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("cancel: ", ctx.Err())
		case <-time.After(1 * time.Second):
			fmt.Println("Time out")
		}
	}(ctx)
	time.Sleep(300 * time.Millisecond)
	cancel()
}

func ctxTimeout() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("time: ", ctx.Err())
		case <-time.After(2 * time.Second):
			fmt.Println("Time out")
		}
	}(ctx)
	time.Sleep(2 * time.Second)
}

func ctxValue() {
	ctx := context.WithValue(context.Background(), "user", "liubo")
	go func(ctx context.Context) {
		v, ok := ctx.Value("user").(string)
		if ok {
			fmt.Println(v)
		}
	}(ctx)
	time.Sleep(time.Second)
}

func ctxExtend() {
	ctx, cancel := context.WithCancel(context.Background())
	go parent(ctx, "parent")
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
func parent(ctx context.Context, name string) {
	go child(ctx, "child")
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " stop")
			return
		default:
			fmt.Println(name, " running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func child(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " stop")
			return
		default:
			fmt.Println(name, " running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
