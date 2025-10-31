# Solid Native Scaffold - Project Structure

## Complete File Tree

```
solid-native/packages/ai_scaffold/
├── README.md                              # Main documentation
├── ARCHITECTURE.md                        # Architecture deep-dive
├── EXAMPLES.md                            # Code examples & patterns
├── PROJECT_STRUCTURE.md                   # This file
├── .gitignore                             # Git ignore rules
├── .env.example                           # Environment template
├── deno.json                              # Deno config & tasks
├── solid-native.config.ts                 # Framework configuration
├── mod.ts                                 # Main export file
│
├── src/
│   ├── app.tsx                           # Application entry point
│   │
│   ├── components/                       # Reusable UI components
│   │   ├── Button.tsx                    # Example button component
│   │   └── index.ts                      # Component exports
│   │
│   ├── services/                         # Business logic layer
│   │   ├── api.ts                        # HTTP client service
│   │   ├── storage.ts                    # Persistence service
│   │   └── index.ts                      # Service exports
│   │
│   ├── stores/                           # State management
│   │   ├── theme.ts                      # Theme store
│   │   └── index.ts                      # Store exports
│   │
│   ├── utils/                            # Utility functions
│   │   ├── platform.ts                   # Platform detection
│   │   ├── logger.ts                     # Logging utility
│   │   └── index.ts                      # Util exports
│   │
│   ├── ai/                               # AI-powered tools (KEY DIFFERENTIATOR)
│   │   ├── client.ts                     # Unified LLM client
│   │   ├── generator.ts                  # Code generator
│   │   ├── index.ts                      # AI exports
│   │   │
│   │   ├── schemas/                      # Generation schemas
│   │   │   ├── component.schema.ts       # Component schema & validation
│   │   │   └── index.ts                  # Schema exports
│   │   │
│   │   └── prompts/                      # System prompts
│   │       ├── component-generator.ts    # Component gen prompts
│   │       └── index.ts                  # Prompt exports
│   │
│   ├── cli/                              # Command-line tools
│   │   ├── dev.ts                        # Development server
│   │   ├── build.ts                      # Production builder
│   │   ├── start.ts                      # App launcher
│   │   ├── scaffold.ts                   # AI scaffolding CLI
│   │   ├── ai.ts                         # AI assistant CLI
│   │   └── doctor.ts                     # Environment checker
│   │
│   └── types/                            # TypeScript definitions
│       └── config.ts                     # Config types
│
└── [build outputs would go in dist/]
```

## File Purposes

### Configuration Files

| File | Purpose |
|------|---------|
| `deno.json` | Deno runtime config, dependencies, tasks, compiler options |
| `solid-native.config.ts` | Framework-specific configuration (platforms, AI, build) |
| `.env.example` | Environment variable template (API keys, etc) |
| `.gitignore` | Files to exclude from version control |

### Entry Points

| File | Purpose |
|------|---------|
| `mod.ts` | Main package export (for `import` from other packages) |
| `src/app.tsx` | Application UI entry point |

### Documentation

| File | Purpose |
|------|---------|
| `README.md` | Getting started, features, usage guide |
| `ARCHITECTURE.md` | Technical architecture, design decisions |
| `EXAMPLES.md` | Code examples, patterns, best practices |
| `PROJECT_STRUCTURE.md` | This file - structure overview |

### Source Code

#### Components (`src/components/`)
Reusable UI building blocks:
- `Button.tsx` - Example button with variants
- More components added as you build

#### Services (`src/services/`)
Business logic and external communication:
- `api.ts` - HTTP REST client
- `storage.ts` - Persistent data storage abstraction
- Add more: auth, analytics, etc.

#### Stores (`src/stores/`)
Application state management:
- `theme.ts` - Global theme state
- Add more: user, cart, settings, etc.

#### Utils (`src/utils/`)
Pure utility functions:
- `platform.ts` - Platform detection & selection
- `logger.ts` - Structured logging
- Add more: date formatting, validation, etc.

#### AI (`src/ai/`) - 🌟 KEY INNOVATION
AI-powered development tools:
- `client.ts` - Unified interface for Anthropic/OpenAI
- `generator.ts` - Code generation engine
- `schemas/` - Zod schemas for validation
- `prompts/` - System prompts for different tasks

#### CLI (`src/cli/`)
Command-line interfaces:
- `dev.ts` - Development server with HMR
- `build.ts` - Production build tool
- `start.ts` - Run built app
- `scaffold.ts` - AI-powered scaffolding
- `ai.ts` - AI assistant commands
- `doctor.ts` - Environment diagnostics

#### Types (`src/types/`)
TypeScript type definitions:
- `config.ts` - Configuration types

## Key Design Patterns

