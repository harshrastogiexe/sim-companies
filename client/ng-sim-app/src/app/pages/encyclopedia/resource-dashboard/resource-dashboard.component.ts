import { Component } from '@angular/core';
import { ResourceBase } from 'src/common/types/ResourceBase';

@Component({
	selector: 'app-resource-dashboard',
	templateUrl: './resource-dashboard.component.html',
	styleUrls: ['./resource-dashboard.component.scss'],
})
export class ResourceDashboardComponent {
	onSelected(resource: ResourceBase) {
		console.log(resource);
	}
}
