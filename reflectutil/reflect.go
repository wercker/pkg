package reflectutil

import "reflect"

// GetMethods uses the reflect package to get the method names on defined on
// in.
func GetMethods(in interface{}) []string {
	value := reflect.TypeOf(in)

	numMethods := value.NumMethod()
	methods := make([]string, numMethods)
	for i := 0; i < numMethods; i++ {
		methods[i] = value.Method(i).Name
	}

	return methods
}
