package visitor

import (
	"fmt"
	"reflect"
)

func add(a, b interface{}) interface{} {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) + b.(float64)
		case reflect.Int64:
			return a.(int64) + b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) + float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) + b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func sub(a, b interface{}) interface{} {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) - b.(float64)
		case reflect.Int64:
			return a.(int64) - b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) - float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) - b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func mul(a, b interface{}) interface{} {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) * b.(float64)
		case reflect.Int64:
			return a.(int64) * b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) * float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) * b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func div(a, b interface{}) interface{} {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) / b.(float64)
		case reflect.Int64:
			return a.(int64) / b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) / float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) / b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func eq(a, b interface{}) bool {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) == b.(float64)
		case reflect.Int64:
			return a.(int64) == b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) == float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) == b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func nq(a, b interface{}) bool {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) != b.(float64)
		case reflect.Int64:
			return a.(int64) != b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) != float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) != b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func gt(a, b interface{}) bool {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) > b.(float64)
		case reflect.Int64:
			return a.(int64) > b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) > float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) > b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func lt(a, b interface{}) bool {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) < b.(float64)
		case reflect.Int64:
			return a.(int64) < b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) < float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) < b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func ge(a, b interface{}) bool {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) >= b.(float64)
		case reflect.Int64:
			return a.(int64) >= b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) >= float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) >= b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}

func le(a, b interface{}) bool {
	aT, bT := reflect.TypeOf(a), reflect.TypeOf(b)
	if aT.Kind() == bT.Kind() {
		switch t := aT.Kind(); t {
		case reflect.Float64:
			return a.(float64) <= b.(float64)
		case reflect.Int64:
			return a.(int64) <= b.(int64)
		default:
			panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
		}
	} else {
		if aT.Kind() == reflect.Float64 && bT.Kind() == reflect.Int64 {
			return a.(float64) <= float64(b.(int64))
		} else if aT.Kind() == reflect.Int64 && bT.Kind() == reflect.Float64 {
			return float64(a.(int64)) <= b.(float64)
		}
	}
	panic(fmt.Errorf("invalid sub type %v vs %v", aT.Kind(), bT.Kind()))
}
