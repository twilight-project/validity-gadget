package validitygadgetchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/twilight-project/validity-gadget-chain/testutil/sample"
	validitygadgetchainsimulation "github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/simulation"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = validitygadgetchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSnapshotProposal = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSnapshotProposal int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	validitygadgetchainGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&validitygadgetchainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSnapshotProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSnapshotProposal, &weightMsgSnapshotProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSnapshotProposal = defaultWeightMsgSnapshotProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSnapshotProposal,
		validitygadgetchainsimulation.SimulateMsgSnapshotProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
