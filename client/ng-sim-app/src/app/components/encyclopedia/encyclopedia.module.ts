import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';
import { BuildingComponent } from './building/building.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { EncyclopediaRoutingModule } from './encyclopedia-routing.module';
import { ResourcesComponent } from './resources/resources.component';
import { ResourceItemComponent } from './resources/resource-item/resource-item.component';

@NgModule({
	declarations: [DashboardComponent, ResourcesComponent, BuildingComponent, ResourceItemComponent],
	imports: [CommonModule, EncyclopediaRoutingModule, SharedModule],
	exports: [DashboardComponent],
})
export class EncyclopediaModule {}
