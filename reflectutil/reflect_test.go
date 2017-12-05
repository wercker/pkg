//-----------------------------------------------------------------------------
// Copyright (c) 2017 Oracle and/or its affiliates.  All rights reserved.
// This program is free software: you can modify it and/or redistribute it
// under the terms of:
//
// (i)  the Universal Permissive License v 1.0 or at your option, any
//      later version (http://oss.oracle.com/licenses/upl); and/or
//
// (ii) the Apache License v 2.0. (http://www.apache.org/licenses/LICENSE-2.0)
//-----------------------------------------------------------------------------

package reflectutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type noMethods struct{}
type someMethods struct{}
type privateMethods struct{}
type nonPointerMethods struct{}
type inheritsMethods struct{ *someMethods }

func (s *someMethods) Method1() {}
func (s *someMethods) Method2() {}
func (s *someMethods) Method3() {}

func (s *privateMethods) Method1() {}
func (s *privateMethods) method3() {}
func (s *privateMethods) method2() {}

func (s nonPointerMethods) Method1() {}
func (s nonPointerMethods) Method3() {}
func (s nonPointerMethods) Method2() {}

func Test_GetMethods(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{"nil", nil, []string{}},
		{"noMethods", noMethods{}, []string{}},
		{"noMethods_pointer", &noMethods{}, []string{}},
		{"someMethods", someMethods{}, []string{"Method1", "Method2", "Method3"}},
		{"someMethods_pointer", &someMethods{}, []string{"Method1", "Method2", "Method3"}},
		{"privateMethods", privateMethods{}, []string{"Method1"}},
		{"privateMethods_pointer", &privateMethods{}, []string{"Method1"}},
		{"nonPointerMethods", nonPointerMethods{}, []string{"Method1", "Method2", "Method3"}},
		{"nonPointerMethods_pointer", &nonPointerMethods{}, []string{"Method1", "Method2", "Method3"}},
		{"inheritsMethods", inheritsMethods{}, []string{"Method1", "Method2", "Method3"}},
		{"inheritsMethods_pointer", &inheritsMethods{}, []string{"Method1", "Method2", "Method3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetMethods(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
