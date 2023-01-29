import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { DropdownItemComponent } from './dropdown-item/dropdown-item.component';
import { DropdownComponent } from './dropdown/dropdown.component';
import { NavbarComponent } from './navbar/navbar.component';
import { SidePanelComponent } from './side-panel/side-panel.component';

@NgModule({
	declarations: [
		NavbarComponent,
		SidePanelComponent,
		DropdownComponent,
		DropdownItemComponent,
	],
	imports: [CommonModule, RouterModule],
	exports: [
		NavbarComponent,
		SidePanelComponent,
		DropdownComponent,
		DropdownItemComponent,
	],
})
export class SharedModule {}
