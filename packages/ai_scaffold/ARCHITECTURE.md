# Solid Native Architecture

## Design Philosophy

Solid Native is built on three core principles:

1. **AI-First Development**: LLMs as first-class development tools
2. **Fine-Grained Reactivity**: SolidJS for optimal performance
3. **Modern Runtime**: Deno for security, speed, and simplicity

## System Architecture

```
┌─────────────────────────────────────────────────────────┐
│                   Developer Interface                    │
│  (Natural Language, CLI, Config Files)                  │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                   AI Layer (This Package)                │
│  - Component Generation                                  │
│  - Code Review                                           │
│  - Type Inference                                        │
│  - Documentation                                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                Application Layer (Your Code)             │
│  - Components (Solid)                                    │
│  - Services (Business Logic)                             │
│  - Stores (State Management)                             │
│  - Utils (Helpers)                                       │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              Framework Layer (@solid-native/core)        │
│  - Native Components (View, Text, etc)                   │
│  - Platform Abstraction                                  │
│  - Event System                                          │
│  - Style Engine                                          │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              Runtime Layer (@solid-native/runtime)       │
│  - Deno Runtime                                          │
│  - FFI Bindings                                          │
│  - Native Bridge                                         │
│  - Platform APIs                                         │
└────────────────────┬────────────────────────────────────┘
                     │
            ┌────────┴────────┐
            ▼                 ▼
    ┌──────────────┐  ┌──────────────┐
    │   iOS APIs   │  │Android APIs  │
    │   (Swift)    │  │  (Kotlin)    │
    └──────────────┘  └──────────────┘
```

## Component Architecture

### Reactivity Model

Solid Native uses SolidJS's fine-grained reactivity system:

```typescript
// Signal - primitive reactive value
const [count, setCount] = createSignal(0);

// Memo - derived value (cached computation)
const doubled = createMemo(() => count() * 2);

// Effect - side effect that runs when dependencies change
createEffect(() => {
  console.log("Count changed:", count());
});
```

Benefits:
- No virtual DOM diffing
- Precise updates (only changed values)
- Better performance than React
- Smaller bundle sizes
- Clearer mental model

### Component Patterns

#### Presentational Components

Pure UI components with no business logic:

```typescript
export function Button(props: ButtonProps) {
  return (
    <Pressable onPress={props.onPress}>
      <Text>{props.title}</Text>
    </Pressable>
  );
}
```

#### Container Components

Components that manage state and logic:

```typescript
export function UserProfile() {
  const { user, loading } = useUser();

  return (
    <Show when={!loading()} fallback={<Loading />}>
      <ProfileCard user={user()} />
    </Show>
  );
}
```

## State Management

### Local State (Signals)

For component-specific state:

```typescript
function Counter() {
  const [count, setCount] = createSignal(0);
  return <button onClick={() => setCount(c => c + 1)}>{count()}</button>;
}
```

### Shared State (Stores)

For app-wide state:

```typescript
// stores/auth.ts
const [user, setUser] = createSignal<User | null>(null);

export function useAuth() {
  return {
    user,
    login: async (creds) => { /* ... */ },
    logout: () => setUser(null),
  };
}
```

### Persistent State (Storage)

For data that survives app restarts:

```typescript
const storage = new StorageService(nativeAdapter);
await storage.set("user", userData);
const user = await storage.get<User>("user");
```

## Service Layer

Services encapsulate business logic and external communication:

### API Service

HTTP communication:

```typescript
const api = new ApiService({ baseUrl: "https://api.example.com" });
const users = await api.get<User[]>("/users");
```

### Storage Service

Data persistence:

```typescript
const storage = new StorageService(adapter);
await storage.set("key", value);
```

Pattern: Services are plain classes, not tied to React/Solid lifecycle

## Platform Abstraction

### Platform Detection

```typescript
if (Platform.isIOS) {
  // iOS-specific code
} else if (Platform.isAndroid) {
  // Android-specific code
}

const value = Platform.select({
  ios: 16,
  android: 12,
  default: 16,
});
```

### FFI Bindings

Direct native code calls via Deno FFI:

```typescript
const lib = Deno.dlopen("./native/libmylib.so", {
  "myFunction": { parameters: ["i32"], result: "i32" },
});

const result = lib.symbols.myFunction(42);
```

## AI Integration

### Code Generation Pipeline

```
Natural Language Description
          ↓
    AI Client (LLM)
          ↓
   Schema Validation (Zod)
          ↓
  Code Template Generation
          ↓
    File System Write
```

### Component Generation Flow

1. User provides description
2. AI client sends to LLM with system prompt
3. LLM returns structured JSON
4. Schema validates response
5. Template engine generates code
6. Code written to file

