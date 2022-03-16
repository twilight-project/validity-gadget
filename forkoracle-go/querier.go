package main

function query_snapshot_proposal(){
	queryClient := types.NewQueryClient(cosmos.Context)

	queryResp, err := queryClient.SnapshotProposals(context.Background(), &types.QuerySnapshotProposalsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// print response from querying all the proposals
	fmt.Print("\n\nAll proposals:\n\n")
	fmt.Println(queryResp)
}