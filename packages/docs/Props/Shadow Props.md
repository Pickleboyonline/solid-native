This roadmap outlines the necessary Shadow style props to implement for SolidNative, along with their descriptions and type signatures.

## iOS Specific Shadow Props

- [ ] **shadowColor**: Sets the color of the shadow.
  - Type: `ColorValue | undefined`
- [ ] **shadowOffset**: Sets the offset for the shadow.
  - Type: `Readonly<{width: number; height: number}> | undefined`
- [ ] **shadowOpacity**: Sets the opacity of the shadow.
  - Type: `AnimatableNumericValue | undefined`
- [ ] **shadowRadius**: Sets the blur radius of the shadow.
  - Type: `number | undefined`
