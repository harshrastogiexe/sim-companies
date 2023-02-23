import {
	Component,
	ElementRef,
	EventEmitter,
	HostListener,
	Input,
	OnChanges,
	Output,
	SimpleChanges,
	ViewChild,
} from '@angular/core';
import { FormControl } from '@angular/forms';
import { DropdownService } from './dropdown.service';

@Component({
	selector: 'app-dropdown',
	templateUrl: './dropdown.component.html',
	providers: [DropdownService],
})
export class DropdownComponent<T> implements OnChanges {
	@Input()
	public items: T[] = [];

	@Input('control')
	public searchFormControl?: FormControl<any>;

	@Input('filter')
	public filterCallback?: (value: T, text: string) => boolean;

	@Input('label')
	labelCb!: (item: T) => string;

	@Output('select')
	selectEventEmitter = new EventEmitter<T>();

	public selected!: T;

	public filtered = this.items;

	@ViewChild('dropdownWrapper')
	private elementRef!: ElementRef;

	public readonly dropdown: DropdownService;

	constructor(dropdown: DropdownService) {
		this.dropdown = dropdown;
	}

	ngOnChanges(changes: SimpleChanges): void {
		console.log(changes);
		if (changes['items']) {
			this.filtered = this.items;
		}
	}

	public onValueChangeHandler(e: Event) {
		const { value } = <HTMLInputElement>e.target;

		if (!value) {
			this.filtered = this.items;
		}

		this.filtered = this.items.filter((item) =>
			this.filterCallback!(item, value)
		);
	}

	public onItemSelectHandler(item: T) {
		this.selected = item;
		this.selectEventEmitter.next(item);
		this.dropdown.toggle(false);
	}

	@HostListener('window:click', ['$event'])
	public onWindowClick(e: PointerEvent) {
		if (this.elementRef.nativeElement.contains(e.target)) return;
		this.dropdown.toggle(false);
	}
}
