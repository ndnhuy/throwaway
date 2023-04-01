package main

import (
	"log"

	"github.com/ndnhuy/proglog/internal/server"
)

func main() {
	src := server.NewHttpServer(":9000")
	log.Fatal(src.ListenAndServe())
}