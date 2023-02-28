import { Component, EventEmitter, Input, Output } from '@angular/core';

const fractionRegex = /^(0|[1-9]\d*)(\.\d+)?$/;

@Component({
	selector: 'app-resource-info-modifier',
	templateUrl: './resource-item-info-modifier.component.html',
	styleUrls: ['./resource-item-info-modifies.scss'],
})
export class ResourceItemInfoModifierComponent {
	@Input()
	label: string = '<NO LABEL>';

	@Input()
	value: number = 0;

	@Output('change')
	valueChange = new EventEmitter<number>();

	onChange(e: Event) {
		const { value } = <HTMLInputElement>e.target;
		const floatValue = parseFloat(value);

		this.value = Number.isNaN(floatValue) ? 0 : floatValue;
		this.valueChange.emit(this.value);
	}
}
