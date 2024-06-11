#  🚧  Text


# Text Props
This roadmap outlines the necessary Text props to implement for SolidNative, along with their descriptions and type signatures.

## iOS Specific Props

- [ ] **adjustsFontSizeToFit**: Specifies whether the font should be scaled down automatically to fit given style constraints.
  - Type: `boolean | undefined`
- [ ] **dynamicTypeRamp**: The Dynamic Type scale ramp to apply to this element on iOS.
  - Type: `'caption2' | 'caption1' | 'footnote' | 'subheadline' | 'callout' | 'body' | 'headline' | 'title3' | 'title2' | 'title1' | 'largeTitle' | undefined`
- [ ] **suppressHighlighting**: When true, no visual change is made when text is pressed down.
  - Type: `boolean | undefined`
- [ ] **lineBreakStrategyIOS**: Set line break strategy on iOS.
  - Type: `'none' | 'standard' | 'hangul-word' | 'push-out' | undefined`

## Android Specific Props

- [ ] **disabled**: Specifies the disabled state of the text view for testing purposes.
  - Type: `boolean | undefined`
- [ ] **selectable**: Lets the user select text to use the native copy and paste functionality.
  - Type: `boolean | undefined`
- [ ] **selectionColor**: The highlight color of the text.
  - Type: `ColorValue | undefined`
- [ ] **textBreakStrategy**: Set text break strategy on Android API Level 23+.
  - Type: `'simple' | 'highQuality' | 'balanced' | undefined`
- [ ] **dataDetectorType**: Determines the types of data converted to clickable URLs in the text element.
  - Type: `null | 'phoneNumber' | 'link' | 'email' | 'none' | 'all' | undefined`
- [ ] **android_hyphenationFrequency**: Hyphenation strategy.
  - Type: `'normal' | 'none' | 'full' | undefined`

## Common Text Props

- [ ] **allowFontScaling**: Specifies whether fonts should scale to respect Text Size accessibility settings.
  - Type: `boolean | undefined`
- [ ] **children**: Child components.
  - Type: `React.ReactNode | undefined`
- [ ] **ellipsizeMode**: Controls how text is truncated when it exceeds the container width.
  - Type: `'head' | 'middle' | 'tail' | 'clip' | undefined`
- [ ] **id**: Used to reference react managed views from native code.
  - Type: `string | undefined`
- [ ] **lineBreakMode**: Controls how text is truncated when it exceeds the container width.
  - Type: `'head' | 'middle' | 'tail' | 'clip' | undefined`
- [ ] **numberOfLines**: Limits the number of lines of text.
  - Type: `number | undefined`
- [ ] **onLayout**: Invoked on mount and layout changes.
  - Type: `((event: LayoutChangeEvent) => void) | undefined`
- [ ] **onTextLayout**: Invoked on text layout.
  - Type: `((event: NativeSyntheticEvent<TextLayoutEventData>) => void) | undefined`
- [ ] **onPress**: This function is called on press.
  - Type: `((event: GestureResponderEvent) => void) | undefined`
- [ ] **onPressIn**: This function is called when a press is initiated.
  - Type: `((event: GestureResponderEvent) => void) | undefined`
- [ ] **onPressOut**: This function is called when a press is released.
  - Type: `((event: GestureResponderEvent) => void) | undefined`
- [ ] **onLongPress**: This function is called on long press.
  - Type: `((event: GestureResponderEvent) => void) | undefined`
- [ ] **style**: Style for the text.
  - Type: `StyleProp<TextStyle> | undefined`
- [ ] **testID**: Used to locate this view in end-to-end tests.
  - Type: `string | undefined`
- [ ] **nativeID**: Used to reference react managed views from native code.
  - Type: `string | undefined`
- [ ] **maxFontSizeMultiplier**: Specifies the largest possible scale a font can reach when allowFontScaling is enabled.
  - Type: `number | null | undefined`
- [ ] **minimumFontScale**: Specifies the smallest possible scale a font can reach when adjustsFontSizeToFit is enabled.
  - Type: `number | undefined`
