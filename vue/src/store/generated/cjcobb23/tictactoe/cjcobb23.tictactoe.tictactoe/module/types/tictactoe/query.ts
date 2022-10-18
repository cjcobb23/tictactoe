/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../tictactoe/params";
import { SystemInfo } from "../tictactoe/system_info";
import { StoredGame } from "../tictactoe/stored_game";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "cjcobb23.tictactoe.tictactoe";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSystemInfoRequest {}

export interface QueryGetSystemInfoResponse {
  SystemInfo: SystemInfo | undefined;
}

export interface QueryGetStoredGameRequest {
  index: number;
}

export interface QueryGetStoredGameResponse {
  storedGame: StoredGame | undefined;
}

export interface QueryAllStoredGameRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredGameResponse {
  storedGame: StoredGame[];
  pagination: PageResponse | undefined;
}

export interface QueryCanPlayMoveRequest {
  gameIndex: number;
  player: string;
  x: number;
  y: number;
}

export interface QueryCanPlayMoveResponse {
  possible: boolean;
  reason: string;
}

export interface QueryStoredGameHumanRequest {
  gameIndex: number;
}

export interface QueryStoredGameHumanResponse {
  gameIndex: number;
  board: string;
  turn: string;
  winner: string;
  draw: boolean;
  x: string;
  o: string;
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

const baseQueryGetSystemInfoRequest: object = {};

export const QueryGetSystemInfoRequest = {
  encode(
    _: QueryGetSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
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

  fromJSON(_: any): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },

  toJSON(_: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetSystemInfoRequest>
  ): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },
};

const baseQueryGetSystemInfoResponse: object = {};

