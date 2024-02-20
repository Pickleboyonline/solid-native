export type ViewProps = {
  children?: JSX.Element;
};

export function View({ children }: ViewProps) {
  return <sn_view> {children}</sn_view>;
}
