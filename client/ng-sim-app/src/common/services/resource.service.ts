import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environment/environment';
import { IResourceService } from '../types';
import { ResourceBase } from '../types/ResourceBase';

@Injectable({ providedIn: 'root' })
export class ResourceService implements IResourceService {
	public readonly http: HttpClient;

	public constructor(http: HttpClient) {
		this.http = http;
	}

	public getAllResource() {
		const uri = environment.server.basePath + '/encyclopedia/resource';
		return this.http.get<ResourceBase[]>(uri);
	}
}
