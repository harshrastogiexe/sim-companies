import { Injectable } from '@angular/core';

@Injectable()
export class DropdownService {
	isOpen = false;

	toggle(open?: boolean) {
		this.isOpen = open || !this.isOpen;
	}
}
