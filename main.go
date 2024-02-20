package main

import (
	"errors"
	"fmt"
	"reflect"
)

func mergeSlices(a interface{}, b interface{}) (interface{}, error) {
	if a == nil && b == nil {
		return nil, errors.New("BOTH SLICES ARE NILL")
	}
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	var result []interface{}

	if reflect.TypeOf(a).Kind() == reflect.Slice {
		slice := reflect.ValueOf(a)
		for i := 0; i < slice.Len(); i++ {
			if reflect.ValueOf(slice.Index(i)).Kind() == reflect.Slice {
				nestedSlice, err := mergeSlices(result, slice.Index(i).Interface())
				if err != nil {
					panic(err)
				}
				result = nestedSlice.([]interface{})
			} else {
				result = append(result, slice.Index(i).Interface())
			}
		}
	} else {
		result = append(result, a)
	}

	if reflect.TypeOf(b).Kind() == reflect.Slice {
		slice := reflect.ValueOf(b)
		for i := 0; i < slice.Len(); i++ {
			if reflect.ValueOf(slice.Index(i).Interface()).Kind() == reflect.Slice {
				nestedSlice, err := mergeSlices(result, slice.Index(i).Interface())
				if err != nil {
					panic(err)
				}
				result = nestedSlice.([]interface{})
			} else {
				result = append(result, slice.Index(i).Interface())
			}
		}
	} else {
		result = append(result, b)
	}

	return result, nil

}

func main() {
	a := []interface{}{[]string{"chetan"}, 9, 10, []int{8, 9}, []interface{}{"a", "xfy", 20.0}}
	b := []interface{}{1, 5, []int{3, 4}, []interface{}{true, "abh", 2}}

	merged, err := mergeSlices(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Merged slice is ", merged)
}
