import { splitProps } from "solid-js";
import type { JSX } from "solid-js";
import { TextStyle } from "./types.ts";

export type TextProps = {
  children?: JSX.Element;
  style?: TextStyle;
};

export function Text(props: TextProps) {
  const [local, rest] = splitProps(props, ["children"]);
  return <sn_text {...rest}>{local.children}</sn_text>;
}
