import {
	Component,
	EventEmitter,
	HostListener,
	Input,
	Output,
} from '@angular/core';
import { ResourceBase } from 'src/app/common/types';

@Component({
	selector: 'app-resource-item',
	templateUrl: './resource-item.component.html',
	styleUrls: ['./resource-item.component.scss'],
})
export class ResourceItemComponent {
	@Input() public resource: ResourceBase | undefined;

	@Output('select') public selectEvent = new EventEmitter<ResourceBase>();

	@HostListener('click') onHostClickHandler() {
		this.selectEvent.emit(this.resource);
	}
}
