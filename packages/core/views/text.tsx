type TextProps = {
  children: JSX.Element;
};

export function Text({ children }: TextProps) {
  return <sn_text>{children}</sn_text>;
}