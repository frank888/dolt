// Copyright 2019 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file incorporates work covered by the following copyright and
// permission notice:
//
// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefInList(t *testing.T) {
	assert := assert.New(t)

	vs := newTestValueStore()

	l, err := NewList(context.Background(), vs)
	assert.NoError(err)
	r, err := NewRef(l, Format_7_18)
	assert.NoError(err)
	l, err = l.Edit().Append(r).List(context.Background())
	assert.NoError(err)
	r2, err := l.Get(context.Background(), 0)
	assert.NoError(err)
	assert.True(r.Equals(r2))
}

func TestRefInSet(t *testing.T) {
	assert := assert.New(t)

	vs := newTestValueStore()

	s, err := NewSet(context.Background(), vs)
	assert.NoError(err)
	r, err := NewRef(s, Format_7_18)
	assert.NoError(err)
	se, err := s.Edit().Insert(r)
	assert.NoError(err)
	s, err = se.Set(context.Background())
	assert.NoError(err)
	r2, err := s.First(context.Background())
	assert.NoError(err)
	assert.True(r.Equals(r2))
}

func TestRefInMap(t *testing.T) {
	assert := assert.New(t)

	vs := newTestValueStore()

	m, err := NewMap(context.Background(), vs)
	assert.NoError(err)
	r, err := NewRef(m, Format_7_18)
	assert.NoError(err)
	m, err = m.Edit().Set(Float(0), r).Set(r, Float(1)).Map(context.Background())
	assert.NoError(err)
	r2, ok, err := m.MaybeGet(context.Background(), Float(0))
	assert.NoError(err)
	assert.True(ok)
	assert.True(r.Equals(r2))

	i, ok, err := m.MaybeGet(context.Background(), r)
	assert.NoError(err)
	assert.True(ok)
	assert.Equal(int32(1), int32(i.(Float)))
}

func TestRefChunks(t *testing.T) {
	assert := assert.New(t)

	vs := newTestValueStore()

	l, err := NewList(context.Background(), vs)
	assert.NoError(err)
	r, err := NewRef(l, Format_7_18)
	assert.NoError(err)
	assert.Len(getChunks(r), 1)
	assert.Equal(r, getChunks(r)[0])
}
