/**
 * Main application entry point
 * This is where your Solid Native app starts
 */

import { createSignal } from "solid-js";
import { View, Text, Button } from "@solid-native/core";
import { useTheme } from "./stores/theme.ts";

export default function App() {
  const [count, setCount] = createSignal(0);
  const { theme, toggleTheme } = useTheme();

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
      <Text style={{ fontSize: 32, marginBottom: 20 }}>
        Welcome to Solid Native
      </Text>

      <Text style={{ fontSize: 18, marginBottom: 10 }}>
        Count: {count()}
      </Text>

      <Button
        onPress={() => setCount(count() + 1)}
        title="Increment"
        style={{ marginBottom: 10 }}
      />

      <Button
        onPress={toggleTheme}
        title={`Theme: ${theme()}`}
      />
    </View>
  );
}
