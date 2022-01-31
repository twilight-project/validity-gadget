import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../validitygadgetchain/params";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Proposal } from "../validitygadgetchain/proposal";
export declare const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QuerySnapshotProposalsRequest {
    /** Adding pagination to request */
    pagination: PageRequest | undefined;
}
export interface QuerySnapshotProposalsResponse {
    /** Returning a list of proposals */
    Proposal: Proposal[];
    /** Adding pagination to response */
    pagination: PageResponse | undefined;
}
export interface QuerySnapshotProposalRequest {
}
export interface QuerySnapshotProposalResponse {
    iteration: string;
    logsnap: string;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QuerySnapshotProposalsRequest: {
    encode(message: QuerySnapshotProposalsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QuerySnapshotProposalsRequest;
    fromJSON(object: any): QuerySnapshotProposalsRequest;
    toJSON(message: QuerySnapshotProposalsRequest): unknown;
    fromPartial(object: DeepPartial<QuerySnapshotProposalsRequest>): QuerySnapshotProposalsRequest;
};
export declare const QuerySnapshotProposalsResponse: {
    encode(message: QuerySnapshotProposalsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QuerySnapshotProposalsResponse;
    fromJSON(object: any): QuerySnapshotProposalsResponse;
    toJSON(message: QuerySnapshotProposalsResponse): unknown;
    fromPartial(object: DeepPartial<QuerySnapshotProposalsResponse>): QuerySnapshotProposalsResponse;
};
export declare const QuerySnapshotProposalRequest: {
    encode(_: QuerySnapshotProposalRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QuerySnapshotProposalRequest;
    fromJSON(_: any): QuerySnapshotProposalRequest;
    toJSON(_: QuerySnapshotProposalRequest): unknown;
    fromPartial(_: DeepPartial<QuerySnapshotProposalRequest>): QuerySnapshotProposalRequest;
};
export declare const QuerySnapshotProposalResponse: {
    encode(message: QuerySnapshotProposalResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QuerySnapshotProposalResponse;
    fromJSON(object: any): QuerySnapshotProposalResponse;
    toJSON(message: QuerySnapshotProposalResponse): unknown;
    fromPartial(object: DeepPartial<QuerySnapshotProposalResponse>): QuerySnapshotProposalResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of SnapshotProposals items. */
    SnapshotProposals(request: QuerySnapshotProposalsRequest): Promise<QuerySnapshotProposalsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    SnapshotProposals(request: QuerySnapshotProposalsRequest): Promise<QuerySnapshotProposalsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
