package main

import "fmt"

func main(){

	// boolean datatype
	var bool1 bool = true
	var bool2 bool = false
	_=bool2
	fmt.Println("Boolean datatype: ", bool1)


	// Integers
	var integer1 uint8 = 5 // 8-bit unsigned integer (0 to 255)
	var integer2 uint16 = 555 // 16-bit unsigned integer (0 to 65535)
	_=integer2 // to remove the error "integer2 declared but not used"
	var integer3 uint32 = 32732487 // 32-bit unsigned integer (0 to 4294967295)
	_=integer3
	var integer4 uint64 = 4757437578575	// 64-bit unsigned integer (0 to 18446744073709551615)
	_=integer4
	var integer5 int8 = -23 // 8-bit signed integer (-128 to 127)
	_=integer5
	var integer6 int16 = 345 // 16-bit signed integer (-32768 to 32767)
	_=integer6
	var integer7 int32 = -74457347 // 32-bit signed integer (-2147483648 to 2147483647)
	_=integer7
	var integer8 int64 = 475787438673 // 64-bit signed integer (-9223372036854775808 to 9223372036854775807)
	_=integer8
	var integer9 byte = 1 // equal to uint8
	_=integer9
	var integer10 rune = -3443 // equal to int32
	_=integer10
	var integer11 int = 213
	_=integer11
	var integer12 uint = 23
	_=integer12
	var integer13 uintptr = 34 // unsigned integer to keep pointer value
	_=integer13
	fmt.Println("Integer datatype: ", integer1)


	// Floating points
	var floating1 float32 = 32.65
	var floating2 float64 = 23443.766
	_=floating2
	fmt.Println("Floating point datatype: ", floating1)

	// Complex
	var complex1 complex64 = complex(5, 6)
	var complex2 complex128 = complex(7, 8)
	_=complex2
	fmt.Println("Complex datatype: ", complex1)

	// String
	var string1 = "Qarabag is Azerbaijan"
	fmt.Println("String datatype: ", string1)
}
