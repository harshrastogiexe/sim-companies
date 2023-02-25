import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environment/environment';
import { BuildingInfo } from '../types';
import { IBuildingService } from '../types/IBuildingService';

@Injectable({ providedIn: 'root' })
export class BuildingService implements IBuildingService {
	private readonly http: HttpClient;

	constructor(http: HttpClient) {
		this.http = http;
	}

	public getBuildingInfo(id: string) {
		const uri = environment.server.basePath + '/encyclopedia/building/' + id;
		return this.http.get<BuildingInfo>(uri);
	}
}
