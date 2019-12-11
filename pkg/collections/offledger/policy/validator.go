/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package policy

import (
	"time"

	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

// ValidateConfig validates the Off-Ledger Collection configuration
func ValidateConfig(config *pb.StaticCollectionConfig) error {
	if config.Type != pb.CollectionType_COL_OFFLEDGER && config.Type != pb.CollectionType_COL_DCAS {
		return errors.Errorf("unsupported off-ledger collection type: %s", config.Type)
	}

	if config.RequiredPeerCount > config.MaximumPeerCount {
		return errors.Errorf("collection-name: %s -- maximum peer count (%d) must be greater than or equal to required peer count (%d)", config.Name, config.MaximumPeerCount, config.RequiredPeerCount)
	}

	if config.BlockToLive != 0 {
		return errors.Errorf("collection-name: %s -- block-to-live not supported", config.Name)
	}

	if config.TimeToLive != "" {
		_, err := time.ParseDuration(config.TimeToLive)
		if err != nil {
			return errors.Errorf("collection-name: %s -- invalid time format for time to live: %s", config.Name, err)
		}
	}

	return nil
}
