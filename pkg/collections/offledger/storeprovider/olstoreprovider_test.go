/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package storeprovider

import (
	"os"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/common/flogging"
	storeapi "github.com/hyperledger/fabric/extensions/collections/api/store"
	"github.com/hyperledger/fabric/extensions/testutil"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreProvider_OpenStore(t *testing.T) {
	channel1 := "channel1"
	channel2 := "channel2"

	f := New(testutil.TestLedgerConf())
	require.NotNil(t, f)

	s1, err := f.OpenStore(channel1)
	require.NotNil(t, s1)
	assert.NoError(t, err)
	assert.Equal(t, "*storeprovider.store", reflect.TypeOf(s1).String())
	assert.Equal(t, s1, f.StoreForChannel(channel1))

	s2, err := f.OpenStore(channel2)
	require.NotNil(t, s2)
	assert.NoError(t, err)
	assert.NotEqual(t, s1, s2)
	assert.Equal(t, s2, f.StoreForChannel(channel2))

	assert.NotPanics(t, func() {
		f.Close()
	})
}

func TestStoreProvider_WithDecorator(t *testing.T) {
	f := New(
		testutil.TestLedgerConf(),
		WithCollectionType(
			common.CollectionType_COL_DCAS,
			WithDecorator(func(key *storeapi.Key, value *storeapi.ExpiringValue) (*storeapi.Key, *storeapi.ExpiringValue, error) {
				return key, value, nil
			}),
			WithKeyDecorator(func(key *storeapi.Key) (*storeapi.Key, error) {
				return key, nil
			}),
		))
	require.NotNil(t, f)
	config, ok := f.collConfigs[common.CollectionType_COL_DCAS]
	assert.True(t, ok)
	assert.NotNil(t, config)
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	_, _, stop := testutil.SetupExtTestEnv()

	//set the logging level to DEBUG to test debug only code
	flogging.ActivateSpec("couchdb=debug")

	viper.Set("coll.offledger.cleanupExpired.Interval", "500ms")

	//run the tests
	code := m.Run()

	//stop couchdb
	stop()

	return code
}
