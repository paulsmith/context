// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Provides a pool of maps of the context type to minimize allocations.

package context

// The maximum size of the free pool, can be set by clients.
var PoolSize int = 10

// pool is the type that holds the free pool of allocated maps.
type pool struct {
	pool []ctxt
}

// newPool returns a new pool object, but note that while the slice that holds
// the free pool is pre-allocated, the individuals maps themselves are not (but
// they are allocated transparently when first requested).
func newPool() *pool {
	p := new(pool)
	p.pool = make([]ctxt, 0, PoolSize)
	return p
}

// alloc returns an allocated context map, retrieving it from the free pool if
// possible.
func (p *pool) alloc() ctxt {
	var m ctxt
	if n := len(p.pool); n > 0 {
		m = p.pool[n-1]
		p.pool = p.pool[:n-1]
	} else {
		m = make(ctxt)
	}
	return m
}

// free returns the map to the free pool, if there is room.
func (p *pool) free(m ctxt) {
	if n := len(p.pool); n < PoolSize {
		for key, _ := range m {
			delete(m, key)
		}
		p.pool = append(p.pool, m)
	}
}
