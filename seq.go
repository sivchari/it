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
package it

import (
	"iter"
	"slices"
)

// Seq is a wrapper around iter.Seq that provides additional methods
// for some common operations like filtering, mapping, and collecting.
type Seq[V any] iter.Seq[V]

// NewSeq creates a new Seq from an existing iter.Seq.
func NewSeq[V any](seq iter.Seq[V]) Seq[V] {
	return Seq[V](seq)
}

// Seq returns the underlying iter.Seq.
func (s Seq[V]) Seq() iter.Seq[V] {
	return iter.Seq[V](s)
}

// Filter returns a new Seq that yields only the elements for which
// the provided function returns true. It stops yielding elements
// as soon as the yield function returns false.
func (s Seq[V]) Filter(f func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range s.Seq() {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// Map returns a new Seq that applies the provided function to each
// element of the original Seq. It stops yielding elements as soon
// as the yield function returns false.
func (s Seq[V]) Map(f func(V) V) Seq[V] {
	return func(yield func(V) bool) {
		for v := range s.Seq() {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Collect collects all elements from the Seq into a slice
// and returns it. This is a convenience method that uses slices.Collect.
func (s Seq[V]) Collect() []V {
	return slices.Collect(s.Seq())
}
