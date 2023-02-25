import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ComponentsModule } from 'src/common/components/components.module';
import { EncyclopediaRouterModule } from './encyclopedia-routing.module';
import { ResourceDashboardComponent } from './resource-dashboard/resource-dashboard.component';
import { ResourceItemInfoComponent } from './resource-info/resource-item-info.component';

@NgModule({
	declarations: [ResourceDashboardComponent, ResourceItemInfoComponent],
	imports: [CommonModule, EncyclopediaRouterModule, ComponentsModule],
})
export class EncyclopediaPageModule {}
