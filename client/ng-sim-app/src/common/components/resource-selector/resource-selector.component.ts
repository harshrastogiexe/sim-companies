import { Component, EventEmitter, OnDestroy, Output } from '@angular/core';
import { FormControl } from '@angular/forms';
import {
	exhaustMap,
	map,
	merge,
	Observable,
	of,
	Subject,
	Subscription,
	tap,
} from 'rxjs';
import { ResourceService } from 'src/common/services/resource.service';
import { ResourceBase } from 'src/common/types/ResourceBase';

let resourcesCache: ResourceBase[] | null = null;

const filterBy = {
	NAME: (resources: ResourceBase[], value: string) =>
		resources.filter((r) => r.name.toLowerCase().includes(value.toLowerCase())),
};

@Component({
	selector: 'app-resource-selector',
	templateUrl: './resource-selector.component.html',
})
export class ResourceSelectorComponent implements OnDestroy {
	private readonly resource: ResourceService;

	private readonly subscription: Subscription[] = [];

	public readonly selected$ = new Subject<ResourceBase>();

	public readonly searchTextFc = new FormControl<string>('');

	public readonly filtered$: Observable<ResourceBase[]>;

	@Output('selected')
	public readonly selectEventEmitter = new EventEmitter<ResourceBase>();

	constructor(resource: ResourceService) {
		this.resource = resource;
		this.filtered$ = merge(
			this.getResources(),
			this.searchTextFc.valueChanges.pipe(
				exhaustMap((value) => this.getFilteredResources(value!))
			)
		);
		const selectSub = this.selected$
			.pipe(
				tap((r) => this.selectEventEmitter.next(r)),
				map((r) => r.name)
			)
			.subscribe((name) => this.searchTextFc.setValue(name));

		this.subscription.push(selectSub);
	}

	ngOnDestroy(): void {
		for (const sub of this.subscription) sub.unsubscribe();
	}

	getResources(): Observable<ResourceBase[]> {
		if (resourcesCache !== null) {
			return of(resourcesCache);
		}

		return this.resource
			.getAllResource()
			.pipe(tap((resources) => (resourcesCache = resources)));
	}

	getFilteredResources(value: string): Observable<ResourceBase[]> {
		return this.getResources().pipe(
			map((resources) =>
				value ? filterBy['NAME'](resources, value) : resourcesCache!.slice()
			)
		);
	}
}
