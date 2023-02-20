import { Pipe, PipeTransform } from '@angular/core';
import { toRelativeTimeString } from '../utils/relativeTimeFormatter';

@Pipe({
	name: 'relativeTime',
})
export class ToRelativeTimeString implements PipeTransform {
	transform(value: Date | string): string {
		return toRelativeTimeString(
			typeof value === 'string' ? new Date(value) : value
		);
	}
}
