package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// Direction represents the direction enum in Yoga.
type Direction int

const (
	DirectionInherit Direction = C.YGDirectionInherit
	DirectionLTR     Direction = C.YGDirectionLTR
	DirectionRTL     Direction = C.YGDirectionRTL
)

// FlexDirection represents the flex direction enum in Yoga.
type FlexDirection int

const (
	FlexDirectionColumn        FlexDirection = C.YGFlexDirectionColumn
	FlexDirectionColumnReverse FlexDirection = C.YGFlexDirectionColumnReverse
	FlexDirectionRow           FlexDirection = C.YGFlexDirectionRow
	FlexDirectionRowReverse    FlexDirection = C.YGFlexDirectionRowReverse
)

// Justify represents the justify content enum in Yoga.
type Justify int

const (
	JustifyFlexStart    Justify = C.YGJustifyFlexStart
	JustifyCenter       Justify = C.YGJustifyCenter
	JustifyFlexEnd      Justify = C.YGJustifyFlexEnd
	JustifySpaceBetween Justify = C.YGJustifySpaceBetween
	JustifySpaceAround  Justify = C.YGJustifySpaceAround
	JustifySpaceEvenly  Justify = C.YGJustifySpaceEvenly
)

// Align represents the align enum in Yoga.
type Align int

const (
	AlignAuto         Align = C.YGAlignAuto
	AlignFlexStart    Align = C.YGAlignFlexStart
	AlignCenter       Align = C.YGAlignCenter
	AlignFlexEnd      Align = C.YGAlignFlexEnd
	AlignStretch      Align = C.YGAlignStretch
	AlignBaseline     Align = C.YGAlignBaseline
	AlignSpaceBetween Align = C.YGAlignSpaceBetween
	AlignSpaceAround  Align = C.YGAlignSpaceAround
	AlignSpaceEvenly  Align = C.YGAlignSpaceEvenly
)

// PositionType represents the position type enum in Yoga.
type PositionType int

const (
	PositionTypeStatic   PositionType = C.YGPositionTypeStatic
	PositionTypeRelative PositionType = C.YGPositionTypeRelative
	PositionTypeAbsolute PositionType = C.YGPositionTypeAbsolute
)

// Wrap represents the wrap enum in Yoga.
type Wrap int

const (
	WrapNoWrap      Wrap = C.YGWrapNoWrap
	WrapWrap        Wrap = C.YGWrapWrap
	WrapWrapReverse Wrap = C.YGWrapWrapReverse
)

// Overflow represents the overflow enum in Yoga.
type Overflow int

const (
	OverflowVisible Overflow = C.YGOverflowVisible
	OverflowHidden  Overflow = C.YGOverflowHidden
	OverflowScroll  Overflow = C.YGOverflowScroll
)

// Display represents the display enum in Yoga.
type Display int

const (
	DisplayFlex Display = C.YGDisplayFlex
	DisplayNone Display = C.YGDisplayNone
)

// Edge represents the edge enum in Yoga.
type Edge int

const (
	EdgeLeft       Edge = C.YGEdgeLeft
	EdgeTop        Edge = C.YGEdgeTop
	EdgeRight      Edge = C.YGEdgeRight
	EdgeBottom     Edge = C.YGEdgeBottom
	EdgeStart      Edge = C.YGEdgeStart
	EdgeEnd        Edge = C.YGEdgeEnd
	EdgeHorizontal Edge = C.YGEdgeHorizontal
	EdgeVertical   Edge = C.YGEdgeVertical
	EdgeAll        Edge = C.YGEdgeAll
)

// Gutter represents the gutter enum in Yoga.
type Gutter int

const (
	GutterColumn Gutter = C.YGGutterColumn
	GutterRow    Gutter = C.YGGutterRow
	GutterAll    Gutter = C.YGGutterAll
)
