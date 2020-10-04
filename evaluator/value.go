package evaluator

// Value is a super type of IntValue and FunValue
type Value interface{}

// AsValue converts an empty interface to a Value
func AsValue(i interface{}) Value {
	return i.(Value)
}

// AsIntValue converts a Value to an IntValue
func AsIntValue(v Value) IntValue {
	return v.(IntValue)
}

// AsFunValue converts a Value to a FunValue
func AsFunValue(v Value) FunValue {
	return v.(FunValue)
}
