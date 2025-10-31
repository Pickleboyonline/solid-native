/**
 * Theme store - manages app-wide theming
 * Uses Solid's createStore for fine-grained reactivity
 */

import { createSignal } from "solid-js";

export type Theme = "light" | "dark";

export interface ThemeColors {
  background: string;
  foreground: string;
  primary: string;
  secondary: string;
  accent: string;
  border: string;
}

const lightTheme: ThemeColors = {
  background: "#FFFFFF",
  foreground: "#000000",
  primary: "#007AFF",
  secondary: "#5856D6",
  accent: "#FF9500",
  border: "#E5E5EA",
};

const darkTheme: ThemeColors = {
  background: "#000000",
  foreground: "#FFFFFF",
  primary: "#0A84FF",
  secondary: "#5E5CE6",
  accent: "#FF9F0A",
  border: "#38383A",
};

const [theme, setTheme] = createSignal<Theme>("light");

export function useTheme() {
  const toggleTheme = () => {
    setTheme((prev) => prev === "light" ? "dark" : "light");
  };

  const colors = (): ThemeColors => {
    return theme() === "light" ? lightTheme : darkTheme;
  };

  return {
    theme,
    setTheme,
    toggleTheme,
    colors,
  };
}
