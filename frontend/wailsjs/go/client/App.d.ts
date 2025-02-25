// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {client} from '../models';
import {config} from '../models';
import {context} from '../models';

export function CheckDatasets():Promise<client.Response>;

export function ChooseFile():Promise<config.ProfileItem>;

export function Debug(arg1:Array<any>):Promise<void>;

export function Error(arg1:Array<any>):Promise<void>;

export function Fatal(arg1:Array<any>):Promise<void>;

export function FetchProxies():Promise<client.Response>;

export function GetProfile():Promise<config.ProfileItem>;

export function Info(arg1:Array<any>):Promise<void>;

export function Shutdown(arg1:context.Context):Promise<void>;

export function UseFetchedDatasets():Promise<client.Response>;

export function Warn(arg1:Array<any>):Promise<void>;
