package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// YGNode represents a Yoga node.
type YGNode struct {
	node C.YGNodeRef
}

// NewNode creates a new Yoga node.
func NewNode() *YGNode {
	return &YGNode{node: C.YGNodeNew()}
}

// Free frees the Yoga node.
func (n *YGNode) Free() {
	C.YGNodeFree(n.node)
}

// SetStyleDirection sets the direction style of the Yoga node.
func (n *YGNode) SetStyleDirection(direction Direction) {
	C.YGNodeStyleSetDirection(n.node, C.YGDirection(direction))
}

// GetStyleDirection gets the direction style of the Yoga node.
func (n *YGNode) GetStyleDirection() Direction {
	return Direction(C.YGNodeStyleGetDirection(n.node))
}

// SetFlexDirection sets the flex direction style of the Yoga node.
func (n *YGNode) SetFlexDirection(flexDirection FlexDirection) {
	C.YGNodeStyleSetFlexDirection(n.node, C.YGFlexDirection(flexDirection))
}

// GetFlexDirection gets the flex direction style of the Yoga node.
func (n *YGNode) GetFlexDirection() FlexDirection {
	return FlexDirection(C.YGNodeStyleGetFlexDirection(n.node))
}

// SetJustifyContent sets the justify content style of the Yoga node.
func (n *YGNode) SetJustifyContent(justify Justify) {
	C.YGNodeStyleSetJustifyContent(n.node, C.YGJustify(justify))
}

// GetJustifyContent gets the justify content style of the Yoga node.
func (n *YGNode) GetJustifyContent() Justify {
	return Justify(C.YGNodeStyleGetJustifyContent(n.node))
}

// SetAlignContent sets the align content style of the Yoga node.
func (n *YGNode) SetAlignContent(alignContent Align) {
	C.YGNodeStyleSetAlignContent(n.node, C.YGAlign(alignContent))
}

// GetAlignContent gets the align content style of the Yoga node.
func (n *YGNode) GetAlignContent() Align {
	return Align(C.YGNodeStyleGetAlignContent(n.node))
}

// SetAlignItems sets the align items style of the Yoga node.
func (n *YGNode) SetAlignItems(alignItems Align) {
	C.YGNodeStyleSetAlignItems(n.node, C.YGAlign(alignItems))
}

// GetAlignItems gets the align items style of the Yoga node.
func (n *YGNode) GetAlignItems() Align {
	return Align(C.YGNodeStyleGetAlignItems(n.node))
}

// SetAlignSelf sets the align self style of the Yoga node.
func (n *YGNode) SetAlignSelf(alignSelf Align) {
	C.YGNodeStyleSetAlignSelf(n.node, C.YGAlign(alignSelf))
}

// GetAlignSelf gets the align self style of the Yoga node.
func (n *YGNode) GetAlignSelf() Align {
	return Align(C.YGNodeStyleGetAlignSelf(n.node))
}

// SetPositionType sets the position type style of the Yoga node.
func (n *YGNode) SetPositionType(positionType PositionType) {
	C.YGNodeStyleSetPositionType(n.node, C.YGPositionType(positionType))
}

// GetPositionType gets the position type style of the Yoga node.
func (n *YGNode) GetPositionType() PositionType {
	return PositionType(C.YGNodeStyleGetPositionType(n.node))
}

// SetFlexWrap sets the flex wrap style of the Yoga node.
func (n *YGNode) SetFlexWrap(flexWrap Wrap) {
	C.YGNodeStyleSetFlexWrap(n.node, C.YGWrap(flexWrap))
}

// GetFlexWrap gets the flex wrap style of the Yoga node.
func (n *YGNode) GetFlexWrap() Wrap {
	return Wrap(C.YGNodeStyleGetFlexWrap(n.node))
}

// SetOverflow sets the overflow style of the Yoga node.
func (n *YGNode) SetOverflow(overflow Overflow) {
	C.YGNodeStyleSetOverflow(n.node, C.YGOverflow(overflow))
}

// GetOverflow gets the overflow style of the Yoga node.
func (n *YGNode) GetOverflow() Overflow {
	return Overflow(C.YGNodeStyleGetOverflow(n.node))
}

// SetDisplay sets the display style of the Yoga node.
func (n *YGNode) SetDisplay(display Display) {
	C.YGNodeStyleSetDisplay(n.node, C.YGDisplay(display))
}

// GetDisplay gets the display style of the Yoga node.
func (n *YGNode) GetDisplay() Display {
	return Display(C.YGNodeStyleGetDisplay(n.node))
}

// SetFlex sets the flex style of the Yoga node.
func (n *YGNode) SetFlex(flex float32) {
	C.YGNodeStyleSetFlex(n.node, C.float(flex))
}

