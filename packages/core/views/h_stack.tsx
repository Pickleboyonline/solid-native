import { SolidNativeElement } from "../types.ts";

export type VStackAlignment = "leading" | "trailing" | "center";

export type VStackProps = {
  children?: (SolidNativeElement | string)[] | SolidNativeElement | string;
  alignment?: VStackAlignment;
  spacing?: number;
};

export function HStack({ children }: VStackProps) {
  return <sn_h_stack> {children}</sn_h_stack>;
}
