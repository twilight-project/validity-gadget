/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../validitygadgetchain/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Proposal } from "../validitygadgetchain/proposal";

export const protobufPackage =
  "twilightproject.validitygadgetchain.validitygadgetchain";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

export interface QuerySnapshotProposalRequest {}

export interface QuerySnapshotProposalResponse {
  iteration: string;
  logsnap: string;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQuerySnapshotProposalsRequest: object = {};

export const QuerySnapshotProposalsRequest = {
  encode(
    message: QuerySnapshotProposalsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QuerySnapshotProposalsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQuerySnapshotProposalsRequest,
    } as QuerySnapshotProposalsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySnapshotProposalsRequest {
    const message = {
      ...baseQuerySnapshotProposalsRequest,
    } as QuerySnapshotProposalsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QuerySnapshotProposalsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QuerySnapshotProposalsRequest>
  ): QuerySnapshotProposalsRequest {
    const message = {
      ...baseQuerySnapshotProposalsRequest,
    } as QuerySnapshotProposalsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQuerySnapshotProposalsResponse: object = {};

export const QuerySnapshotProposalsResponse = {
  encode(
    message: QuerySnapshotProposalsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Proposal) {
      Proposal.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QuerySnapshotProposalsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQuerySnapshotProposalsResponse,
    } as QuerySnapshotProposalsResponse;
    message.Proposal = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Proposal.push(Proposal.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySnapshotProposalsResponse {
    const message = {
      ...baseQuerySnapshotProposalsResponse,
    } as QuerySnapshotProposalsResponse;
    message.Proposal = [];
    if (object.Proposal !== undefined && object.Proposal !== null) {
      for (const e of object.Proposal) {
        message.Proposal.push(Proposal.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QuerySnapshotProposalsResponse): unknown {
    const obj: any = {};
    if (message.Proposal) {
      obj.Proposal = message.Proposal.map((e) =>
        e ? Proposal.toJSON(e) : undefined
      );
    } else {
      obj.Proposal = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QuerySnapshotProposalsResponse>
  ): QuerySnapshotProposalsResponse {
    const message = {
      ...baseQuerySnapshotProposalsResponse,
    } as QuerySnapshotProposalsResponse;
    message.Proposal = [];
    if (object.Proposal !== undefined && object.Proposal !== null) {
      for (const e of object.Proposal) {
        message.Proposal.push(Proposal.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQuerySnapshotProposalRequest: object = {};

export const QuerySnapshotProposalRequest = {
  encode(
    _: QuerySnapshotProposalRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QuerySnapshotProposalRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQuerySnapshotProposalRequest,
    } as QuerySnapshotProposalRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QuerySnapshotProposalRequest {
    const message = {
      ...baseQuerySnapshotProposalRequest,
    } as QuerySnapshotProposalRequest;
    return message;
  },

  toJSON(_: QuerySnapshotProposalRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QuerySnapshotProposalRequest>
  ): QuerySnapshotProposalRequest {
    const message = {
      ...baseQuerySnapshotProposalRequest,
    } as QuerySnapshotProposalRequest;
    return message;
  },
};

const baseQuerySnapshotProposalResponse: object = {
  iteration: "",
  logsnap: "",
};

export const QuerySnapshotProposalResponse = {
  encode(
    message: QuerySnapshotProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.iteration !== "") {
      writer.uint32(10).string(message.iteration);
    }
    if (message.logsnap !== "") {
      writer.uint32(18).string(message.logsnap);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QuerySnapshotProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQuerySnapshotProposalResponse,
    } as QuerySnapshotProposalResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.iteration = reader.string();
          break;
        case 2:
          message.logsnap = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySnapshotProposalResponse {
    const message = {
      ...baseQuerySnapshotProposalResponse,
    } as QuerySnapshotProposalResponse;
    if (object.iteration !== undefined && object.iteration !== null) {
      message.iteration = String(object.iteration);
    } else {
      message.iteration = "";
    }
    if (object.logsnap !== undefined && object.logsnap !== null) {
      message.logsnap = String(object.logsnap);
    } else {
      message.logsnap = "";
    }
    return message;
  },

  toJSON(message: QuerySnapshotProposalResponse): unknown {
    const obj: any = {};
    message.iteration !== undefined && (obj.iteration = message.iteration);
    message.logsnap !== undefined && (obj.logsnap = message.logsnap);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QuerySnapshotProposalResponse>
  ): QuerySnapshotProposalResponse {
    const message = {
      ...baseQuerySnapshotProposalResponse,
    } as QuerySnapshotProposalResponse;
    if (object.iteration !== undefined && object.iteration !== null) {
      message.iteration = object.iteration;
    } else {
      message.iteration = "";
    }
    if (object.logsnap !== undefined && object.logsnap !== null) {
      message.logsnap = object.logsnap;
    } else {
      message.logsnap = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of SnapshotProposals items. */
  SnapshotProposals(
    request: QuerySnapshotProposalsRequest
  ): Promise<QuerySnapshotProposalsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "twilightproject.validitygadgetchain.validitygadgetchain.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SnapshotProposals(
    request: QuerySnapshotProposalsRequest
  ): Promise<QuerySnapshotProposalsResponse> {
    const data = QuerySnapshotProposalsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "twilightproject.validitygadgetchain.validitygadgetchain.Query",
      "SnapshotProposals",
      data
    );
    return promise.then((data) =>
      QuerySnapshotProposalsResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