### 1. Barrel Exports (`index.ts`)
Each directory exports via `index.ts` for clean imports:

```typescript
// Instead of:
import { Button } from "./components/Button.tsx";
import { Card } from "./components/Card.tsx";

// You can:
import { Button, Card } from "./components";
```

### 2. Service Layer Pattern
Business logic separated from UI:

```
UI (Components) → Services → External APIs/Storage
```

### 3. Store Pattern
Global state management:

```typescript
// Define once
const [state, setState] = createSignal(initial);

// Use everywhere
export function useStore() {
  return { state, actions };
}
```

### 4. Platform Abstraction
Write once, run anywhere:

```typescript
Platform.select({
  ios: iosImplementation,
  android: androidImplementation,
  default: fallback,
});
```

### 5. AI-First Workflow
Natural language → Structured code:

```
Description → LLM → Schema → Template → Code
```

## Dependencies (from deno.json)

### Core Framework
- `@solid-native/core` - UI components (to be built)
- `@solid-native/runtime` - Native runtime (to be built)
- `solid-js` - Reactivity engine

### AI & ML
- `@anthropic/sdk` - Claude AI client
- `openai` - OpenAI client
- `zod` - Schema validation

### Tooling
- `cliffy` - CLI framework
- `ts-morph` - TypeScript AST manipulation

### Standard Library
- `std/*` - Deno standard library

## Tasks (from deno.json)

| Task | Command | Purpose |
|------|---------|---------|
| `dev` | `deno task dev` | Start development server |
| `build` | `deno task build` | Build for production |
| `start` | `deno task start` | Run built app |
| `scaffold` | `deno task scaffold "..."` | Generate code with AI |
| `ai` | `deno task ai generate component "..."` | AI assistant |
| `doctor` | `deno task doctor` | Check environment |
| `test` | `deno task test` | Run tests |
| `check` | `deno task check` | Type check |
| `fmt` | `deno task fmt` | Format code |
| `lint` | `deno task lint` | Lint code |

## What Makes This Special

### 1. LLM-Native Development
- Component generation from descriptions
- Code review automation
- Type inference
- Documentation generation

### 2. Deno-First
- No `node_modules` bloat
- Built-in TypeScript
- Secure by default
- Fast startup

### 3. SolidJS Performance
- No virtual DOM
- Fine-grained reactivity
- Smaller bundles
- Faster updates

### 4. Mobile-First Design
- Touch-optimized components
- Platform-aware styling
- Native performance
- Cross-platform by default

### 5. Developer Experience
- Convention over configuration
- Minimal boilerplate
- Clear structure
- Type-safe

## Next Steps for Development

### Phase 1: Foundation (Current)
- ✅ Project structure
- ✅ Configuration system
- ✅ AI tooling skeleton
- ✅ Example components
- ✅ Documentation

### Phase 2: Core Framework
- [ ] Implement `@solid-native/core` components
- [ ] Implement `@solid-native/runtime` bridge
- [ ] FFI bindings for iOS/Android
- [ ] Navigation system
- [ ] Gesture handling

### Phase 3: Development Tools
- [ ] Working dev server with HMR
- [ ] Fast refresh for Solid
- [ ] Source maps
- [ ] DevTools integration

### Phase 4: Build System
- [ ] Production bundler
- [ ] Tree shaking
- [ ] Code splitting
- [ ] Asset optimization
- [ ] Platform-specific builds

### Phase 5: AI Enhancement
- [ ] Context-aware generation
- [ ] Multi-file refactoring
- [ ] Automated testing
- [ ] Performance optimization suggestions

### Phase 6: Ecosystem
- [ ] Component library
- [ ] Plugin system
- [ ] Templates
- [ ] Examples gallery
- [ ] Community tools

## Philosophy

This scaffold embodies:

1. **Simplicity**: Clear structure, obvious patterns
2. **Power**: AI assistance, modern tools
3. **Performance**: Solid + Deno + Native
4. **Safety**: TypeScript everywhere
5. **Joy**: Delightful DX

## Comparison: File Count

**Expo/React Native starter:**
- ~500+ files (with node_modules)
- Complex configuration
- Multiple config files

**Solid Native:**
- ~30 source files
- Single config file
- No node_modules (Deno)
- Everything you need, nothing you don't

## Contributing Guide

When adding new features:

1. **Components** → `src/components/`
2. **Business Logic** → `src/services/`
3. **State** → `src/stores/`
4. **Utils** → `src/utils/`
5. **AI Tools** → `src/ai/`
6. **CLI Tools** → `src/cli/`

Always:
- Add TypeScript types
- Export via `index.ts`
- Document with JSDoc
- Follow naming conventions
- Test your code

---

**This structure is designed to scale from prototypes to production apps while keeping complexity low and developer joy high.**
