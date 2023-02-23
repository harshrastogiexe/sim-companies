import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environment/environment';
import { IMarketService, MarketOrder } from '../types';

@Injectable({ providedIn: 'root' })
export class MarketService implements IMarketService {
	private readonly http: HttpClient;

	public constructor(http: HttpClient) {
		this.http = http;
	}

	public getMarketOrder(resourceId: string): Observable<MarketOrder[]> {
		const uri = environment.server.basePath + '/market/' + resourceId;
		return this.http.get<MarketOrder[]>(uri);
	}
}
