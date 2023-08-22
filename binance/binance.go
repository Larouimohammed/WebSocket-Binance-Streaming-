package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Trade struct {
	Exchange  string  `json:"exchange"`
	Base      string  `json:"base"`
	Quote     string  `json:"quote"`
	Direction string  `json:"direction"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	Timestamp int64   `json:"timestamp"`
	PriceUsd  float64 `json:"priceUsd"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Oput = make(chan Trade, 1024)

func Bstream(wg *sync.WaitGroup) {

	// websocket client connection
	c, _, err := websocket.DefaultDialer.Dial("wss://ws.coincap.io/trades/binance", nil)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	//defer c.Close()

	var input = make(chan Trade, 1024) // 1️⃣
	// producer: read data stream from websocket and send to channel
	wg.Add(1)
	go func() {
		// read from the websocket
		for {
			_, message, err := c.ReadMessage() // 3️⃣
			//fmt.Println(message)
			if err != nil {
				fmt.Print(err)
				log.Fatal(err)
				break

			}
			// unmarshal the message
			var trade Trade

			json.Unmarshal(message, &trade) // 4️⃣
			// send the trade to the channel
			input <- trade // 5️⃣

			x := <-input
		    Oput <- x
			defer wg.Done()	
		}
		
		defer close(input)
	}()
wg.Wait()
}
