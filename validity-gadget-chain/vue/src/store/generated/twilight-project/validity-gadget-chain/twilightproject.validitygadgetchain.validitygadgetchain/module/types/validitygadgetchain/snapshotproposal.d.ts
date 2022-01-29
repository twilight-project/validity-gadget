import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
export interface SnapshotProposal {
    creator: string;
    id: number;
    iteration: string;
    logSnap: string;
}
export declare const SnapshotProposal: {
    encode(message: SnapshotProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SnapshotProposal;
    fromJSON(object: any): SnapshotProposal;
    toJSON(message: SnapshotProposal): unknown;
    fromPartial(object: DeepPartial<SnapshotProposal>): SnapshotProposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
