package main

import (
	"context"
	"fmt"
)

func main() {
	ctx0 := context.Background()
	fmt.Println(ctx0)

	ctx1 := context.WithValue(ctx0, "hello", "world1")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("hello"))

	ctx2 := context.WithValue(ctx1, "hello", "world2")
	fmt.Println(ctx2)
	fmt.Println(ctx2.Value("hello"))
	fmt.Println(ctx1.Value("hello"))
}
