import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ToRelativeTimeString } from '../services';
import { HeaderComponent } from './header/header.component';
import { MarketItemComponent } from './market-item/market-item.component';
import { DropdownInputDirective } from './shared/dropdown/dropdown-input/dropdown-input.directive';
import { DropdownComponent } from './shared/dropdown/dropdown.component';

@NgModule({
	declarations: [
		MarketItemComponent,
		ToRelativeTimeString,
		DropdownComponent,
		DropdownInputDirective,
		HeaderComponent,
	],
	imports: [CommonModule, ReactiveFormsModule],
	exports: [
		MarketItemComponent,
		ToRelativeTimeString,
		DropdownComponent,
		DropdownInputDirective,
		HeaderComponent,
	],
})
export class ComponentsModule {}
