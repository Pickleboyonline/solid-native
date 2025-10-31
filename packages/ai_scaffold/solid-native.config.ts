import type { SolidNativeConfig } from "./src/types/config.ts";

/**
 * Solid Native Configuration
 *
 * This is the main configuration file for your Solid Native application.
 * It defines build settings, platform configurations, and AI-assisted development options.
 */
export default {
  // App metadata
  app: {
    name: "SolidNativeApp",
    displayName: "Solid Native App",
    version: "1.0.0",
    bundleId: "com.solidnative.app",
    orientation: "portrait" as const,
  },

  // Entry point for your application
  entry: "./src/app.tsx",

  // Platform-specific configurations
  platforms: {
    ios: {
      enabled: true,
      deploymentTarget: "14.0",
      swiftVersion: "5.0",
      bundleId: "com.solidnative.app.ios",
    },
    android: {
      enabled: true,
      minSdkVersion: 24,
      targetSdkVersion: 34,
      compileSdkVersion: 34,
      bundleId: "com.solidnative.app.android",
    },
  },

  // Build configuration
  build: {
    // Output directory for bundled code
    outDir: "./dist",

    // Source maps for debugging
    sourceMaps: true,

    // Minify production builds
    minify: true,

    // Tree shaking
    treeShaking: true,

    // Asset handling
    assets: ["./assets/**/*"],

    // External dependencies (not bundled)
    external: [],
  },

  // Development server configuration
  dev: {
    port: 8081,
    host: "localhost",

    // Hot module replacement
    hmr: true,

    // Fast refresh for Solid components
    fastRefresh: true,

    // Open dev tools on start
    devTools: true,
  },

  // Native module configuration
  native: {
    // Path to native modules directory
    modulesDir: "./native",

    // FFI bindings
    ffi: {
      enabled: true,
      libsDir: "./native/libs",
    },

    // Native dependencies
    dependencies: [],
  },

  // AI-assisted development configuration
  ai: {
    // Enable AI features
    enabled: true,

    // AI provider (anthropic, openai, local)
    provider: "anthropic" as const,

    // Model to use
    model: "claude-sonnet-4-5-20250929",

    // Features
    features: {
      // AI-powered component generation
      componentGeneration: true,

      // Intelligent code completion
      codeCompletion: true,

      // Automatic type inference and generation
      typeGeneration: true,

      // Code review and suggestions
      codeReview: true,

      // Natural language commands
      naturalLanguageCommands: true,

      // Documentation generation
      docGeneration: true,
    },

    // Context awareness
    context: {
      // Include project structure in context
      includeProjectStructure: true,

      // Include recent changes
      includeGitHistory: true,

      // Include dependencies
      includeDependencies: true,

      // Maximum context size (in tokens)
      maxContextSize: 50000,
    },

    // Custom prompts directory
    promptsDir: "./src/ai/prompts",

    // Component schemas directory
    schemasDir: "./src/ai/schemas",
  },

  // TypeScript configuration overrides
  typescript: {
    strict: true,
    jsxImportSource: "solid-js",
  },

  // Plugins for extending functionality
  plugins: [],

  // Experimental features
  experimental: {
    // Use Deno's native HTTP server
    denoServer: true,

    // Concurrent rendering
    concurrentRendering: false,

    // Streaming SSR
    streamingSSR: false,
  },
} satisfies SolidNativeConfig;
