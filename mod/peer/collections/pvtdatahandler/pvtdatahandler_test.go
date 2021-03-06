/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pvtdatahandler

import (
	"testing"

	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/stretchr/testify/assert"
)

func TestHandler_HandleGetPrivateData(t *testing.T) {
	h := New("testchannel", nil)

	config := &pb.StaticCollectionConfig{
		Name: "coll1",
	}

	value, handled, err := h.HandleGetPrivateData("tx1", "ns1", config, "key1")
	assert.NoError(t, err)
	assert.False(t, handled)
	assert.Nil(t, value)
}

func TestHandler_HandleGetPrivateDataMultipleKeys(t *testing.T) {
	h := New("testchannel", nil)

	config := &pb.StaticCollectionConfig{
		Name: "coll1",
	}

	value, handled, err := h.HandleGetPrivateDataMultipleKeys("tx1", "ns1", config, []string{"key1", "key2"})
	assert.NoError(t, err)
	assert.False(t, handled)
	assert.Nil(t, value)
}
