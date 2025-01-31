// Copyright 2021, Chef.  All rights reserved.
// https://github.com/forkiss/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazasync

import (
	"sync"
	"testing"

	"github.com/forkiss/naza/pkg/assert"

	"github.com/forkiss/naza/pkg/nazalog"
)

func TestCurGoroutineId(t *testing.T) {
	max := 5

	gid, err := CurGoroutineId()
	assert.Equal(t, nil, err)
	nazalog.Infof("> main. gid=%d", gid)
	var wg sync.WaitGroup
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func(ii int) {
			gid, err := CurGoroutineId()
			assert.Equal(t, nil, err)
			nazalog.Infof("> %d. gid=%d", ii, gid)
			wg.Done()
		}(i)
	}
	wg.Wait()
	gid, err = CurGoroutineId()
	assert.Equal(t, nil, err)
	nazalog.Infof("< main. gid=%d", gid)
}
