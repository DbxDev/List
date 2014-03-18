package List

import (
	"DbxDev/Logger"
	"strconv"
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
func TestShuffleInt(i *testing.T) {
	size := 10
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = i
	}
	Logger.Debugf("Original : %v", a)
	ShuffleInt(a)
	Logger.Debugf("Shuffled : %v", a)
}
func TestShuffleString(i *testing.T) {
	size := 10
	a := make([]string, size)
	for i := 0; i < size; i++ {
		a[i] = "Element-" + strconv.Itoa(i)
	}
	Logger.Debugf("Original : %v", a)
	ShuffleString(a)
	Logger.Debugf("Shuffled : %v", a)
}
