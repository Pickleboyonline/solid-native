package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// YGConfig represents a Yoga configuration.
type YGConfig struct {
	config C.YGConfigRef
}

// NewConfig creates a new Yoga configuration.
func NewConfig() *YGConfig {
	return &YGConfig{config: C.YGConfigNew()}
}

// Free frees the Yoga configuration.
func (c *YGConfig) Free() {
	C.YGConfigFree(c.config)
}

// SetUseWebDefaults sets whether to use web defaults for the Yoga configuration.
func (c *YGConfig) SetUseWebDefaults(enabled bool) {
	C.YGConfigSetUseWebDefaults(c.config, C.bool(enabled))
}

// GetUseWebDefaults gets whether to use web defaults for the Yoga configuration.
func (c *YGConfig) GetUseWebDefaults() bool {
	return bool(C.YGConfigGetUseWebDefaults(c.config))
}

// SetPointScaleFactor sets the point scale factor for the Yoga configuration.
func (c *YGConfig) SetPointScaleFactor(pixelsInPoint float32) {
	C.YGConfigSetPointScaleFactor(c.config, C.float(pixelsInPoint))
}

// GetPointScaleFactor gets the point scale factor for the Yoga configuration.
func (c *YGConfig) GetPointScaleFactor() float32 {
	return float32(C.YGConfigGetPointScaleFactor(c.config))
}

// SetLogger sets a custom logger for the Yoga configuration.
func (c *YGConfig) SetLogger(logger func(config *YGConfig, node *YGNode, level LogLevel, format string, args ...interface{})) {
	// Implement logger setting logic here.
}

// SetContext sets an arbitrary context pointer on the Yoga configuration.
func (c *YGConfig) SetContext(context interface{}) {
	// Implement context setting logic here.
}

// GetContext gets the currently set context of the Yoga configuration.
func (c *YGConfig) GetContext() interface{} {
	// Implement context getting logic here.
	return nil
}

// LogLevel represents the log level enum in Yoga.
type LogLevel int

const (
	LogLevelError   LogLevel = C.YGLogLevelError
	LogLevelWarn    LogLevel = C.YGLogLevelWarn
	LogLevelInfo    LogLevel = C.YGLogLevelInfo
	LogLevelDebug   LogLevel = C.YGLogLevelDebug
	LogLevelVerbose LogLevel = C.YGLogLevelVerbose
	LogLevelFatal   LogLevel = C.YGLogLevelFatal
)
