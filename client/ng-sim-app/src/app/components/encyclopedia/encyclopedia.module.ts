import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';
import { DashboardComponent } from './dashboard/dashboard.component';
import { EncyclopediaRoutingModule } from './encyclopedia-routing.module';
import { ResourcesComponent } from './resources/resources.component';
import { BuildingComponent } from './building/building.component';

@NgModule({
	declarations: [DashboardComponent, ResourcesComponent, BuildingComponent],
	imports: [CommonModule, EncyclopediaRoutingModule, SharedModule],
	exports: [DashboardComponent],
})
export class EncyclopediaModule {}
