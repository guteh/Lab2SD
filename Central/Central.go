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
		time.Sleep(5 * time.Second) //Espera 5 segundos

		recursos["AT"] += 10  //Genera recursos AT y MP
		recursos["MP"] += 5

		
		if recursos["AT"] + 10 > 50 {  //Si los recursos generados son mayor a 50, se asigna 50
			recursos["AT"] = 50
		}
		
		if recursos["MP"] > 20 { //Si los recursos generados son mayor a 20, se asigna 20
			recursos["MP"] = 20
		}
		
		//fmt.Printf("Recursos generados: AT: %d, MP: %d\n", recursos["AT"], recursos["MP"])
	}
}

type server struct {  //Crea el servidor rcp
    pb.UnimplementedServicioRecursosServer
    recursos map[string]int
}

//Implementa la funcion SolicitarM de la interfaz ServicioRecursos de RCP

func (s *server) SolicitarM(ctx context.Context, req *pb.ResourceRequest) (*pb.ResourceResponse, error) {

	
    if s.recursos["AT"] >= int(req.GetAT()) && s.recursos["MP"] >= int(req.GetMP()) {  //Si los recursos son suficientes

        s.recursos["AT"] -= int(req.GetAT())  //Se restan los recursos asignados
        s.recursos["MP"] -= int(req.GetMP())

		fmt.Printf("Recepcion de solicitud desde equipo %d, %d AT y %d MP -- APROBADA --\nAT EN SISTEMA: %d ; MP EN SISTEMA: %d \n", req.GetID(), req.GetAT(), req.GetMP(), s.recursos["AT"], s.recursos["MP"])
		fmt.Println("\n")
        return &pb.ResourceResponse{Message: 1}, nil //Se retorna en funcion SolicitarM un mensaje de aprobacion
    } else {
        // No hay suficientes recursos
		fmt.Printf("Recepcion de solicitud desde equipo %d, %d AT y %d MP -- DENEGADA --\nAT EN SISTEMA: %d ; MP EN SISTEMA: %d \n", req.GetID(), req.GetAT(), req.GetMP(), s.recursos["AT"], s.recursos["MP"])
		fmt.Println("\n")
        return &pb.ResourceResponse{Message: 0}, nil  //Se retorna en funcion SolicitarM un mensaje de denegacion
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
        log.Fatalf("Fallo al escuchar %v", err)
    }

	if err := grpcServer.Serve(lis); err != nil {  //Se inicia el servidor
        log.Fatalf("Fallo al crear servidor: %s", err)
    }



}
