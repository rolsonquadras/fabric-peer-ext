/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package transientstore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStoreProvider(t *testing.T) {
	p, err := NewStoreProvider("")
	require.NoError(t, err)
	require.Empty(t, p)
}
