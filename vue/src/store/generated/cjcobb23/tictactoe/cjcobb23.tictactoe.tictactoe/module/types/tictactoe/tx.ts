/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "cjcobb23.tictactoe.tictactoe";

export interface MsgInvite {
  creator: string;
}

export interface MsgInviteResponse {
  gameIndex: number;
}

export interface MsgAccept {
  creator: string;
  gameIndex: number;
}

export interface MsgAcceptResponse {
  x: string;
  o: string;
}

export interface MsgMove {
  creator: string;
  gameIndex: number;
  x: number;
  y: number;
}

export interface MsgMoveResponse {
  winner: string;
}

const baseMsgInvite: object = { creator: "" };

export const MsgInvite = {
  encode(message: MsgInvite, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInvite {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInvite } as MsgInvite;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInvite {
    const message = { ...baseMsgInvite } as MsgInvite;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgInvite): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInvite>): MsgInvite {
    const message = { ...baseMsgInvite } as MsgInvite;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgInviteResponse: object = { gameIndex: 0 };

export const MsgInviteResponse = {
  encode(message: MsgInviteResponse, writer: Writer = Writer.create()): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInviteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInviteResponse } as MsgInviteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInviteResponse {
    const message = { ...baseMsgInviteResponse } as MsgInviteResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    return message;
  },

  toJSON(message: MsgInviteResponse): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInviteResponse>): MsgInviteResponse {
    const message = { ...baseMsgInviteResponse } as MsgInviteResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    return message;
  },
};

const baseMsgAccept: object = { creator: "", gameIndex: 0 };

export const MsgAccept = {
  encode(message: MsgAccept, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== 0) {
      writer.uint32(16).uint64(message.gameIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAccept {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAccept } as MsgAccept;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAccept {
    const message = { ...baseMsgAccept } as MsgAccept;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    return message;
  },

  toJSON(message: MsgAccept): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAccept>): MsgAccept {
    const message = { ...baseMsgAccept } as MsgAccept;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    return message;
  },
};

const baseMsgAcceptResponse: object = { x: "", o: "" };

export const MsgAcceptResponse = {
  encode(message: MsgAcceptResponse, writer: Writer = Writer.create()): Writer {
    if (message.x !== "") {
      writer.uint32(10).string(message.x);
    }
    if (message.o !== "") {
      writer.uint32(18).string(message.o);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAcceptResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAcceptResponse } as MsgAcceptResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.x = reader.string();
          break;
        case 2:
          message.o = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptResponse {
    const message = { ...baseMsgAcceptResponse } as MsgAcceptResponse;
    if (object.x !== undefined && object.x !== null) {
      message.x = String(object.x);
    } else {
      message.x = "";
    }
    if (object.o !== undefined && object.o !== null) {
      message.o = String(object.o);
    } else {
      message.o = "";
    }
    return message;
  },

  toJSON(message: MsgAcceptResponse): unknown {
    const obj: any = {};
    message.x !== undefined && (obj.x = message.x);
    message.o !== undefined && (obj.o = message.o);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAcceptResponse>): MsgAcceptResponse {
    const message = { ...baseMsgAcceptResponse } as MsgAcceptResponse;
    if (object.x !== undefined && object.x !== null) {
      message.x = object.x;
    } else {
      message.x = "";
    }
    if (object.o !== undefined && object.o !== null) {
      message.o = object.o;
    } else {
      message.o = "";
    }
    return message;
  },
};

const baseMsgMove: object = { creator: "", gameIndex: 0, x: 0, y: 0 };

export const MsgMove = {
  encode(message: MsgMove, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== 0) {
      writer.uint32(16).uint64(message.gameIndex);
    }
    if (message.x !== 0) {
      writer.uint32(24).uint64(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(32).uint64(message.y);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMove {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMove } as MsgMove;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.x = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.y = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMove {
    const message = { ...baseMsgMove } as MsgMove;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    if (object.x !== undefined && object.x !== null) {
      message.x = Number(object.x);
    } else {
      message.x = 0;
    }
    if (object.y !== undefined && object.y !== null) {
      message.y = Number(object.y);
    } else {
      message.y = 0;
    }
    return message;
  },

  toJSON(message: MsgMove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.x !== undefined && (obj.x = message.x);
    message.y !== undefined && (obj.y = message.y);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMove>): MsgMove {
    const message = { ...baseMsgMove } as MsgMove;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    if (object.x !== undefined && object.x !== null) {
      message.x = object.x;
    } else {
      message.x = 0;
    }
    if (object.y !== undefined && object.y !== null) {
      message.y = object.y;
    } else {
      message.y = 0;
    }
    return message;
  },
};

const baseMsgMoveResponse: object = { winner: "" };

export const MsgMoveResponse = {
  encode(message: MsgMoveResponse, writer: Writer = Writer.create()): Writer {
    if (message.winner !== "") {
      writer.uint32(10).string(message.winner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMoveResponse } as MsgMoveResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMoveResponse {
    const message = { ...baseMsgMoveResponse } as MsgMoveResponse;
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    return message;
  },

  toJSON(message: MsgMoveResponse): unknown {
    const obj: any = {};
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMoveResponse>): MsgMoveResponse {
    const message = { ...baseMsgMoveResponse } as MsgMoveResponse;
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Invite(request: MsgInvite): Promise<MsgInviteResponse>;
  Accept(request: MsgAccept): Promise<MsgAcceptResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Move(request: MsgMove): Promise<MsgMoveResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Invite(request: MsgInvite): Promise<MsgInviteResponse> {
    const data = MsgInvite.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Msg",
      "Invite",
      data
    );
    return promise.then((data) => MsgInviteResponse.decode(new Reader(data)));
  }

  Accept(request: MsgAccept): Promise<MsgAcceptResponse> {
    const data = MsgAccept.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Msg",
      "Accept",
      data
    );
    return promise.then((data) => MsgAcceptResponse.decode(new Reader(data)));
  }

  Move(request: MsgMove): Promise<MsgMoveResponse> {
    const data = MsgMove.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Msg",
      "Move",
      data
    );
    return promise.then((data) => MsgMoveResponse.decode(new Reader(data)));
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
