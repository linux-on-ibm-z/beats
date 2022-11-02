// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package prometheus

import (
	"math"
	"runtime"
	"testing"
)

type testCase struct {
	f        func(uint64) bool
	v        uint64
	expected bool
	name     string
}

// On amd64, uint64(math.Inf(1)) == uint64(math.Inf(-1))
// == uint64(math.NaN()) == 0x80000000_00000000.
// All of the above are considered valid.
var amd64NaNTestCases = []testCase{
	{IsNaN, uint64(math.Inf(1)), true, "IsNaN(uint64(+Inf))"},
	{IsNaN, uint64(math.Inf(-1)), true, "IsNaN(uint64(-Inf))"},
	{IsNaN, uint64(math.NaN()), true, "IsNaN(uint64(NaN))"},
	{IsNaN, uint64(0), false, "IsNaN(0)"},
	{IsNaN, uint64(math.MaxUint64), false, "IsNaN(MaxUint64)"},
	{IsNaN, uint64(math.MaxInt64), false, "IsNaN(MaxInt64)"},
	{IsInf, uint64(math.Inf(1)), true, "IsInf(uint64(+Inf))"},
	{IsInf, uint64(math.Inf(-1)), true, "IsInf(uint64(-Inf))"},
	{IsInf, uint64(math.NaN()), true, "IsInf(uint64(NaN))"},
	{IsInf, uint64(0), false, "IsInf(0)"},
	{IsInf, uint64(math.MaxUint64), false, "IsInf(MaxUint64)"},
	{IsInf, uint64(math.MaxInt64), false, "IsInf(MaxInt64)"},
}

// On s390x, uint64(math.Inf(1)) == math.MaxUint64
// and uint64(math.Inf(-1)) == uint64(math.NaN()) == 0.
// Only uint64(math.Inf(1)) is considered valid.
var s390xNaNTestCases = []testCase{
	{IsNaN, uint64(math.Inf(1)), false, "IsNaN(uint64(+Inf))"},
	{IsNaN, uint64(math.Inf(-1)), false, "IsNaN(uint64(-Inf))"},
	{IsNaN, uint64(math.NaN()), false, "IsNaN(uint64(NaN))"},
	{IsNaN, uint64(0), false, "IsNaN(0)"},
	{IsNaN, uint64(math.MaxUint64), false, "IsNaN(MaxUint64)"},
	{IsNaN, uint64(math.MaxInt64), false, "IsNaN(MaxInt64)"},
	{IsInf, uint64(math.Inf(1)), true, "IsInf(uint64(+Inf))"},
	{IsInf, uint64(math.Inf(-1)), false, "IsInf(uint64(-Inf))"},
	{IsInf, uint64(math.NaN()), false, "IsInf(uint64(NaN))"},
	{IsInf, uint64(0), false, "IsInf(0)"},
	{IsInf, uint64(math.MaxUint64), true, "IsInf(MaxUint64)"},
	{IsInf, uint64(math.MaxInt64), false, "IsInf(MaxInt64)"},
}

func TestIsNaNIsInf(t *testing.T) {
	testCases := amd64NaNTestCases
	switch runtime.GOARCH {
	case "s390x":
		testCases = s390xNaNTestCases
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.f(tt.v)
			if r != tt.expected {
				t.Errorf("%s\n  want: %v\n  got:  %v", tt.name, tt.expected, r)
			}
		})
	}
}
