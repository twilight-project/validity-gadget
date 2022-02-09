package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tendermint/starport/starport/pkg/cosmosclient"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
	"github.com/ybbus/jsonrpc/v2"
)

type ChainTip struct {
	Block           string `json:"block"`
	Height          int64  `json:"height"`
	ID              int64  `json:"id"`
	Node            int64  `json:"node"`
	Parent_chaintip *int64 `json:"parent_chaintip,omitempty"`
	Status          string `json:"status"`
}

type LogSnap struct {
	BlockHash string `json:"block_hash"`
	Height    int64  `json:"height"`
}

func main() {

	chaintip := get_active_chaintip()

	logSnap := &LogSnap{
		BlockHash: chaintip.Block,
		Height:    chaintip.Height,
	}

	logSnapJson, err := json.Marshal(logSnap)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	homePath := filepath.Join(home, ".validity-gadget-chain")

	cosmosOptions := []cosmosclient.Option{
		cosmosclient.WithHome(homePath),
	}

	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(context.Background(), cosmosOptions...)
	if err != nil {
		log.Fatal(err)
	}

	// account `alice` was initialized during `starport chain serve`
	accountName := "alice"

	// get account from the keyring by account name and return a bech32 address
	address, err := cosmos.Address(accountName)
	if err != nil {
		log.Fatal(err)
	}

	// define a message to create a post
	msg := &types.MsgSnapshotProposal{
		Creator:   address.String(),
		Iteration: "0",
		LogSnap:   string(logSnapJson),
	}

	// broadcast a transaction from account `alice` with the message to create a post
	// store response in txResp
	txResp, err := cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatal(err)
	}

	// print response from broadcasting a transaction
	fmt.Print("MsgSnapshotProposal:\n\n")
	fmt.Println(txResp)

	queryClient := types.NewQueryClient(cosmos.Context)

	queryResp, err := queryClient.SnapshotProposals(context.Background(), &types.QuerySnapshotProposalsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// print response from querying all the posts
	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)

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
