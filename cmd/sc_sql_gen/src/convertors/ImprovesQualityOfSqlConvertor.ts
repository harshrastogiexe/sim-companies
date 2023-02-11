import { ISqlConvertor } from "../core/convertors";
import { Resource } from "../Resource";
import { SqlConvertor } from "./SqlConvertor";

export class ImprovesQualityOfSqlConvertor
  extends SqlConvertor
  implements ISqlConvertor<Resource[]>
{
  convert(data: Resource[]): string {
    return (
      "INSERT INTO improves_quality_of VALUES\n" +
      data
        .map((resource) => {
          return resource.improvesQualityOf
            .map((item) => [
              this.convertValue(resource.db_letter),
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
