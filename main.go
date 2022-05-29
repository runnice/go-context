package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// contexto raiz. contexto pai
	ctx := context.Background()

	// Ele vai retornar uma função de cancelamento e quando ele executar a função de cancelamento, ele vai cancelar todo mundo que tiver com aquele contexto.
	ctx, cancel := context.WithCancel(ctx)

	// Defer espera tudo ser executado e executa ele por último.
	defer cancel()
	//
	go func() {
		time.Sleep(time.Second * 10)
		cancel()
	}()

	bookHotel(ctx)

	//Quando trabalhamos com contextos, ele deve ser passado como primeiro parâmetro por boas práticas

}
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // Se o contexto for cancelado, ele retorna um erro.
		fmt.Println("Time exceed. Hotel Booking Cancelled")
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel Booked")
	}

}
