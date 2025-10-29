import { createSignal, onMount } from "solid-js";
import { View, Text, Button } from "solid-native/core";
import { createEffect } from "solid-js";
import { QueryClient, QueryClientProvider } from "@tanstack/solid-query";
import { createNavigation } from "./navigators/stub.ts";
import { RootStack } from "./navigators/root_stack.ts";

const queryClient = new QueryClient();

const Navigation = createNavigation(RootStack) as () => string;

export function App() {
  /**
   * TODO: Want to build a General TODO App. Features a:
   * - Screen Navigation
   * - Account creation
   * -
   */

  /**
   * Screen Navigation is gonna be difficult, because it's a large lib.
   *
   */
  return (
    <QueryClientProvider client={queryClient}>
      <Navigation />
    </QueryClientProvider>
  );
}
