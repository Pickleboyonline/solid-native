/**
 * Helper function to grab and type anything on the global object
 */
export function getNativeModule<ModuleType>(moduleName: string): ModuleType {
  // deno-lint-ignore no-explicit-any
  const mod = (globalThis as any)[moduleName];

  return mod;
}
