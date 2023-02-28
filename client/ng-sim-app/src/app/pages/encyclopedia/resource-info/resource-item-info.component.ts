import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { Observable, Subject, combineLatest, map } from 'rxjs';
import { BuildingService, ResourceService } from 'src/common/services';
import {
	BuildingInfo,
	IBuildingService,
	IResourceService,
	ResourceBase,
	ResourceInfo,
} from 'src/common/types';

@Component({
	selector: 'app-resource-item-info',
	templateUrl: './resource-item-info.component.html',
	styleUrls: ['./resource-info.component.scss'],
})
export class ResourceItemInfoComponent implements OnChanges {
	private readonly building: IBuildingService;

	private readonly resourceService: IResourceService;

	private readonly buildingInfo$ = new Subject<BuildingInfo>();

	@Input()
	resource!: ResourceInfo;

	public buildingName$?: Observable<string> = this.buildingInfo$.pipe(
		map(({ name }) => name)
	);

	public requiredFor$?: Observable<ResourceBase[]>;

	public producedFrom$?: Observable<
		{ amount: number; resource: ResourceInfo }[]
	>;

	public productionBonus = 0;

	public buildingLevel = 1;

	public admin = 47.65;

	public get adminCost() {
		return (this.resource.base_salary * this.admin) / 100;
	}

	public get labourCost() {
		return this.resource.base_salary + this.adminCost;
	}

	get production() {
		return (
			this.resource.produced_an_hour *
			this.buildingLevel *
			(1 + this.productionBonus / 100)
		);
	}

	public constructor(building: BuildingService, resource: ResourceService) {
		this.building = building;
		this.resourceService = resource;
	}

	public ngOnChanges(changes: SimpleChanges): void {
		if ('resource' in changes) {
			this.onResourceChange(<ResourceInfo>changes['resource'].currentValue);
		}
	}

	private onResourceChange(resource: ResourceInfo) {
		const producedAt = (<ResourceInfo>resource).produced_at;
		const getResourceInfo = (id: string) =>
			this.resourceService.getResourceInfo(id);

		this.building
			.getBuildingInfo(producedAt!)
			.subscribe((building) => this.buildingInfo$.next(building));

		this.requiredFor$ = combineLatest(
			resource.needed_for.map((id) => getResourceInfo(id.toString()))
		);

		this.producedFrom$ = combineLatest(
			resource.producedFrom.map((item) =>
				this.resourceService
					.getResourceInfo(item.ResourceId.toString())
					.pipe(map((resource) => ({ amount: item.Amount, resource })))
			)
		);
	}
}
