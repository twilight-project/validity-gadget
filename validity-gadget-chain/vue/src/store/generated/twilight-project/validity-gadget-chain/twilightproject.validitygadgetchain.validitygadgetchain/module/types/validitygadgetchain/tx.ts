/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage =
  "twilightproject.validitygadgetchain.validitygadgetchain";

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
  validatorAddress: string;
  orchestratorAddress: string;
  ethAddress: string;
  btcPublicKey: string;
}

export interface MsgSetDelegateAddressesResponse {}

const baseMsgSnapshotProposal: object = {
  creator: "",
  iteration: "",
  logSnap: "",
};

export const MsgSnapshotProposal = {
  encode(
    message: MsgSnapshotProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.iteration !== "") {
      writer.uint32(18).string(message.iteration);
    }
    if (message.logSnap !== "") {
      writer.uint32(26).string(message.logSnap);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSnapshotProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSnapshotProposal } as MsgSnapshotProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.iteration = reader.string();
          break;
        case 3:
          message.logSnap = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSnapshotProposal {
    const message = { ...baseMsgSnapshotProposal } as MsgSnapshotProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.iteration !== undefined && object.iteration !== null) {
      message.iteration = String(object.iteration);
    } else {
      message.iteration = "";
    }
    if (object.logSnap !== undefined && object.logSnap !== null) {
      message.logSnap = String(object.logSnap);
    } else {
      message.logSnap = "";
    }
    return message;
  },

  toJSON(message: MsgSnapshotProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.iteration !== undefined && (obj.iteration = message.iteration);
    message.logSnap !== undefined && (obj.logSnap = message.logSnap);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSnapshotProposal>): MsgSnapshotProposal {
    const message = { ...baseMsgSnapshotProposal } as MsgSnapshotProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.iteration !== undefined && object.iteration !== null) {
      message.iteration = object.iteration;
    } else {
      message.iteration = "";
    }
    if (object.logSnap !== undefined && object.logSnap !== null) {
      message.logSnap = object.logSnap;
    } else {
      message.logSnap = "";
    }
    return message;
  },
};

const baseMsgSnapshotProposalResponse: object = { id: 0 };

