package main

import (
	"fmt"
	"github.com/mkdemkov/storage/internal/storage"
)

func main() {

	st := storage.NewStorage()

	fmt.Println("Hello from make", st)
}
