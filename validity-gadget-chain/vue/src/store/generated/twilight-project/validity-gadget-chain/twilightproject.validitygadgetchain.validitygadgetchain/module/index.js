// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSetDelegateAddresses } from "./types/validitygadgetchain/tx";
import { MsgSnapshotProposal } from "./types/validitygadgetchain/tx";
const types = [
    ["/twilightproject.validitygadgetchain.validitygadgetchain.MsgSetDelegateAddresses", MsgSetDelegateAddresses],
    ["/twilightproject.validitygadgetchain.validitygadgetchain.MsgSnapshotProposal", MsgSnapshotProposal],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSetDelegateAddresses: (data) => ({ typeUrl: "/twilightproject.validitygadgetchain.validitygadgetchain.MsgSetDelegateAddresses", value: MsgSetDelegateAddresses.fromPartial(data) }),
        msgSnapshotProposal: (data) => ({ typeUrl: "/twilightproject.validitygadgetchain.validitygadgetchain.MsgSnapshotProposal", value: MsgSnapshotProposal.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
