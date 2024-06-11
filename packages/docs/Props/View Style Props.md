This roadmap outlines the necessary View style props to implement for SolidNative, along with their descriptions and type signatures.

## Visibility Props

- [ ] **backfaceVisibility**: Controls whether the back face of the view is visible when rotated.
  - Type: `'visible' | 'hidden' | undefined`
- [ ] **opacity**: Sets the opacity of the view.
  - Type: `AnimatableNumericValue | undefined`

## Color Props

- [ ] **backgroundColor**: Sets the background color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderBlockColor**: Sets the block border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderBlockEndColor**: Sets the block end border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderBlockStartColor**: Sets the block start border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderBottomColor**: Sets the bottom border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderColor**: Sets the border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderEndColor**: Sets the end border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderLeftColor**: Sets the left border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderRightColor**: Sets the right border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderStartColor**: Sets the start border color of the view.
  - Type: `ColorValue | undefined`
- [ ] **borderTopColor**: Sets the top border color of the view.
  - Type: `ColorValue | undefined`

## Border Radius Props

- [ ] **borderBottomEndRadius**: Sets the bottom end border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderBottomLeftRadius**: Sets the bottom left border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderBottomRightRadius**: Sets the bottom right border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderBottomStartRadius**: Sets the bottom start border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderCurve**: Sets the corner curve of the borders (iOS 13+).
  - Type: `'circular' | 'continuous' | undefined`
- [ ] **borderEndEndRadius**: Sets the end end border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderEndStartRadius**: Sets the end start border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderRadius**: Sets the border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderStartEndRadius**: Sets the start end border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderStartStartRadius**: Sets the start start border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderTopEndRadius**: Sets the top end border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderTopLeftRadius**: Sets the top left border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderTopRightRadius**: Sets the top right border radius of the view.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **borderTopStartRadius**: Sets the top start border radius of the view.
  - Type: `AnimatableNumericValue | undefined`

## Border Style Props

- [ ] **borderStyle**: Sets the style of the border (solid, dotted, dashed).
  - Type: `'solid' | 'dotted' | 'dashed' | undefined`

## Elevation and Z-Index Props

- [ ] **elevation**: Sets the elevation of the view (Android only).
  - Type: `number | undefined`

## Interaction Props

- [ ] **pointerEvents**: Controls whether the view can be the target of touch events.
  - Type: `'box-none' | 'none' | 'box-only' | 'auto' | undefined`
- [ ] **cursor**: Sets the cursor type when hovering over the view.
  - Type: `CursorValue | undefined`
