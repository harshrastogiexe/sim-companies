import { Observable } from 'rxjs';
import { BuildingInfo } from './BuildingInfo';

export interface IBuildingService {
	getBuildingInfo(id: string): Observable<BuildingInfo>;
}
