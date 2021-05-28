package zero_test

import (
	"testing"

	"github.com/youngbloood/zero"
)

var (
	str  = "str"
	str1 = "str1"
	str2 = "str2"
	str3 = "str3"
	str4 = "str4"
	str5 = "str5"
	str6 = "str6"

	str7  = "str7"
	str8  = "str8"
	str9  = "str9"
	str10 = "str10"
	str11 = "str11"
	str12 = "str12"
	str13 = "str13"

	int1 = 1000
	ptr1 *int
	ptr2 *int = &int1
)

type zeroV struct {
	Str *string
	Int int
}

type zeroT struct {
	Bool bool

	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	UInt   uint
	UInt8  uint8
	UInt16 uint16
	UInt32 uint32
	UInt63 uint64

	Str string

	Slice1 []int
	Slice2 []*string

	Array1 [3]int
	Array2 [3]*string

	Map1 map[string]int
	Map2 map[string]*string

	Chan1 chan interface{}
	Chan2 chan struct{}

	Float32 float32
	Float64 float64

	Complex64  complex64
	Complex128 complex128

	Interface1 interface{}
	Interface2 interface{}

	Struct1 zeroV
	Struct2 *zeroV
}

func newZeroT() *zeroT {
	return &zeroT{
		Bool:  true,
		Int:   1,
		Int8:  2,
		Int16: 3,
		Int32: 4,
		Int64: 5,

		UInt:   6,
		UInt8:  7,
		UInt16: 8,
		UInt32: 9,
		UInt63: 10,

		Str: str,

		Slice1: []int{11, 12, 13},
		Slice2: []*string{&str1, &str2, &str3},

		Array1: [3]int{14, 15, 16},
		Array2: [3]*string{&str4, &str5, &str6},

		Map1: map[string]int{
			str7: 17,
			str8: 18,
			str9: 19,
		},
		Map2: map[string]*string{
			str10: &str10,
			str11: &str11,
			str12: &str12,
		},

		Chan1: make(chan interface{}),
		Chan2: make(chan struct{}, 10),

		Float32: float32(20.0),
		Float64: float64(21.0),

		Complex64:  complex(22, 23),
		Complex128: complex(24, 25),

		Interface1: (interface{})(ptr1),
		Interface2: ptr2,
	}
}

func TestZeroStruct(t *testing.T) {
	z := newZeroT()

	v := zeroV{
		Str: &str13,
		Int: 26,
	}
	z.Struct1 = v
	z.Struct2 = &v

	t.Log("z1 = ", z)
	zero.Reset(*z)
	t.Log("z2 = ", z)
}

func TestZeroStructPtr(t *testing.T) {
	z := newZeroT()

	v := zeroV{
		Str: &str13,
		Int: 26,
	}
	z.Struct1 = v
	z.Struct2 = &v

	t.Log("z1 = ", z)
	zero.Reset(z)
	t.Log("z2 = ", z)
}

func TestZeroInt(t *testing.T) {
	var Int int = 1
	zero.Reset(Int)
	t.Log("Int = ", Int)
	zero.Reset(&Int)
	t.Log("Int = ", Int)

	var Int8 int8 = 8
	zero.Reset(Int8)
	t.Log("Int8 = ", Int8)
	zero.Reset(&Int8)
	t.Log("Int8 = ", Int8)

	var Int16 int16 = 16
	zero.Reset(Int16)
	t.Log("Int16 = ", Int16)
	zero.Reset(&Int16)
	t.Log("Int16 = ", Int16)

	var Int32 int32 = 32
	zero.Reset(Int32)
	t.Log("Int32 = ", Int32)
	zero.Reset(&Int32)
	t.Log("Int32 = ", Int32)

	var Int64 int64 = 32
	zero.Reset(Int64)
	t.Log("Int64 = ", Int64)
	zero.Reset(&Int64)
	t.Log("Int64 = ", Int64)

	var Uint uint = 1
	zero.Reset(Uint)
	t.Log("Uint = ", Uint)
	zero.Reset(&Uint)
	t.Log("Uint = ", Uint)

	var Uint8 uint8 = 6
	zero.Reset(Uint8)
	t.Log("Uint8 = ", Uint8)
	zero.Reset(&Uint8)
	t.Log("Uint8 = ", Uint8)

	var Uint16 uint16 = 16
	zero.Reset(Uint16)
	t.Log("Uint16 = ", Uint16)
	zero.Reset(&Uint16)
	t.Log("Uint16 = ", Uint16)

	var Uint32 uint32 = 32
	zero.Reset(Uint32)
	t.Log("Uint32 = ", Uint32)
	zero.Reset(&Uint32)
	t.Log("Uint32 = ", Uint32)

	var Uint64 uint64 = 64
	zero.Reset(Uint64)
	t.Log("Uint64 = ", Uint64)
	zero.Reset(&Uint64)
	t.Log("Uint64 = ", Uint64)

}

func TestZeroString(t *testing.T) {
	var Str = "str"
	zero.Reset(Str)
	t.Log("Str = ", Str)
	zero.Reset(&Str)
	t.Log("Str = ", Str)
	Str2 := &Str
	Str3 := &Str2
	zero.Reset(&Str3)
	t.Log("Str3 = ", Str3)
}

func TestZeroArray(t *testing.T) {
	var array1 = [3]int{2, 3, 4}
	zero.Reset(array1)
	t.Log("array1 = ", array1)
	zero.Reset(&array1)
	t.Log("array1 = ", array1)

	a, b, c := 5, 6, 7
	var array2 = [3]*int{&a, &b, &c}
	zero.Reset(array2)
	t.Log("array2 = ", array2)
	zero.Reset(&array2)
	t.Log("array2 = ", array2)
}

func TestZeroSlice(t *testing.T) {
	var slice1 = []int{2, 3, 4}
	zero.Reset(slice1)
	t.Log("slice1 = ", slice1)
	zero.Reset(&slice1)
	t.Log("slice1 = ", slice1)

	a, b, c := 5, 6, 7
	var slice2 = []*int{&a, &b, &c}
	zero.Reset(slice2)
	t.Log("slice2 = ", slice2)
	zero.Reset(&slice2)
	t.Log("slice2 = ", slice2)

	var slice3 = []int{8, 9, 10}
	slice33 := &slice3
	slice34 := &slice33
	zero.Reset(slice34)
	t.Log("slice34 = ", slice3)
}

func TestZeroMap(t *testing.T) {
	map1 := map[string]int{
		"aa": 11,
		"bb": 22,
	}
	zero.Reset(map1)
	t.Log("map1 = ", map1)
	zero.Reset(&map1)
	t.Log("map1 = ", map1)

}

func TestZeroFunction(t *testing.T) {
	func1 := func(v interface{}) {
		t.Logf("v = %+v\n", v)
	}
	zero.Reset(func1)
	func1(11)
	zero.Reset(&func1)
	if func1 == nil {
		t.Log("func1 is nil")
	} else {
		func1(22)
	}
}

func TestZeroComplex(t *testing.T) {
	var cmplx64 complex64 = complex(1, 2)
	zero.Reset(cmplx64)
	t.Log("cmplx64 = ", cmplx64)
	zero.Reset(&cmplx64)
	t.Log("cmplx64 = ", cmplx64)

	var cmplx128 complex64 = complex(3, 4)
	zero.Reset(cmplx128)
	t.Log("cmplx128 = ", cmplx128)
	zero.Reset(&cmplx128)
	t.Log("cmplx128 = ", cmplx128)
}
