# Solid Native Scaffold

> AI-powered scaffolding and development tools for Solid Native - a cross-platform mobile framework built on Deno and SolidJS

## Overview

Solid Native Scaffold is the intelligent development toolkit for building cross-platform mobile apps with:

- **Deno Runtime**: Modern, secure JavaScript/TypeScript runtime with built-in tooling
- **SolidJS**: Fine-grained reactive UI library with superior performance
- **AI-First Development**: LLM-powered code generation, scaffolding, and assistance
- **Type Safety**: Full TypeScript support with strict type checking
- **Native Performance**: Direct FFI bindings for platform features

## Key Features

### 🤖 AI-Powered Development

- **Natural Language Scaffolding**: Generate components from descriptions
- **Intelligent Code Review**: AI-assisted code quality checks
- **Type Generation**: Automatic TypeScript type inference
- **Documentation**: Auto-generated docs from code

### ⚡ Modern Developer Experience

- **Fast Refresh**: Instant updates during development
- **Hot Module Replacement**: Preserve state during updates
- **Built-in CLI**: No separate build tools needed
- **Deno Tasks**: Simple task runner integration

### 🎯 Best Practices Built-in

- Convention over configuration
- Component-driven architecture
- Service-oriented design pattern
- Reactive state management
- Platform abstraction layers

## Project Structure

```
solid-native-app/
├── deno.json                    # Deno configuration and tasks
├── solid-native.config.ts       # Framework configuration
├── src/
│   ├── app.tsx                  # Application entry point
│   ├── components/              # Reusable UI components
│   │   ├── Button.tsx
│   │   └── index.ts
│   ├── services/                # Business logic and APIs
│   │   ├── api.ts              # HTTP client
│   │   ├── storage.ts          # Data persistence
│   │   └── index.ts
│   ├── stores/                  # State management
│   │   ├── theme.ts            # Theme store
│   │   └── index.ts
│   ├── utils/                   # Utility functions
│   │   ├── platform.ts         # Platform detection
│   │   ├── logger.ts           # Logging utility
│   │   └── index.ts
│   ├── ai/                      # AI tooling
│   │   ├── client.ts           # LLM client
│   │   ├── generator.ts        # Code generator
│   │   ├── schemas/            # Generation schemas
│   │   └── prompts/            # System prompts
│   ├── cli/                     # CLI commands
│   │   ├── dev.ts              # Development server
│   │   ├── build.ts            # Production build
│   │   ├── scaffold.ts         # AI scaffolding
│   │   ├── ai.ts               # AI assistant
│   │   └── doctor.ts           # Environment check
│   └── types/                   # TypeScript types
│       └── config.ts
└── mod.ts                       # Main export
```

## Getting Started

### Prerequisites

