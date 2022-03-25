package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ModuleName defines the module name
	ModuleName = "validitygadgetchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_validitygadgetchain"

	// Keep track of the index of snapshot-proposals
	SnapshotProposalKey      = "Snapshot-Proposal-value-"
	SnapshotProposalCountKey = "Snapshot-Proposal-count-"
)

var (
	// EthAddressByValidatorKey indexes cosmos validator account addresses
	// i.e. gravity1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm
	EthAddressByValidatorKey = "EthAddressValidatorKey"

	// ValidatorByEthAddressKey indexes ethereum addresses
	// i.e. 0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B
	ValidatorByEthAddressKey = "ValidatorByEthAddressKey"

	// ValsetRequestKey indexes valset requests by nonce
	ValsetRequestKey = "ValsetRequestKey"

	// ValsetConfirmKey indexes valset confirmations by nonce and the validator account address
	// i.e gravity1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm
	ValsetConfirmKey = "ValsetConfirmKey"

	// OracleClaimKey Claim details by nonce and validator address
	// i.e. gravityvaloper1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm
	// A claim is named more intuitively than an Attestation, it is literally
	// a validator making a claim to have seen something happen. Claims are
	// attached to attestations which can be thought of as 'the event' that
	// will eventually be executed.
	OracleClaimKey = "OracleClaimKey"

	// OracleAttestationKey attestation details by nonce and validator address
	// i.e. gravityvaloper1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm
	// An attestation can be thought of as the 'event to be executed' while
	// the Claims are an individual validator saying that they saw an event
	// occur the Attestation is 'the event' that multiple claims vote on and
	// eventually executes
	OracleAttestationKey = "OracleAttestationKey"

	// OutgoingTXPoolKey indexes the last nonce for the outgoing tx pool
	OutgoingTXPoolKey = "OutgoingTXPoolKey"

	// DenomiatorPrefix indexes token contract addresses from ETH on gravity
	DenomiatorPrefix = "DenomiatorPrefix"

	// OutgoingTXBatchKey indexes outgoing tx batches under a nonce and token address
	OutgoingTXBatchKey = "OutgoingTXBatchKey"

	// BatchConfirmKey indexes validator confirmations by token contract address
	BatchConfirmKey = "BatchConfirmKey"

	// SecondIndexNonceByClaimKey indexes latest nonce for a given claim type
	SecondIndexNonceByClaimKey = "SecondIndexNonceByClaimKey"

	// LastEventNonceByValidatorKey indexes lateset event nonce by validator
	LastEventNonceByValidatorKey = "LastEventNonceByValidatorKey"

	// LastObservedEventNonceKey indexes the latest event nonce
	LastObservedEventNonceKey = "LastObservedEventNonceKey"

	// SequenceKeyPrefix indexes different txids
	SequenceKeyPrefix = "SequenceKeyPrefix"

	// KeyLastTXPoolID indexes the lastTxPoolID
	KeyLastTXPoolID = SequenceKeyPrefix + "lastTxPoolId"

	// KeyLastOutgoingBatchID indexes the lastBatchID
	KeyLastOutgoingBatchID = SequenceKeyPrefix + "lastBatchId"

	// KeyOrchestratorAddress indexes the validator keys for an orchestrator
	KeyOrchestratorAddress = "KeyOrchestratorAddress"

	// KeyOutgoingLogicCall indexes the outgoing logic calls
	KeyOutgoingLogicCall = "KeyOutgoingLogicCall"

	// KeyOutgoingLogicConfirm indexes the outgoing logic confirms
	KeyOutgoingLogicConfirm = "KeyOutgoingLogicConfirm"

	// LastObservedEthereumBlockHeightKey indexes the latest Ethereum block height
	LastObservedEthereumBlockHeightKey = "LastObservedEthereumBlockHeightKey"

	// DenomToERC20Key prefixes the index of Cosmos originated asset denoms to ERC20s
	DenomToERC20Key = "DenomToERC20Key"

	// ERC20ToDenomKey prefixes the index of Cosmos originated assets ERC20s to denoms
	ERC20ToDenomKey = "ERC20ToDenomKey"

	// LastSlashedValsetNonce indexes the latest slashed valset nonce
	LastSlashedValsetNonce = "LastSlashedValsetNonce"

	// LatestValsetNonce indexes the latest valset nonce
	LatestValsetNonce = "LatestValsetNonce"

	// LastSlashedBatchBlock indexes the latest slashed batch block height
	LastSlashedBatchBlock = "LastSlashedBatchBlock"

	// LastSlashedLogicCallBlock indexes the latest slashed logic call block height
	LastSlashedLogicCallBlock = "LastSlashedLogicCallBlock"

	// LastUnBondingBlockHeight indexes the last validator unbonding block height
	LastUnBondingBlockHeight = "LastUnBondingBlockHeight"

	// LastObservedValsetNonceKey indexes the latest observed valset nonce
	// HERE THERE BE DRAGONS, do not use this value as an up to date validator set
	// on Ethereum it will always lag significantly and may be totally wrong at some
	// times.
	LastObservedValsetKey = "LastObservedValsetKey"

	// PastEthSignatureCheckpointKey indexes eth signature checkpoints that have existed
	PastEthSignatureCheckpointKey = "PastEthSignatureCheckpointKey"
)

// GetOrchestratorAddressKey returns the following key format
// prefix
// [0xe8][gravity1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm]
func GetOrchestratorAddressKey(orc sdk.AccAddress) string {
	if err := sdk.VerifyAddressFormat(orc); err != nil {
		panic(sdkerrors.Wrap(err, "invalid orchestrator address"))
	}
	return KeyOrchestratorAddress + string(orc.Bytes())
}

// GetEthAddressByValidatorKey returns the following key format
// prefix              cosmos-validator
// [0x0][gravityvaloper1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm]
func GetEthAddressByValidatorKey(validator sdk.ValAddress) string {
	if err := sdk.VerifyAddressFormat(validator); err != nil {
		panic(sdkerrors.Wrap(err, "invalid validator address"))
	}
	return EthAddressByValidatorKey + string(validator.Bytes())
}

// GetValidatorByEthAddressKey returns the following key format
// prefix              cosmos-validator
// [0xf9][0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B]
func GetValidatorByEthAddressKey(ethAddress EthAddress) string {
	return ValidatorByEthAddressKey + string([]byte(ethAddress.GetAddress()))
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
