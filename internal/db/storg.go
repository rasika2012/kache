/*
 * MIT License
 *
 * Copyright (c)  2018 Kasun Vithanage
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package db

import (
	"sync"
	"time"
)

// DataType is a enum type which holds the type of data
type DataType int

const (
	// TypeString string type
	TypeString = DataType(iota + 1)

	// TypeList list type
	TypeList

	// TypeHashMap hashmap type
	TypeHashMap

	// TypeSet set type
	TypeSet
)

// DataNode holds data node which used to store in db
type DataNode struct {
	// Type of the data
	Type DataType

	// ExpiresAt for the data(unix timestamp)
	ExpiresAt int64

	// Value of the data
	Value interface{}

	// rw mux
	mux sync.RWMutex
}

// NewDataNode creates a new *DataNode
func NewDataNode(t DataType, exp int64, val interface{}) *DataNode {
	return &DataNode{Type: t, ExpiresAt: exp, Value: val}
}

// IsExpired will be true when node expired
func (node *DataNode) IsExpired() bool {
	node.mux.RLock()
	expired := node.ExpiresAt != -1 && node.ExpiresAt < time.Now().Unix()
	node.mux.RUnlock()
	return expired
}

// SetExpiration will set the expiration for the node
func (node *DataNode) SetExpiration(ttl int64) {
	node.mux.Lock()
	node.ExpiresAt = ttl
	node.mux.Unlock()
}
