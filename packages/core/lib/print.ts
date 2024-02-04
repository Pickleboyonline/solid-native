/**
 * Prints to Swift Console
 *
 * TODO: Remove and put with native module (like a logger)
 */
// deno-lint-ignore no-explicit-any
export const print = (globalThis as any)._print as (str: string) => void;
