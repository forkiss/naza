// Copyright 2019, Chef.  All rights reserved.
// https://github.com/forkiss/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package slicebytepool

import (
	"testing"

	"github.com/forkiss/naza/pkg/assert"
)

func TestNewSharedSliceByte(t *testing.T) {
	pool := NewSliceBytePool(StrategyMultiSlicePoolBucket)
	assert.IsNotNil(t, pool)
	ssb := NewSharedSliceByte(1000, WithPool(pool))
	assert.IsNotNil(t, ssb)
	ssb2 := ssb.Ref()
	assert.IsNotNil(t, ssb2)

	ssb.ReleaseIfNeeded()
	ssb2.ReleaseIfNeeded()
}

func TestWrapSharedSliceByte(t *testing.T) {
	b := make([]byte, 1000)
	pool := NewSliceBytePool(StrategyMultiSlicePoolBucket)
	assert.IsNotNil(t, pool)
	ssb := WrapSharedSliceByte(b, WithPool(pool))
	assert.IsNotNil(t, ssb)
}
