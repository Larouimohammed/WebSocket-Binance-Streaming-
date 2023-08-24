package connector

import (
	"fmt"
	"hh/binance"
	"sync"
)

func Connector(wg *sync.WaitGroup) {

	wg.Add(1)
	go func() {
		for {
			x := <-binance.Oput
			fmt.Println(x)

			defer wg.Done()
		}
	}()
	wg.Wait()
}
