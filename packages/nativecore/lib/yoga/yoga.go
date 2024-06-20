package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../../external/yoga
#cgo LDFLAGS: ${SRCDIR}/../../external/yoga/tests/build/yoga/libyogacore.a -lc++
#include <yoga/Yoga.h>
*/
import "C"

// Header is only needed to link againts tests. Currently no pure go flexbox sadly and
// cgo does not like to link against static libs. So just make sure to have the Yoga library
// imported for both iOS and Android so the C symbols are available

// For Go tests, the linker is setup annd will build for host
