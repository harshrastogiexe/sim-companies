import {
	Component,
	ElementRef,
	HostListener,
	Input,
	OnInit,
} from '@angular/core';
import { Subject } from 'rxjs';
import { InputFilterListService } from './input-filter-list.service';

@Component({
	selector: 'app-input-filter-list',
	templateUrl: './input-filter-list.component.html',
	styleUrls: ['./input-filter-list.component.scss'],
	providers: [InputFilterListService],
})
export class InputFilterListComponent implements OnInit {
	public readonly inputFilterService: InputFilterListService;

	private readonly elementRef: ElementRef<HTMLElement>;

	@Input()
	public subject?: Subject<any>;

	constructor(
		inputFilterService: InputFilterListService,
		elementRef: ElementRef
	) {
		this.inputFilterService = inputFilterService;
		this.elementRef = elementRef;
	}

	ngOnInit(): void {
		if (this.subject) this.inputFilterService.useSubject(this.subject);
	}

	@HostListener('window:click', ['$event'])
	onClick(e: PointerEvent) {
		const isInsideElm = this.elementRef.nativeElement.contains(
			e.target as Node
		);

		if (isInsideElm || this.inputFilterService.hidden) return;
		this.inputFilterService.toggle();
	}
}
