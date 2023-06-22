import { createSignal } from "solid-js";

export function App() {
  const [count, setCount] = createSignal(0);

  setInterval(() => setCount(count() + 1), 1000);

  return <>{`Hello from Solid!\nCounter: ${count()}`}</>;
}
