import { Observable } from 'rxjs';
import { ResourceBase } from './ResourceBase';
import { ResourceInfo } from './ResourceInfo';

export interface IResourceService {
	/**
	 * fetch all resource basic detail
	 */
	getAllResource(): Observable<ResourceBase[]>;

	/**
	 * fetches the detailed info about the resource
	 * @param id resource id
	 */
	getResourceInfo(id: string): Observable<ResourceInfo>;
}
