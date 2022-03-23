use cosmos_sdk_proto::cosmos::bank::v1beta1::MsgSend;
use deep_space::address::Address;
use deep_space::{Coin, client::Contact, Fee, MessageArgs, Msg, PrivateKey};
use cosmos_sdk_proto::cosmos::base::abci::v1beta1::TxResponse;
use deep_space::error::CosmosGrpcError;
use crate::proto::MsgSetDelegateAddresses;
use std::{collections::HashMap, time::Duration};

pub const MEMO: &str = "Bridge Orchestrator";
pub const TIMEOUT: Duration = Duration::from_secs(60);

async fn send_transaction(contact: &Contact) -> Result<TxResponse, CosmosGrpcError> {

let private_key = PrivateKey::from_secret("mySecret".as_bytes());
let public_key = private_key.to_public_key("cosmospub").unwrap();

println!("{:?}", private_key);
println!("{:?}", public_key);

let address = public_key.to_address();

let coin = Coin {
    denom: "token".to_string(),
    amount: 1u32.into(),
};
let send = MsgSend {
    amount: vec![coin.clone().into()],
    from_address: address.to_string(),
    to_address: "cosmos1qftkfqaa9jgeapgq9yvsfw5tqna270namrdyxm".to_string(),
};
let msg = Msg::new("/cosmos.bank.v1beta1.MsgSend", send);

// future must be awaited in tokio runtime
contact.send_message(&vec![msg], None, &[coin], None, private_key).await

}

/// Send a transaction updating the eth address for the sending
/// Cosmos address. The sending Cosmos address should be a validator
/// this can only be called once! Key rotation code is possible but
/// not currently implemented
pub async fn set_delegate_addresses(
    contact: &Contact,
    delegate_cosmos_address: Address,
    private_key: PrivateKey,
    fee: Coin,
) -> Result<TxResponse, CosmosGrpcError> {
    let our_valoper_address = private_key
        .to_address(&contact.get_prefix())
        .unwrap()
        // This works so long as the format set by the cosmos hub is maintained
        // having a main prefix followed by a series of titles for specific keys
        // this will not work if that convention is broken. This will be resolved when
        // GRPC exposes prefix endpoints (coming to upstream cosmos sdk soon)
        .to_bech32(format!("{}", contact.get_prefix()))
        .unwrap();

    let msg_set_orch_address = MsgSetDelegateAddresses {
        validator: our_valoper_address.to_string(),
        btc_oracle: delegate_cosmos_address.to_string(),
        eth_oracle: delegate_cosmos_address.to_string(),
        watch_tower: delegate_cosmos_address.to_string(),
    };

    let msg = Msg::new(
        "/twilightproject.validitygadgetchain.validitygadgetchain.MsgSetDelegateAddresses",
        msg_set_orch_address,
    );
    contact
        .send_message(
            &[msg],
            Some(MEMO.to_string()),
            &[fee],
            Some(TIMEOUT),
            private_key,
        )
        .await
}

#[cfg(test)]
mod tests {

use super::*;
//use tokio_test;
use deep_space::{client::Contact};
use log::error;
use std::{process::exit};
pub const TIMEOUT: Duration = Duration::from_secs(60);

#[tokio::test]
async fn send_transaction_test(){
    let contact = Contact::new("http://0.0.0.0:9090", Duration::from_secs(5), "cosmos").unwrap();
    let result = send_transaction(&contact).await
    .expect("Failed to send a transcation");
    let res = contact.wait_for_tx(result, TIMEOUT).await;
    println!("{:?}", res);

    if let Err(e) = res {
        error!("Failed trying to register delegate addresses error {:?}, correct the error and try again", e);
        exit(1);
    }
}

#[tokio::test]
async fn set_delegate_addresses_test(){
    let contact = Contact::new("http://0.0.0.0:9090", Duration::from_secs(5), "cosmos").unwrap();
    let private_key = PrivateKey::from_secret("mySecret".as_bytes());
    let public_key = private_key.to_public_key("cosmospub").unwrap();

    println!("{:?}", private_key);
    println!("{:?}", public_key);

    let fee = Coin {
        denom: "token".to_string(),
        amount: 1u32.into(),
    };

    let result = set_delegate_addresses(&contact, public_key.to_address(), private_key, fee).await
    .expect("Failed to set orchestrator address");
    let res = contact.wait_for_tx(result, TIMEOUT).await;
    println!("{:?}", res);

    if let Err(e) = res {
        error!("Failed trying to register delegate addresses error {:?}, correct the error and try again", e);
        exit(1);
    }
}
}
