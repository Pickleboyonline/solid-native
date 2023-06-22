// import { Text } from "../core/views/text.tsx";
import { createSignal } from "solid-js";
import { Button } from "../core/views/button.tsx";

export function App() {
  const [count, setCount] = createSignal(0);

  setInterval(() => setCount(count() + 1), 1000);

  return <Button title="Click me!" />;
}
