import { Component } from '@angular/core';
import { BehaviorSubject, Observable, take, tap } from 'rxjs';
import { EncyclopediaService } from 'src/app/common/encyclopedia.service';
import { ResourceBase } from 'src/app/common/types';

@Component({
	selector: 'app-resources',
	templateUrl: './resources.component.html',
	styles: [],
})
export class ResourcesComponent {
	resources$: Observable<ResourceBase[]> = this.encyclopedia
		.fetchResources()
		.pipe(tap((t) => console.log(t)));

	resource$ = new BehaviorSubject<ResourceBase | null>(null);

	constructor(private readonly encyclopedia: EncyclopediaService) {
		this.resources$
			.pipe(take(1))
			.subscribe((resource) => this.resource$.next(resource[0]));
	}
}
