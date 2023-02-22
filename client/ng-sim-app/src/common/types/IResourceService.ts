import { Observable } from 'rxjs';
import { ResourceBase } from './ResourceBase';

export interface IResourceService {
	/**
	 * fetch all resource basic detail
	 */
	getAllResource(): Observable<ResourceBase[]>;
}
