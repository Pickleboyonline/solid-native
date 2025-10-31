/**
 * Example reusable button component
 */

import { JSX } from "solid-js";
import { Pressable, Text } from "@solid-native/core";

export interface ButtonProps {
  title: string;
  onPress: () => void;
  variant?: "primary" | "secondary" | "outline";
  disabled?: boolean;
  style?: JSX.CSSProperties;
}

export function Button(props: ButtonProps) {
  const getButtonStyles = (): JSX.CSSProperties => {
    const base = {
      padding: 12,
      borderRadius: 8,
      alignItems: "center" as const,
      justifyContent: "center" as const,
      minWidth: 100,
    };

    switch (props.variant ?? "primary") {
      case "primary":
        return { ...base, backgroundColor: "#007AFF" };
      case "secondary":
        return { ...base, backgroundColor: "#5856D6" };
      case "outline":
        return { ...base, borderWidth: 1, borderColor: "#007AFF" };
      default:
        return base;
    }
  };

  const getTextStyles = (): JSX.CSSProperties => {
    const base = {
      fontSize: 16,
      fontWeight: "600" as const,
    };

    if (props.variant === "outline") {
      return { ...base, color: "#007AFF" };
    }

    return { ...base, color: "#FFFFFF" };
  };

  return (
    <Pressable
      onPress={props.onPress}
      disabled={props.disabled}
      style={{ ...getButtonStyles(), ...props.style }}
    >
      <Text style={getTextStyles()}>
        {props.title}
      </Text>
    </Pressable>
  );
}
