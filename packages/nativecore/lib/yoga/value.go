package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// YGValue represents a Yoga value.
type YGValue struct {
	Value float32
	Unit  YGUnit
}

func convertCYGValueToGo(value C.YGValue) YGValue {
	return YGValue{
		Value: float32(value.value),
		Unit:  YGUnit(value.value),
	}
}

// YGUnit represents the unit enum in Yoga.
type YGUnit int

const (
	YGUnitUndefined YGUnit = C.YGUnitUndefined
	YGUnitPoint     YGUnit = C.YGUnitPoint
	YGUnitPercent   YGUnit = C.YGUnitPercent
	YGUnitAuto      YGUnit = C.YGUnitAuto
)

// GetUnit returns the unit of the YGValue.
func (v YGValue) GetUnit() YGUnit {
	return v.Unit
}

// GetValue returns the value of the YGValue.
func (v YGValue) GetValue() float32 {
	return v.Value
}

// Predefined constants for YGValue.
var (
	YGValueAuto      = YGValue{Value: float32(C.YGValueAuto.value), Unit: YGUnit(C.YGValueAuto.unit)}
	YGValueUndefined = YGValue{Value: float32(C.YGValueUndefined.value), Unit: YGUnit(C.YGValueUndefined.unit)}
	YGValueZero      = YGValue{Value: float32(C.YGValueZero.value), Unit: YGUnit(C.YGValueZero.unit)}
	// TODO: Figure out if this works lol
	// YGUndefined any = C.YGUndefined
)
