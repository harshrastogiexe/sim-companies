export class SqlConvertor {
  convertValue(value: string | number | null | undefined | boolean): string {
    if (value === null || value === undefined) return "NULL";
    switch (typeof value) {
      case "string":
        return `'${value}'`;
      case "boolean":
        return (value ? 1 : 0).toString();
      default:
        return value.toString();
    }
  }

  select<T>(value: T, ...keys: (keyof T)[]): string[] {
    const values: string[] = [];
    keys.forEach((key) => values.push(this.convertValue(<any>value[key])));
    return values;
  }
}
