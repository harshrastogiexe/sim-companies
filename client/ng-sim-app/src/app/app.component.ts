import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { toRelativeTimeString } from '@utils/*';
import { Observable } from 'rxjs';
import { MarketService } from 'src/common/services/market.service';
import { IMarketService, MarketOrder } from 'src/common/types';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: [],
})
export class AppComponent implements OnInit {
	public searchFg = new FormGroup(
		{ search: new FormControl<string>('') },
		{ updateOn: 'submit' }
	);

	public orders$!: Observable<MarketOrder[]>;

	private readonly market: IMarketService;

	constructor(market: MarketService) {
		this.market = market;
	}

	ngOnInit(): void {
		const resourceId: string = '21';
		this.orders$ = this.market.getMarketOrder(resourceId);
	}

	handleSearch() {
		const { search } = this.searchFg.value;

		if (!search) {
			return;
		}
	}

	relativeTime(time: string): string {
		return toRelativeTimeString(new Date(time));
	}
}