export const MsgSnapshotProposalResponse = {
  encode(
    message: MsgSnapshotProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSnapshotProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSnapshotProposalResponse,
    } as MsgSnapshotProposalResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSnapshotProposalResponse {
    const message = {
      ...baseMsgSnapshotProposalResponse,
    } as MsgSnapshotProposalResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgSnapshotProposalResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSnapshotProposalResponse>
  ): MsgSnapshotProposalResponse {
    const message = {
      ...baseMsgSnapshotProposalResponse,
    } as MsgSnapshotProposalResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgSetDelegateAddresses: object = {
  creator: "",
  validatorAddress: "",
  orchestratorAddress: "",
  ethAddress: "",
  btcPublicKey: "",
};

export const MsgSetDelegateAddresses = {
  encode(
    message: MsgSetDelegateAddresses,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.validatorAddress !== "") {
      writer.uint32(18).string(message.validatorAddress);
    }
    if (message.orchestratorAddress !== "") {
      writer.uint32(26).string(message.orchestratorAddress);
    }
    if (message.ethAddress !== "") {
      writer.uint32(34).string(message.ethAddress);
    }
    if (message.btcPublicKey !== "") {
      writer.uint32(42).string(message.btcPublicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetDelegateAddresses {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetDelegateAddresses,
    } as MsgSetDelegateAddresses;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.validatorAddress = reader.string();
          break;
        case 3:
          message.orchestratorAddress = reader.string();
          break;
        case 4:
          message.ethAddress = reader.string();
          break;
        case 5:
          message.btcPublicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetDelegateAddresses {
    const message = {
      ...baseMsgSetDelegateAddresses,
    } as MsgSetDelegateAddresses;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.validatorAddress !== undefined &&
      object.validatorAddress !== null
    ) {
      message.validatorAddress = String(object.validatorAddress);
    } else {
      message.validatorAddress = "";
    }
    if (
      object.orchestratorAddress !== undefined &&
      object.orchestratorAddress !== null
    ) {
      message.orchestratorAddress = String(object.orchestratorAddress);
    } else {
      message.orchestratorAddress = "";
    }
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = String(object.ethAddress);
    } else {
      message.ethAddress = "";
    }
    if (object.btcPublicKey !== undefined && object.btcPublicKey !== null) {
      message.btcPublicKey = String(object.btcPublicKey);
    } else {
      message.btcPublicKey = "";
    }
    return message;
  },

  toJSON(message: MsgSetDelegateAddresses): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.validatorAddress !== undefined &&
      (obj.validatorAddress = message.validatorAddress);
    message.orchestratorAddress !== undefined &&
      (obj.orchestratorAddress = message.orchestratorAddress);
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    message.btcPublicKey !== undefined &&
      (obj.btcPublicKey = message.btcPublicKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetDelegateAddresses>
  ): MsgSetDelegateAddresses {
    const message = {
      ...baseMsgSetDelegateAddresses,
    } as MsgSetDelegateAddresses;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.validatorAddress !== undefined &&
      object.validatorAddress !== null
    ) {
      message.validatorAddress = object.validatorAddress;
    } else {
      message.validatorAddress = "";
    }
    if (
      object.orchestratorAddress !== undefined &&
      object.orchestratorAddress !== null
    ) {
      message.orchestratorAddress = object.orchestratorAddress;
    } else {
      message.orchestratorAddress = "";
    }
    if (object.ethAddress !== undefined && object.ethAddress !== null) {
      message.ethAddress = object.ethAddress;
    } else {
      message.ethAddress = "";
    }
    if (object.btcPublicKey !== undefined && object.btcPublicKey !== null) {
      message.btcPublicKey = object.btcPublicKey;
    } else {
      message.btcPublicKey = "";
    }
    return message;
  },
};

const baseMsgSetDelegateAddressesResponse: object = {};

export const MsgSetDelegateAddressesResponse = {
  encode(
    _: MsgSetDelegateAddressesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetDelegateAddressesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetDelegateAddressesResponse,
    } as MsgSetDelegateAddressesResponse;
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

  fromJSON(_: any): MsgSetDelegateAddressesResponse {
    const message = {
      ...baseMsgSetDelegateAddressesResponse,
    } as MsgSetDelegateAddressesResponse;
    return message;
  },

  toJSON(_: MsgSetDelegateAddressesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSetDelegateAddressesResponse>
  ): MsgSetDelegateAddressesResponse {
    const message = {
      ...baseMsgSetDelegateAddressesResponse,
    } as MsgSetDelegateAddressesResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SnapshotProposal(
    request: MsgSnapshotProposal
  ): Promise<MsgSnapshotProposalResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetDelegateAddresses(
    request: MsgSetDelegateAddresses
  ): Promise<MsgSetDelegateAddressesResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SnapshotProposal(
    request: MsgSnapshotProposal
  ): Promise<MsgSnapshotProposalResponse> {
    const data = MsgSnapshotProposal.encode(request).finish();
    const promise = this.rpc.request(
      "twilightproject.validitygadgetchain.validitygadgetchain.Msg",
      "SnapshotProposal",
      data
    );
    return promise.then((data) =>
      MsgSnapshotProposalResponse.decode(new Reader(data))
    );
  }

  SetDelegateAddresses(
    request: MsgSetDelegateAddresses
  ): Promise<MsgSetDelegateAddressesResponse> {
    const data = MsgSetDelegateAddresses.encode(request).finish();
    const promise = this.rpc.request(
      "twilightproject.validitygadgetchain.validitygadgetchain.Msg",
      "SetDelegateAddresses",
      data
    );
    return promise.then((data) =>
      MsgSetDelegateAddressesResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
