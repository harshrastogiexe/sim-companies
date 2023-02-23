import { Directive, HostListener, Input } from '@angular/core';
import { InputFilterListService } from './input-filter-list.service';

@Directive({ selector: '[app-input-filter-list-item]' })
export class InputFilterListItemDirective<T> {
	@Input()
	value!: T;

	private readonly inputFilterService: InputFilterListService<T>;

	constructor(inputFilterService: InputFilterListService) {
		this.inputFilterService = inputFilterService;
	}

	@HostListener('click')
	onClick() {
		if (!this.value) throw new Error('value required');
		this.inputFilterService.select(this.value);
		this.inputFilterService.toggle();
	}
}
