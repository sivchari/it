/*
Copyright 2025 sivchari.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package it_test

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sivchari/it"
)

func Example() {
	seq := it.NewSeq(slices.Values([]int{1, 2, 3, 4, 5})).
		Filter(
			func(v int) bool {
				return v%2 == 0
			}).
		Map(func(v int) int {
			return v * v
		}).
		Seq()

	for v := range seq {
		fmt.Println(v)
	}

	seq2 := it.NewSeq2(slices.All([]string{"a", "b", "c"})).
		Filter(
			func(i int, v string) bool {
				return v != "b"
			},
		).
		Map(
			func(i int, v string) (int, string) {
				return i * 2, strings.Repeat(v, 2)
			},
		).
		Seq2()
	for i, v := range seq2 {
		fmt.Println(i, v)
	}

	// Output:
	// 4
	// 16
	// 0 aa
	// 4 cc
}
