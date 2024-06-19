package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// RoundValueToPixelGrid rounds a point value to the nearest whole pixel.
func RoundValueToPixelGrid(value float64, pointScaleFactor float64, forceCeil bool, forceFloor bool) float64 {
	return float64(C.YGRoundValueToPixelGrid(C.double(value), C.double(pointScaleFactor), C.bool(forceCeil), C.bool(forceFloor)))
}
