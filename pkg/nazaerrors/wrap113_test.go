// Copyright 2021, Chef.  All rights reserved.
// https://github.com/forkiss/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

//go:build go1.13
// +build go1.13

package nazaerrors

import (
	"errors"
	"io"
	"testing"

	"github.com/forkiss/naza/pkg/assert"

	"github.com/forkiss/naza/pkg/nazalog"
)

func TestWrap(t *testing.T) {
	err := Wrap(io.EOF)
	nazalog.Debugf("%+v", err)
	assert.Equal(t, true, errors.Is(err, io.EOF))
	err = Wrap(err)
	nazalog.Debugf("%+v", err)
	assert.Equal(t, true, errors.Is(err, io.EOF))

	err = Wrap(io.EOF, "aaa")
	nazalog.Debugf("%+v", err)
}
