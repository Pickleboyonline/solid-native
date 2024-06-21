package snmobile

import (
	"fmt"
	"log"
	"nativecore/lib/yoga"
)

type Set map[string]struct{}

// UpdateNodeStyle updates the YGNode based on the JSValue style object.
func UpdateNodeStyle(node *yoga.YGNode, styleJSValueMap map[string]JSValue, prevKeys Set) Set {
	// Define default value functions
	// TODO: Determine if 0 is OK, and if i need to use YGUndefined or something
	defaultValueSetters := map[string]func(){
		"alignContent": func() { node.SetAlignContent(yoga.AlignStretch) },
		"alignItems":   func() { node.SetAlignItems(yoga.AlignStretch) },
		"alignSelf":    func() { node.SetAlignSelf(yoga.AlignAuto) },
		"aspectRatio":  func() { node.SetAspectRatio(yoga.YGUndefined) },
		"borderBottomWidth": func() {
			node.SetBorder(yoga.EdgeBottom,
				yoga.YGUndefined,
				// 0
			)
		},
		"borderEndWidth":    func() { node.SetBorder(yoga.EdgeEnd, yoga.YGUndefined) },
		"borderLeftWidth":   func() { node.SetBorder(yoga.EdgeLeft, yoga.YGUndefined) },
		"borderRightWidth":  func() { node.SetBorder(yoga.EdgeRight, yoga.YGUndefined) },
		"borderStartWidth":  func() { node.SetBorder(yoga.EdgeStart, yoga.YGUndefined) },
		"borderTopWidth":    func() { node.SetBorder(yoga.EdgeTop, yoga.YGUndefined) },
		"borderWidth":       func() { node.SetBorder(yoga.EdgeAll, yoga.YGUndefined) },
		"bottom":            func() { node.SetPosition(yoga.EdgeBottom, yoga.YGUndefined) },
		"display":           func() { node.SetDisplay(yoga.DisplayFlex) },
		"end":               func() { node.SetPosition(yoga.EdgeEnd, yoga.YGUndefined) },
		"flex":              func() { node.SetFlex(yoga.YGUndefined) },
		"flexBasis":         func() { node.SetFlexBasisAuto() },
		"flexDirection":     func() { node.SetFlexDirection(yoga.FlexDirectionColumn) },
		"flexGrow":          func() { node.SetFlexGrow(yoga.YGUndefined) },
		"flexShrink":        func() { node.SetFlexShrink(yoga.YGUndefined) },
		"flexWrap":          func() { node.SetFlexWrap(yoga.WrapNoWrap) },
		"gap":               func() { node.SetGap(yoga.GutterAll, yoga.YGUndefined) },
		"height":            func() { node.SetHeightAuto() },
		"justifyContent":    func() { node.SetJustifyContent(yoga.JustifyFlexStart) },
		"left":              func() { node.SetPosition(yoga.EdgeLeft, yoga.YGUndefined) },
		"margin":            func() { node.SetMargin(yoga.EdgeAll, yoga.YGUndefined) },
		"marginBottom":      func() { node.SetMargin(yoga.EdgeBottom, yoga.YGUndefined) },
		"marginEnd":         func() { node.SetMargin(yoga.EdgeEnd, yoga.YGUndefined) },
		"marginHorizontal":  func() { node.SetMargin(yoga.EdgeHorizontal, yoga.YGUndefined) },
		"marginLeft":        func() { node.SetMargin(yoga.EdgeLeft, yoga.YGUndefined) },
		"marginRight":       func() { node.SetMargin(yoga.EdgeRight, yoga.YGUndefined) },
		"marginStart":       func() { node.SetMargin(yoga.EdgeStart, yoga.YGUndefined) },
		"marginTop":         func() { node.SetMargin(yoga.EdgeTop, yoga.YGUndefined) },
		"marginVertical":    func() { node.SetMargin(yoga.EdgeVertical, yoga.YGUndefined) },
		"maxHeight":         func() { node.SetMaxHeight(yoga.YGValueUndefined.GetValue()) },
		"maxWidth":          func() { node.SetMaxWidth(yoga.YGValueUndefined.GetValue()) },
		"minHeight":         func() { node.SetMinHeight(yoga.YGValueUndefined.GetValue()) },
		"minWidth":          func() { node.SetMinWidth(yoga.YGValueUndefined.GetValue()) },
		"overflow":          func() { node.SetOverflow(yoga.OverflowVisible) },
		"padding":           func() { node.SetPadding(yoga.EdgeAll, yoga.YGUndefined) },
		"paddingBottom":     func() { node.SetPadding(yoga.EdgeBottom, yoga.YGUndefined) },
		"paddingEnd":        func() { node.SetPadding(yoga.EdgeEnd, yoga.YGUndefined) },
		"paddingHorizontal": func() { node.SetPadding(yoga.EdgeHorizontal, yoga.YGUndefined) },
		"paddingLeft":       func() { node.SetPadding(yoga.EdgeLeft, yoga.YGUndefined) },
		"paddingRight":      func() { node.SetPadding(yoga.EdgeRight, yoga.YGUndefined) },
		"paddingStart":      func() { node.SetPadding(yoga.EdgeStart, yoga.YGUndefined) },
		"paddingTop":        func() { node.SetPadding(yoga.EdgeTop, yoga.YGUndefined) },
		"paddingVertical":   func() { node.SetPadding(yoga.EdgeVertical, yoga.YGUndefined) },
		"position":          func() { node.SetPositionType(yoga.PositionTypeStatic) },
		"right":             func() { node.SetPosition(yoga.EdgeRight, yoga.YGUndefined) },
		"start":             func() { node.SetPosition(yoga.EdgeStart, yoga.YGUndefined) },
		"top":               func() { node.SetPosition(yoga.EdgeTop, yoga.YGUndefined) },
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
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPosition(yoga.EdgeBottom, 0) },
				onPercent: func(percent float32) { node.SetPositionPercent(yoga.EdgeBottom, percent) },
				onNumber:  func(num float32) { node.SetPosition(yoga.EdgeBottom, num) },
			})
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
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPosition(yoga.EdgeEnd, 0) },
				onPercent: func(percent float32) { node.SetPositionPercent(yoga.EdgeEnd, percent) },
				onNumber:  func(num float32) { node.SetPosition(yoga.EdgeEnd, num) },
			})
		case "flex":
			if value.IsNumber() {
				node.SetFlex(float32(value.GetNumber()))
			}
		case "flexBasis":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    node.SetFlexBasisAuto,
				onPercent: node.SetFlexBasisPercent,
				onNumber:  node.SetFlexBasis,
			})
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
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPosition(yoga.EdgeLeft, 0) },
				onPercent: func(percent float32) { node.SetPositionPercent(yoga.EdgeLeft, percent) },
				onNumber:  func(num float32) { node.SetPosition(yoga.EdgeLeft, num) },
			})
		case "margin":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto: func() {
					node.SetMarginAuto(yoga.EdgeAll)
				},
				onPercent: func(percent float32) {
					node.SetMarginPercent(yoga.EdgeAll, percent)
				},
				onNumber: func(num float32) {
					node.SetMargin(yoga.EdgeAll, num)
				},
			})
		case "marginBottom":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeBottom, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeBottom, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeBottom, num) },
			})
		case "marginEnd":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeEnd, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeEnd, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeEnd, num) },
			})
		case "marginHorizontal":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeHorizontal, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeHorizontal, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeHorizontal, num) },
			})
		case "marginLeft":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeLeft, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeLeft, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeLeft, num) },
			})
		case "marginRight":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeRight, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeRight, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeRight, num) },
			})
		case "marginStart":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeStart, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeStart, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeStart, num) },
			})
		case "marginTop":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeTop, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeTop, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeTop, num) },
			})
		case "marginVertical":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMargin(yoga.EdgeVertical, 0) },
				onPercent: func(percent float32) { node.SetMarginPercent(yoga.EdgeVertical, percent) },
				onNumber:  func(num float32) { node.SetMargin(yoga.EdgeVertical, num) },
			})
		case "maxHeight":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMaxHeight(yoga.YGValueAuto.GetValue()) },
				onPercent: node.SetMaxHeightPercent,
				onNumber:  node.SetMaxHeight,
			})
		case "maxWidth":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMaxWidth(yoga.YGValueAuto.GetValue()) },
				onPercent: node.SetMaxWidthPercent,
				onNumber:  node.SetMaxWidth,
			})
		case "minHeight":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMinHeight(yoga.YGValueAuto.GetValue()) },
				onPercent: node.SetMinHeightPercent,
				onNumber:  node.SetMinHeight,
			})
		case "minWidth":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetMinWidth(yoga.YGValueAuto.GetValue()) },
				onPercent: node.SetMinWidthPercent,
				onNumber:  node.SetMinWidth,
			})
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
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeAll, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeAll, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeAll, num) },
			})
		case "paddingBottom":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeBottom, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeBottom, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeBottom, num) },
			})
		case "paddingEnd":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeEnd, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeEnd, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeEnd, num) },
			})
		case "paddingHorizontal":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeHorizontal, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeHorizontal, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeHorizontal, num) },
			})
		case "paddingLeft":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeLeft, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeLeft, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeLeft, num) },
			})
		case "paddingRight":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeRight, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeRight, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeRight, num) },
			})
		case "paddingStart":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeStart, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeStart, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeStart, num) },
			})
		case "paddingTop":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeTop, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeTop, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeTop, num) },
			})
		case "paddingVertical":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPadding(yoga.EdgeVertical, 0) },
				onPercent: func(percent float32) { node.SetPaddingPercent(yoga.EdgeVertical, percent) },
				onNumber:  func(num float32) { node.SetPadding(yoga.EdgeVertical, num) },
			})
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
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPosition(yoga.EdgeRight, 0) },
				onPercent: func(percent float32) { node.SetPositionPercent(yoga.EdgeRight, percent) },
				onNumber:  func(num float32) { node.SetPosition(yoga.EdgeRight, num) },
			})
		case "start":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPosition(yoga.EdgeStart, 0) },
				onPercent: func(percent float32) { node.SetPositionPercent(yoga.EdgeStart, percent) },
				onNumber:  func(num float32) { node.SetPosition(yoga.EdgeStart, num) },
			})
		case "top":
			handleDimensionValue(value, dimensionValueHandlers{
				onAuto:    func() { node.SetPosition(yoga.EdgeTop, 0) },
				onPercent: func(percent float32) { node.SetPositionPercent(yoga.EdgeTop, percent) },
				onNumber:  func(num float32) { node.SetPosition(yoga.EdgeTop, num) },
			})
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

// parsePercent is a helper function to parse percentage strings.
func parsePercent(percentStr string) float32 {
	var percentVal float32
	fmt.Sscanf(percentStr, "%f%%", &percentVal)
	return percentVal
}
