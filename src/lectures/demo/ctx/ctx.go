package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v Operation Completed", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v Abortado", msg)
				return
			}
		}
	}()
	return out
}

// Usar contet quando trabalhar com multiplas goroutines, pois é possível cancelar todas usando uma única chamada do cancelamento do contexto.
func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "Webserver", 200)
	microservice := sampleOperation(ctx, "Microservice", 500)
	database := sampleOperation(ctx, "Database", 900)
Mainloop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microservice:
			fmt.Println(val)
			fmt.Println("Cancel Context")
			cancelCtx()
			break Mainloop
		case val := <-database:
			fmt.Println(val)
		}
	}
	fmt.Println(<-database)
}
