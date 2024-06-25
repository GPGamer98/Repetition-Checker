export namespace main {
	
	export class EntryCount {
	
	
	    static createFrom(source: any = {}) {
	        return new EntryCount(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

