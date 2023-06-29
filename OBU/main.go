/*
	On Board Unit - This will send Location Coordinate to receiver service.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

type OBUData struct {
	OBUID int     `json:"obuId"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func main() {
	obuIDS := generateOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := generateLocation()
			data := OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			err := conn.WriteJSON(data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("OBU ID: %d \nLat: %f \nLong: %f \n\n", data.OBUID, data.Lat, data.Long)
			time.Sleep(time.Second * 1)
		}
	}
}

func init() {
	/* If this is not done then program will generate same set of random numbers. */
	rand.New(rand.NewSource(time.Now().UnixMilli()))
}

func generateCoord() float64 {
	n := float64(rand.Intn(100)) + 1
	f := rand.Float64()
	return n + f
}

func generateLocation() (float64, float64) {
	return generateCoord(), generateCoord()
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)

	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}
