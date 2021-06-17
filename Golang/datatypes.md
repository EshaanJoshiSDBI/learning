var keyword for variable definition
`var` variable_name `datatype` = data
# Types
Integers
### u stands for unsigned which means no negative values
- uint8.byte(0-255)
- uint16(0-65535)
- uint32(0 - 4294967295)
- uint64(0 - 18446744073709551615)
- int8.byte(-128 to 127)
- int16(-32768 to 32767)
- int32/rune(-2147483648 to 2147483647)
- int64(-9223372036854775808 to 9223372036854775807)
### If I use int64 to store a single digit, it will still take 64 bits
- uint(32 or 64 bits)
	- based on os takes 32/64
- int(same as uint but can take negative values)
- uintptr(an unsigned variable to store the uninterpreted bits of a pointer value)
## Floating
### If I use float64 to store a single digit, it will still take 64 bits
- float32
- float64
## Complex
### If I use complex128 to store a single digit, it will still take 128 bits
- complex64
- complex128
## Strings
- "Hello world"
## Boolean
- true
- false
# If we don't use a variable then that's an error
## for variable type `%T`
## a variable cannot be redefined
