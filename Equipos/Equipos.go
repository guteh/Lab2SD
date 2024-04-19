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

	// Create a channel to receive the results
	results := make(chan string)

	// Launch four goroutines, each representing a team
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go solicitarM(results, &wg)
		fmt.Printf("Team %d has been launched\n", i+1)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Print the results
	for result := range results {
		fmt.Println(result)
	}
}

func solicitarM(results chan<- string, wg *sync.WaitGroup) {
	// Decrement the wait group counter when the function finishes
	defer wg.Done()

	// Generate random quantities of AT and MP
	RandAT := rand.Intn(11) + 20
	RandMP := rand.Intn(6) + 10
	results <- fmt.Sprintf("Team selected %d AT and %d MP", RandAT, RandMP)

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
		response, err := c.PedirRecursos(context.Background(), &pb.ResourceRequest{ID: 1, AT: int32(RandAT), MP: int32(RandMP)})
		if err != nil {
			fmt.Println("Error al enviar el mensaje al servidor central:", err)
		}
		if response.Message == 1 {
			fmt.Println("Recursos obtenidos exitosamente")
			break
		} else {
			fmt.Println("No hay recursos suficientes, esperando 3 segundos para volver a intentar")
			time.Sleep(3 * time.Second)
		}
	}
}
