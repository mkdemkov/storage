package main

import (
	"fmt"
	"github.com/mkdemkov/storage/internal/storage"
	"log"
)

func main() {

	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded", file)
}
