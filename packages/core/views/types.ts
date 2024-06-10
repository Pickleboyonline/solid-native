export interface FlexStyle {
  justifyContent?:
    | "flex-start"
    | "center"
    | "flex-end"
    | "space-between"
    | "space-around"
    | "space-evenly";
  alignItems?: "flex-start" | "center" | "flex-end" | "stretch" | "baseline";
  alignSelf?:
    | "auto"
    | "flex-start"
    | "center"
    | "flex-end"
    | "stretch"
    | "baseline";
  flexDirection?: "row" | "row-reverse" | "column" | "column-reverse";
  // TODO: Impliment this one Swift side. This is shorthand for other flex props, see notes.
  flex?: number;
  flexWrap?: "nowrap" | "wrap" | "wrap-reverse";
  flexBasis?: string | number;
  flexGrow?: number;
  flexShrink?: number;
  width?: string | number;
  maxWidth?: string | number;
  minWidth?: string | number;
  height?: string | number;
  maxHeight?: string | number;
  minHeight?: string | number;
  position?: "absolute" | "relative";
  marginTop?: string | number;
  marginRight?: string | number;
  marginBottom?: string | number;
  marginLeft?: string | number;
  paddingTop?: string | number;
  paddingRight?: string | number;
  paddingBottom?: string | number;
  paddingLeft?: string | number;
  borderWidth?: number;
}
