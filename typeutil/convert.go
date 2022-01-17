package typeutil

var TypeBoolInteger map[bool]uint16

func ConvertBoolInteger(enable bool) uint16 {
	return TypeBoolInteger[enable]
}

func ConvertIntegerBool(val uint16) bool {
	if val == 0 {
		return false
	}
	return true
}

func init() {
	TypeBoolInteger = make(map[bool]uint16)
	TypeBoolInteger[true] = 1
	TypeBoolInteger[false] = 0
}
