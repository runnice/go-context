package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Quando fizer a chamada Http e essa chamada demorar mais de 5s ele vai cancelar a request.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Preparando Request 
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Executando Request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer res.Body.Close()

	// Pegando o conteúdo dessa request e jogando no terminal.
	io.Copy(os.Stdout, res.Body)

}
