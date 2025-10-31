# Quick Start Guide

Get up and running with Solid Native in 5 minutes.

## Prerequisites

```bash
# Install Deno (if not already installed)
curl -fsSL https://deno.land/install.sh | sh

# Verify installation
deno --version
```

## Setup

### 1. Configure Environment

```bash
# Copy environment template
cp .env.example .env

# Add your API key
echo "ANTHROPIC_API_KEY=sk-ant-your-key-here" >> .env
```

### 2. Verify Setup

```bash
# Check environment
deno task doctor
```

Expected output:
```
🔍 Checking Solid Native environment...

✅ Deno version
✅ Configuration file
✅ Dependencies
✅ AI configuration

✨ Environment check complete!
```

## Your First Component

### Using AI Scaffolding

```bash
# Generate a component with natural language
deno task scaffold "a profile card with avatar, name, email, and a follow button"
```

This will output a complete, typed component:

```typescript
/**
 * Profile card component
 */

import { View, Text, Image, Button } from "@solid-native/core";

export interface ProfileCardProps {
  avatar: string;
  name: string;
  email: string;
  isFollowing: boolean;
  onFollowPress: () => void;
}

export function ProfileCard(props: ProfileCardProps) {
  return (
    <View style={{ padding: 16, backgroundColor: "#fff", borderRadius: 12 }}>
      <Image source={{ uri: props.avatar }} style={{ width: 64, height: 64 }} />
      <Text style={{ fontSize: 18, fontWeight: "bold" }}>{props.name}</Text>
      <Text style={{ color: "#666" }}>{props.email}</Text>
      <Button
        title={props.isFollowing ? "Following" : "Follow"}
        onPress={props.onFollowPress}
      />
    </View>
  );
}
```

### Manual Component

Create `src/components/Counter.tsx`:

```typescript
import { createSignal } from "solid-js";
import { View, Text, Button } from "@solid-native/core";

export function Counter() {
  const [count, setCount] = createSignal(0);

  return (
    <View style={{ padding: 20 }}>
      <Text style={{ fontSize: 32 }}>{count()}</Text>
      <Button title="+" onPress={() => setCount(c => c + 1)} />
    </View>
  );
}
```

Export it in `src/components/index.ts`:

```typescript
export { Counter } from "./Counter.tsx";
```

## Development Workflow

### 1. Start Dev Server

```bash
deno task dev
```

This will:
- Start development server on `http://localhost:8081`
- Enable hot module replacement
- Watch for file changes
- Show logs in console

### 2. Make Changes

Edit `src/app.tsx`:

```typescript
import { Counter } from "./components/Counter.tsx";

export default function App() {
  return (
    <View style={{ flex: 1, justifyContent: "center" }}>
      <Counter />
    </View>
  );
}
```

Changes will hot-reload automatically!

### 3. Use AI Assistant

```bash
# Generate a component
deno task ai generate component "a loading spinner with different sizes"

# Review your code
deno task ai review src/components/Counter.tsx

# Generate types
deno task ai generate types src/services/api.ts
```

## Common Tasks

### Add a New Service

```typescript
// src/services/auth.ts
import { ApiService } from "./api.ts";

export interface User {
  id: string;
  email: string;
  name: string;
}

class AuthService extends ApiService {
  async login(email: string, password: string): Promise<User> {
    return this.post<User>("/auth/login", { email, password });
  }

  async logout(): Promise<void> {
    await this.post("/auth/logout");
  }
}

export const authService = new AuthService({
  baseUrl: "https://api.example.com",
});
```

Export in `src/services/index.ts`:

```typescript
export { authService } from "./auth.ts";
export type { User } from "./auth.ts";
```

### Add a New Store

```typescript
// src/stores/user.ts
import { createSignal } from "solid-js";
import type { User } from "../services/index.ts";

const [user, setUser] = createSignal<User | null>(null);
const [isAuthenticated, setIsAuthenticated] = createSignal(false);

export function useUser() {
  const login = (userData: User) => {
    setUser(userData);
    setIsAuthenticated(true);
  };

  const logout = () => {
    setUser(null);
    setIsAuthenticated(false);
  };

  return {
    user,
    isAuthenticated,
    login,
    logout,
  };
}
```

Export in `src/stores/index.ts`:

```typescript
export { useUser } from "./user.ts";
```

### Add Platform-Specific Code

```typescript
import { Platform } from "./utils/platform.ts";

const styles = {
  container: {
    padding: Platform.select({
      ios: 20,
      android: 16,
      default: 18,
    }),
    marginTop: Platform.isIOS ? 44 : 0, // Status bar height
  },
};
```

## Building for Production

```bash
# Build optimized bundle
deno task build

# Run production build
deno task start
```

This will:
- Bundle and minify code
- Tree shake unused code
- Optimize assets
- Generate source maps
- Create platform-specific packages

## Testing

### Unit Tests

Create `src/utils/platform.test.ts`:

