/**
 * Type definitions for Solid Native configuration
 */

export interface SolidNativeConfig {
  app: AppConfig;
  entry: string;
  platforms: PlatformConfig;
  build: BuildConfig;
  dev: DevConfig;
  native: NativeConfig;
  ai: AIConfig;
  typescript?: TypeScriptConfig;
  plugins?: Plugin[];
  experimental?: ExperimentalConfig;
}

export interface AppConfig {
  name: string;
  displayName: string;
  version: string;
  bundleId: string;
  orientation: "portrait" | "landscape" | "auto";
  description?: string;
  author?: string;
}

export interface PlatformConfig {
  ios: IOSConfig;
  android: AndroidConfig;
}

export interface IOSConfig {
  enabled: boolean;
  deploymentTarget: string;
  swiftVersion: string;
  bundleId: string;
  teamId?: string;
  provisioningProfile?: string;
}

export interface AndroidConfig {
  enabled: boolean;
  minSdkVersion: number;
  targetSdkVersion: number;
  compileSdkVersion: number;
  bundleId: string;
  signingConfig?: {
    keystore: string;
    keystorePassword: string;
    keyAlias: string;
    keyPassword: string;
  };
}

export interface BuildConfig {
  outDir: string;
  sourceMaps: boolean;
  minify: boolean;
  treeShaking: boolean;
  assets: string[];
  external: string[];
  target?: "es2020" | "es2021" | "es2022" | "esnext";
}

export interface DevConfig {
  port: number;
  host: string;
  hmr: boolean;
  fastRefresh: boolean;
  devTools: boolean;
  https?: boolean;
}

export interface NativeConfig {
  modulesDir: string;
  ffi: {
    enabled: boolean;
    libsDir: string;
  };
  dependencies: string[];
}

export interface AIConfig {
  enabled: boolean;
  provider: "anthropic" | "openai" | "local";
  model: string;
  apiKey?: string;
  features: {
    componentGeneration: boolean;
    codeCompletion: boolean;
    typeGeneration: boolean;
    codeReview: boolean;
    naturalLanguageCommands: boolean;
    docGeneration: boolean;
  };
  context: {
    includeProjectStructure: boolean;
    includeGitHistory: boolean;
    includeDependencies: boolean;
    maxContextSize: number;
  };
  promptsDir: string;
  schemasDir: string;
}

export interface TypeScriptConfig {
  strict?: boolean;
  jsxImportSource?: string;
  [key: string]: unknown;
}

export interface Plugin {
  name: string;
  setup: (config: SolidNativeConfig) => void | Promise<void>;
}

export interface ExperimentalConfig {
  denoServer?: boolean;
  concurrentRendering?: boolean;
  streamingSSR?: boolean;
  [key: string]: unknown;
}
