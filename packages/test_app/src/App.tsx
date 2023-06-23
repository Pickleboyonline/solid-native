import { createSignal, onMount } from "solid-js";
import { print, Button, VStack } from "solid-native/core";

export function App() {
  const [count, setCount] = createSignal(0);

  setInterval(() => setCount(count() + 1), 1000);

  onMount(() => {
    print("App Mounted!");
  });

  return (
    <VStack>
      {`Hello World! Count: ${count()}`}
      <Button title="Click me!" />
    </VStack>
  );
}
