import { ISqlConvertor } from "../core/convertors";
import { Resource } from "../core/resource.types";
import { SqlConvertor } from "./SqlConvertor";

export class ProducedFromSqlConvertors
  extends SqlConvertor
  implements ISqlConvertor<Resource[]>
{
  convert(data: Resource[]): string {
    const header = `INSERT INTO produced_froms\nVALUES\n\t`;
    const values = data
      .filter((item) => item !== null && item.producedFrom !== null)
      .map(({ producedFrom, db_letter }) =>
        producedFrom.map((from) => ({ id: db_letter, from }))
      )
      .flat()
      .map((item) =>
        [
          this.convertValue(item.id),
          this.convertValue(item.from.resource.db_letter),
          this.convertValue(item.from.amount),
        ].join(", ")
      )
      .map((value) => `(${value})`)
      .join(",\n\t");
    return header + values;
  }
}
