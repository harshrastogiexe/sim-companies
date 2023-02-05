export interface ResourceBase {
	name: string;
	image: string;
	db_letter: number;
	transportation: number;
	retailable: boolean | null;
	research: boolean;
	exchangeTradable: boolean;
	realmAvailable: boolean;
}

export interface Resource extends ResourceBase {
	producedFrom: any[];
	soldAt: null;
	soldAtRestaurant: null;
	producedAt: BuildingBase;
	neededFor: ResourceBase[];
	transportNeeded: number;
	producedAnHour: number;
	baseSalary: number;
	averageRetailPrice: number | null;
	marketSaturation: number | null;
	marketSaturationLabel: number | null;
	retailModeling: number | null;
	storeBaseSalary: number | null;
	retailData: any[];
	improvesQualityOf: ResourceBase[];
}

export interface BuildingBase {
	dbLetter: string;
	image: string;
	images: string[];
	name: string;
	category: string;
	cost: number;
	robotsNeeded: number;
	realmAvailable: boolean;
}
