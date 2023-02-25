import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { isNotUndefined } from '@utils/*';
import { Observable, of, tap } from 'rxjs';
import { environment } from 'src/environment/environment';
import { IResourceService, ResourceInfo } from '../types';
import { ResourceBase } from '../types/ResourceBase';

type CacheResource =
	| { base: ResourceBase; info?: ResourceInfo }
	| { base?: ResourceBase; info: ResourceInfo };

@Injectable({ providedIn: 'root' })
export class ResourceService implements IResourceService {
	public readonly http: HttpClient;

	public constructor(http: HttpClient) {
		this.http = http;
	}

	/**
	 * return list of all resources of present in cache data, if cache size is zero fetched data from api and
	 * store data to cache
	 * @returns list of all resource base details
	 */
	public getAllResource() {
		if (ResourceService.resources.size !== 0) {
			const resources = [...ResourceService.resources.values()]
				.map((r) => r.base)
				.filter(isNotUndefined);
			return of(resources);
		}

		const uri = environment.server.basePath + '/encyclopedia/resource';
		const save = (resources: ResourceBase[]) =>
			resources.forEach((resource) => ResourceService.save({ base: resource }));

		return this.http
			.get<ResourceBase[]>(uri)
			.pipe(tap((resources) => save(resources)));
	}

	/**
	 * checks whether item is in cache and returns it, else fetch form api then cache it
	 * @param id resource db_letter
	 * @returns resource more detailed info
	 */
	public getResourceInfo(id: string): Observable<ResourceInfo> {
		const cached = ResourceService.resources.get(id);
		if (cached?.info) {
			return of(cached.info);
		}
		const uri = environment.server.basePath + '/encyclopedia/resource/' + id;
		const save = (resource: ResourceInfo) =>
			ResourceService.save({ info: resource });

		return this.http
			.get<ResourceInfo>(uri)
			.pipe(tap((resource) => save(resource)));
	}

	/**
	 * Cached resource data
	 */
	private static resources = new Map<string, CacheResource>();

	/**
	 * save the data to store
	 * @param data base and info of resource
	 */
	private static save(data: CacheResource) {
		const item =
			data.base?.db_letter ||
			data.info?.id ||
			new Error('expected base or info details, got undefined');

		if (item instanceof Error) throw item;

		const resource = this.resources.get(item.toString());
		this.resources.set(item.toString(), { ...resource, ...data });
	}

	private static get(id: string) {
		return this.resources.get(id);
	}
}
