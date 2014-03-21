package List

import (
	"DbxDev/Logger"
	"strconv"
	"testing"
)

func TestInitArrayTools(t *testing.T) {
	Logger.Init()

}
func TestAppendInt(t *testing.T) {
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
func TestAppendStr(t *testing.T) {

	a := make([]string, 5, 6)
	b := make([]string, 10)
	for i := 0; i < len(a); i++ {
		a[i] = "Elem" + strconv.Itoa(i)
	}
	for i := 0; i < len(b); i++ {
		b[i] = "Elem" + strconv.Itoa(10+i)
	}
	Logger.Debugf("a=%v , len=%v , cap=%v", a, len(a), cap(a))
	Logger.Debugf("b=%v , len=%v , cap=%v", b, len(b), cap(b))
	a = AppendString(a, "New value1")
	Logger.Debugf("%v , len=%v , cap=%v", a, len(a), cap(a))
	a = AppendString(a, "New value2")
	Logger.Debugf("%v , len=%v , cap=%v", a, len(a), cap(a))
	a = AppendString(a, b...)
	Logger.Debugf("%v , len=%v , cap=%v", a, len(a), cap(a))
	if cap(a) != 27 || len(a) != 17 {
		t.Errorf("Unexpected format for the array. Found len=%v, cap=%v. Expected len=%v, cap=%v.", len(a), cap(a), 17, 27)
	}
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
