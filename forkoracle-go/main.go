package main

import (
	"fmt"
	"log"

	"github.com/ybbus/jsonrpc/v2"
)

type Chaintip struct {
	block           string `json:"block"`
	height          int64  `json:"height"`
	id              int64  `json:"id"`
	node            int64  `json:"node"`
	parent_chaintip string `json:"parent_chaintip,omitempty"`
	status          string `json:"status"`
}

func main() {

	get_active_chaintip()

	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// homePath := filepath.Join(home, ".validity-gadget-chain")

	// cosmosOptions := []cosmosclient.Option{
	// 	cosmosclient.WithHome(homePath),
	// }

	// // create an instance of cosmosclient
	// cosmos, err := cosmosclient.New(context.Background(), cosmosOptions...)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // account `alice` was initialized during `starport chain serve`
	// accountName := "alice"

	// // get account from the keyring by account name and return a bech32 address
	// address, err := cosmos.Address(accountName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // define a message to create a post
	// msg := &types.MsgSnapshotProposal{
	// 	Creator:   address.String(),
	// 	Iteration: "Hello!",
	// 	LogSnap:   "This is the first post",
	// }

	// // broadcast a transaction from account `alice` with the message to create a post
	// // store response in txResp
	// txResp, err := cosmos.BroadcastTx(accountName, msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // print response from broadcasting a transaction
	// fmt.Print("MsgSnapshotProposal:\n\n")
	// fmt.Println(txResp)

	// queryClient := types.NewQueryClient(cosmos.Context)

	// queryResp, err := queryClient.SnapshotProposals(context.Background(), &types.QuerySnapshotProposalsRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // print response from querying all the posts
	// fmt.Print("\n\nAll posts:\n\n")
	// fmt.Println(queryResp)

}

func get_active_chaintip() {
	rpcClient := jsonrpc.NewClient("http://127.0.0.1:8339")

	var chaintip *Chaintip
	err := rpcClient.CallFor(&chaintip, "getforks")

	if err != nil || chaintip == nil {
		fmt.Println("err")
		log.Fatal(err)
	}

	// if chaintip.Block == "" {
	// 	// json rpc response did not provide a field "name" in the result object
	// 	fmt.Println("no result")
	// }

	//fmt.Println(chaintip.Block)

}
