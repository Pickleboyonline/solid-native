import { createEnum } from "../lib/createEnum.ts";

export const FontWeight = createEnum([
  "ultraLight",
  "thin",
  "light",
  "regular",
  "medium",
  "semibold",
  "bold",
  "heavy",
  "black",
] as const);

type TextProps = {
  // children: JSX.Element;
  fontWeight?: keyof typeof FontWeight;
  text?: string;
  fontSize?: number;
  strikethrough?: unknown;
  underline?: unknown;
  bold?: boolean;
  foregroundColor?: string;
};

/**
 * Turn this into a shadow text node.
 * @param param0
 * @returns
 */
export function Text(props: TextProps) {
  return <sn_text {...props} />;
}
