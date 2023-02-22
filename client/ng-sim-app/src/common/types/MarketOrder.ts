export type MarketOrder = {
	id: number;
	kind: number;
	quantity: number;
	quality: number;
	price: number;
	seller: {
		id: number;
		company: string;
		realmId: number;
		logo: string;
		certificates: number;
		contest_wins: number;
		npc: boolean;
		courseId: null;
		ip: string;
	};
	posted: string;
	fees: number;
};
