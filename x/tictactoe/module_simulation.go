package tictactoe

import (
	"math/rand"

	"github.com/cjcobb23/tictactoe/testutil/sample"
	tictactoesimulation "github.com/cjcobb23/tictactoe/x/tictactoe/simulation"
	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = tictactoesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgInvite = "op_weight_msg_invite"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInvite int = 100

	opWeightMsgAccept = "op_weight_msg_accept"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccept int = 100

	opWeightMsgMove = "op_weight_msg_move"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMove int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tictactoeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tictactoeGenesis)
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

	var weightMsgInvite int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInvite, &weightMsgInvite, nil,
		func(_ *rand.Rand) {
			weightMsgInvite = defaultWeightMsgInvite
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInvite,
		tictactoesimulation.SimulateMsgInvite(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccept int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccept, &weightMsgAccept, nil,
		func(_ *rand.Rand) {
			weightMsgAccept = defaultWeightMsgAccept
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccept,
		tictactoesimulation.SimulateMsgAccept(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMove int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMove, &weightMsgMove, nil,
		func(_ *rand.Rand) {
			weightMsgMove = defaultWeightMsgMove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMove,
		tictactoesimulation.SimulateMsgMove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
