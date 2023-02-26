import { ResourceBase } from './ResourceBase';

export interface ResourceInfo extends ResourceBase {
	id: string;
	sold_at: string | null;
	sold_at_restaurant: string | null;
	produced_at: string | null;
	needed_for: number[];
	improves_quality_of: number[];
	transport_needed: number;
	produced_an_hour: number;
	base_salary: number;
	average_retail_price: number | null;
	market_saturation: number | null;
	market_saturation_label: string | null;
	retail_modeling: string | null;
	store_base_salary: number | null;
	producedFrom: { Amount: number; ResourceId: number }[];
}
