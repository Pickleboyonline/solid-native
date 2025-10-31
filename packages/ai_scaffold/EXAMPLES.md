# Solid Native Examples

## Quick Start Examples

### Basic Counter App

```typescript
// src/app.tsx
import { createSignal } from "solid-js";
import { View, Text, Button } from "@solid-native/core";

export default function App() {
  const [count, setCount] = createSignal(0);

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
      <Text style={{ fontSize: 48, marginBottom: 20 }}>
        {count()}
      </Text>

      <Button
        title="Increment"
        onPress={() => setCount(c => c + 1)}
      />
    </View>
  );
}
```

### Todo List with State Management

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
    setTodos([
      ...todos(),
      { id: crypto.randomUUID(), text, completed: false },
    ]);
  };

  const toggleTodo = (id: string) => {
    setTodos(todos().map(todo =>
      todo.id === id ? { ...todo, completed: !todo.completed } : todo
    ));
  };

  const removeTodo = (id: string) => {
    setTodos(todos().filter(todo => todo.id !== id));
  };

  return {
    todos,
    addTodo,
    toggleTodo,
    removeTodo,
  };
}
```

```typescript
// src/components/TodoList.tsx
import { For } from "solid-js";
import { View, Text, Button } from "@solid-native/core";
import { useTodos } from "../stores/todos.ts";

export function TodoList() {
  const { todos, toggleTodo, removeTodo } = useTodos();

  return (
    <View>
      <For each={todos()}>
        {(todo) => (
          <View style={{ flexDirection: "row", padding: 12 }}>
            <Text
              style={{
                flex: 1,
                textDecoration: todo.completed ? "line-through" : "none",
              }}
              onPress={() => toggleTodo(todo.id)}
            >
              {todo.text}
            </Text>

            <Button
              title="Delete"
              onPress={() => removeTodo(todo.id)}
            />
          </View>
        )}
      </For>
    </View>
  );
}
```

### API Integration

```typescript
// src/services/users.ts
import { ApiService } from "./api.ts";

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
}

class UserService extends ApiService {
  constructor() {
    super({ baseUrl: "https://api.example.com" });
  }

  async getUsers(): Promise<User[]> {
    return this.get<User[]>("/users");
  }

  async getUser(id: string): Promise<User> {
    return this.get<User>(`/users/${id}`);
  }

  async createUser(data: Omit<User, "id">): Promise<User> {
    return this.post<User>("/users", data);
  }

  async updateUser(id: string, data: Partial<User>): Promise<User> {
    return this.put<User>(`/users/${id}`, data);
  }

  async deleteUser(id: string): Promise<void> {
    await this.delete(`/users/${id}`);
  }
}

export const userService = new UserService();
```

```typescript
// src/components/UserList.tsx
import { createSignal, onMount, For, Show } from "solid-js";
import { View, Text } from "@solid-native/core";
import { userService, type User } from "../services/users.ts";

