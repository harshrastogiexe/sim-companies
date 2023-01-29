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
