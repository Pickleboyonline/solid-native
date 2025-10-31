/**
 * Platform utilities
 * Detect and handle platform-specific functionality
 */

export type Platform = "ios" | "android" | "web";

export interface PlatformInfo {
  os: Platform;
  version: string;
  isNative: boolean;
}

/**
 * Get current platform information
 * In a real implementation, this would use FFI to call native platform APIs
 */
export function getPlatform(): PlatformInfo {
  // This is a placeholder - real implementation would use FFI
  const userAgent = globalThis.navigator?.userAgent || "";

  if (userAgent.includes("iPhone") || userAgent.includes("iPad")) {
    return {
      os: "ios",
      version: "14.0",
      isNative: true,
    };
  }

  if (userAgent.includes("Android")) {
    return {
      os: "android",
      version: "12",
      isNative: true,
    };
  }

  return {
    os: "web",
    version: "1.0.0",
    isNative: false,
  };
}

export const Platform = {
  OS: getPlatform().os,
  Version: getPlatform().version,
  isIOS: getPlatform().os === "ios",
  isAndroid: getPlatform().os === "android",
  isWeb: getPlatform().os === "web",

  select<T>(
    specifics: { ios?: T; android?: T; web?: T; native?: T; default?: T },
  ): T | undefined {
    const platform = getPlatform();

    // Check for platform-specific value first
    if (platform.os in specifics) {
      return specifics[platform.os as keyof typeof specifics];
    }

    // Check for native fallback
    if (platform.isNative && "native" in specifics) {
      return specifics.native;
    }

    // Return default
    return specifics.default;
  },
};
