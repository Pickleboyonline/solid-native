package snmobile

import (
	"fmt"
	"log"
	"nativecore/lib/yoga"
)

type Set map[string]struct{}

// UpdateNodeStyle updates the YGNode based on the JSValue style object.
func updateNodeStyleAndReturnNewStyleKeys(node *yoga.YGNode, styleJSValueMap map[string]JSValue, prevKeys Set) Set {
	// Define default value functions
	defaultValueSetters := map[string]func(){
		"alignContent":      func() { node.SetAlignContent(yoga.AlignStretch) },
		"alignItems":        func() { node.SetAlignItems(yoga.AlignStretch) },
		"alignSelf":         func() { node.SetAlignSelf(yoga.AlignAuto) },
		"aspectRatio":       func() { node.SetAspectRatio(0) },
		"borderBottomWidth": func() { node.SetBorder(yoga.EdgeBottom, 0) },
		"borderEndWidth":    func() { node.SetBorder(yoga.EdgeEnd, 0) },
		"borderLeftWidth":   func() { node.SetBorder(yoga.EdgeLeft, 0) },
		"borderRightWidth":  func() { node.SetBorder(yoga.EdgeRight, 0) },
		"borderStartWidth":  func() { node.SetBorder(yoga.EdgeStart, 0) },
		"borderTopWidth":    func() { node.SetBorder(yoga.EdgeTop, 0) },
		"borderWidth":       func() { node.SetBorder(yoga.EdgeAll, 0) },
		"bottom":            func() { node.SetPosition(yoga.EdgeBottom, 0) },
		"display":           func() { node.SetDisplay(yoga.DisplayFlex) },
		"end":               func() { node.SetPosition(yoga.EdgeEnd, 0) },
		"flex":              func() { node.SetFlex(0) },
		"flexBasis":         func() { node.SetFlexBasisAuto() },
		"flexDirection":     func() { node.SetFlexDirection(yoga.FlexDirectionColumn) },
		"flexGrow":          func() { node.SetFlexGrow(0) },
		"flexShrink":        func() { node.SetFlexShrink(0) },
		"flexWrap":          func() { node.SetFlexWrap(yoga.WrapNoWrap) },
		"gap":               func() { node.SetGap(yoga.GutterAll, 0) },
		"height":            func() { node.SetHeightAuto() },
		"justifyContent":    func() { node.SetJustifyContent(yoga.JustifyFlexStart) },
		"left":              func() { node.SetPosition(yoga.EdgeLeft, 0) },
		"margin":            func() { node.SetMargin(yoga.EdgeAll, 0) },
		"marginBottom":      func() { node.SetMargin(yoga.EdgeBottom, 0) },
		"marginEnd":         func() { node.SetMargin(yoga.EdgeEnd, 0) },
		"marginHorizontal":  func() { node.SetMargin(yoga.EdgeHorizontal, 0) },
		"marginLeft":        func() { node.SetMargin(yoga.EdgeLeft, 0) },
		"marginRight":       func() { node.SetMargin(yoga.EdgeRight, 0) },
		"marginStart":       func() { node.SetMargin(yoga.EdgeStart, 0) },
		"marginTop":         func() { node.SetMargin(yoga.EdgeTop, 0) },
		"marginVertical":    func() { node.SetMargin(yoga.EdgeVertical, 0) },
		"maxHeight":         func() { node.SetMaxHeight(yoga.YGValueUndefined.GetValue()) },
		"maxWidth":          func() { node.SetMaxWidth(yoga.YGValueUndefined.GetValue()) },
		"minHeight":         func() { node.SetMinHeight(yoga.YGValueUndefined.GetValue()) },
		"minWidth":          func() { node.SetMinWidth(yoga.YGValueUndefined.GetValue()) },
		"overflow":          func() { node.SetOverflow(yoga.OverflowVisible) },
		"padding":           func() { node.SetPadding(yoga.EdgeAll, 0) },
		"paddingBottom":     func() { node.SetPadding(yoga.EdgeBottom, 0) },
		"paddingEnd":        func() { node.SetPadding(yoga.EdgeEnd, 0) },
		"paddingHorizontal": func() { node.SetPadding(yoga.EdgeHorizontal, 0) },
		"paddingLeft":       func() { node.SetPadding(yoga.EdgeLeft, 0) },
		"paddingRight":      func() { node.SetPadding(yoga.EdgeRight, 0) },
		"paddingStart":      func() { node.SetPadding(yoga.EdgeStart, 0) },
		"paddingTop":        func() { node.SetPadding(yoga.EdgeTop, 0) },
		"paddingVertical":   func() { node.SetPadding(yoga.EdgeVertical, 0) },
		"position":          func() { node.SetPositionType(yoga.PositionTypeStatic) },
		"right":             func() { node.SetPosition(yoga.EdgeRight, 0) },
		"start":             func() { node.SetPosition(yoga.EdgeStart, 0) },
		"top":               func() { node.SetPosition(yoga.EdgeTop, 0) },
		"width":             func() { node.SetWidthAuto() },
		"zIndex":            func() { /* No direct Yoga equivalent */ },
		"direction":         func() { node.SetStyleDirection(yoga.DirectionInherit) },
	}

	newKeys := make(map[string]struct{})

	// Process the style object
	for key, value := range styleJSValueMap {
		newKeys[key] = struct{}{}
		log.Println("Key Set: ", key, value.valueType)
		switch key {
		case "alignContent":
			if value.IsString() {
				switch value.GetString() {
				case "flex-start":
					node.SetAlignContent(yoga.AlignFlexStart)
				case "flex-end":
					node.SetAlignContent(yoga.AlignFlexEnd)
				case "center":
					node.SetAlignContent(yoga.AlignCenter)
				case "stretch":
					node.SetAlignContent(yoga.AlignStretch)
				case "space-between":
					node.SetAlignContent(yoga.AlignSpaceBetween)
				case "space-around":
					node.SetAlignContent(yoga.AlignSpaceAround)
				case "space-evenly":
					node.SetAlignContent(yoga.AlignSpaceEvenly)
				}
			}
		case "alignItems":
			if value.IsString() {
				switch value.GetString() {
				case "flex-start":
					node.SetAlignItems(yoga.AlignFlexStart)
				case "flex-end":
					node.SetAlignItems(yoga.AlignFlexEnd)
				case "center":
					node.SetAlignItems(yoga.AlignCenter)
				case "stretch":
					node.SetAlignItems(yoga.AlignStretch)
				case "baseline":
					node.SetAlignItems(yoga.AlignBaseline)
				}
			}
		case "alignSelf":
			if value.IsString() {
				switch value.GetString() {
				case "auto":
					node.SetAlignSelf(yoga.AlignAuto)
				case "flex-start":
					node.SetAlignSelf(yoga.AlignFlexStart)
				case "flex-end":
					node.SetAlignSelf(yoga.AlignFlexEnd)
				case "center":
					node.SetAlignSelf(yoga.AlignCenter)
				case "stretch":
					node.SetAlignSelf(yoga.AlignStretch)
				case "baseline":
					node.SetAlignSelf(yoga.AlignBaseline)
				}
			}
		case "aspectRatio":
			if value.IsNumber() {
				node.SetAspectRatio(float32(value.GetNumber()))
			}
		case "borderBottomWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeBottom, float32(value.GetNumber()))
			}
		case "borderEndWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeEnd, float32(value.GetNumber()))
			}
		case "borderLeftWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeLeft, float32(value.GetNumber()))
			}
		case "borderRightWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeRight, float32(value.GetNumber()))
			}
		case "borderStartWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeStart, float32(value.GetNumber()))
			}
		case "borderTopWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeTop, float32(value.GetNumber()))
			}
		case "borderWidth":
			if value.IsNumber() {
				node.SetBorder(yoga.EdgeAll, float32(value.GetNumber()))
			}
		case "bottom":
			handleDimension(node, value, yoga.EdgeBottom)
		case "display":
			if value.IsString() {
				switch value.GetString() {
				case "none":
					node.SetDisplay(yoga.DisplayNone)
				case "flex":
					node.SetDisplay(yoga.DisplayFlex)
				}
			}
		case "end":
			handleDimension(node, value, yoga.EdgeEnd)
		case "flex":
			if value.IsNumber() {
				node.SetFlex(float32(value.GetNumber()))
			}
		case "flexBasis":
			handleDimension(node, value, yoga.EdgeAll)
		case "flexDirection":
			if value.IsString() {
				switch value.GetString() {
				case "row":
					node.SetFlexDirection(yoga.FlexDirectionRow)
				case "column":
					node.SetFlexDirection(yoga.FlexDirectionColumn)
				case "row-reverse":
					node.SetFlexDirection(yoga.FlexDirectionRowReverse)
				case "column-reverse":
					node.SetFlexDirection(yoga.FlexDirectionColumnReverse)
				}
			}
		case "flexGrow":
			if value.IsNumber() {
				node.SetFlexGrow(float32(value.GetNumber()))
			}
		case "flexShrink":
			if value.IsNumber() {
				node.SetFlexShrink(float32(value.GetNumber()))
			}
		case "flexWrap":
			if value.IsString() {
				switch value.GetString() {
				case "nowrap":
					node.SetFlexWrap(yoga.WrapNoWrap)
				case "wrap":
					node.SetFlexWrap(yoga.WrapWrap)
				case "wrap-reverse":
					node.SetFlexWrap(yoga.WrapWrapReverse)
				}
			}
		case "gap":
			if value.IsNumber() {
				node.SetGap(yoga.GutterAll, float32(value.GetNumber()))
			}
		case "height":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    node.SetHeightAuto,
				onPercent: node.SetHeightPercent,
				onNumber:  node.SetHeight,
			})
		case "justifyContent":
			if value.IsString() {
				switch value.GetString() {
				case "flex-start":
					node.SetJustifyContent(yoga.JustifyFlexStart)
				case "flex-end":
					node.SetJustifyContent(yoga.JustifyFlexEnd)
				case "center":
					node.SetJustifyContent(yoga.JustifyCenter)
				case "space-between":
					node.SetJustifyContent(yoga.JustifySpaceBetween)
				case "space-around":
					node.SetJustifyContent(yoga.JustifySpaceAround)
				case "space-evenly":
					node.SetJustifyContent(yoga.JustifySpaceEvenly)
				}
			}
		case "left":
			handleDimension(node, value, yoga.EdgeLeft)
		case "margin":
			handleDimension(node, value, yoga.EdgeAll)
		case "marginBottom":
			handleDimension(node, value, yoga.EdgeBottom)
		case "marginEnd":
			handleDimension(node, value, yoga.EdgeEnd)
		case "marginHorizontal":
			handleDimension(node, value, yoga.EdgeHorizontal)
		case "marginLeft":
			handleDimension(node, value, yoga.EdgeLeft)
		case "marginRight":
			handleDimension(node, value, yoga.EdgeRight)
		case "marginStart":
			handleDimension(node, value, yoga.EdgeStart)
		case "marginTop":
			handleDimension(node, value, yoga.EdgeTop)
		case "marginVertical":
			handleDimension(node, value, yoga.EdgeVertical)
		case "maxHeight":
			handleDimension(node, value, yoga.EdgeAll)
		case "maxWidth":
			handleDimension(node, value, yoga.EdgeAll)
		case "minHeight":
			handleDimension(node, value, yoga.EdgeAll)
		case "minWidth":
			handleDimension(node, value, yoga.EdgeAll)
		case "overflow":
			if value.IsString() {
				switch value.GetString() {
				case "visible":
					node.SetOverflow(yoga.OverflowVisible)
				case "hidden":
					node.SetOverflow(yoga.OverflowHidden)
				case "scroll":
					node.SetOverflow(yoga.OverflowScroll)
				}
			}
		case "padding":
			handleDimension(node, value, yoga.EdgeAll)
		case "paddingBottom":
			handleDimension(node, value, yoga.EdgeBottom)
		case "paddingEnd":
			handleDimension(node, value, yoga.EdgeEnd)
		case "paddingHorizontal":
			handleDimension(node, value, yoga.EdgeHorizontal)
		case "paddingLeft":
			handleDimension(node, value, yoga.EdgeLeft)
		case "paddingRight":
			handleDimension(node, value, yoga.EdgeRight)
		case "paddingStart":
			handleDimension(node, value, yoga.EdgeStart)
		case "paddingTop":
			handleDimension(node, value, yoga.EdgeTop)
		case "paddingVertical":
			handleDimension(node, value, yoga.EdgeVertical)
		case "position":
			if value.IsString() {
				switch value.GetString() {
				case "absolute":
					node.SetPositionType(yoga.PositionTypeAbsolute)
				case "relative":
					node.SetPositionType(yoga.PositionTypeRelative)
				case "static":
					node.SetPositionType(yoga.PositionTypeStatic)
				}
			}
		case "right":
			handleDimension(node, value, yoga.EdgeRight)
		case "start":
			handleDimension(node, value, yoga.EdgeStart)
		case "top":
			handleDimension(node, value, yoga.EdgeTop)
		case "width":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    node.SetWidthAuto,
				onPercent: node.SetWidthPercent,
				onNumber:  node.SetWidth,
			})
		case "direction":
			if value.IsString() {
				switch value.GetString() {
				case "inherit":
					node.SetStyleDirection(yoga.DirectionInherit)
				case "ltr":
					node.SetStyleDirection(yoga.DirectionLTR)
				case "rtl":
					node.SetStyleDirection(yoga.DirectionRTL)
				}
			}
		}
	}

	// Reset styles that are no longer present
	for key := range prevKeys {
		if _, exists := styleJSValueMap[key]; !exists {
			if setter, found := defaultValueSetters[key]; found {
				setter()
			}
		}
	}

	return newKeys
}

