import { Directive, HostListener } from '@angular/core';
import { InputFilterListService } from './input-filter-list.service';

@Directive({
	selector: 'input[app-input-text-filter-search]',
})
export class InputTextFilterSearchDirective {
	private readonly inputListService: InputFilterListService;

	public constructor(inputListService: InputFilterListService) {
		this.inputListService = inputListService;
	}

	@HostListener('focus')
	private onFocus() {
		this.inputListService.toggle();
	}
}
