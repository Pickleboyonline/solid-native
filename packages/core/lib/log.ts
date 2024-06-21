/**
 * Prints to Swift Console
 *
 * TODO: Remove and put with native module (like a logger)
 */
// deno-lint-ignore no-explicit-any
export const log = (globalThis as any).log as (str: string) => void;
