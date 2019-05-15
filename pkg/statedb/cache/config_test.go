/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package statecahe

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/require"
)

func TestStateCacheConfig(t *testing.T) {
	// test for default value
	require.Equal(t, 100, GetStateDataCacheSize())

	// set value and verify
	viper.Set("ledger.state.cache.size", 200)
	require.Equal(t, 200, GetStateDataCacheSize())
}
