package ptr

import (
	"testing"

	"github.com/ukfast/sdk-go/pkg/connection"
)

func TestByte(t *testing.T) {
	b := []byte{'a'}

	if *Byte(b[0]) != b[0] {
		t.Error("expected pointer to byte")
	}
}

func TestInt(t *testing.T) {
	i := 1

	if *Int(i) != i {
		t.Error("expected pointer to int")
	}
}

func TestInt8(t *testing.T) {
	var i int8
	i = 1

	if *Int8(i) != i {
		t.Error("expected pointer to int8")
	}
}

func TestInt16(t *testing.T) {
	var i int16
	i = 1

	if *Int16(i) != i {
		t.Error("expected pointer to int16")
	}
}

func TestInt32(t *testing.T) {
	var i int32
	i = 1

	if *Int32(i) != i {
		t.Error("expected pointer to int32")
	}
}

func TestInt64(t *testing.T) {
	var i int64
	i = 1

	if *Int64(i) != i {
		t.Error("expected pointer to int64")
	}
}

func TestUInt(t *testing.T) {
	var i uint
	i = 1

	if *UInt(i) != i {
		t.Error("expected pointer to uint")
	}
}

func TestUInt8(t *testing.T) {
	var i uint8
	i = 1

	if *UInt8(i) != i {
		t.Error("expected pointer to uint8")
	}
}

func TestUInt16(t *testing.T) {
	var i uint16
	i = 1

	if *UInt16(i) != i {
		t.Error("expected pointer to uint16")
	}
}

func TestUInt32(t *testing.T) {
	var i uint32
	i = 1

	if *UInt32(i) != i {
		t.Error("expected pointer to uint32")
	}
}

func TestUInt64(t *testing.T) {
	var i uint64
	i = 1

	if *UInt64(i) != i {
		t.Error("expected pointer to uint64")
	}
}

func TestUIntPtr(t *testing.T) {
	var i uintptr
	i = 1

	if *UIntPtr(i) != i {
		t.Error("expected pointer to uintptr")
	}
}

func TestFloat32(t *testing.T) {
	var i float32
	i = 1

	if *Float32(i) != i {
		t.Error("expected pointer to float32")
	}
}

func TestFloat64(t *testing.T) {
	var i float64
	i = 1

	if *Float64(i) != i {
		t.Error("expected pointer to float64")
	}
}

func TestComplex64(t *testing.T) {
	var i complex64
	i = 1

	if *Complex64(i) != i {
		t.Error("expected pointer to complex64")
	}
}

func TestComplex128(t *testing.T) {
	var i complex128
	i = 1

	if *Complex128(i) != i {
		t.Error("expected pointer to complex128")
	}
}

func TestString(t *testing.T) {
	s := "a"

	if *String(s) != s {
		t.Error("expected pointer to string")
	}
}

func TestBool(t *testing.T) {
	b := true

	if *Bool(b) != b {
		t.Error("expected pointer to bool")
	}
}

func TestDate(t *testing.T) {
	var d connection.Date = "2019-01-13"

	if *Date(d) != d {
		t.Error("expected pointer to Date")
	}
}

func TestDateTime(t *testing.T) {
	var d connection.DateTime = "2019-01-13T15:04:05-0700"

	if *DateTime(d) != d {
		t.Error("expected pointer to DateTime")
	}
}

func TestIPAddress(t *testing.T) {
	var d connection.IPAddress = "1.2.3.4"

	if *IPAddress(d) != d {
		t.Error("expected pointer to IPAddress")
	}
}

func TestToIntOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		i := 1

		if ToIntOrDefault(&i) != i {
			t.Error("expected int")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *int
		if ToIntOrDefault(i) != 0 {
			t.Error("expected int")
		}
	})
}

func TestToInt8OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i int8
		i = 1

		if ToInt8OrDefault(&i) != i {
			t.Error("expected int8")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *int8

		if ToInt8OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToInt16OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i int16
		i = 1

		if ToInt16OrDefault(&i) != i {
			t.Error("expected int16")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *int16

		if ToInt16OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToInt32OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i int32
		i = 1

		if ToInt32OrDefault(&i) != i {
			t.Error("expected int32")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *int32

		if ToInt32OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToInt64OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i int64
		i = 1

		if ToInt64OrDefault(&i) != i {
			t.Error("expected int64")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *int64

		if ToInt64OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToUIntOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i uint
		i = 1

		if ToUIntOrDefault(&i) != i {
			t.Error("expected uint")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *uint

		if ToUIntOrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToUInt8OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i uint8
		i = 1

		if ToUInt8OrDefault(&i) != i {
			t.Error("expected uint8")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *uint8

		if ToUInt8OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToUInt16OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i uint16
		i = 1

		if ToUInt16OrDefault(&i) != i {
			t.Error("expected uint16")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *uint16

		if ToUInt16OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToUInt32OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i uint32
		i = 1

		if ToUInt32OrDefault(&i) != i {
			t.Error("expected uint32")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *uint32

		if ToUInt32OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToUInt64OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i uint64
		i = 1

		if ToUInt64OrDefault(&i) != i {
			t.Error("expected uint64")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *uint64

		if ToUInt64OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToUIntPtrOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i uintptr
		i = 1

		if ToUIntPtrOrDefault(&i) != i {
			t.Error("expected uintptr")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *uintptr

		if ToUIntPtrOrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToFloat32OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i float32
		i = 1

		if ToFloat32OrDefault(&i) != i {
			t.Error("expected float32")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *float32

		if ToFloat32OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToFloat64OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i float64
		i = 1

		if ToFloat64OrDefault(&i) != i {
			t.Error("expected float64")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *float64

		if ToFloat64OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToComplex64OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i complex64
		i = 1

		if ToComplex64OrDefault(&i) != i {
			t.Error("expected complex64")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *complex64

		if ToComplex64OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToComplex128OrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var i complex128
		i = 1

		if ToComplex128OrDefault(&i) != i {
			t.Error("expected complex128")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var i *complex128

		if ToComplex128OrDefault(i) != 0 {
			t.Error("expected default of 0")
		}
	})
}

func TestToStringOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		s := "a"

		if ToStringOrDefault(&s) != s {
			t.Error("expected string")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var s *string

		if ToStringOrDefault(s) != "" {
			t.Error("expected default of ''")
		}
	})
}

func TestToBoolOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		b := true

		if ToBoolOrDefault(&b) != b {
			t.Error("expected bool")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var b *bool

		if ToBoolOrDefault(b) != false {
			t.Error("expected default value of false")
		}
	})
}

func TestToDateOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var d connection.Date = "2019-01-13"

		if ToDateOrDefault(&d) != d {
			t.Error("expected Date")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var d *connection.Date

		if ToDateOrDefault(d) != "" {
			t.Error("expected default of ''")
		}
	})
}

func TestToDateTimeOrDefault(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var d connection.DateTime = "2019-01-13T15:04:05-0700"

		if ToDateTimeOrDefault(&d) != d {
			t.Error("expected DateTime")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var d *connection.DateTime

		if ToDateTimeOrDefault(d) != "" {
			t.Error("expected default of ''")
		}
	})
}

func TestToIPAddressOrDefault_NotNil(t *testing.T) {
	t.Run("NotNil", func(t *testing.T) {
		var d connection.IPAddress = "1.2.3.4"

		if ToIPAddressOrDefault(&d) != d {
			t.Error("expected IPAddress")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var d *connection.IPAddress

		if ToIPAddressOrDefault(d) != "" {
			t.Error("expected default of ''")
		}
	})
}
