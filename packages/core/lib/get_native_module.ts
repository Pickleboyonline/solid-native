/**
 * Returns module from Solid Native module manager
 */
export function getNativeModule<ModuleType>(moduleName: string): ModuleType {
  // deno-lint-ignore no-explicit-any
  const mod = (globalThis as any)._getNativeModule(moduleName);

  return mod;
}
