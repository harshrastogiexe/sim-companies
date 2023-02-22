import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { IResourceService } from '../types';
import { ResourceBase } from '../types/ResourceBase';

@Injectable({ providedIn: 'root' })
export class ResourceService implements IResourceService {
	public readonly http: HttpClient;

	public constructor(http: HttpClient) {
		this.http = http;
	}

	public getAllResource() {
		const uri = 'http://localhost:8080/encyclopedia/resource';
		return this.http.get<ResourceBase[]>(uri);
	}
}
