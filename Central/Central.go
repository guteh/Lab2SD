package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func generarRecursos(recursos map[string]int) {
	for {
		time.Sleep(5 * time.Second)

		recursos["AT"] += 10
		recursos["MP"] += 5

		fmt.Printf("Recursos generados: AT: %d, MP: %d\n", recursos["AT"], recursos["MP"])

		if recursos["AT"] > 50 {
			recursos["AT"] = 50
		}

		if recursos["MP"] > 20 {
			recursos["MP"] = 20
		}

	}
}

type server struct{}

func (s *server) PedirRecursos(ctx context.Context, in *ResourceRequest) (*ResourceResponse, error) {
	if recursos["AT"] >= in.At() && recursos["MP"] >= in.Mp() {
		recursos["AT"] -= in.At()
		recursos["MP"] -= in.Mp()
		return &ResourceResponse{Status: 1}, nil
	} else {
		//No hay recursos suficientes
		return &ResourceResponse{Status: 0}, nil
	}
}

func main() {
	var recursos map[string]int
	recursos = make(map[string]int)

	recursos["AT"] = 0
	recursos["MP"] = 0

	// Cada 5 segundos se genereran 10 AT y 5 MP los cuales seran almacenados en bodega, el maximo que se puede almacenar es 50 AT y 20 MP
	go generarRecursos(recursos)

	time.Sleep(17 * time.Second)

	addr := "0.0.0.0:8080"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error al crear el socket:", err)
		return
	}
	s := grpc.NewServer()
	RegisterResourceServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
