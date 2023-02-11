import { ISqlConvertor } from "../core/convertors";
import { Resource } from "../Resource";
import { SqlConvertor } from "./SqlConvertor";

export class ResourceBaseSqlConvertor
  extends SqlConvertor
  implements ISqlConvertor<Resource[]>
{
  convert(data: Resource[]): string {
    return (
      "INSERT INTO resource_base VALUES\n" +
      data
        .map((resource) =>
          this.select(
            resource,
            "db_letter",
            "name",
            "image",
            "transportation",
            "retailable",
            "research",
            "exchangeTradable",
            "realmAvailable"
          ).join(", ")
        )
        .map((value) => `(${value})`)
        .join(",\n")
    );
  }
}