export function UserList() {
  const [users, setUsers] = createSignal<User[]>([]);
  const [loading, setLoading] = createSignal(true);
  const [error, setError] = createSignal<string | null>(null);

  onMount(async () => {
    try {
      const data = await userService.getUsers();
      setUsers(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load users");
    } finally {
      setLoading(false);
    }
  });

  return (
    <View>
      <Show when={loading()}>
        <Text>Loading...</Text>
      </Show>

      <Show when={error()}>
        <Text style={{ color: "red" }}>{error()}</Text>
      </Show>

      <For each={users()}>
        {(user) => (
          <View style={{ padding: 12 }}>
            <Text style={{ fontSize: 18, fontWeight: "bold" }}>
              {user.name}
            </Text>
            <Text style={{ color: "#666" }}>{user.email}</Text>
          </View>
        )}
      </For>
    </View>
  );
}
```

### Form Handling

```typescript
// src/components/LoginForm.tsx
import { createSignal } from "solid-js";
import { View, Text, TextInput, Button } from "@solid-native/core";

export interface LoginFormProps {
  onSubmit: (email: string, password: string) => Promise<void>;
}

export function LoginForm(props: LoginFormProps) {
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [loading, setLoading] = createSignal(false);
  const [error, setError] = createSignal<string | null>(null);

  const handleSubmit = async () => {
    if (!email() || !password()) {
      setError("Please fill in all fields");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      await props.onSubmit(email(), password());
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <View style={{ padding: 20 }}>
      <TextInput
        value={email()}
        onChangeText={setEmail}
        placeholder="Email"
        keyboardType="email-address"
        autoCapitalize="none"
        style={{ marginBottom: 12 }}
      />

      <TextInput
        value={password()}
        onChangeText={setPassword}
        placeholder="Password"
        secureTextEntry
        style={{ marginBottom: 12 }}
      />

      <Show when={error()}>
        <Text style={{ color: "red", marginBottom: 12 }}>
          {error()}
        </Text>
      </Show>

      <Button
        title={loading() ? "Logging in..." : "Login"}
        onPress={handleSubmit}
        disabled={loading()}
      />
    </View>
  );
}
```

### Theme System

```typescript
// src/stores/theme.ts
import { createSignal, createMemo } from "solid-js";
import { StorageService } from "../services/storage.ts";

export type Theme = "light" | "dark" | "auto";

const [theme, setTheme] = createSignal<Theme>("auto");
const [systemTheme, setSystemTheme] = createSignal<"light" | "dark">("light");

export function useTheme() {
  const activeTheme = createMemo(() => {
    return theme() === "auto" ? systemTheme() : theme();
  });

  const colors = createMemo(() => {
    const isDark = activeTheme() === "dark";

    return {
      background: isDark ? "#000000" : "#FFFFFF",
      foreground: isDark ? "#FFFFFF" : "#000000",
      primary: isDark ? "#0A84FF" : "#007AFF",
      secondary: isDark ? "#5E5CE6" : "#5856D6",
      border: isDark ? "#38383A" : "#E5E5EA",
      card: isDark ? "#1C1C1E" : "#F2F2F7",
      error: isDark ? "#FF453A" : "#FF3B30",
      success: isDark ? "#32D74B" : "#34C759",
    };
  });

  const setAndSaveTheme = async (newTheme: Theme) => {
    setTheme(newTheme);
    // await storage.set("theme", newTheme);
  };

  return {
    theme,
    activeTheme,
    colors,
    setTheme: setAndSaveTheme,
    isDark: () => activeTheme() === "dark",
  };
}
```

### Platform-Specific Code

```typescript
// src/components/StatusBar.tsx
import { Platform } from "../utils/platform.ts";
import { StatusBar as IOSStatusBar } from "@solid-native/core/ios";
import { StatusBar as AndroidStatusBar } from "@solid-native/core/android";

export interface StatusBarProps {
  style: "light" | "dark";
}

export function StatusBar(props: StatusBarProps) {
  return Platform.select({
    ios: <IOSStatusBar barStyle={props.style} />,
    android: <AndroidStatusBar barStyle={props.style} />,
    default: null,
  });
}
```

### AI Component Generation

```bash
# Generate a profile card component
deno task scaffold "a user profile card with avatar image, name, bio text, and follow button"
```

Generated output:

```typescript
/**
 * User profile card component
 */

import { Show } from "solid-js";
import { View, Text, Image, Button } from "@solid-native/core";

export interface UserProfileCardProps {
  /** User's avatar image URL */
  avatar: string;

  /** User's display name */
  name: string;

  /** User's bio text */
  bio: string;

  /** Whether the user is currently followed */
  isFollowing: boolean;

  /** Callback when follow button is pressed */
  onFollowPress: () => void;
}

export function UserProfileCard(props: UserProfileCardProps) {
  return (
    <View
      style={{
        backgroundColor: "#FFFFFF",
        borderRadius: 12,
        padding: 16,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
      }}
    >
      <View style={{ flexDirection: "row", alignItems: "center" }}>
        <Image
          source={{ uri: props.avatar }}
          style={{
            width: 64,
            height: 64,
            borderRadius: 32,
            marginRight: 16,
          }}
        />

        <View style={{ flex: 1 }}>
          <Text
            style={{
              fontSize: 18,
              fontWeight: "bold",
              marginBottom: 4,
            }}
          >
            {props.name}
          </Text>

          <Text
            style={{
              fontSize: 14,
              color: "#666",
              marginBottom: 12,
            }}
          >
            {props.bio}
          </Text>

          <Button
            title={props.isFollowing ? "Following" : "Follow"}
            onPress={props.onFollowPress}
            variant={props.isFollowing ? "outline" : "primary"}
          />
        </View>
      </View>
    </View>
  );
}
```

### Custom Hooks

```typescript
// src/utils/hooks/useDebounce.ts
import { createSignal, createEffect, onCleanup } from "solid-js";

export function useDebounce<T>(value: () => T, delay: number) {
  const [debouncedValue, setDebouncedValue] = createSignal<T>(value());

  createEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedValue(value());
    }, delay);

    onCleanup(() => clearTimeout(timer));
  });

  return debouncedValue;
}
```

```typescript
// Usage
import { createSignal } from "solid-js";
import { useDebounce } from "./utils/hooks/useDebounce.ts";

