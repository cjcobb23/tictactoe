// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgMove } from "./types/tictactoe/tx";
import { MsgAccept } from "./types/tictactoe/tx";
import { MsgInvite } from "./types/tictactoe/tx";


const types = [
  ["/cjcobb23.tictactoe.tictactoe.MsgMove", MsgMove],
  ["/cjcobb23.tictactoe.tictactoe.MsgAccept", MsgAccept],
  ["/cjcobb23.tictactoe.tictactoe.MsgInvite", MsgInvite],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgMove: (data: MsgMove): EncodeObject => ({ typeUrl: "/cjcobb23.tictactoe.tictactoe.MsgMove", value: MsgMove.fromPartial( data ) }),
    msgAccept: (data: MsgAccept): EncodeObject => ({ typeUrl: "/cjcobb23.tictactoe.tictactoe.MsgAccept", value: MsgAccept.fromPartial( data ) }),
    msgInvite: (data: MsgInvite): EncodeObject => ({ typeUrl: "/cjcobb23.tictactoe.tictactoe.MsgInvite", value: MsgInvite.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