export const QueryGetSystemInfoResponse = {
  encode(
    message: QueryGetSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SystemInfo !== undefined) {
      SystemInfo.encode(message.SystemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SystemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromJSON(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.SystemInfo !== undefined &&
      (obj.SystemInfo = message.SystemInfo
        ? SystemInfo.toJSON(message.SystemInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoResponse>
  ): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromPartial(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetStoredGameRequest: object = { index: 0 };

export const QueryGetStoredGameRequest = {
  encode(
    message: QueryGetStoredGameRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== 0) {
      writer.uint32(8).uint64(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredGameRequest,
    } as QueryGetStoredGameRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredGameRequest {
    const message = {
      ...baseQueryGetStoredGameRequest,
    } as QueryGetStoredGameRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index);
    } else {
      message.index = 0;
    }
    return message;
  },

  toJSON(message: QueryGetStoredGameRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredGameRequest>
  ): QueryGetStoredGameRequest {
    const message = {
      ...baseQueryGetStoredGameRequest,
    } as QueryGetStoredGameRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = 0;
    }
    return message;
  },
};

const baseQueryGetStoredGameResponse: object = {};

export const QueryGetStoredGameResponse = {
  encode(
    message: QueryGetStoredGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storedGame !== undefined) {
      StoredGame.encode(message.storedGame, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredGameResponse,
    } as QueryGetStoredGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedGame = StoredGame.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredGameResponse {
    const message = {
      ...baseQueryGetStoredGameResponse,
    } as QueryGetStoredGameResponse;
    if (object.storedGame !== undefined && object.storedGame !== null) {
      message.storedGame = StoredGame.fromJSON(object.storedGame);
    } else {
      message.storedGame = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStoredGameResponse): unknown {
    const obj: any = {};
    message.storedGame !== undefined &&
      (obj.storedGame = message.storedGame
        ? StoredGame.toJSON(message.storedGame)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredGameResponse>
  ): QueryGetStoredGameResponse {
    const message = {
      ...baseQueryGetStoredGameResponse,
    } as QueryGetStoredGameResponse;
    if (object.storedGame !== undefined && object.storedGame !== null) {
      message.storedGame = StoredGame.fromPartial(object.storedGame);
    } else {
      message.storedGame = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredGameRequest: object = {};

export const QueryAllStoredGameRequest = {
  encode(
    message: QueryAllStoredGameRequest,
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
  ): QueryAllStoredGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredGameRequest,
    } as QueryAllStoredGameRequest;
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

  fromJSON(object: any): QueryAllStoredGameRequest {
    const message = {
      ...baseQueryAllStoredGameRequest,
    } as QueryAllStoredGameRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredGameRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredGameRequest>
  ): QueryAllStoredGameRequest {
    const message = {
      ...baseQueryAllStoredGameRequest,
    } as QueryAllStoredGameRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredGameResponse: object = {};

export const QueryAllStoredGameResponse = {
  encode(
    message: QueryAllStoredGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.storedGame) {
      StoredGame.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllStoredGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredGameResponse,
    } as QueryAllStoredGameResponse;
    message.storedGame = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedGame.push(StoredGame.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllStoredGameResponse {
    const message = {
      ...baseQueryAllStoredGameResponse,
    } as QueryAllStoredGameResponse;
    message.storedGame = [];
    if (object.storedGame !== undefined && object.storedGame !== null) {
      for (const e of object.storedGame) {
        message.storedGame.push(StoredGame.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredGameResponse): unknown {
    const obj: any = {};
    if (message.storedGame) {
      obj.storedGame = message.storedGame.map((e) =>
        e ? StoredGame.toJSON(e) : undefined
      );
    } else {
      obj.storedGame = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredGameResponse>
  ): QueryAllStoredGameResponse {
    const message = {
      ...baseQueryAllStoredGameResponse,
    } as QueryAllStoredGameResponse;
    message.storedGame = [];
    if (object.storedGame !== undefined && object.storedGame !== null) {
      for (const e of object.storedGame) {
        message.storedGame.push(StoredGame.fromPartial(e));
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

const baseQueryCanPlayMoveRequest: object = {
  gameIndex: 0,
  player: "",
  x: 0,
  y: 0,
};

export const QueryCanPlayMoveRequest = {
  encode(
    message: QueryCanPlayMoveRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.x !== 0) {
      writer.uint32(24).uint64(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(32).uint64(message.y);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCanPlayMoveRequest,
    } as QueryCanPlayMoveRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.player = reader.string();
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

  fromJSON(object: any): QueryCanPlayMoveRequest {
    const message = {
      ...baseQueryCanPlayMoveRequest,
    } as QueryCanPlayMoveRequest;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
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

  toJSON(message: QueryCanPlayMoveRequest): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.player !== undefined && (obj.player = message.player);
    message.x !== undefined && (obj.x = message.x);
    message.y !== undefined && (obj.y = message.y);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCanPlayMoveRequest>
  ): QueryCanPlayMoveRequest {
    const message = {
      ...baseQueryCanPlayMoveRequest,
    } as QueryCanPlayMoveRequest;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
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

const baseQueryCanPlayMoveResponse: object = { possible: false, reason: "" };

export const QueryCanPlayMoveResponse = {
  encode(
    message: QueryCanPlayMoveResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.possible === true) {
      writer.uint32(8).bool(message.possible);
    }
    if (message.reason !== "") {
      writer.uint32(18).string(message.reason);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCanPlayMoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCanPlayMoveResponse,
    } as QueryCanPlayMoveResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.possible = reader.bool();
          break;
        case 2:
          message.reason = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCanPlayMoveResponse {
    const message = {
      ...baseQueryCanPlayMoveResponse,
    } as QueryCanPlayMoveResponse;
    if (object.possible !== undefined && object.possible !== null) {
      message.possible = Boolean(object.possible);
    } else {
      message.possible = false;
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = String(object.reason);
    } else {
      message.reason = "";
    }
    return message;
  },

  toJSON(message: QueryCanPlayMoveResponse): unknown {
    const obj: any = {};
    message.possible !== undefined && (obj.possible = message.possible);
    message.reason !== undefined && (obj.reason = message.reason);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCanPlayMoveResponse>
  ): QueryCanPlayMoveResponse {
    const message = {
      ...baseQueryCanPlayMoveResponse,
    } as QueryCanPlayMoveResponse;
    if (object.possible !== undefined && object.possible !== null) {
      message.possible = object.possible;
    } else {
      message.possible = false;
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = object.reason;
    } else {
      message.reason = "";
    }
    return message;
  },
};

const baseQueryStoredGameHumanRequest: object = { gameIndex: 0 };

export const QueryStoredGameHumanRequest = {
  encode(
    message: QueryStoredGameHumanRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryStoredGameHumanRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryStoredGameHumanRequest,
    } as QueryStoredGameHumanRequest;
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

  fromJSON(object: any): QueryStoredGameHumanRequest {
    const message = {
      ...baseQueryStoredGameHumanRequest,
    } as QueryStoredGameHumanRequest;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    return message;
  },

  toJSON(message: QueryStoredGameHumanRequest): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryStoredGameHumanRequest>
  ): QueryStoredGameHumanRequest {
    const message = {
      ...baseQueryStoredGameHumanRequest,
    } as QueryStoredGameHumanRequest;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    return message;
  },
};

const baseQueryStoredGameHumanResponse: object = {
  gameIndex: 0,
  board: "",
  turn: "",
  winner: "",
  draw: false,
  x: "",
  o: "",
};

export const QueryStoredGameHumanResponse = {
  encode(
    message: QueryStoredGameHumanResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    if (message.board !== "") {
      writer.uint32(18).string(message.board);
    }
    if (message.turn !== "") {
      writer.uint32(26).string(message.turn);
    }
    if (message.winner !== "") {
      writer.uint32(34).string(message.winner);
    }
    if (message.draw === true) {
      writer.uint32(40).bool(message.draw);
    }
    if (message.x !== "") {
      writer.uint32(50).string(message.x);
    }
    if (message.o !== "") {
      writer.uint32(58).string(message.o);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryStoredGameHumanResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryStoredGameHumanResponse,
    } as QueryStoredGameHumanResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.board = reader.string();
          break;
        case 3:
          message.turn = reader.string();
          break;
        case 4:
          message.winner = reader.string();
          break;
        case 5:
          message.draw = reader.bool();
          break;
        case 6:
          message.x = reader.string();
          break;
        case 7:
          message.o = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryStoredGameHumanResponse {
    const message = {
      ...baseQueryStoredGameHumanResponse,
    } as QueryStoredGameHumanResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = String(object.board);
    } else {
      message.board = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = String(object.turn);
    } else {
      message.turn = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    if (object.draw !== undefined && object.draw !== null) {
      message.draw = Boolean(object.draw);
    } else {
      message.draw = false;
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
    return message;
  },

  toJSON(message: QueryStoredGameHumanResponse): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.board !== undefined && (obj.board = message.board);
    message.turn !== undefined && (obj.turn = message.turn);
    message.winner !== undefined && (obj.winner = message.winner);
    message.draw !== undefined && (obj.draw = message.draw);
    message.x !== undefined && (obj.x = message.x);
    message.o !== undefined && (obj.o = message.o);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryStoredGameHumanResponse>
  ): QueryStoredGameHumanResponse {
    const message = {
      ...baseQueryStoredGameHumanResponse,
    } as QueryStoredGameHumanResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = object.board;
    } else {
      message.board = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = object.turn;
    } else {
      message.turn = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    if (object.draw !== undefined && object.draw !== null) {
      message.draw = object.draw;
    } else {
      message.draw = false;
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
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SystemInfo by index. */
  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse>;
  /** Queries a StoredGame by index. */
  StoredGame(
    request: QueryGetStoredGameRequest
  ): Promise<QueryGetStoredGameResponse>;
  /** Queries a list of StoredGame items. */
  StoredGameAll(
    request: QueryAllStoredGameRequest
  ): Promise<QueryAllStoredGameResponse>;
  /** Queries a list of CanPlayMove items. */
  CanPlayMove(
    request: QueryCanPlayMoveRequest
  ): Promise<QueryCanPlayMoveResponse>;
  /** Queries a list of StoredGameHuman items. */
  StoredGameHuman(
    request: QueryStoredGameHumanRequest
  ): Promise<QueryStoredGameHumanResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Query",
      "SystemInfo",
      data
    );
    return promise.then((data) =>
      QueryGetSystemInfoResponse.decode(new Reader(data))
    );
  }

  StoredGame(
    request: QueryGetStoredGameRequest
  ): Promise<QueryGetStoredGameResponse> {
    const data = QueryGetStoredGameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Query",
      "StoredGame",
      data
    );
    return promise.then((data) =>
      QueryGetStoredGameResponse.decode(new Reader(data))
    );
  }

  StoredGameAll(
    request: QueryAllStoredGameRequest
  ): Promise<QueryAllStoredGameResponse> {
    const data = QueryAllStoredGameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Query",
      "StoredGameAll",
      data
    );
    return promise.then((data) =>
      QueryAllStoredGameResponse.decode(new Reader(data))
    );
  }

  CanPlayMove(
    request: QueryCanPlayMoveRequest
  ): Promise<QueryCanPlayMoveResponse> {
    const data = QueryCanPlayMoveRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Query",
      "CanPlayMove",
      data
    );
    return promise.then((data) =>
      QueryCanPlayMoveResponse.decode(new Reader(data))
    );
  }

  StoredGameHuman(
    request: QueryStoredGameHumanRequest
  ): Promise<QueryStoredGameHumanResponse> {
    const data = QueryStoredGameHumanRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cjcobb23.tictactoe.tictactoe.Query",
      "StoredGameHuman",
      data
    );
    return promise.then((data) =>
      QueryStoredGameHumanResponse.decode(new Reader(data))
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