function SearchComponent() {
  const [search, setSearch] = createSignal("");
  const debouncedSearch = useDebounce(search, 500);

  createEffect(() => {
    // This only runs 500ms after user stops typing
    console.log("Searching for:", debouncedSearch());
  });

  return <TextInput value={search()} onChangeText={setSearch} />;
}
```

### Async Data Loading

```typescript
// src/utils/hooks/useAsync.ts
import { createSignal, createEffect } from "solid-js";

export function useAsync<T>(
  asyncFn: () => Promise<T>,
  dependencies: unknown[] = [],
) {
  const [data, setData] = createSignal<T | null>(null);
  const [loading, setLoading] = createSignal(true);
  const [error, setError] = createSignal<Error | null>(null);

  createEffect(() => {
    // Re-run when dependencies change
    dependencies.forEach((dep) => dep);

    setLoading(true);
    setError(null);

    asyncFn()
      .then(setData)
      .catch(setError)
      .finally(() => setLoading(false));
  });

  return { data, loading, error };
}
```

```typescript
// Usage
function UserProfile(props: { userId: string }) {
  const { data: user, loading, error } = useAsync(
    () => userService.getUser(props.userId),
    [props.userId],
  );

  return (
    <Show when={!loading()} fallback={<Loading />}>
      <Show when={!error()} fallback={<Error error={error()!} />}>
        <UserCard user={user()!} />
      </Show>
    </Show>
  );
}
```

## Best Practices

### 1. Component Composition

Break down complex UIs into smaller, reusable components:

```typescript
// Good
<UserProfile>
  <Avatar user={user} />
  <UserInfo user={user} />
  <FollowButton userId={user.id} />
</UserProfile>

// Avoid
<UserProfileWithEverything user={user} />
```

### 2. Derive State

Use `createMemo` for computed values:

```typescript
// Good
const fullName = createMemo(() => `${firstName()} ${lastName()}`);

// Avoid
const [fullName, setFullName] = createSignal("");
createEffect(() => setFullName(`${firstName()} ${lastName()}`));
```

### 3. Cleanup Effects

Always cleanup subscriptions and timers:

```typescript
createEffect(() => {
  const timer = setInterval(() => {
    console.log("tick");
  }, 1000);

  onCleanup(() => clearInterval(timer));
});
```

### 4. Type Everything

Use TypeScript for all code:

```typescript
// Good
interface User {
  id: string;
  name: string;
}

function getUser(id: string): Promise<User> {
  // ...
}

// Avoid
function getUser(id) {
  // ...
}
```

### 5. Platform Abstraction

Write platform-agnostic code when possible:

```typescript
// Good
const padding = Platform.select({ ios: 16, android: 12 });

// Avoid
if (Platform.OS === "ios") {
  // iOS code
} else {
  // Android code
}
```

## Common Patterns

### Loading States

```typescript
<Show when={!loading()} fallback={<Spinner />}>
  <Content />
</Show>
```

### Error Handling

```typescript
<Show when={!error()} fallback={<ErrorMessage error={error()!} />}>
  <Content />
</Show>
```

### Conditional Rendering

```typescript
<Show when={isLoggedIn()} fallback={<LoginPrompt />}>
  <UserDashboard />
</Show>
```

### List Rendering

```typescript
<For each={items()}>
  {(item, index) => <ItemCard item={item} index={index()} />}
</For>
```

### Dynamic Components

```typescript
<Dynamic component={componentMap[type()]} {...props} />
```
