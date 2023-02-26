import {
	Component,
	Input,
	OnChanges,
	OnInit,
	SimpleChanges,
} from '@angular/core';
import { combineLatest, map, Observable, Subject } from 'rxjs';
import { BuildingService, ResourceService } from 'src/common/services';
import {
	BuildingInfo,
	IResourceService,
	ResourceBase,
	ResourceInfo,
} from 'src/common/types';
import { IBuildingService } from 'src/common/types/IBuildingService';

@Component({
	selector: 'app-resource-item-info',
	templateUrl: './resource-item-info.component.html',
	styleUrls: ['./resource-info.component.scss'],
})
export class ResourceItemInfoComponent implements OnInit, OnChanges {
	private readonly building: IBuildingService;

	private readonly resourceService: IResourceService;

	private readonly buildingInfo$ = new Subject<BuildingInfo>();

	@Input()
	resource!: ResourceInfo;

	buildingName$?: Observable<string> = this.buildingInfo$.pipe(
		map(({ name }) => name)
	);

	requiredFor$?: Observable<ResourceBase[]>;

	producedFrom$?: Observable<{ amount: number; resource: ResourceInfo }[]>;

	productionBonus = 0;

	buildingLevel = 1;

	constructor(building: BuildingService, resource: ResourceService) {
		this.building = building;
		this.resourceService = resource;
	}

	ngOnInit() {}

	ngOnChanges(changes: SimpleChanges): void {
		if ('resource' in changes) {
			const resource = <ResourceInfo>changes['resource'].currentValue;
			const producedAt = (<ResourceInfo>resource).produced_at;
			this.building
				.getBuildingInfo(producedAt!)
				.subscribe((building) => this.buildingInfo$.next(building));

			this.requiredFor$ = combineLatest(
				resource.needed_for.map((id) =>
					this.resourceService.getResourceInfo(id.toString())
				)
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
}
