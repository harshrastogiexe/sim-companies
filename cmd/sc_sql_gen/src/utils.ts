import { access, constants, mkdir, writeFile } from "fs/promises";
import { dirname } from "path";
export async function isFileExist(path: string) {
  try {
    await access(path, constants.F_OK);
    return true;
  } catch {
    return false;
  }
}

export async function createFile(filepath: string, data?: string | Buffer) {
  await mkdir(dirname(filepath), { recursive: true });
  await writeFile(filepath, data || "");
}
