import { FlexStyle } from "./types.ts";

export interface ImageStyle extends FlexStyle {
  resizeMode?: "cover" | "contain" | "stretch" | "center";
}

export type ImageSource = {
  uri: string;
};

export type ImageProps = {
  source: ImageSource;
  style?: ImageStyle;
  resizeMode?: "cover" | "contain" | "stretch" | "center";
};

export function Image(props: ImageProps) {
  return <sn_image {...props} />;
}
