// Copyright 2021, Chef.  All rights reserved.
// https://github.com/forkiss/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package chartbar

import "math"

// TODO(chef): 移入nazamath中

func isInteger(v float64) bool {
	return math.Ceil(v) == math.Floor(v)
}
