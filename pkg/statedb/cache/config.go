/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package statecahe

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/spf13/viper"
)

const (
	stateDataCacheSize        = "ledger.state.cache.size"
	defaultStateDataCacheSize = 100
)

var logger = flogging.MustGetLogger("statedb")

// GetStateDataCacheSize returns the state data cache size value. If not set, the default value is 100.
func GetStateDataCacheSize() int {
	cap := viper.GetInt(stateDataCacheSize)
	if cap <= 0 {
		cap = defaultStateDataCacheSize
	}
	logger.Infof("State data cache size is set to %v", cap)
	return cap
}
