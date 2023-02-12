import { createWriteStream } from "fs";
import { readFile, writeFile } from "fs/promises";
import { Writable } from "stream";
import { cmdArgs } from "./args";
import { ConvertorTypes } from "./ConvertorTypes";
import { Resource } from "./core/resource.types";
import { fetchAllBuildings, fetchAllResources } from "./simcompanies";
import { createFile, isFileExist } from "./utils";

const fetchItem: { [item: string]: () => Promise<any> } = {
  ["resource"]: fetchAllResources,
  ["building"]: fetchAllBuildings,
};

async function main() {
  const args = await cmdArgs;
  const { table } = args;
  if (!table) throw new Error("table name required, use --help");

  const convertor = ConvertorTypes[table];
  if (!convertor) {
    throw new Error(`no sql convertor attached for table ${table}`);
  }

  const { filepath } = args;
  if (!filepath) {
    throw new Error("no filepath provided to read or write data, use --help");
  }

  if (!args.item) {
    throw new Error("--item value required, use --help to know more");
  }

  let data: Resource[] | any;
  if (args.fetch) {
    [data] = await Promise.all([fetchItem[args.item](), createFile(filepath)]);
    writeFile(filepath, JSON.stringify(data));
  } else {
    const blob = await readFile(filepath);
    data = JSON.parse(blob.toString());
  }
  const value = convertor.convert(data);
  let stream: Writable;

  if (args.save && !(await isFileExist(args.save))) {
    await createFile(args.save);
  }

  stream = args.save
    ? createWriteStream(args.save, { start: 0 })
    : process.stdout;

  stream.write(value);
  if (args.save) {
    console.log(`content written to ${args.save} successfully`);
  }
}

console.time("Time Elapsed");
main().then(() => console.timeEnd("Time Elapsed"));
