type TextProps = {
  children: JSX.Element;
  style?: TextStyle;
  bold?: boolean;
};

type TextStyle = {
  color?: "red" | "black";
};

export function Text({ children, ...rest }: TextProps) {
  return <sn_text {...rest}>{children}</sn_text>;
}
