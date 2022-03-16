import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
export interface ProposalExtractor {
    observed: boolean;
    votes: string[];
    height: number;
}
export declare const ProposalExtractor: {
    encode(message: ProposalExtractor, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ProposalExtractor;
    fromJSON(object: any): ProposalExtractor;
    toJSON(message: ProposalExtractor): unknown;
    fromPartial(object: DeepPartial<ProposalExtractor>): ProposalExtractor;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
