/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cjcobb23.tictactoe.tictactoe";

export interface StoredGame {
  index: number;
  state: string;
  x: string;
  o: string;
  prevIndex: number;
  nextIndex: number;
  blockHeightExpiration: number;
}

const baseStoredGame: object = {
  index: 0,
  state: "",
  x: "",
  o: "",
  prevIndex: 0,
  nextIndex: 0,
  blockHeightExpiration: 0,
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== 0) {
      writer.uint32(8).uint64(message.index);
    }
    if (message.state !== "") {
      writer.uint32(18).string(message.state);
    }
    if (message.x !== "") {
      writer.uint32(26).string(message.x);
    }
    if (message.o !== "") {
      writer.uint32(34).string(message.o);
    }
    if (message.prevIndex !== 0) {
      writer.uint32(40).uint64(message.prevIndex);
    }
    if (message.nextIndex !== 0) {
      writer.uint32(48).uint64(message.nextIndex);
    }
    if (message.blockHeightExpiration !== 0) {
      writer.uint32(56).uint64(message.blockHeightExpiration);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredGame } as StoredGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.state = reader.string();
          break;
        case 3:
          message.x = reader.string();
          break;
        case 4:
          message.o = reader.string();
          break;
        case 5:
          message.prevIndex = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.nextIndex = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.blockHeightExpiration = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index);
    } else {
      message.index = 0;
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
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
    if (object.prevIndex !== undefined && object.prevIndex !== null) {
      message.prevIndex = Number(object.prevIndex);
    } else {
      message.prevIndex = 0;
    }
    if (object.nextIndex !== undefined && object.nextIndex !== null) {
      message.nextIndex = Number(object.nextIndex);
    } else {
      message.nextIndex = 0;
    }
    if (
      object.blockHeightExpiration !== undefined &&
      object.blockHeightExpiration !== null
    ) {
      message.blockHeightExpiration = Number(object.blockHeightExpiration);
    } else {
      message.blockHeightExpiration = 0;
    }
    return message;
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.state !== undefined && (obj.state = message.state);
    message.x !== undefined && (obj.x = message.x);
    message.o !== undefined && (obj.o = message.o);
    message.prevIndex !== undefined && (obj.prevIndex = message.prevIndex);
    message.nextIndex !== undefined && (obj.nextIndex = message.nextIndex);
    message.blockHeightExpiration !== undefined &&
      (obj.blockHeightExpiration = message.blockHeightExpiration);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredGame>): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = 0;
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
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
    if (object.prevIndex !== undefined && object.prevIndex !== null) {
      message.prevIndex = object.prevIndex;
    } else {
      message.prevIndex = 0;
    }
    if (object.nextIndex !== undefined && object.nextIndex !== null) {
      message.nextIndex = object.nextIndex;
    } else {
      message.nextIndex = 0;
    }
    if (
      object.blockHeightExpiration !== undefined &&
      object.blockHeightExpiration !== null
    ) {
      message.blockHeightExpiration = object.blockHeightExpiration;
    } else {
      message.blockHeightExpiration = 0;
    }
    return message;
  },
};

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
