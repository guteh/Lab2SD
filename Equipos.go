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
	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup



	// Launch four goroutines, each representing a team
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go solicitarM(i + 1, &wg)
		fmt.Printf("Team %d has been launched\n", i+1)
	}

	// Wait for all goroutines to finish
	wg.Wait()

}

func solicitarM(id int,wg *sync.WaitGroup) {
	// Decrement the wait group counter when the function finishes

	defer wg.Done()
	time.Sleep(10 * time.Second)

	// Generate random quantities of AT and MP
	RandAT := rand.Intn(11) + 20
	RandMP := rand.Intn(6) + 10

	serverAddr := "0.0.0.0:8080"

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error al conectar al servidor central:", err)
		return
	}
	defer conn.Close()

	c := pb.NewServicioRecursosClient(conn)
	// Send the result to the results channel

	for {
		response, err := c.PedirRecursos(context.Background(), &pb.ResourceRequest{ID: int32(id), AT: int32(RandAT), MP: int32(RandMP)})
		if err != nil {
			fmt.Println("Error al enviar el mensaje al servidor central:", err)
		}
		if response.Message == 1 {
			fmt.Println("Recursos obtenidos exitosamente")
			wg.Done()
			break
		} else {
			fmt.Println("No hay recursos suficientes, esperando 3 segundos para volver a intentar")
			time.Sleep(3 * time.Second)
		}
	}
}
