package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou minha request")
	defer log.Println("Finalizou minha request")
	// Se o contexto for cancelado, ele retorna um erro.
	select {
	case <-time.After(time.Second * 5):
		fmt.Fprint(w, "Requisição processada com sucesso")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Requisição processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada")
		http.Error(w, "Requisição cancelada", http.StatusRequestTimeout)
	}
}
