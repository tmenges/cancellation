package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go l1(ctx, "L1-A")

	time.Sleep(time.Duration(5) * time.Second)

	fmt.Printf("L0 calling cancel\n")

	cancel()

	fmt.Printf("L0 sleeping\n")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Printf("L0 done\n")
}

func l1(ctx context.Context, name string) {
	fmt.Printf("%s start\n", name)
	go l2a(ctx, "L2-A")

	ctx2, cancel2 := context.WithCancel(ctx)
	go l2b(ctx2, "L2-B")

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Printf("%s calling cancel2\n", name)
	cancel2()

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Printf("%s waiting for cancel\n", name)

	<-ctx.Done()

	fmt.Printf("%s done\n", name)
}

func l2a(ctx context.Context, name string) {
	fmt.Printf("%s start\n", name)
	<-ctx.Done()

	fmt.Printf("%s done\n", name)
}

func l2b(ctx context.Context, name string) {
	fmt.Printf("%s start\n", name)
	<-ctx.Done()

	fmt.Printf("%s done\n", name)
}
