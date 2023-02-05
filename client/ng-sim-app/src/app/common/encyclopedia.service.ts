import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Resource, ResourceBase } from './types';

@Injectable({ providedIn: 'root' })
export class EncyclopediaService {
	private readonly http: HttpClient;

	constructor(http: HttpClient) {
		this.http = http;
	}

	public fetchResources(): Observable<ResourceBase[]> {
		console.count('fetch resource');
		return this.http.get<ResourceBase[]>('http://localhost:8080/resource/');
	}

	public fetchResourceById(id: string): Observable<Resource> {
		const uri = 'http://localhost:8080/resource/' + id;
		return this.http.get<Resource>(uri);
	}
}
