import { readFile } from "fs/promises";
import { dirname, join } from "path/posix";
import { Resource } from "../Resource";

export function* getResourcesId(): Generator<string, void, unknown> {
  for (let i = 1; i <= 145; i++) {
    if (i > 35 && i < 40) continue;
    yield i.toString();
  }
}

export async function* getBuildingId() {
  console.log();
  const path = join(dirname(__filename), "buildingIDS.txt");
  const blob = await readFile(path);
  const ids = blob.toString().split("\n");

  for (const id of ids) {
    if (!id) continue;
    yield id;
  }
}

/**
 * fetch resource from simcompnaines api, if not found returns null
 * @param id db_letter
 * @returns
 */
export async function fetchResource(id: string): Promise<Resource | null> {
  const uri = `https://www.simcompanies.com/api/v4/en/0/encyclopedia/resources/2/${id}`;
  const response = await fetch(uri);
  if (!response.status.toString().startsWith("2")) return null;

  return response.json();
}

export async function fetchAllResources() {
  const promises: Promise<Resource>[] = [];
  for (const resourceID of getResourcesId()) {
    promises.push(<Promise<Resource>>fetchResource(resourceID));
  }
  return Promise.all(promises);
}

export async function fetchBuilding(id: string): Promise<any> {
  const uri = `https://www.simcompanies.com/api/v3/0/encyclopedia/buildings/${id}`;

  const response = await fetch(uri);
  if (!response.status.toString().startsWith("2")) return null;

  return response.json();
}

export async function fetchAllBuildings() {
  const promises: Promise<any>[] = [];
  for await (const buildingID of getBuildingId()) {
    promises.push(<Promise<any>>fetchBuilding(buildingID));
  }
  return Promise.all(promises);
}
