type ButtonProps = {
  title?: string;
  onPress?: () => void;
};

export function Button(props: ButtonProps) {
  return <sn_button {...props} />;
}
