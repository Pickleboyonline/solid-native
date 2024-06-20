package yoga

// Rect represents a rectangle with origin and size.
type Rect struct {
	Origin Point
	Size   Size
}

// Point represents a point in 2D space.
type Point struct {
	X, Y float32
}

// Size represents the dimensions of a rectangle.
type Size struct {
	Width, Height float32
}

// EdgeInsets represents the padding or border widths.
type EdgeInsets struct {
	Top, Left, Bottom, Right float32
}

// LayoutMetrics represents the layout metrics for a Yoga node.
type LayoutMetrics struct {
	Frame         Rect
	BorderWidth   EdgeInsets
	ContentFrame  Rect
	ContentInsets EdgeInsets
	DisplayType   Display
	Direction     Direction
}

// UIEdgeInsetsInsetRect insets a rectangle by the given edge insets.
func uiEdgeInsetsInsetRect(rect Rect, insets EdgeInsets) Rect {
	return Rect{
		Origin: Point{
			X: rect.Origin.X + insets.Left,
			Y: rect.Origin.Y + insets.Top,
		},
		Size: Size{
			Width:  rect.Size.Width - insets.Left - insets.Right,
			Height: rect.Size.Height - insets.Top - insets.Bottom,
		},
	}
}

// ConvertYogaNodeToLayoutMetrics converts a Yoga node to layout metrics.
func NewLayoutMetrics(node *YGNode) LayoutMetrics {
	var layoutMetrics LayoutMetrics

	frame := Rect{
		Origin: Point{
			X: (node.GetLayoutLeft()),
			Y: (node.GetLayoutTop()),
		},
		Size: Size{
			Width:  (node.GetLayoutWidth()),
			Height: (node.GetLayoutHeight()),
		},
	}

	padding := EdgeInsets{
		Top:    (node.GetLayoutPadding(EdgeTop)),
		Left:   (node.GetLayoutPadding(EdgeLeft)),
		Bottom: (node.GetLayoutPadding(EdgeBottom)),
		Right:  (node.GetLayoutPadding(EdgeRight)),
	}

	borderWidth := EdgeInsets{
		Top:    (node.GetLayoutBorder(EdgeTop)),
		Left:   (node.GetLayoutBorder(EdgeLeft)),
		Bottom: (node.GetLayoutBorder(EdgeBottom)),
		Right:  (node.GetLayoutBorder(EdgeRight)),
	}

	compoundInsets := EdgeInsets{
		Top:    borderWidth.Top + padding.Top,
		Left:   borderWidth.Left + padding.Left,
		Bottom: borderWidth.Bottom + padding.Bottom,
		Right:  borderWidth.Right + padding.Right,
	}

	bounds := Rect{
		Origin: Point{X: 0, Y: 0},
		Size:   frame.Size,
	}
	contentFrame := uiEdgeInsetsInsetRect(bounds, compoundInsets)

	contentInsets := EdgeInsets{
		Left:   borderWidth.Left + node.GetLayoutPadding(EdgeLeft),
		Top:    borderWidth.Top + node.GetLayoutPadding(EdgeTop),
		Right:  borderWidth.Right + node.GetLayoutPadding(EdgeRight),
		Bottom: borderWidth.Left + node.GetLayoutPadding(EdgeBottom),
	}

	layoutMetrics.Frame = frame
	layoutMetrics.BorderWidth = borderWidth
	layoutMetrics.ContentFrame = contentFrame
	layoutMetrics.DisplayType = node.GetDisplay()
	layoutMetrics.Direction = node.GetLayoutDirection()
	layoutMetrics.ContentInsets = contentInsets

	return layoutMetrics
}
