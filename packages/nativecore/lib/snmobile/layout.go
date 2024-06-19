package snmobile

import "nativecore/lib/yoga"

type LayoutMetrics struct {
	X                  float32
	Y                  float32
	Width              float32
	Height             float32
	ContentLeftInset   float32
	ContentTopInset    float32
	ContentRightInset  float32
	ContentBottomInset float32
}

func convertYogaLayoutMetricToSNLayoutMetrics(l yoga.LayoutMetrics) *LayoutMetrics {
	return &LayoutMetrics{
		X:                  l.Frame.Origin.X,
		Y:                  l.Frame.Origin.Y,
		Width:              l.Frame.Size.Width,
		Height:             l.Frame.Size.Height,
		ContentLeftInset:   l.ContentInsets.Left,
		ContentTopInset:    l.BorderWidth.Top,
		ContentRightInset:  l.ContentInsets.Right,
		ContentBottomInset: l.ContentInsets.Bottom,
	}
}
