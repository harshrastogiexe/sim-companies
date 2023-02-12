import yargs from "yargs";

export const cmdArgs = yargs
  .options("filepath", {
    alias: "f",
    string: true,
    desc: "file from which values will be read",
  })
  .options("table", {
    alias: "t",
    string: true,
    desc: "table name for which query will be generated",
  })
  .options("fetch", {
    desc: "fetch data from api",
    default: false,
    boolean: true,
  })
  .options("save", {
    alias: "s",
    string: true,
    desc: "filepath to which content will be saved",
  })
  .options("item", { choices: ["resource", "building"] }).argv;
