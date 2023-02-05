import { Component, EventEmitter, OnInit } from '@angular/core';
import {
	BehaviorSubject as Subject,
	map,
	merge,
	Observable,
	switchMap,
	take,
	tap,
} from 'rxjs';
import { EncyclopediaService } from 'src/app/common/encyclopedia.service';
import { Resource, ResourceBase } from 'src/app/common/types';

@Component({
	selector: 'app-resources',
	templateUrl: './resources.component.html',
	styles: [],
})
export class ResourcesComponent implements OnInit {
	resources$: Observable<ResourceBase[]> = this.encyclopedia.fetchResources();

	resourceId$ = new EventEmitter<string>();

	resource$ = new Subject<Resource | null>(null);

	constructor(private readonly encyclopedia: EncyclopediaService) {}

	ngOnInit(): void {
		const initialResourceId$ = this.resources$.pipe(
			take(1),
			map((resources) => resources[0]),
			map((resource) => resource.db_letter)
		);
		const resourceId$ = merge(this.resourceId$, initialResourceId$);

		resourceId$
			.pipe(
				tap((id) => console.log(id)),
				switchMap((id) => {
					return this.fetchResource(id!.toString());
				})
			)
			.subscribe((resource) => this.resource$.next(resource));
	}

	fetchResource(id: string): Observable<Resource> {
		return this.encyclopedia.fetchResourceById(id);
	}

	onResourceSelectHandler(id: string) {
		this.fetchResource(id).subscribe((resource) =>
			this.resource$.next(resource)
		);
	}
}
