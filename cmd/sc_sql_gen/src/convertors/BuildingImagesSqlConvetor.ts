import { ISqlConvertor } from "../core/convertors";
import { Building } from "../core/resource.types";
import { SqlConvertor } from "./SqlConvertor";

export class BuildingImagesSqlConvertor
  extends SqlConvertor
  implements ISqlConvertor<Building[]>
{
  convert(data: Building[]): string {
    return (
      `INSERT INTO [building_level_images]\nVALUES\n` +
      data
        .map((building) => ({
          images: building.images,
          buildingID: building.db_letter,
        }))
        .map(({ images, buildingID }) =>
          images
            .map((image, index) =>
              [buildingID, index + 1, image]
                .map((value) => this.convertValue(value))
                .join(", ")
            )
            .map((value) => `\t(${value})`)
            .join(",\n")
        )
        .join(",\n")
    );
  }
}
