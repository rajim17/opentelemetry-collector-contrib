// Copyright The OpenTelemetry Authors
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

package receivercreator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
)

func TestReceiverMap(t *testing.T) {
	rm := receiverMap{}
	assert.Equal(t, 0, rm.Size())

	r1 := &nopWithEndpointReceiver{}
	r2 := &nopWithEndpointReceiver{}
	r3 := &nopWithEndpointReceiver{}

	rm.Put("a", r1)
	assert.Equal(t, 1, rm.Size())

	rm.Put("a", r2)
	assert.Equal(t, 2, rm.Size())

	rm.Put("b", r3)
	assert.Equal(t, 3, rm.Size())

	assert.Equal(t, []component.Component{r1, r2}, rm.Get("a"))
	assert.Nil(t, rm.Get("missing"))

	rm.RemoveAll("missing")
	assert.Equal(t, 3, rm.Size())

	rm.RemoveAll("b")
	assert.Equal(t, 2, rm.Size())

	rm.RemoveAll("a")
	assert.Equal(t, 0, rm.Size())

	rm.Put("a", r1)
	rm.Put("b", r2)
	assert.Equal(t, 2, rm.Size())
	assert.Equal(t, []component.Component{r1, r2}, rm.Values())
}
