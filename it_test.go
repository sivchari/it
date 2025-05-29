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
	"maps"
	"slices"
	"testing"

	"github.com/sivchari/it"
)

func TestSeqMap(t *testing.T) {
	mapped := it.NewSeq(slices.Values([]int{1, 2, 3, 4, 5})).
		Map(func(x int) int { return x * 2 }).
		Collect()
	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(mapped, expected) {
		t.Errorf("Expected %v but got %v", expected, mapped)
	}
}

func TestSeqFilter(t *testing.T) {
	filtered := it.NewSeq(slices.Values([]int{1, 2, 3, 4, 5})).
		Filter(func(x int) bool { return x%2 == 0 }).
		Collect()
	expected := []int{2, 4}
	if !slices.Equal(filtered, expected) {
		t.Errorf("Expected %v but got %v", expected, filtered)
	}
}

func TestSeqCollect(t *testing.T) {
	collected := it.NewSeq(slices.Values([]int{1, 2, 3, 4, 5})).Collect()
	expected := []int{1, 2, 3, 4, 5}
	if !slices.Equal(collected, expected) {
		t.Errorf("Expected %v but got %v", expected, collected)
	}
}

func TestSeq2Map(t *testing.T) {
	mapped := maps.Collect(
		it.NewSeq2(maps.All(map[string]int{"a": 1, "b": 2, "c": 3})).
			Map(func(k string, v int) (string, int) { return k, v * 2 }).
			Seq2())
	expected := map[string]int{"a": 2, "b": 4, "c": 6}
	if !maps.Equal(mapped, expected) {
		t.Errorf("Expected %v but got %v", expected, mapped)
	}
}

func TestSeq2Filter(t *testing.T) {
	filtered := maps.Collect(
		it.NewSeq2(maps.All(map[string]int{"a": 1, "b": 2, "c": 3})).
			Filter(func(k string, v int) bool { return v%2 == 0 }).
			Seq2())
	expected := map[string]int{"b": 2}
	if !maps.Equal(filtered, expected) {
		t.Errorf("Expected %v but got %v", expected, filtered)
	}
}
