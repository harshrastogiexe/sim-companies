import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BuildingComponent } from './building/building.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { ResourcesComponent } from './resources/resources.component';

const routes: Routes = [
	{
		path: '',
		component: DashboardComponent,
		children: [
			{ path: '', component: ResourcesComponent },
			{ path: 'building', component: BuildingComponent },
		],
	},
];

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule],
})
export class EncyclopediaRoutingModule {}
