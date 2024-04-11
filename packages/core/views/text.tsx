import { splitProps } from "solid-js";
type TextProps = {
  children: JSX.Element;
  style?: SolidNativeTextStyle;
};

type SolidNativeTextStyle = {
  color?: string; // Represented as a hex string, e.g., "#FFFFFF"
  fontFamily?: string; // The name of the font
  fontSize?: number; // Font size in points
  fontStyle?: "normal" | "italic"; // Only 'normal' and 'italic' are generally supported in most systems
  fontWeight?:
    | "normal"
    | "bold"
    | "100"
    | "200"
    | "300"
    | "400"
    | "500"
    | "600"
    | "700"
    | "800"
    | "900";
  fontVariant?: string[]; // Not directly supported in SwiftUI. Might need special handling or be ignored.
  letterSpacing?: number; // Space between characters in points
  lineHeight?: number; // Not directly supported in SwiftUI Text views
  textAlign?: "auto" | "left" | "right" | "center" | "justify"; // 'justify' might need special handling
  textDecorationColor?: string; // Represented as a hex string
  textDecorationLine?:
    | "none"
    | "underline"
    | "line-through"
    | "underline line-through";
  textDecorationStyle?: "solid" | "double" | "dotted" | "dashed"; // Not directly supported in SwiftUI, consider alternatives
  textShadowColor?: string; // Represented as a hex string
  textShadowOffset?: { width: number; height: number };
  textShadowRadius?: number;
  textTransform?: "none" | "capitalize" | "uppercase" | "lowercase";
  writingDirection?: "auto" | "ltr" | "rtl"; // Not directly supported in SwiftUI Text views
  userSelect?: "none" | "text" | "all" | "auto"; // Not applicable in native development
};

// Example of using the interface
// const exampleStyle: SolidNativeTextStyle = {
//   color: "#FF00FF",
//   fontFamily: "Helvetica",
//   fontSize: 16,
//   fontWeight: "bold",
//   textAlign: "center",
//   textDecorationLine: "underline",
//   textShadowColor: "#000000",
//   textShadowOffset: { width: 2, height: 2 },
//   textShadowRadius: 3,
// };

export function Text(props: TextProps) {
  const [local, rest] = splitProps(props, ["children"]);
  return <sn_text {...rest}>{local.children}</sn_text>;
}
