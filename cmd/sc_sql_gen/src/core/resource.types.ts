export interface Resource {
  name: string;
  image: string;
  db_letter: number;
  transportation: number;
  retailable: boolean | null;
  research: boolean;
  exchangeTradable: boolean;
  realmAvailable: boolean;
  producedFrom: ProducedFrom[];
  soldAt: BuildingBase | null;
  soldAtRestaurant: BuildingBase | null;
  producedAt: BuildingBase;
  neededFor: ResourceBase[];
  transportNeeded: number;
  producedAnHour: number;
  baseSalary: number;
  averageRetailPrice: number | null;
  marketSaturation: number | null;
  marketSaturationLabel: Label | null;
  retailModeling: null | string;
  storeBaseSalary: number | null;
  retailData: RetailDatum[];
  improvesQualityOf: ResourceBase[];
  season?: number[];
}

export interface ResourceBase {
  name: string;
  image: string;
  db_letter: number;
  transportation: number;
  retailable: boolean | null;
  research: boolean;
  exchangeTradable: boolean;
  realmAvailable: boolean;
  season?: number[];
}

export enum Label {
  AlmostNone = "almost none",
  Extreme = "extreme",
  High = "high",
  Low = "low",
  Normal = "normal",
  VeryHigh = "very high",
  VeryLow = "very low",
}

export interface ProducedFrom {
  resource: ResourceBase;
  amount: number;
}

export interface RetailDatum {
  averagePrice?: number;
  amountSoldRestaurant?: number;
  demand: number;
  date: Date;
  label: Label;
  amountSold?: number;
}

export interface Building {
  db_letter: string;
  image: string;
  images: string[];
  name: string;
  category: Category;
  cost: number;
  robotsNeeded: number;
  realmAvailable: boolean;
  doesProduce: BuildingBase[];
  doesSell: BuildingBase[];
  costUnits: number;
  hours: number;
  wages: number;
}

export enum Category {
  Bank = "bank",
  Production = "production",
  Recreation = "recreation",
  Research = "research",
  Sales = "sales",
}

export interface BuildingBase {
  name: string;
  image: string;
  db_letter: number;
  transportation: number;
  retailable: boolean | null;
  research: boolean;
  exchangeTradable: boolean;
  realmAvailable: boolean;
  season?: number[];
}
