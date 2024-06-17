package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// Edge represents the edges for margin and padding
type Edge int

const (
	EdgeLeft   Edge = C.YGEdgeLeft
	EdgeTop    Edge = C.YGEdgeTop
	EdgeRight  Edge = C.YGEdgeRight
	EdgeBottom Edge = C.YGEdgeBottom
	EdgeStart  Edge = C.YGEdgeStart
	EdgeEnd    Edge = C.YGEdgeEnd
	EdgeAll    Edge = C.YGEdgeAll
)

// Direction represents the layout direction
type Direction int

const (
	DirectionInherit Direction = C.YGDirectionInherit
	DirectionLTR     Direction = C.YGDirectionLTR
	DirectionRTL     Direction = C.YGDirectionRTL
)
