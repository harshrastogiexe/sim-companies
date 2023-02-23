import { Component, OnInit } from '@angular/core';
import { toRelativeTimeString } from '@utils/*';
import { map, merge, Observable, Subject, switchMap } from 'rxjs';
import { MarketService } from 'src/common/services/market.service';
import { ResourceService } from 'src/common/services/resource.service';
import {
	IMarketService,
	IResourceService,
	MarketOrder,
} from 'src/common/types';
import { ResourceBase } from 'src/common/types/ResourceBase';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
	public selected$ = new Subject<ResourceBase>();

	public orders$ = new Subject<MarketOrder[]>();

	public resources$!: Observable<ResourceBase[]>;

	private readonly market: IMarketService;

	private readonly resource: IResourceService;

	constructor(market: MarketService, resource: ResourceService) {
		this.market = market;
		this.resource = resource;

		this.resources$ = this.resource.getAllResource();
	}

	ngOnInit(): void {
		const INITIAL_RESOURCE_TO_DISPLAY = '1';
		const orders$ = merge(
			this.market.getMarketOrder(INITIAL_RESOURCE_TO_DISPLAY),
			this.selected$.pipe(
				map((r) => r.db_letter.toString()),
				switchMap((id) => this.market.getMarketOrder(id))
			)
		);

		orders$.subscribe((orders) => this.orders$.next(orders));
	}

	relativeTime(time: string): string {
		return toRelativeTimeString(new Date(time));
	}

	labelCb = (r: ResourceBase) => r.name;
	filterCb = (r: ResourceBase, text: string) =>
		r.name.toLowerCase().includes(text.toLowerCase());
}
