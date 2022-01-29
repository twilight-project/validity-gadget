import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
export interface Proposal {
    creator: string;
    id: number;
    iteration: string;
    logSnap: string;
}
export declare const Proposal: {
    encode(message: Proposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Proposal;
    fromJSON(object: any): Proposal;
    toJSON(message: Proposal): unknown;
    fromPartial(object: DeepPartial<Proposal>): Proposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
