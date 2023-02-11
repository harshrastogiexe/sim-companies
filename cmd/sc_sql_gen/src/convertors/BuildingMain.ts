import { ISqlConvertor } from "../core/convertors";
import { Building } from "../Resource";
import { SqlConvertor } from "./SqlConvertor";

const TABLE_MAIN = "building_main";

export class BuildingMainSqlConvertor
  extends SqlConvertor
  implements ISqlConvertor<Building[]>
{
  convert(data: Building[]): string {
    const values = data
      .map((building) => {
        const items = this.select(
          building,
          "db_letter",
          "costUnits",
          "hours",
          "wages"
        );
        return items.join(", ");
      })
      .map((value) => `(${value})`)
      .join(",\n\t");

    return `INSERT INTO ${TABLE_MAIN}
VALUES
    ${values}`;
  }
}
