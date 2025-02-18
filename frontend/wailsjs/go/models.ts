export namespace client {
	
	export class Response {
	    Code: number;
	    Message: string;
	    Data: any;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Code = source["Code"];
	        this.Message = source["Message"];
	        this.Data = source["Data"];
	    }
	}

}

export namespace config {
	
	export class ProfileItem {
	    CoroutineCount: number;
	    LiveProxies: number;
	    AllProxies: number;
	    LiveProxyLists: string[];
	    Timeout: number;
	    SocksAddress: string;
	    FilePath: string;
	    Status: number;
	    Code: number;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new ProfileItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CoroutineCount = source["CoroutineCount"];
	        this.LiveProxies = source["LiveProxies"];
	        this.AllProxies = source["AllProxies"];
	        this.LiveProxyLists = source["LiveProxyLists"];
	        this.Timeout = source["Timeout"];
	        this.SocksAddress = source["SocksAddress"];
	        this.FilePath = source["FilePath"];
	        this.Status = source["Status"];
	        this.Code = source["Code"];
	        this.Error = source["Error"];
	    }
	}

}

