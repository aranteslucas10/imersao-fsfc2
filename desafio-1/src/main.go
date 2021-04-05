package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Aplicação iniciada na porta 8000")
	fmt.Println("Acesse: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
