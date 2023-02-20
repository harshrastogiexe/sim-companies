import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ToRelativeTimeString } from '../services';
import { MarketItemComponent } from './market-item/market-item.component';

@NgModule({
	declarations: [MarketItemComponent, ToRelativeTimeString],
	imports: [CommonModule],
	exports: [MarketItemComponent, ToRelativeTimeString],
})
export class ComponentsModule {}