type dimensionValueHandlers struct {
	onAuto    func()
	onPercent func(percent float32)
	onNumber  func(num float32)
}

func handleDimensionValue(value JSValue, h dimensionValueHandlers) {
	if value.IsString() {
		strVal := value.GetString()
		if strVal == "auto" {
			h.onAuto()
		} else if len(strVal) > 0 && strVal[len(strVal)-1] == '%' {
			percentVal := parsePercent(strVal)
			h.onPercent(percentVal)
		}
	} else if value.IsNumber() {
		numVal := float32(value.GetNumber())
		h.onNumber(numVal)
	}
}

// handleDimension is a helper function to handle dimension values.
func handleDimension(node *yoga.YGNode, value JSValue, edge yoga.Edge) {
	if value.IsString() {
		strVal := value.GetString()
		if strVal == "auto" {
			switch edge {
			case yoga.EdgeAll:
				node.SetFlexBasisAuto()
			case yoga.EdgeLeft, yoga.EdgeTop, yoga.EdgeRight, yoga.EdgeBottom, yoga.EdgeStart, yoga.EdgeEnd:
				node.SetPosition(edge, 0) // Yoga does not support 'auto' for position, setting default
			}
		} else if len(strVal) > 0 && strVal[len(strVal)-1] == '%' {
			percentVal := parsePercent(strVal)
			switch edge {
			case yoga.EdgeAll:
				node.SetFlexBasisPercent(percentVal)
			default:
				node.SetPositionPercent(edge, percentVal)
			}
		}
	} else if value.IsNumber() {
		numVal := float32(value.GetNumber())
		switch edge {
		case yoga.EdgeAll:
			node.SetFlexBasis(numVal)
		default:
			node.SetPosition(edge, numVal)
		}
	}
}

// parsePercent is a helper function to parse percentage strings.
func parsePercent(percentStr string) float32 {
	var percentVal float32
	fmt.Sscanf(percentStr, "%f%%", &percentVal)
	return percentVal
}
