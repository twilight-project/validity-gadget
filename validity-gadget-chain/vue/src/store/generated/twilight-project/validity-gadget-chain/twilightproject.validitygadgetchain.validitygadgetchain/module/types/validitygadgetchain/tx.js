/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
export const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
const baseMsgSnapshotProposal = {
    creator: "",
    iteration: "",
    logSnap: "",
};
export const MsgSnapshotProposal = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSnapshotProposal };
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
    fromJSON(object) {
        const message = { ...baseMsgSnapshotProposal };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.iteration !== undefined && object.iteration !== null) {
            message.iteration = String(object.iteration);
        }
        else {
            message.iteration = "";
        }
        if (object.logSnap !== undefined && object.logSnap !== null) {
            message.logSnap = String(object.logSnap);
        }
        else {
            message.logSnap = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.iteration !== undefined && (obj.iteration = message.iteration);
        message.logSnap !== undefined && (obj.logSnap = message.logSnap);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSnapshotProposal };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.iteration !== undefined && object.iteration !== null) {
            message.iteration = object.iteration;
        }
        else {
            message.iteration = "";
        }
        if (object.logSnap !== undefined && object.logSnap !== null) {
            message.logSnap = object.logSnap;
        }
        else {
            message.logSnap = "";
        }
        return message;
    },
};
const baseMsgSnapshotProposalResponse = { id: 0 };
export const MsgSnapshotProposalResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSnapshotProposalResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgSnapshotProposalResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgSnapshotProposalResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgSetDelegateAddresses = {
    creator: "",
    validator: "",
    btcOracle: "",
    ethOracle: "",
    watchTower: "",
};
export const MsgSetDelegateAddresses = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.validator !== "") {
            writer.uint32(18).string(message.validator);
        }
        if (message.btcOracle !== "") {
            writer.uint32(26).string(message.btcOracle);
        }
        if (message.ethOracle !== "") {
            writer.uint32(34).string(message.ethOracle);
        }
        if (message.watchTower !== "") {
            writer.uint32(42).string(message.watchTower);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSetDelegateAddresses,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.validator = reader.string();
                    break;
                case 3:
                    message.btcOracle = reader.string();
                    break;
                case 4:
                    message.ethOracle = reader.string();
                    break;
                case 5:
                    message.watchTower = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgSetDelegateAddresses,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.validator !== undefined && object.validator !== null) {
            message.validator = String(object.validator);
        }
        else {
            message.validator = "";
        }
        if (object.btcOracle !== undefined && object.btcOracle !== null) {
            message.btcOracle = String(object.btcOracle);
        }
        else {
            message.btcOracle = "";
        }
        if (object.ethOracle !== undefined && object.ethOracle !== null) {
            message.ethOracle = String(object.ethOracle);
        }
        else {
            message.ethOracle = "";
        }
        if (object.watchTower !== undefined && object.watchTower !== null) {
            message.watchTower = String(object.watchTower);
        }
        else {
            message.watchTower = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.validator !== undefined && (obj.validator = message.validator);
        message.btcOracle !== undefined && (obj.btcOracle = message.btcOracle);
        message.ethOracle !== undefined && (obj.ethOracle = message.ethOracle);
        message.watchTower !== undefined && (obj.watchTower = message.watchTower);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgSetDelegateAddresses,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.validator !== undefined && object.validator !== null) {
            message.validator = object.validator;
        }
        else {
            message.validator = "";
        }
        if (object.btcOracle !== undefined && object.btcOracle !== null) {
            message.btcOracle = object.btcOracle;
        }
        else {
            message.btcOracle = "";
        }
        if (object.ethOracle !== undefined && object.ethOracle !== null) {
            message.ethOracle = object.ethOracle;
        }
        else {
            message.ethOracle = "";
        }
        if (object.watchTower !== undefined && object.watchTower !== null) {
            message.watchTower = object.watchTower;
        }
        else {
            message.watchTower = "";
        }
        return message;
    },
};
const baseMsgSetDelegateAddressesResponse = {};
export const MsgSetDelegateAddressesResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSetDelegateAddressesResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgSetDelegateAddressesResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgSetDelegateAddressesResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    SnapshotProposal(request) {
        const data = MsgSnapshotProposal.encode(request).finish();
        const promise = this.rpc.request("twilightproject.validitygadgetchain.validitygadgetchain.Msg", "SnapshotProposal", data);
        return promise.then((data) => MsgSnapshotProposalResponse.decode(new Reader(data)));
    }
    SetDelegateAddresses(request) {
        const data = MsgSetDelegateAddresses.encode(request).finish();
        const promise = this.rpc.request("twilightproject.validitygadgetchain.validitygadgetchain.Msg", "SetDelegateAddresses", data);
        return promise.then((data) => MsgSetDelegateAddressesResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
