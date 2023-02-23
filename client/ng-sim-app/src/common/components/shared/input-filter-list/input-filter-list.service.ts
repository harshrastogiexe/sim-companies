import { Injectable } from '@angular/core';
import { BehaviorSubject, Subject } from 'rxjs';

@Injectable()
export class InputFilterListService<T = any> {
	private hiddenSub$ = new BehaviorSubject<boolean>(true);

	hidden$ = this.hiddenSub$.asObservable();

	private selectedSub$?: Subject<T>;

	get hidden() {
		return this.hiddenSub$.value;
	}

	public toggle() {
		const { value } = this.hiddenSub$;
		this.hiddenSub$.next(!value);
	}

	public show() {
		this.hiddenSub$.next(false);
	}

	public hide() {
		this.hiddenSub$.next(true);
	}

	public useSubject(sub: Subject<T>) {
		this.selectedSub$ = sub;
	}

	public select(value: T) {
		this.selectedSub$?.next(value);
	}
}
