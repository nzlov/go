package utils

var MathUtils_NanoToSec float64 = 1 / 1000000000
var MathUtils_PI float64 = 3.1415927
var MathUtils_PI2 float64 = MathUtils_PI * 2

var MathUtils_SIN_BITS uint = 14 // 16KB. Adjust for accuracy.
var MathUtils_SIN_MASK int = ^(1 << MathUtils_SIN_BITS)
var MathUtils_SIN_COUNT int = MathUtils_SIN_MASK + 1

var MathUtils_radFull float64 = MathUtils_PI * 2
var MathUtils_degFull float64 = 360
var MathUtils_radToIndex float64 = float64(MathUtils_SIN_COUNT) / MathUtils_radFull
var MathUtils_degToIndex float64 = float64(MathUtils_SIN_COUNT) / MathUtils_degFull

var MathUtils_radiansToDegrees float64 = 180 / MathUtils_PI
var MathUtils_radDeg float64 = MathUtils_radiansToDegrees
var MathUtils_degreesToRadians float64 = MathUtils_PI / 180
var MathUtils_degRad float64 = MathUtils_degreesToRadians

func MathUtils_ToRadians(angdeg float64) float64 {
	return angdeg / 180.0 * MathUtils_PI
}
func MathUtils_Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
