/// MsgSetDelegateAddress
/// this message allows validators to delegate their voting responsibilities
/// to a given key. This key is then used as an optional authentication method
/// for sigining oracle claims
/// VALIDATOR
/// The validator field is a cosmosvaloper1... string (i.e. sdk.ValAddress)
/// that references a validator in the active set
/// ETH_ORACLE
/// The eth_oracle field is a cosmos1... string  (i.e. sdk.AccAddress) that
/// references the key that is being delegated to
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetDelegateAddresses {
    #[prost(string, tag="1")]
    pub validator_address: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub orchestrator_address: ::prost::alloc::string::String,
    #[prost(string, tag="3")]
    pub eth_address: ::prost::alloc::string::String,
    #[prost(string, tag="4")]
    pub btc_public_key: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetDelegateAddressesResponse {
}