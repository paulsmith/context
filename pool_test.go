// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context

import "testing"

func TestPool(t *testing.T) {
	p := newPool()
	if cap(p.pool) != PoolSize {
		t.Errorf("capacity of pool should equal PoolSize, was %v", cap(p.pool))
	}
	if len(p.pool) != 0 {
		t.Errorf("len of new pool should be 0, was %v", len(p.pool))
	}
	var maps []ctxt
	for i := 0; i < PoolSize; i++ {
		if len(p.pool) != 0 {
			t.Errorf("len(pool): expected 0, was %v", len(p.pool))
		}
		maps = append(maps, p.alloc())
	}
	for i, m := range maps {
		if len(p.pool) != i {
			t.Errorf("len(pool): expected %v, was %v", i, len(p.pool))
		}
		p.free(m)
	}
	// Allocate more than the size of the free pool
	for i := 0; i < PoolSize*2; i++ {
		maps = append(maps, p.alloc())
	}
	// Try to return more to the free pool than its size
	for i, m := range maps {
		p.free(m)
		if i >= PoolSize && len(p.pool) != PoolSize {
			t.Errorf("expected free pool to max out at %v, was %v", PoolSize, len(p.pool))
		}
	}
}
