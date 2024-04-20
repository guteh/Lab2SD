package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	pb "Lab2SD/Proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var wg sync.WaitGroup



	// Se crean los cuatro grupos
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go InicioEquipo(i + 1, &wg)  //Empieza ejecucion de equipo
		time.Sleep(2 * time.Second)  //Espero 2 segundos para no colapsar
		fmt.Printf("Equipo %d ha empezado su mision!\n", i+1)
	}

	// Espera que todas las ejecuciones terminen para finalizar la ejecucion del codigo.
	wg.Wait()

}

func InicioEquipo(id int,wg *sync.WaitGroup) { //Toma como parametros el id del equipo y el grupo de espera

	defer wg.Done()
	time.Sleep(10 * time.Second)  //Espera 10 segundos por enunciado

	// Genera cantidades random de recursos
	RandAT := rand.Intn(11) + 20
	RandMP := rand.Intn(6) + 10

	serverAddr := "0.0.0.0:8080"

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))  //Se conecta al servidor central
	if err != nil {
		fmt.Println("Error al conectar al servidor central:", err)
		return
	}
	defer conn.Close()

	c := pb.NewServicioRecursosClient(conn) 

	for {
		response, err := c.SolicitarM(context.Background(), &pb.ResourceRequest{ID: int32(id), AT: int32(RandAT), MP: int32(RandMP)})  //Envia peticion de recursos a servidor central
		if err != nil {
			fmt.Println("Error al enviar el mensaje al servidor central:", err)
		}
		if response.Message == 1 {
			fmt.Printf("EQUIPO %d: Solicitando %d AT y %d MP -- APROBADA -- ;\n Conquista existosa! Cerrando comunicacion.\n",id, RandAT, RandMP)
			break  //Si la respuesta es 1, se cierra la comunicacion
		} else {
			fmt.Printf("EQUIPO %d: Solicitando %d AT y %d MP -- DENEGADA -- ;\n Reintentando en 3 segundos...\n",id, RandAT, RandMP)
			time.Sleep(3 * time.Second)  //Si la respuesta es 0, se espera 3 segundos y se vuelve a enviar la peticion
		}
	}
}
