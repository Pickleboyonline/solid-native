import { FlexStyle } from './types.ts'

export interface ViewStyle extends FlexStyle {
  backgroundColor?: string
}

export type ViewProps = {
  style?: ViewStyle
  children?: JSX.Element;
};

export function View(props: ViewProps) {
  return <sn_view {...props} />;
}