// GetFlex gets the flex style of the Yoga node.
func (n *YGNode) GetFlex() float32 {
	return float32(C.YGNodeStyleGetFlex(n.node))
}

// SetFlexGrow sets the flex grow style of the Yoga node.
func (n *YGNode) SetFlexGrow(flexGrow float32) {
	C.YGNodeStyleSetFlexGrow(n.node, C.float(flexGrow))
}

// GetFlexGrow gets the flex grow style of the Yoga node.
func (n *YGNode) GetFlexGrow() float32 {
	return float32(C.YGNodeStyleGetFlexGrow(n.node))
}

// SetFlexShrink sets the flex shrink style of the Yoga node.
func (n *YGNode) SetFlexShrink(flexShrink float32) {
	C.YGNodeStyleSetFlexShrink(n.node, C.float(flexShrink))
}

// GetFlexShrink gets the flex shrink style of the Yoga node.
func (n *YGNode) GetFlexShrink() float32 {
	return float32(C.YGNodeStyleGetFlexShrink(n.node))
}

// SetFlexBasis sets the flex basis style of the Yoga node.
func (n *YGNode) SetFlexBasis(flexBasis float32) {
	C.YGNodeStyleSetFlexBasis(n.node, C.float(flexBasis))
}

// SetFlexBasisPercent sets the flex basis style of the Yoga node in percent.
func (n *YGNode) SetFlexBasisPercent(flexBasis float32) {
	C.YGNodeStyleSetFlexBasisPercent(n.node, C.float(flexBasis))
}

// SetFlexBasisAuto sets the flex basis style of the Yoga node to auto.
func (n *YGNode) SetFlexBasisAuto() {
	C.YGNodeStyleSetFlexBasisAuto(n.node)
}

// GetFlexBasis gets the flex basis style of the Yoga node.
func (n *YGNode) GetFlexBasis() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetFlexBasis(n.node))
}

// SetPosition sets the position style of the Yoga node for the given edge.
func (n *YGNode) SetPosition(edge Edge, position float32) {
	C.YGNodeStyleSetPosition(n.node, C.YGEdge(edge), C.float(position))
}

// SetPositionPercent sets the position style of the Yoga node in percent for the given edge.
func (n *YGNode) SetPositionPercent(edge Edge, position float32) {
	C.YGNodeStyleSetPositionPercent(n.node, C.YGEdge(edge), C.float(position))
}

// GetPosition gets the position style of the Yoga node for the given edge.
func (n *YGNode) GetPosition(edge Edge) YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetPosition(n.node, C.YGEdge(edge)))
}

// SetMargin sets the margin style of the Yoga node for the given edge.
func (n *YGNode) SetMargin(edge Edge, margin float32) {
	C.YGNodeStyleSetMargin(n.node, C.YGEdge(edge), C.float(margin))
}

// SetMarginPercent sets the margin style of the Yoga node in percent for the given edge.
func (n *YGNode) SetMarginPercent(edge Edge, margin float32) {
	C.YGNodeStyleSetMarginPercent(n.node, C.YGEdge(edge), C.float(margin))
}

// SetMarginAuto sets the margin style of the Yoga node to auto for the given edge.
func (n *YGNode) SetMarginAuto(edge Edge) {
	C.YGNodeStyleSetMarginAuto(n.node, C.YGEdge(edge))
}

// GetMargin gets the margin style of the Yoga node for the given edge.
func (n *YGNode) GetMargin(edge Edge) YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetMargin(n.node, C.YGEdge(edge)))
}

// SetPadding sets the padding style of the Yoga node for the given edge.
func (n *YGNode) SetPadding(edge Edge, padding float32) {
	C.YGNodeStyleSetPadding(n.node, C.YGEdge(edge), C.float(padding))
}

// SetPaddingPercent sets the padding style of the Yoga node in percent for the given edge.
func (n *YGNode) SetPaddingPercent(edge Edge, padding float32) {
	C.YGNodeStyleSetPaddingPercent(n.node, C.YGEdge(edge), C.float(padding))
}

// GetPadding gets the padding style of the Yoga node for the given edge.
func (n *YGNode) GetPadding(edge Edge) YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetPadding(n.node, C.YGEdge(edge)))
}

// SetBorder sets the border style of the Yoga node for the given edge.
func (n *YGNode) SetBorder(edge Edge, border float32) {
	C.YGNodeStyleSetBorder(n.node, C.YGEdge(edge), C.float(border))
}

// GetBorder gets the border style of the Yoga node for the given edge.
func (n *YGNode) GetBorder(edge Edge) float32 {
	return float32(C.YGNodeStyleGetBorder(n.node, C.YGEdge(edge)))
}

