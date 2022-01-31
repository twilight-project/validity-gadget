/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../validitygadgetchain/params";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { Proposal } from "../validitygadgetchain/proposal";
export const protobufPackage = "twilightproject.validitygadgetchain.validitygadgetchain";
const baseQueryParamsRequest = {};
export const QueryParamsRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsRequest };
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
        const message = { ...baseQueryParamsRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
};
const baseQueryParamsResponse = {};
export const QueryParamsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
};
const baseQuerySnapshotProposalsRequest = {};
export const QuerySnapshotProposalsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySnapshotProposalsRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQuerySnapshotProposalsRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQuerySnapshotProposalsRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQuerySnapshotProposalsResponse = {};
export const QuerySnapshotProposalsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Proposal) {
            Proposal.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySnapshotProposalsResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQuerySnapshotProposalsResponse,
        };
        message.Proposal = [];
        if (object.Proposal !== undefined && object.Proposal !== null) {
            for (const e of object.Proposal) {
                message.Proposal.push(Proposal.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Proposal) {
            obj.Proposal = message.Proposal.map((e) => e ? Proposal.toJSON(e) : undefined);
        }
        else {
            obj.Proposal = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQuerySnapshotProposalsResponse,
        };
        message.Proposal = [];
        if (object.Proposal !== undefined && object.Proposal !== null) {
            for (const e of object.Proposal) {
                message.Proposal.push(Proposal.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQuerySnapshotProposalRequest = {};
export const QuerySnapshotProposalRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySnapshotProposalRequest,
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
            ...baseQuerySnapshotProposalRequest,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseQuerySnapshotProposalRequest,
        };
        return message;
    },
};
const baseQuerySnapshotProposalResponse = {
    iteration: "",
    logsnap: "",
};
export const QuerySnapshotProposalResponse = {
    encode(message, writer = Writer.create()) {
        if (message.iteration !== "") {
            writer.uint32(10).string(message.iteration);
        }
        if (message.logsnap !== "") {
            writer.uint32(18).string(message.logsnap);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySnapshotProposalResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQuerySnapshotProposalResponse,
        };
        if (object.iteration !== undefined && object.iteration !== null) {
            message.iteration = String(object.iteration);
        }
        else {
            message.iteration = "";
        }
        if (object.logsnap !== undefined && object.logsnap !== null) {
            message.logsnap = String(object.logsnap);
        }
        else {
            message.logsnap = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.iteration !== undefined && (obj.iteration = message.iteration);
        message.logsnap !== undefined && (obj.logsnap = message.logsnap);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQuerySnapshotProposalResponse,
        };
        if (object.iteration !== undefined && object.iteration !== null) {
            message.iteration = object.iteration;
        }
        else {
            message.iteration = "";
        }
        if (object.logsnap !== undefined && object.logsnap !== null) {
            message.logsnap = object.logsnap;
        }
        else {
            message.logsnap = "";
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Params(request) {
        const data = QueryParamsRequest.encode(request).finish();
        const promise = this.rpc.request("twilightproject.validitygadgetchain.validitygadgetchain.Query", "Params", data);
        return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
    }
    SnapshotProposals(request) {
        const data = QuerySnapshotProposalsRequest.encode(request).finish();
        const promise = this.rpc.request("twilightproject.validitygadgetchain.validitygadgetchain.Query", "SnapshotProposals", data);
        return promise.then((data) => QuerySnapshotProposalsResponse.decode(new Reader(data)));
    }
}
