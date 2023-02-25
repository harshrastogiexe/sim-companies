export interface BuildingInfo {
	id: string;
	name: string;
	category: string;
	robots_needed: number;
	realm_available: boolean;
	images: string[];
	doesSell: number[];
	cost_units: number;
	hours: number;
	wages: number;
}
