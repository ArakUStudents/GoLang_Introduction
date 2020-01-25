package main

import "fmt"

func typesSample() {
	var b bool
	b2 := true
	var b3, b4 = false, false
	const b5 = true
	b = b2 && b3 || b4 && b5
	b = !b

	var unsignedInt8 uint8 = 255 + 0 // same as 'byte'
	var unsignedInt16 uint16 = 65_535 + 0
	var unsignedInt32 uint32 = 4_294_967_259 + 0
	var unsignedInt64 uint64 = 18_446_744_073_709_551_615 + 0
	var unsignedInt uint = 18_446_744_073_709_551_615 + 0

	fmt.Printf("%T\t max:%d\n", unsignedInt8, unsignedInt8)
	fmt.Printf("%T\t max:%d\n", unsignedInt16, unsignedInt16)
	fmt.Printf("%T\t max:%d\n", unsignedInt32, unsignedInt32) // same as 'rune'
	fmt.Printf("%T\t max:%d\n", unsignedInt64, unsignedInt64)
	fmt.Printf("%T\t max:%d\n", unsignedInt, unsignedInt)

	var signedInt8 int8 = -128 + 127
	var signedInt16 int16 = -32_768 + 32_767
	var signedInt32 int32 = -2_147_483_648 + 2_147_483_647
	var signedInt64 int64 = -9_223_372_036_854_775_808 + 9_223_372_036_854_775_807
	var signedInt int = -9_223_372_036_854_775_808 + 9_223_372_036_854_775_807

	fmt.Printf("%T\t min:%d,\t\t\t max:%d\n", signedInt8, -128, 127)
	fmt.Printf("%T\t min:%d,\t\t\t max:%d\n", signedInt16, -32_768, 32_767)
	fmt.Printf("%T\t min:%d,\t\t max:%d\n", signedInt32, -2_147_483_648, 2_147_483_647)
	fmt.Printf("%T\t min:%d,\t max:%d\n", signedInt64, -9_223_372_036_854_775_808, 9_223_372_036_854_775_807)
	fmt.Printf("%T\t min:%d,\t max:%d\n", signedInt, -9_223_372_036_854_775_808, 9_223_372_036_854_775_807)

	var f32 float32 = 1.0
	var f64 float64 = float64(f32)
	f64 = f64 + 0.2

	var c64 complex64 = complex(12, 13)
	var c128 complex128 = complex128(c64)
	c128 = 12 + 123i + c128

	var str string
	str = "this is a text"
	println(str)
	str = str + " too."
	println(str)
	str = `this

	is \ string (O_o)`
	println(str)
}