- [Deno](https://deno.land) 1.40 or higher
- iOS development: Xcode 14+
- Android development: Android Studio & SDK

### Installation

```bash
# Clone or create a new project
deno init my-app
cd my-app

# Copy scaffold structure (when published)
# deno install jsr:@solid-native/scaffold
```

### Configuration

Configure your app in `solid-native.config.ts`:

```typescript
export default {
  app: {
    name: "MyApp",
    displayName: "My Awesome App",
    version: "1.0.0",
    bundleId: "com.mycompany.myapp",
  },

  ai: {
    enabled: true,
    provider: "anthropic",
    model: "claude-sonnet-4-5-20250929",
  },

  platforms: {
    ios: { enabled: true },
    android: { enabled: true },
  },
} satisfies SolidNativeConfig;
```

### Development

```bash
# Start development server with hot reload
deno task dev

# Generate component with AI
deno task scaffold "a button component with loading state and variants"

# Run AI assistant
deno task ai generate component "user profile card with avatar and bio"

# Review code
deno task ai review src/components/Button.tsx

# Check environment
deno task doctor

# Build for production
deno task build

# Start built app
deno task start
```

## AI-Powered Features

### Component Generation

Generate production-ready components from natural language:

```bash
deno task scaffold "a card component with image, title, description and action button"
```

This generates:
- Typed component with props interface
- Styled JSX with mobile-first design
- Proper imports and exports
- Documentation comments

### Code Review

Get AI feedback on your code:

```bash
deno task ai review src/components/MyComponent.tsx
```

Reviews for:
- Best practices
- Performance optimizations
- Type safety
- Accessibility
- Potential bugs

### Type Generation

Automatically infer and generate TypeScript types:

```bash
deno task ai generate types src/services/api.ts
```

## Core Concepts

### Components

Build UI with SolidJS reactive components:

```typescript
import { createSignal } from "solid-js";
import { View, Text } from "@solid-native/core";

export function Counter() {
  const [count, setCount] = createSignal(0);

  return (
    <View>
      <Text>{count()}</Text>
      <Button onPress={() => setCount(c => c + 1)} />
    </View>
  );
}
```

### Services

Encapsulate business logic in services:

```typescript
import { ApiService } from "./services";

const api = new ApiService({
  baseUrl: "https://api.example.com"
});

const users = await api.get<User[]>("/users");
```

### Stores

Manage state with Solid stores:

```typescript
import { createSignal } from "solid-js";

export function useAuth() {
  const [user, setUser] = createSignal<User | null>(null);

  const login = async (credentials: Credentials) => {
    const user = await api.post("/auth/login", credentials);
    setUser(user);
  };

  return { user, login };
}
```

### Platform Utilities

Write platform-specific code:

```typescript
import { Platform } from "./utils";

const styles = {
  padding: Platform.select({
    ios: 16,
    android: 12,
    default: 16,
  }),
};
```

## Architecture Decisions

### Why Deno?

- Built-in TypeScript support (no compilation needed)
- Modern standard library
- Secure by default (explicit permissions)
- Fast startup and execution
- Native HTTP server and fetch
- FFI for native bindings

### Why SolidJS?

- True reactivity (no virtual DOM)
- Fine-grained updates (only what changes)
- Smaller bundle sizes
- Better performance than React
- Simpler mental model
- Perfect for mobile constraints

### Why AI-First?

- Accelerates development
- Reduces boilerplate
- Enforces best practices
- Lowers barrier to entry
- Generates documentation
- Catches bugs early

## Configuration Reference

### App Config

```typescript
app: {
  name: string;              // Internal name
  displayName: string;       // User-facing name
  version: string;           // Semantic version
  bundleId: string;          // Unique identifier
  orientation: "portrait" | "landscape" | "auto";
}
```

### AI Config

```typescript
ai: {
  enabled: boolean;
  provider: "anthropic" | "openai" | "local";
  model: string;
  features: {
    componentGeneration: boolean;
    codeCompletion: boolean;
    typeGeneration: boolean;
    codeReview: boolean;
    naturalLanguageCommands: boolean;
    docGeneration: boolean;
  };
}
```

### Platform Config

```typescript
platforms: {
  ios: {
    enabled: boolean;
    deploymentTarget: string;   // Min iOS version
    swiftVersion: string;
    bundleId: string;
  };
  android: {
    enabled: boolean;
    minSdkVersion: number;
    targetSdkVersion: number;
    compileSdkVersion: number;
    bundleId: string;
  };
}
```

## Environment Variables

```bash
# AI Configuration
ANTHROPIC_API_KEY=sk-ant-...
OPENAI_API_KEY=sk-...

# Logging
LOG_LEVEL=debug|info|warn|error

# Development
DEV_PORT=8081
```

## Contributing

This is a scaffold/template for Solid Native projects. Contributions welcome!

## Roadmap

- [ ] Implement actual dev server with HMR
- [ ] Native FFI bindings for iOS/Android
- [ ] Component library
- [ ] Navigation system
- [ ] Animation framework
- [ ] Testing utilities
- [ ] CI/CD templates
- [ ] Plugin system
- [ ] Streaming SSR

## License

MIT

## Links

- [Deno Documentation](https://deno.land/manual)
- [SolidJS Documentation](https://www.solidjs.com)
- [Anthropic Claude](https://www.anthropic.com)

---

Built with ❤️ for the future of mobile development
