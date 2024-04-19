package main

import (
	pb "Lab2SD/Proto"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func generarRecursos(recursos map[string]int) {  //Genera recursos cada 5 segundos
	for {
		time.Sleep(5 * time.Second)

		recursos["AT"] += 10
		recursos["MP"] += 5

		
		if recursos["AT"] + 10 > 50 {
			recursos["AT"] = 50
		}
		
		if recursos["MP"] > 20 {
			recursos["MP"] = 20
		}
		
		fmt.Printf("Recursos generados: AT: %d, MP: %d\n", recursos["AT"], recursos["MP"])
	}
}

type server struct {  //Crea el servidor
    pb.UnimplementedServicioRecursosServer
    recursos map[string]int
}

//Implementa servidor de recursos, no esta probado el funcionamiento

func (s *server) PedirRecursos(ctx context.Context, req *pb.ResourceRequest) (*pb.ResourceResponse, error) {
    // Implement your method logic here
    // For example, check if there are enough resources and return a response
	fmt.Printf("Recursos solicitados: AT: %d, MP: %d por grupo de ID: %d\n", req.GetAT(), req.GetMP(), req.GetID())
    if s.recursos["AT"] >= int(req.GetAT()) && s.recursos["MP"] >= int(req.GetMP()) {
        s.recursos["AT"] -= int(req.GetAT())
        s.recursos["MP"] -= int(req.GetMP())
        return &pb.ResourceResponse{Message: 1}, nil
    } else {
        // Not enough resources
        return &pb.ResourceResponse{Message: 0}, nil
    }
}




func main() {
	recursos := make(map[string]int) //Se inicializan recursos

	recursos["AT"] = 0
	recursos["MP"] = 0

	// Cada 5 segundos se genereran 10 AT y 5 MP los cuales seran almacenados en bodega, el maximo que se puede almacenar es 50 AT y 20 MP
	go generarRecursos(recursos)

	grpcServer := grpc.NewServer() //Se crea el servidor

	s := &server{
        recursos: recursos, //Se le asignan los recursos al servidor
    }

	pb.RegisterServicioRecursosServer(grpcServer, s) //Se registra el servidor

	addr := "0.0.0.0:8080"  //Se asigna la direccion del servidor
	lis, err := net.Listen("tcp", addr) //Se crea el listener
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

	if err := grpcServer.Serve(lis); err != nil {  //Se inicia el servidor
        log.Fatalf("failed to serve: %s", err)
    }



}
