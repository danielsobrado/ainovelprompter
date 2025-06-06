export namespace main {
	
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

