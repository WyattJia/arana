/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package reduce

import (
	"time"
)

import (
	gxbig "github.com/dubbogo/gost/math/big"
)

type Reducer interface {
	Int64(prev, next int64) (int64, error)
	Float64(prev, next float64) (float64, error)
	Decimal(prev, next *gxbig.Decimal) (*gxbig.Decimal, error)
	Time(prev, next time.Time) (time.Time, error)
}

func Max() Reducer {
	return maxReducer{}
}

func Min() Reducer {
	return minReducer{}
}

func Sum() Reducer {
	return sumReducer{}
}
