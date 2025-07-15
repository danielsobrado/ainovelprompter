export namespace main {
	
	export class MCPCharacterData {
	    id: string;
	    name: string;
	    description: string;
	    traits?: Record<string, any>;
	    notes?: string;
	    createdAt?: string;
	    updatedAt?: string;
	
	    static createFrom(source: any = {}) {
	        return new MCPCharacterData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.traits = source["traits"];
	        this.notes = source["notes"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	    }
	}
	export class PromptVariant {
	    variantLabel?: string;
	    targetModelFamilies?: string[];
	    targetModels?: string[];
	    promptText: string;
	
	    static createFrom(source: any = {}) {
	        return new PromptVariant(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.variantLabel = source["variantLabel"];
	        this.targetModelFamilies = source["targetModelFamilies"];
	        this.targetModels = source["targetModels"];
	        this.promptText = source["promptText"];
	    }
	}
	export class PromptDefinition {
	    id: string;
	    label: string;
	    category: string;
	    order: number;
	    description?: string;
	    defaultPromptText: string;
	    variants?: PromptVariant[];
	
	    static createFrom(source: any = {}) {
	        return new PromptDefinition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	        this.category = source["category"];
	        this.order = source["order"];
	        this.description = source["description"];
	        this.defaultPromptText = source["defaultPromptText"];
	        this.variants = this.convertValues(source["variants"], PromptVariant);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace storage {
	
	export class StorageStats {
	    totalFiles: number;
	    totalSize: number;
	    entitiesByType: Record<string, number>;
	    versionsByType: Record<string, number>;
	    // Go type: time
	    oldestTimestamp: any;
	    // Go type: time
	    newestTimestamp: any;
	
	    static createFrom(source: any = {}) {
	        return new StorageStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.totalFiles = source["totalFiles"];
	        this.totalSize = source["totalSize"];
	        this.entitiesByType = source["entitiesByType"];
	        this.versionsByType = source["versionsByType"];
	        this.oldestTimestamp = this.convertValues(source["oldestTimestamp"], null);
	        this.newestTimestamp = this.convertValues(source["newestTimestamp"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

