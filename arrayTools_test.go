package List

import (
	"DbxDev/Logger"
	"testing"
)

func TestInitArrayTools(t *testing.T) {
	Logger.Init()

}
func TestAppend(t *testing.T) {
	t.SkipNow()
	a := make([]int, 5, 6)
	b := make([]int, 10)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	for i := 0; i < len(b); i++ {
		b[i] = 10 + i
	}
	Logger.Debugf("a=%v , len=%v , cap=%v", a, len(a), cap(a))
	Logger.Debugf("b=%v , len=%v , cap=%v", b, len(b), cap(b))
	a = AppendInt(a, 88)
	Logger.Debugf("%v , len=%v , cap=%v", a, len(a), cap(a))
	a = AppendInt(a, 99)
	Logger.Debugf("%v , len=%v , cap=%v", a, len(a), cap(a))
	a = AppendInt(a, b...)
	Logger.Debugf("%v , len=%v , cap=%v", a, len(a), cap(a))
}
