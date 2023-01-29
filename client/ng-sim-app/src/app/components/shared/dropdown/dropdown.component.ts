import {
	Component,
	ContentChild,
	ElementRef,
	HostListener,
	Input,
	ViewChild,
} from '@angular/core';
import { DropdownItemComponent } from '../dropdown-item/dropdown-item.component';

@Component({
	selector: 'app-dropdown',
	templateUrl: './dropdown.component.html',
	styles: [],
})
export class DropdownComponent {
	/**
	 * Text appear on button of component
	 */
	@Input('label') buttonText: string = '<NO BUTTON TITLE>';

	visible: boolean = false;

	@ContentChild(DropdownItemComponent)
	dropdownItem!: DropdownItemComponent;

	@ViewChild('dropdown')
	dropdownElmRef?: ElementRef<HTMLElement>;

	@HostListener('window:click', ['$event'])
	onWindowClick(event: Event) {
		console.log(this.dropdownItem);
		if (this.visible == false) return;

		if (!this.dropdownElmRef?.nativeElement) {
			throw new Error('dropdown reference not found');
		}

		this.visible = false;
	}
}
