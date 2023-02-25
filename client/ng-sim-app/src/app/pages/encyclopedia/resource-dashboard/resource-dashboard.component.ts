import { Component } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { ResourceService } from 'src/common/services/resource.service';
import { IResourceService, ResourceBase, ResourceInfo } from 'src/common/types';

@Component({
	selector: 'app-resource-dashboard',
	templateUrl: './resource-dashboard.component.html',
	styleUrls: ['./resource-dashboard.component.scss'],
})
export class ResourceDashboardComponent {
	private readonly resource: IResourceService;

	public readonly resourceInfo$ = new BehaviorSubject<ResourceInfo | null>(
		null
	);

	constructor(resource: ResourceService) {
		this.resource = resource;
	}

	public onSelected(resource: ResourceBase) {
		this.resource
			.getResourceInfo(resource.db_letter.toString())
			.subscribe((resourceData) => this.resourceInfo$.next(resourceData));
	}
}
