export interface ISqlConvertor<T> {
  convert(data: T): string;
}
