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

type Seq[V any] iter.Seq[V]

func NewSeq[V any](seq iter.Seq[V]) Seq[V] {
	return Seq[V](seq)
}

func (s Seq[V]) Seq() iter.Seq[V] {
	return iter.Seq[V](s)
}

func (s Seq[V]) Filter(f func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range s.Seq() {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func (s Seq[V]) Map(f func(V) V) Seq[V] {
	return func(yield func(V) bool) {
		for v := range s.Seq() {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func (s Seq[V]) Collect() []V {
	return slices.Collect(s.Seq())
}

type Seq2[K, V any] iter.Seq2[K, V]

func NewSeq2[K, V any](seq iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](seq)
}

func (s Seq2[K, V]) Seq2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}

func (s Seq2[K, V]) Filter(f func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range s.Seq2() {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

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
