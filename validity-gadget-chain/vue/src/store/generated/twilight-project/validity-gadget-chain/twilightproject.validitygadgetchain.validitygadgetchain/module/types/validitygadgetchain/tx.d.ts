import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
export interface MsgSnapshotProposal {
    creator: string;
    iteration: string;
    logSnap: string;
}
export interface MsgSnapshotProposalResponse {
    id: number;
}
export interface MsgSetDelegateAddresses {
    creator: string;
    validator: string;
    btcOracle: string;
    ethOracle: string;
    watchTower: string;
}
export interface MsgSetDelegateAddressesResponse {
}
export declare const MsgSnapshotProposal: {
    encode(message: MsgSnapshotProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSnapshotProposal;
    fromJSON(object: any): MsgSnapshotProposal;
    toJSON(message: MsgSnapshotProposal): unknown;
    fromPartial(object: DeepPartial<MsgSnapshotProposal>): MsgSnapshotProposal;
};
export declare const MsgSnapshotProposalResponse: {
    encode(message: MsgSnapshotProposalResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSnapshotProposalResponse;
    fromJSON(object: any): MsgSnapshotProposalResponse;
    toJSON(message: MsgSnapshotProposalResponse): unknown;
    fromPartial(object: DeepPartial<MsgSnapshotProposalResponse>): MsgSnapshotProposalResponse;
};
export declare const MsgSetDelegateAddresses: {
    encode(message: MsgSetDelegateAddresses, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetDelegateAddresses;
    fromJSON(object: any): MsgSetDelegateAddresses;
    toJSON(message: MsgSetDelegateAddresses): unknown;
    fromPartial(object: DeepPartial<MsgSetDelegateAddresses>): MsgSetDelegateAddresses;
};
export declare const MsgSetDelegateAddressesResponse: {
    encode(_: MsgSetDelegateAddressesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetDelegateAddressesResponse;
    fromJSON(_: any): MsgSetDelegateAddressesResponse;
    toJSON(_: MsgSetDelegateAddressesResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetDelegateAddressesResponse>): MsgSetDelegateAddressesResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    SnapshotProposal(request: MsgSnapshotProposal): Promise<MsgSnapshotProposalResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SetDelegateAddresses(request: MsgSetDelegateAddresses): Promise<MsgSetDelegateAddressesResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SnapshotProposal(request: MsgSnapshotProposal): Promise<MsgSnapshotProposalResponse>;
    SetDelegateAddresses(request: MsgSetDelegateAddresses): Promise<MsgSetDelegateAddressesResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
