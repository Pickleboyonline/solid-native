/**
 * Storage service - persistent data storage
 * Abstracts platform-specific storage implementations
 */

export interface StorageAdapter {
  getItem(key: string): Promise<string | null>;
  setItem(key: string, value: string): Promise<void>;
  removeItem(key: string): Promise<void>;
  clear(): Promise<void>;
  getAllKeys(): Promise<string[]>;
}

/**
 * Generic storage service that works with any storage adapter
 */
export class StorageService {
  constructor(private adapter: StorageAdapter) {}

  async get<T>(key: string): Promise<T | null> {
    const value = await this.adapter.getItem(key);
    if (value === null) return null;

    try {
      return JSON.parse(value) as T;
    } catch {
      return value as T;
    }
  }

  async set<T>(key: string, value: T): Promise<void> {
    const serialized = typeof value === "string"
      ? value
      : JSON.stringify(value);

    await this.adapter.setItem(key, serialized);
  }

  async remove(key: string): Promise<void> {
    await this.adapter.removeItem(key);
  }

  async clear(): Promise<void> {
    await this.adapter.clear();
  }

  async getAllKeys(): Promise<string[]> {
    return this.adapter.getAllKeys();
  }

  async has(key: string): Promise<boolean> {
    const value = await this.adapter.getItem(key);
    return value !== null;
  }
}

// Example platform-specific adapter (would be in @solid-native/runtime)
// export class NativeStorageAdapter implements StorageAdapter {
//   async getItem(key: string): Promise<string | null> {
//     return await Deno.ffi.nativeStorage.get(key);
//   }
//   // ... other methods
// }
