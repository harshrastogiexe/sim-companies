import {} from "buffer";
import { createReadStream } from "fs";
import * as fs from "fs/promises";

const BaseUri = "https://www.simcompanies.com/api/v3/0/encyclopedia/buildings/";

interface Building {
  db_letter: string;
  image: string;
  images: string[];
  name: string;
  category: string;
  cost: number;
  robotsNeeded: number;
  realmAvailable: boolean;
}

class Encyclopedia {
  constructor() {}

  async getBuilding(id: string): Promise<Building | null> {
    const uri = BaseUri + id;
    const resp = await fetch(uri);

    if (!resp.status.toString().startsWith("2")) return null;

    const body = await resp.json();
    return body;
  }

  async *getAllBuilding() {
    const encyclopedia = new Encyclopedia();
    const promises: Promise<Building | null>[] = [];
    for (let i = 48; i; i++) {
      const building = await encyclopedia.getBuilding(String.fromCharCode(i));
      if (building) yield building;
      switch (i) {
        case 57:
          i = 65;
          break;
        case 90:
          i = 97;
          break;
        case 127:
          return;
      }
    }

    const buildings = await Promise.all(promises);
    for (const building of buildings) {
      if (building) yield building;
    }
  }

  static async BuildingToCSV(filePath: string) {
    const encyclopedia = new Encyclopedia();
    const file = await fs.open(filePath, 0o666);
    const ws = file.createWriteStream();

    for await (const building of encyclopedia.getAllBuilding()) {
      const data =
        [
          building.db_letter,
          building.name,
          building.image,
          building.category,
          building.cost,
          building.robotsNeeded,
          building.realmAvailable ? 1 : 0,
        ].join(", ") + "\n";

      ws.write(data);
    }

    await file.close();
  }

  static CSVToSqlCreateCmd(filename: string, tableName: string) {
    const header = `INSERT INTO ${tableName} VALUES\n`;
    const rs = createReadStream(filename);
    rs.on("data", (data) => {
      const lines = data
        .toString()
        .split("\n")
        .filter((line) => line.length > 0)
        .map((line) => line.split(", "))
        .map((lineWords) => {
          return `(${[
            ...lineWords.slice(0, 4),
            parseInt(lineWords[4]),
            parseFloat(lineWords[5]),
            parseInt(lineWords[6]),
          ]
            .map((value) => (typeof value === "string" ? `'${value}'` : value))
            .join(", ")})`;
        })
        .join(",\n");
      console.log(header + lines);
    });
  }
}

// Encyclopedia.BuildingToCSV("data.csv");
Encyclopedia.CSVToSqlCreateCmd("data.csv", "resource_base");
