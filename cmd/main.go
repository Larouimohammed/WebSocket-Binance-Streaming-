package main

import (
	//connector "hh/Connector"
	//	connector "hh/Connector"

	"hh/binance"
	"hh/server"

	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go binance.Bstream(wg)
	//go connector.Connector(wg)
	server.New().Run()
	wg.Wait()
}
