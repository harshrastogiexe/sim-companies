import { ResourceBaseSqlConvertor } from "./convertors";
import { BuildingImagesSqlConvertor } from "./convertors/BuildingImagesSqlConvetor";
import { BuildingMainSqlConvertor } from "./convertors/BuildingMain";
import { DoesAction } from "./convertors/DoesProduce";
import { ImprovesQualityOfSqlConvertor } from "./convertors/ImprovesQualityOfSqlConvertor";
import { ISqlConvertor } from "./core/convertors";

export const ConvertorTypes: { [key: string]: ISqlConvertor<any> | undefined } =
  {
    ["resource_base"]: new ResourceBaseSqlConvertor(),
    ["improves_quality_of"]: new ImprovesQualityOfSqlConvertor(),
    ["does_sell"]: new DoesAction("sell"),
    ["does_produce"]: new DoesAction("produce"),
    ["building_main"]: new BuildingMainSqlConvertor(),
    ["images"]: new BuildingImagesSqlConvertor(),
  };
