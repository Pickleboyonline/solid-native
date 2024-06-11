# 🚧 View

The base View Component. I'm tracking progress on this by its props.

(Also uses Style, Layout and Shadow Props)

# View Props:

This roadmap outlines the necessary View props to implement for SolidNative, along with their descriptions and type signatures.

## iOS Specific Props

- [ ] **shouldRasterizeIOS**: Whether this view should be rendered as a bitmap before compositing. Useful for animations and interactions that do not modify this component's dimensions nor its children.
  - Type: `boolean | undefined`

## Apple TV Specific Props

- [ ] **isTVSelectable**: When set to true, this view will be focusable and navigable using the Apple TV remote.
  - Type: `boolean | undefined`
- [ ] **hasTVPreferredFocus**: May be set to true to force the Apple TV focus engine to move focus to this view.
  - Type: `boolean | undefined`
- [ ] **tvParallaxProperties**: Object with properties to control Apple TV parallax effects.
  - Type: `TVParallaxProperties | undefined`
- [ ] **tvParallaxShiftDistanceX**: Used to change the appearance of the Apple TV parallax effect when this view goes in or out of focus.
  - Type: `number | undefined`
- [ ] **tvParallaxShiftDistanceY**: Used to change the appearance of the Apple TV parallax effect when this view goes in or out of focus.
  - Type: `number | undefined`
- [ ] **tvParallaxTiltAngle**: Used to change the appearance of the Apple TV parallax effect when this view goes in or out of focus.
  - Type: `number | undefined`
- [ ] **tvParallaxMagnification**: Used to change the appearance of the Apple TV parallax effect when this view goes in or out of focus.
  - Type: `number | undefined`

## Android Specific Props

- [ ] **collapsable**: Views that are only used to layout their children or otherwise don't draw anything may be automatically removed from the native hierarchy as an optimization.
  - Type: `boolean | undefined`
- [ ] **renderToHardwareTextureAndroid**: Whether this view should render itself (and all of its children) into a single hardware texture on the GPU. Useful for animations and interactions that only modify opacity, rotation, translation, and/or scale.
  - Type: `boolean | undefined`
- [ ] **focusable**: Whether this `View` should be focusable with a non-touch input device, e.g., receive focus with a hardware keyboard.
  - Type: `boolean | undefined`
- [ ] **tabIndex**: Indicates whether this `View` should be focusable with a non-touch input device, e.g., receive focus with a hardware keyboard.
  - Type: `0 | -1 | undefined`

## Common View Props

- [ ] **children**: Child components.
  - Type: `React.ReactNode | undefined`
- [ ] **hitSlop**: Defines how far a touch event can start away from the view.
  - Type: `null | Insets | number | undefined`
- [ ] **id**: Used to reference react managed views from native code.
  - Type: `string | undefined`
- [ ] **needsOffscreenAlphaCompositing**: Whether this view needs to be rendered offscreen and composited with an alpha to preserve correct colors and blending behavior.
  - Type: `boolean | undefined`
- [ ] **onLayout**: Invoked on mount and layout changes.
  - Type: `((event: LayoutChangeEvent) => void) | undefined`
- [ ] **pointerEvents**: Controls how this view receives pointer events.
  - Type: `'box-none' | 'none' | 'box-only' | 'auto' | undefined`
- [ ] **removeClippedSubviews**: Special performance property useful for scrolling content when there are many subviews.
  - Type: `boolean | undefined`
- [ ] **style**: Style for the view.
  - Type: `StyleProp<ViewStyle> | undefined`
- [ ] **testID**: Used to locate this view in end-to-end tests.
  - Type: `string | undefined`
- [ ] **nativeID**: Used to reference react managed views from native code.
  - Type: `string | undefined`
