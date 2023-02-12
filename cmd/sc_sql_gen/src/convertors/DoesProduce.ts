import { ISqlConvertor } from "../core/convertors";
import { Building } from "../core/resource.types";
import { SqlConvertor } from "./SqlConvertor";

export class DoesAction
  extends SqlConvertor
  implements ISqlConvertor<Building[]>
{
  constructor(private readonly type: "sell" | "produce") {
    super();
  }

  convert(data: Building[]): string {
    return (
      `INSERT INTO does_${this.type} VALUES\n` +
      data
        .map((building) => {
          building.doesProduce;
          const value =
            this.type == "sell"
              ? ("doesProduce" as const)
              : ("doesSell" as const);

          return building[value]
            .map((item) => [
              this.convertValue(building.db_letter),
              ...this.select(item, "db_letter"),
            ])
            .map((items) => items.join(", "))
            .map((value) => `(${value})`)
            .join(",\n");
        })
        .filter((value) => value)
        .join(",\n")
    );
  }
}
