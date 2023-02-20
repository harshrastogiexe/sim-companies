import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { toRelativeTimeString } from '@utils/*';
import { of } from 'rxjs';
import { market } from './market-data';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: [],
})
export class AppComponent {
	searchFg = new FormGroup(
		{ search: new FormControl<string>('') },
		{ updateOn: 'submit' }
	);

	market = of(
		market.sort((a, b) => (new Date(a.posted) < new Date(b.posted) ? 1 : -1))
	);

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