// SetGap sets the gap style of the Yoga node for the given gutter.
func (n *YGNode) SetGap(gutter Gutter, gapLength float32) {
	C.YGNodeStyleSetGap(n.node, C.YGGutter(gutter), C.float(gapLength))
}

// GetGap gets the gap style of the Yoga node for the given gutter.
func (n *YGNode) GetGap(gutter Gutter) float32 {
	return float32(C.YGNodeStyleGetGap(n.node, C.YGGutter(gutter)))
}

// SetWidth sets the width style of the Yoga node.
func (n *YGNode) SetWidth(width float32) {
	C.YGNodeStyleSetWidth(n.node, C.float(width))
}

// SetWidthPercent sets the width style of the Yoga node in percent.
func (n *YGNode) SetWidthPercent(width float32) {
	C.YGNodeStyleSetWidthPercent(n.node, C.float(width))
}

// SetWidthAuto sets the width style of the Yoga node to auto.
func (n *YGNode) SetWidthAuto() {
	C.YGNodeStyleSetWidthAuto(n.node)
}

// GetWidth gets the width style of the Yoga node.
func (n *YGNode) GetWidth() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetWidth(n.node))
}

// SetHeight sets the height style of the Yoga node.
func (n *YGNode) SetHeight(height float32) {
	C.YGNodeStyleSetHeight(n.node, C.float(height))
}

// SetHeightPercent sets the height style of the Yoga node in percent.
func (n *YGNode) SetHeightPercent(height float32) {
	C.YGNodeStyleSetHeightPercent(n.node, C.float(height))
}

// SetHeightAuto sets the height style of the Yoga node to auto.
func (n *YGNode) SetHeightAuto() {
	C.YGNodeStyleSetHeightAuto(n.node)
}

// GetHeight gets the height style of the Yoga node.
func (n *YGNode) GetHeight() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetHeight(n.node))
}

// SetMinWidth sets the min width style of the Yoga node.
func (n *YGNode) SetMinWidth(minWidth float32) {
	C.YGNodeStyleSetMinWidth(n.node, C.float(minWidth))
}

// SetMinWidthPercent sets the min width style of the Yoga node in percent.
func (n *YGNode) SetMinWidthPercent(minWidth float32) {
	C.YGNodeStyleSetMinWidthPercent(n.node, C.float(minWidth))
}

// GetMinWidth gets the min width style of the Yoga node.
func (n *YGNode) GetMinWidth() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetMinWidth(n.node))
}

// SetMinHeight sets the min height style of the Yoga node.
func (n *YGNode) SetMinHeight(minHeight float32) {
	C.YGNodeStyleSetMinHeight(n.node, C.float(minHeight))
}

// SetMinHeightPercent sets the min height style of the Yoga node in percent.
func (n *YGNode) SetMinHeightPercent(minHeight float32) {
	C.YGNodeStyleSetMinHeightPercent(n.node, C.float(minHeight))
}

// GetMinHeight gets the min height style of the Yoga node.
func (n *YGNode) GetMinHeight() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetMinHeight(n.node))
}

// SetMaxWidth sets the max width style of the Yoga node.
func (n *YGNode) SetMaxWidth(maxWidth float32) {
	C.YGNodeStyleSetMaxWidth(n.node, C.float(maxWidth))
}

// SetMaxWidthPercent sets the max width style of the Yoga node in percent.
func (n *YGNode) SetMaxWidthPercent(maxWidth float32) {
	C.YGNodeStyleSetMaxWidthPercent(n.node, C.float(maxWidth))
}

// GetMaxWidth gets the max width style of the Yoga node.
func (n *YGNode) GetMaxWidth() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetMaxWidth(n.node))
}

// SetMaxHeight sets the max height style of the Yoga node.
func (n *YGNode) SetMaxHeight(maxHeight float32) {
	C.YGNodeStyleSetMaxHeight(n.node, C.float(maxHeight))
}

// SetMaxHeightPercent sets the max height style of the Yoga node in percent.
func (n *YGNode) SetMaxHeightPercent(maxHeight float32) {
	C.YGNodeStyleSetMaxHeightPercent(n.node, C.float(maxHeight))
}

// GetMaxHeight gets the max height style of the Yoga node.
func (n *YGNode) GetMaxHeight() YGValue {
	return convertCYGValueToGo(C.YGNodeStyleGetMaxHeight(n.node))
}

// SetAspectRatio sets the aspect ratio style of the Yoga node.
func (n *YGNode) SetAspectRatio(aspectRatio float32) {
	C.YGNodeStyleSetAspectRatio(n.node, C.float(aspectRatio))
}

// GetAspectRatio gets the aspect ratio style of the Yoga node.
func (n *YGNode) GetAspectRatio() float32 {
	return float32(C.YGNodeStyleGetAspectRatio(n.node))
}
