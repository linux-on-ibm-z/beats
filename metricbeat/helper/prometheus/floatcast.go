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
)

var (
	uint64NaN               = uint64(math.NaN())
	uint64PlusInf           = uint64(math.Inf(0))
	uint64MinusInf          = uint64(math.Inf(-1))
	haveValidUint64NaN      = true
	haveValidUint64PlusInf  = true
	haveValidUint64MinusInf = true
)

func init() {
	if uint64NaN == 0 {
		haveValidUint64NaN = false
	}
	if uint64PlusInf == 0 {
		haveValidUint64PlusInf = false
	}
	if uint64MinusInf == 0 {
		haveValidUint64MinusInf = false
	}
}

// IsNaN returns whether v is equal to the arch dependent
// value of uint64(math.NaN()) or false if that arch
// dependent value equals 0.
//
// This avoids filtering out all prometheus metric values of
// 0 on some platforms.
func IsNaN(v uint64) bool {
	return haveValidUint64NaN && v == uint64NaN
}

// IsInf returns whether v is equal to the arch dependent
// value of uint64(math.Inf(0)) or uint64(math.Inf(-1)).
// If either of those values is 0 on the current arch then
// false is used for that part of the comparision instead.
//
// This avoids filtering out all prometheus metric values of
// 0 on some platforms.
func IsInf(v uint64) bool {
	return (haveValidUint64PlusInf && v == uint64PlusInf) ||
		(haveValidUint64MinusInf && v == uint64MinusInf)
}
