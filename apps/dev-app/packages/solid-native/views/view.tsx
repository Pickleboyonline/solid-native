import { ViewStyle } from "./types.ts";
import type { JSX } from "solid-js";

export type ViewProps = {
  style?: ViewStyle;
  children?: JSX.Element;
};

export function View(props: ViewProps) {
  return <sn_view {...props} />;
}
