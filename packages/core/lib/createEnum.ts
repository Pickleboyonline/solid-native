export function createEnum<T extends string | number>(
  values: readonly T[]
): {
  [key in T]: key;
} {
  // deno-lint-ignore no-explicit-any
  const enumMap: any = {};

  for (const value of values) {
    enumMap[value] = value;
  }

  return enumMap;
}