```typescript
import { assertEquals } from "std/assert/mod.ts";
import { Platform } from "./platform.ts";

Deno.test("Platform.select returns correct value", () => {
  const result = Platform.select({
    ios: "iOS",
    android: "Android",
    default: "Web",
  });

  // Will depend on environment
  assertEquals(typeof result, "string");
});
```

Run tests:

```bash
deno task test
```

## Code Quality

```bash
# Format code
deno task fmt

# Lint code
deno task lint

# Type check
deno task check
```

## Example: Complete Feature

Let's build a todo list feature:

### 1. Generate Component with AI

```bash
deno task scaffold "a todo list item with checkbox, text, and delete button"
```

### 2. Create Store

```typescript
// src/stores/todos.ts
import { createSignal } from "solid-js";

export interface Todo {
  id: string;
  text: string;
  completed: boolean;
}

const [todos, setTodos] = createSignal<Todo[]>([]);

export function useTodos() {
  const addTodo = (text: string) => {
    setTodos([...todos(), {
      id: crypto.randomUUID(),
      text,
      completed: false,
    }]);
  };

  const toggleTodo = (id: string) => {
    setTodos(todos().map(todo =>
      todo.id === id ? { ...todo, completed: !todo.completed } : todo
    ));
  };

  const deleteTodo = (id: string) => {
    setTodos(todos().filter(todo => todo.id !== id));
  };

  return { todos, addTodo, toggleTodo, deleteTodo };
}
```

### 3. Create Components

```typescript
// src/components/TodoItem.tsx
import { View, Text, Button, Checkbox } from "@solid-native/core";
import type { Todo } from "../stores/todos.ts";

export interface TodoItemProps {
  todo: Todo;
  onToggle: () => void;
  onDelete: () => void;
}

export function TodoItem(props: TodoItemProps) {
  return (
    <View style={{ flexDirection: "row", padding: 12 }}>
      <Checkbox
        checked={props.todo.completed}
        onToggle={props.onToggle}
      />
      <Text
        style={{
          flex: 1,
          textDecoration: props.todo.completed ? "line-through" : "none",
        }}
      >
        {props.todo.text}
      </Text>
      <Button title="Delete" onPress={props.onDelete} />
    </View>
  );
}
```

```typescript
// src/components/TodoList.tsx
import { For } from "solid-js";
import { View } from "@solid-native/core";
import { TodoItem } from "./TodoItem.tsx";
import { useTodos } from "../stores/todos.ts";

export function TodoList() {
  const { todos, toggleTodo, deleteTodo } = useTodos();

  return (
    <View>
      <For each={todos()}>
        {(todo) => (
          <TodoItem
            todo={todo}
            onToggle={() => toggleTodo(todo.id)}
            onDelete={() => deleteTodo(todo.id)}
          />
        )}
      </For>
    </View>
  );
}
```

### 4. Use in App

```typescript
// src/app.tsx
import { createSignal } from "solid-js";
import { View, TextInput, Button } from "@solid-native/core";
import { TodoList } from "./components/TodoList.tsx";
import { useTodos } from "./stores/todos.ts";

export default function App() {
  const [input, setInput] = createSignal("");
  const { addTodo } = useTodos();

  const handleAdd = () => {
    if (input().trim()) {
      addTodo(input());
      setInput("");
    }
  };

  return (
    <View style={{ flex: 1, padding: 20 }}>
      <View style={{ flexDirection: "row", marginBottom: 20 }}>
        <TextInput
          value={input()}
          onChangeText={setInput}
          placeholder="Add todo..."
          style={{ flex: 1, marginRight: 8 }}
        />
        <Button title="Add" onPress={handleAdd} />
      </View>

      <TodoList />
    </View>
  );
}
```

Done! You now have a complete, working todo list with:
- ✅ Type-safe components
- ✅ Reactive state management
- ✅ Proper code organization
- ✅ Platform-aware styling

## Next Steps

- Read [ARCHITECTURE.md](./ARCHITECTURE.md) for deep dive
- Check [EXAMPLES.md](./EXAMPLES.md) for more patterns
- Review [README.md](./README.md) for full documentation
- Explore [PROJECT_STRUCTURE.md](./PROJECT_STRUCTURE.md) for organization

## Tips

1. **Use AI generously** - It's built in for a reason!
2. **Keep components small** - Single responsibility principle
3. **Type everything** - TypeScript is your friend
4. **Platform-agnostic first** - Use Platform.select only when needed
5. **Derive state** - Use createMemo instead of duplicate state

## Troubleshooting

### "Module not found"
```bash
# Re-cache dependencies
deno cache --reload src/app.tsx
```

### "API key not configured"
```bash
# Check environment
cat .env

# Set key
export ANTHROPIC_API_KEY=sk-ant-...
```

### "Permission denied"
```bash
# Deno is secure by default, grant permissions:
deno task dev --allow-all
```

## Getting Help

- Check documentation files in this directory
- Review examples in EXAMPLES.md
- Use AI assistant: `deno task ai`
- Read Solid docs: https://solidjs.com
- Read Deno docs: https://deno.land

---

Happy coding with Solid Native! 🚀
