import { Component, Input } from '@angular/core';

@Component({
	selector: 'app-resource-info-field',
	templateUrl: './resource-info-field.component.html',
	styleUrls: ['./resource-info-field.component.scss'],
})
export class ResourceInfoFieldComponent<T = any> {
	@Input()
	value!: T;

	@Input()
	label!: string;
}
