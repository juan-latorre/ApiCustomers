package main

import (
	//crypto "ApiCustomers/crypto"
	//customer_service "ApiCustomers/services/customer.service"
	"ApiCustomers/HTTP_Server/server"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	//"net/http"
	//"time"
)

func main() {

	fmt.Println("Listado Clientes")

	ctx := context.Background()

	serverDoneChannel := make(chan os.Signal, 1)
	signal.Notify(serverDoneChannel,os.Interrupt,syscall.SIGTERM)

    srv := server.New(":8080")
	
	go func () {

		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}() 
	
    log.Println("Servidor iniciado...")
	
	<-serverDoneChannel

	srv.Shutdown(ctx)
	log.Println("Servidor Detenido...")

}
