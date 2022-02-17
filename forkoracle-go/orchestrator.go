package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ybbus/jsonrpc/v2"
)

func orchestrator() {
	tick := time.Tick(50000 * time.Millisecond)
	for range tick {
		fmt.Println("Tick")
		chaintip := get_active_chaintip()
		send_transaction((chaintip))
	}

}

func get_active_chaintip() *ChainTip {

	rpcClient := jsonrpc.NewClient("http://127.0.0.1:8339")
	response, err := rpcClient.Call("getforks")

	if err != nil {
		log.Fatal(err)
	}

	var c []*ChainTip
	response.GetObject(&c)

	//fmt.Printf("%v\n", c[0].Block)
	return c[0]
}
