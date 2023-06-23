import { createSignal } from "solid-js";
import { print, Button, VStack } from "solid-native/core";

export function App() {
  const [count, setCount] = createSignal(0);

  setInterval(() => setCount(count() + 1), 1000);

  return (
    <VStack>
      <Button
        title={"Click me! " + count()}
        onPress={() => {
          print("Hello world!");
        }}
      />
      <Button title="Click me 2" />
    </VStack>
  );
}
