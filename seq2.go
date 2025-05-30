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

import "iter"

// Seq2 is a wrapper around iter.Seq2 that provides additional methods
// for some common operations like filtering, mapping, and collecting.
type Seq2[K, V any] iter.Seq2[K, V]

// NewSeq2 creates a new Seq2 from an existing iter.Seq2.
func NewSeq2[K, V any](seq iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](seq)
}

// Seq2 returns the underlying iter.Seq2.
func (s Seq2[K, V]) Seq2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}

// Filter returns a new Seq2 that yields only the key-value pairs
// for which the provided function returns true. It stops yielding
// pairs as soon as the yield function returns false.
func (s Seq2[K, V]) Filter(f func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range s.Seq2() {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// Map returns a new Seq2 that applies the provided function to each
// key-value pair of the original Seq2. It stops yielding pairs
// as soon as the yield function returns false.
func (s Seq2[K, V]) Map(f func(K, V) (K, V)) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range s.Seq2() {
			newK, newV := f(k, v)
			if !yield(newK, newV) {
				return
			}
		}
	}
}
