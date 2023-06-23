import { For } from "solid-js";
import { SolidNativeElement } from "../types.ts";

export type VStackProps = {
  children?: (SolidNativeElement | string)[] | SolidNativeElement | string;
};

export function VStack({ children }: VStackProps) {
  return <sn_v_stack> {children}</sn_v_stack>;
}
