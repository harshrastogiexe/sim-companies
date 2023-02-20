// formats the date relative to current time
export function toRelativeTimeString(date: Date): string {
	const minutes = (new Date().getTime() - date.getTime()) / 1000 / 60;
	const formatter = new Intl.RelativeTimeFormat('en', { numeric: 'auto' });
	let value: number, unit: Intl.RelativeTimeFormatUnit;
	switch (true) {
		case minutes < 60:
			[value, unit] = [minutes, 'minutes'];
			break;
		case minutes / 60 < 24:
			[value, unit] = [minutes / 60, 'hours'];
			break;
		default:
			[value, unit] = [minutes / 60 / 24, 'days'];
	}
	return formatter.format(-Math.trunc(value), unit);
}
