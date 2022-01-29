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
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SnapshotProposal(request: MsgSnapshotProposal): Promise<MsgSnapshotProposalResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SnapshotProposal(request: MsgSnapshotProposal): Promise<MsgSnapshotProposalResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
