import { Component, Input } from '@angular/core';
import { MarketOrder } from 'src/common/types';

const DEFAULT_IMAGE =
	'https://scontent.fdel3-1.fna.fbcdn.net/v/t39.30808-6/302188984_506252578173092_3139193369050851442_n.png?_nc_cat=100&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=khb2iK623jgAX_oyDX6&_nc_ht=scontent.fdel3-1.fna&oh=00_AfDe-mO0C0W2QMRoL9zsgSsilhzfRSkyA44IVwa_8w71Dg&oe=63F7700B';

@Component({
	selector: 'app-market-item',
	templateUrl: './market-item.component.html',
	styles: [
		`
			:host {
				@apply gap-1 bg-dark-0 flex overflow-hidden rounded-xl;
			}
		`,
	],
})
export class MarketItemComponent {
	@Input() item!: MarketOrder;

	get img() {
		return this.item.seller.logo || DEFAULT_IMAGE;
	}
}
