/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privacyenabledstate

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/stateleveldb"
	"github.com/stretchr/testify/require"
)

func TestNewProvider(t *testing.T) {
	require.NotEmpty(t, NewVersionedDBProvider(stateleveldb.NewVersionedDBProvider()))

	require.Empty(t, NewVersionedDBProvider(nil))
}
