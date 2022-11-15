import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/newareyui/dex/tx";
import { MsgUpdatePool } from "./types/newareyui/dex/tx";
import { MsgDeletePool } from "./types/newareyui/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/newareyui.dex.MsgCreatePool", MsgCreatePool],
    ["/newareyui.dex.MsgUpdatePool", MsgUpdatePool],
    ["/newareyui.dex.MsgDeletePool", MsgDeletePool],
    
];

export { msgTypes }