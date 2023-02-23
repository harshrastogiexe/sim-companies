import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ToRelativeTimeString } from '../services';
import { HeaderComponent } from './header/header.component';
import { MarketItemComponent } from './market-item/market-item.component';
import { DropdownInputDirective } from './shared/dropdown/dropdown-input/dropdown-input.directive';
import { DropdownComponent } from './shared/dropdown/dropdown.component';
import { InputFilterListItemDirective } from './shared/input-filter-list/input-filter-list-item.directive';
import { InputFilterListComponent } from './shared/input-filter-list/input-filter-list.component';
import { InputTextFilterSearchDirective } from './shared/input-filter-list/input-text-filter-search.directive';

@NgModule({
	declarations: [
		MarketItemComponent,
		ToRelativeTimeString,
		DropdownComponent,
		DropdownInputDirective,
		HeaderComponent,
		InputFilterListComponent,
		InputTextFilterSearchDirective,
		InputFilterListItemDirective,
	],
	imports: [CommonModule, ReactiveFormsModule],
	exports: [
		MarketItemComponent,
		ToRelativeTimeString,
		DropdownComponent,
		DropdownInputDirective,
		HeaderComponent,
		InputFilterListComponent,
		InputFilterListItemDirective,
		InputTextFilterSearchDirective,
	],
})
export class ComponentsModule {}
