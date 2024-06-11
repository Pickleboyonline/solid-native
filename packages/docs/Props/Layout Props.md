This roadmap outlines the necessary Flexbox props to implement for SolidNative, along with their descriptions and type signatures.

## Flex Container Props

- [ ] **alignContent**: Controls the alignment of the lines in a flex container.
  - Type: `'flex-start' | 'flex-end' | 'center' | 'stretch' | 'space-between' | 'space-around' | 'space-evenly' | undefined`
- [ ] **alignItems**: Aligns flex items along the cross axis.
  - Type: `FlexAlignType | undefined`
- [ ] **flexDirection**: Defines the main axis direction (row or column).
  - Type: `'row' | 'column' | 'row-reverse' | 'column-reverse' | undefined`
- [ ] **flexWrap**: Controls whether the flex container is single-line or multi-line.
  - Type: `'wrap' | 'nowrap' | 'wrap-reverse' | undefined`
- [ ] **justifyContent**: Aligns flex items along the main axis.
  - Type: `'flex-start' | 'flex-end' | 'center' | 'space-between' | 'space-around' | 'space-evenly' | undefined`

## Flex Item Props

- [ ] **alignSelf**: Overrides the `alignItems` value for a specific flex item.
  - Type: `'auto' | FlexAlignType | undefined`
- [ ] **flex**: Shorthand for `flexGrow`, `flexShrink`, and `flexBasis`.
  - Type: `number | undefined`
- [ ] **flexBasis**: Sets the initial size of a flex item.
  - Type: `DimensionValue | undefined`
- [ ] **flexGrow**: Defines the ability for a flex item to grow if necessary.
  - Type: `number | undefined`
- [ ] **flexShrink**: Defines the ability for a flex item to shrink if necessary.
  - Type: `number | undefined`

## Sizing Props

- [ ] **width**: Sets the width of an element.
  - Type: `DimensionValue | undefined`
- [ ] **height**: Sets the height of an element.
  - Type: `DimensionValue | undefined`
- [ ] **maxWidth**: Sets the maximum width of an element.
  - Type: `DimensionValue | undefined`
- [ ] **maxHeight**: Sets the maximum height of an element.
  - Type: `DimensionValue | undefined`
- [ ] **minWidth**: Sets the minimum width of an element.
  - Type: `DimensionValue | undefined`
- [ ] **minHeight**: Sets the minimum height of an element.
  - Type: `DimensionValue | undefined`
- [ ] **aspectRatio**: Defines the aspect ratio for an element.
  - Type: `number | string | undefined`

## Margin Props

- [ ] **margin**: Sets the margin on all sides of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginBottom**: Sets the bottom margin of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginEnd**: Sets the end margin of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginHorizontal**: Sets the horizontal margins of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginLeft**: Sets the left margin of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginRight**: Sets the right margin of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginStart**: Sets the start margin of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginTop**: Sets the top margin of an element.
  - Type: `DimensionValue | undefined`
- [ ] **marginVertical**: Sets the vertical margins of an element.
  - Type: `DimensionValue | undefined`

## Padding Props

- [ ] **padding**: Sets the padding on all sides of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingBottom**: Sets the bottom padding of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingEnd**: Sets the end padding of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingHorizontal**: Sets the horizontal paddings of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingLeft**: Sets the left padding of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingRight**: Sets the right padding of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingStart**: Sets the start padding of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingTop**: Sets the top padding of an element.
  - Type: `DimensionValue | undefined`
- [ ] **paddingVertical**: Sets the vertical paddings of an element.
  - Type: `DimensionValue | undefined`

## Positioning Props

- [ ] **position**: Sets the position type of an element.
  - Type: `'absolute' | 'relative' | 'static' | undefined`
- [ ] **top**: Sets the top position of an element.
  - Type: `DimensionValue | undefined`
- [ ] **right**: Sets the right position of an element.
  - Type: `DimensionValue | undefined`
- [ ] **bottom**: Sets the bottom position of an element.
  - Type: `DimensionValue | undefined`
- [ ] **left**: Sets the left position of an element.
  - Type: `DimensionValue | undefined`
- [ ] **start**: Sets the start position of an element.
  - Type: `DimensionValue | undefined`
- [ ] **end**: Sets the end position of an element.
  - Type: `DimensionValue | undefined`

## Other Props

- [ ] **borderWidth**: Sets the border width on all sides of an element.
  - Type: `number | undefined`
- [ ] **borderBottomWidth**: Sets the bottom border width of an element.
  - Type: `number | undefined`
- [ ] **borderEndWidth**: Sets the end border width of an element.
  - Type: `number | undefined`
- [ ] **borderLeftWidth**: Sets the left border width of an element.
  - Type: `number | undefined`
- [ ] **borderRightWidth**: Sets the right border width of an element.
  - Type: `number | undefined`
- [ ] **borderStartWidth**: Sets the start border width of an element.
  - Type: `number | undefined`
- [ ] **borderTopWidth**: Sets the top border width of an element.
  - Type: `number | undefined`
- [ ] **display**: Controls the display behavior of an element.
  - Type: `'none' | 'flex' | undefined`
- [ ] **overflow**: Controls the overflow behavior of an element.
  - Type: `'visible' | 'hidden' | 'scroll' | undefined`
- [ ] **zIndex**: Controls the stack order of an element.
  - Type: `number | undefined`
- [ ] **direction**: Sets the directionality of text and elements.
  - Type: `'inherit' | 'ltr' | 'rtl' | undefined`

## Gap Props

- [ ] **rowGap**: Sets the gap between rows in a flex container.
  - Type: `number | undefined`
- [ ] **gap**: Sets the gap between items in a flex container.
  - Type: `number | undefined`
- [ ] **columnGap**: Sets the gap between columns in a flex container.
  - Type: `number | undefined`
