import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { ResourceDashboardComponent } from './resource-dashboard/resource-dashboard.component';

@NgModule({
	imports: [
		RouterModule.forChild([
			{ path: 'encyclopedia/resource', component: ResourceDashboardComponent },
		]),
	],
})
export class EncyclopediaRouterModule {}
