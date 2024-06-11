#  🚧  Button

This roadmap outlines the necessary Button props to implement for SolidNative, along with their descriptions and type signatures.

## Common Button Props

- [ ] **title**: Text to display inside the button. On Android, the given title will be converted to the uppercased form.
  - Type: `string`
- [ ] **color**: Color of the text (iOS) or background color of the button (Android).
  - Type: `ColorValue | undefined`

## Accessibility Props

- [ ] **accessibilityLabel**: Overrides the text that's read by the screen reader when the user interacts with the element.
  - Type: `string | undefined`
- [ ] **accessibilityState**: Describes the current state of the element to the screen reader.
  - Type: `{ disabled?: boolean | undefined; selected?: boolean | undefined; checked?: boolean | "mixed" | undefined; busy?: boolean | undefined; expanded?: boolean | undefined; } | undefined`

## TV Focus Props

- [ ] **hasTVPreferredFocus**: May be set to true to force the Apple TV focus engine to move focus to this view.
  - Type: `boolean | undefined`
- [ ] **nextFocusDown**: Designates the next view to receive focus when the user navigates down.
  - Type: `number | undefined`
- [ ] **nextFocusForward**: Designates the next view to receive focus when the user navigates forward.
  - Type: `number | undefined`
- [ ] **nextFocusLeft**: Designates the next view to receive focus when the user navigates left.
  - Type: `number | undefined`
- [ ] **nextFocusRight**: Designates the next view to receive focus when the user navigates right.
  - Type: `number | undefined`
- [ ] **nextFocusUp**: Designates the next view to receive focus when the user navigates up.
  - Type: `number | undefined`

## Interaction Props

- [ ] **testID**: Used to locate this view in end-to-end tests.
  - Type: `string | undefined`
- [ ] **disabled**: If true, disable all interactions for this component.
  - Type: `boolean | undefined`
- [ ] **onPress**: This function is called when the button is pressed.
  - Type: `((event: GestureResponderEvent) => void) | undefined`
- [ ] **touchSoundDisabled**: If true, disable the touch sound for this component.
  - Type: `boolean | undefined`
