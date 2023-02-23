import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { map, merge, Observable, Subject, switchMap, tap } from 'rxjs';
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
	private resources: ResourceBase[] = [];

	private readonly market: IMarketService;

	private readonly resource: IResourceService;

	public readonly selected$ = new Subject<ResourceBase>();

	public readonly orders$ = new Subject<MarketOrder[]>();

	public readonly resources$: Observable<ResourceBase[]>;

	public readonly searchTextControl = new FormControl<string>('');

	public readonly filteredResources$: Observable<ResourceBase[]>;

	constructor(market: MarketService, resource: ResourceService) {
		this.market = market;
		this.resource = resource;
		this.resources$ = this.resource
			.getAllResource()
			.pipe(tap((resource) => (this.resources = resource)));
		this.filteredResources$ = merge(
			this.resources$,
			this.searchTextControl.valueChanges.pipe(
				map((value) => (value ? this.filterByName(value) : this.resources))
			)
		);
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

	filterByName(value: string): ResourceBase[] {
		return this.resources.filter((r) =>
			r.name.toLowerCase().includes(value.toLowerCase())
		);
	}
}