### Context Building

For AI to understand your codebase:

```typescript
const context = {
  projectStructure: await scanFiles(),
  dependencies: await parseDependencies(),
  recentChanges: await getGitDiff(),
  existingTypes: await extractTypes(),
};
```

## Build System

### Development Mode

1. Load configuration
2. Start file watcher
3. Bundle with source maps
4. Inject HMR client
5. Start WebSocket server
6. Connect to native runtime

### Production Build

1. Load configuration
2. Bundle all code
3. Tree shake unused code
4. Minify JavaScript
5. Optimize assets
6. Generate platform bundles
7. Package native apps

### Hot Module Replacement

```
File Change Detected
        ↓
   Bundle Module
        ↓
  Compute Diff
        ↓
 Send via WebSocket
        ↓
 Runtime Updates
        ↓
Preserve Component State
```

## Native Bridge

### Message Passing

```typescript
// JavaScript → Native
bridge.send({
  type: "NATIVE_API_CALL",
  api: "camera",
  method: "takePicture",
  params: { quality: 0.8 },
});

// Native → JavaScript
bridge.on("NATIVE_EVENT", (event) => {
  console.log("Native event:", event);
});
```

### FFI Direct Calls

For performance-critical paths:

```typescript
const result = nativeLib.symbols.fastComputation(data);
```

## Security Model

### Deno Permissions

- `--allow-read`: File system read access
- `--allow-write`: File system write access
- `--allow-net`: Network access
- `--allow-env`: Environment variables
- `--allow-ffi`: Foreign function interface
- `--allow-run`: Subprocess execution

Each permission can be scoped:

```bash
deno run --allow-read=./src --allow-net=api.example.com app.ts
```

### API Key Management

- Never commit API keys
- Use environment variables
- Validate at runtime
- Separate dev/prod keys

## Performance Considerations

### Bundle Size

- Tree shaking removes unused code
- Solid has no runtime overhead (no VDOM)
- Code splitting for lazy loading
- Asset optimization

### Runtime Performance

- Fine-grained updates (no re-renders)
- Direct FFI for native calls
- Efficient event handling
- Minimal JavaScript ↔ Native bridge traffic

### Memory Management

- Cleanup effects on component unmount
- Dispose subscriptions
- Clear timers
- Release FFI resources

## Testing Strategy

### Unit Tests

Test pure functions and utilities:

```typescript
Deno.test("Platform detection", () => {
  assertEquals(Platform.OS, "ios");
});
```

### Component Tests

Test component behavior:

```typescript
Deno.test("Button calls onPress", () => {
  let called = false;
  render(() => <Button onPress={() => called = true} />);
  fireEvent.click(screen.getByRole("button"));
  assert(called);
});
```

### Integration Tests

Test service integration:

```typescript
Deno.test("API service fetches users", async () => {
  const api = new ApiService({ baseUrl: mockServer.url });
  const users = await api.get("/users");
  assertEquals(users.length, 2);
});
```

## Future Enhancements

### Streaming SSR

Server-side rendering with streaming:

```typescript
const stream = renderToStream(() => <App />);
return new Response(stream, { headers: { "Content-Type": "text/html" } });
```

### Concurrent Rendering

Prioritize important updates:

```typescript
startTransition(() => {
  setSearchResults(newResults);
});
```

### Plugin System

Extend framework capabilities:

```typescript
export default {
  plugins: [
    analyticsPlugin(),
    crashReportingPlugin(),
    customBridgePlugin(),
  ],
};
```

## Comparison to Alternatives

### vs Expo/React Native

| Feature | Solid Native | Expo |
|---------|-------------|------|
| Runtime | Deno | Node.js |
| UI Library | SolidJS | React |
| Reactivity | Fine-grained | Virtual DOM |
| TypeScript | Built-in | Via Babel |
| Bundler | Deno/esbuild | Metro |
| FFI | Native Deno FFI | Limited |
| AI Tools | First-class | Third-party |

### vs Flutter

| Feature | Solid Native | Flutter |
|---------|-------------|---------|
| Language | TypeScript/JavaScript | Dart |
| Ecosystem | npm/JSR | pub.dev |
| Web Support | Native | Yes |
| Learning Curve | Lower (JS) | Higher (Dart) |
| Bundle Size | Smaller | Larger |

## References

- [SolidJS Reactivity](https://www.solidjs.com/guides/reactivity)
- [Deno FFI](https://deno.land/manual/runtime/ffi_api)
- [Anthropic Claude](https://docs.anthropic.com/claude/docs)
- [Mobile First Design](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps/Responsive/Mobile_first)
