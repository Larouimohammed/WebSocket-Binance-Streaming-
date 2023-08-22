package connector

import (
	"encoding/json"
	"fmt"
	"hh/binance"
	"sync"
)

func Connector(wg *sync.WaitGroup) {

	wg.Add(1)
	go func() {
		for {
			x := <-binance.Oput
			out, _ := json.Marshal(&x)
			fmt.Println(string(out))
		
			defer wg.Done()
		}
	}()
	wg.Wait()
}
