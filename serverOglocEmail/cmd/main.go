package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido al servicio de email de Ogloc")
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
