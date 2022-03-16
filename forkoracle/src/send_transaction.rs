use cosmos_sdk_proto::cosmos::bank::v1beta1::MsgSend;
use deep_space::{Coin, client::Contact, Fee, MessageArgs, Msg, PrivateKey};
use cosmos_sdk_proto::cosmos::base::abci::v1beta1::TxResponse;
use deep_space::error::CosmosGrpcError;

async fn send_transaction(contact: &Contact) -> Result<TxResponse, CosmosGrpcError> {

let private_key = PrivateKey::from_secret("mySecret".as_bytes());
let public_key = private_key.to_public_key("cosmospub").unwrap();

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
    .expect("Failed to update Eth address");
    let res = contact.wait_for_tx(result, TIMEOUT).await;

    if let Err(e) = res {
        error!("Failed trying to register delegate addresses error {:?}, correct the error and try again", e);
        exit(1);
    }
}
}
