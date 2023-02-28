import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ComponentsModule } from 'src/common/components/components.module';
import { EncyclopediaRouterModule } from './encyclopedia-routing.module';
import { ResourceDashboardComponent } from './resource-dashboard/resource-dashboard.component';
import { ResourceInfoFieldComponent } from './resource-info-field/resource-info-field.component';
import { ResourceItemInfoComponent } from './resource-info/resource-item-info.component';
import { ResourceItemInfoModifierComponent } from './resource-item-info-modifier/resource-item-info-modifier.component';

@NgModule({
	declarations: [
		ResourceDashboardComponent,
		ResourceItemInfoComponent,
		ResourceInfoFieldComponent,
		ResourceItemInfoModifierComponent,
	],
	imports: [CommonModule, EncyclopediaRouterModule, ComponentsModule],
})
export class EncyclopediaPageModule {}
