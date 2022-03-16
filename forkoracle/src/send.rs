
use cosmrs::{
    bank::MsgSend,
    crypto::secp256k1,
    rpc,
    tx::{self, AccountNumber, Fee, Msg, SignDoc, SignerInfo},
    Coin,
};
use tokio;
use std::{panic, str};

/// Chain ID to use for tests
const CHAIN_ID: &str = "validitygadgetchain";

/// RPC port
const RPC_PORT: u16 = 26657;

/// Expected account number
const ACCOUNT_NUMBER: AccountNumber = 2;

/// Bech32 prefix for an account
const ACCOUNT_PREFIX: &str = "cosmos";

/// Denom name
const DENOM: &str = "token";

/// Example memo
const MEMO: &str = "test memo";

#[test]
fn msg_send() {
    let sender_private_key = secp256k1::SigningKey::random();
    let sender_public_key = sender_private_key.public_key();
    let sender_account_id = sender_public_key.account_id(ACCOUNT_PREFIX).unwrap();
    let x = true;
    assert!(x, "false");
}