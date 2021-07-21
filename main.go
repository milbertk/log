package main

import (
	"log"

	"github.com/milbertk/log/bd"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
}
