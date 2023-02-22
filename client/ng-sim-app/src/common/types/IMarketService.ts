import { Observable } from 'rxjs';
import { MarketOrder } from './MarketOrder';

export interface IMarketService {
	/**
	 * fetches the current market orders for the particular resource
	 * @param resourceId id for resource from which data is fetched
	 * @returns observable for list market item orders
	 */
	getMarketOrder(resourceId: string): Observable<MarketOrder[]>;
}
