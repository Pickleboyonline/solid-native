This roadmap outlines the necessary Text style props to implement for SolidNative, along with their descriptions and type signatures.

## iOS Specific Props

- [ ] **fontVariant**: Specifies font variants for text.
  - Type: `FontVariant[] | undefined`
- [ ] **textDecorationColor**: Sets the color of the text decoration.
  - Type: `ColorValue | undefined`
- [ ] **textDecorationStyle**: Sets the style of the text decoration (solid, double, dotted, dashed).
  - Type: `'solid' | 'double' | 'dotted' | 'dashed' | undefined`
- [ ] **writingDirection**: Sets the writing direction for text.
  - Type: `'auto' | 'ltr' | 'rtl' | undefined`

## Android Specific Props

- [ ] **textAlignVertical**: Sets the vertical alignment of text.
  - Type: `'auto' | 'top' | 'bottom' | 'center' | undefined`
- [ ] **verticalAlign**: Sets the vertical alignment within the parent element.
  - Type: `'auto' | 'top' | 'bottom' | 'middle' | undefined`
- [ ] **includeFontPadding**: Controls whether the text includes extra font padding.
  - Type: `boolean | undefined`

## Common Text Style Props

- [ ] **color**: Sets the color of the text.
  - Type: `ColorValue | undefined`
- [ ] **fontFamily**: Sets the font family for the text.
  - Type: `string | undefined`
- [ ] **fontSize**: Sets the font size of the text.
  - Type: `number | undefined`
- [ ] **fontStyle**: Sets the style of the font (normal, italic).
  - Type: `'normal' | 'italic' | undefined`
- [ ] **fontWeight**: Specifies the font weight.
  - Type: `'normal' | 'bold' | '100' | '200' | '300' | '400' | '500' | '600' | '700' | '800' | '900' | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | 'ultralight' | 'thin' | 'light' | 'medium' | 'regular' | 'semibold' | 'condensedBold' | 'condensed' | 'heavy' | 'black' | undefined`
- [ ] **letterSpacing**: Sets the spacing between characters.
  - Type: `number | undefined`
- [ ] **lineHeight**: Sets the height of a line of text.
  - Type: `number | undefined`
- [ ] **textAlign**: Sets the alignment of text (auto, left, right, center, justify).
  - Type: `'auto' | 'left' | 'right' | 'center' | 'justify' | undefined`
- [ ] **textDecorationLine**: Sets the decoration line for text (none, underline, line-through, underline line-through).
  - Type: `'none' | 'underline' | 'line-through' | 'underline line-through' | undefined`
- [ ] **textDecorationStyle**: Sets the style of the text decoration (solid, double, dotted, dashed).
  - Type: `'solid' | 'double' | 'dotted' | 'dashed' | undefined`
- [ ] **textDecorationColor**: Sets the color of the text decoration.
  - Type: `ColorValue | undefined`
- [ ] **textShadowColor**: Sets the color of the text shadow.
  - Type: `ColorValue | undefined`
- [ ] **textShadowOffset**: Sets the offset for the text shadow.
  - Type: `{width: number; height: number} | undefined`
- [ ] **textShadowRadius**: Sets the radius of the text shadow.
  - Type: `number | undefined`
- [ ] **textTransform**: Controls the capitalization of text (none, capitalize, uppercase, lowercase).
  - Type: `'none' | 'capitalize' | 'uppercase' | 'lowercase' | undefined`
- [ ] **userSelect**: Controls whether the text can be selected.
  - Type: `'auto' | 'none' | 'text' | 'contain' | 'all' | undefined`
