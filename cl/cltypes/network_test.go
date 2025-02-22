package cltypes_test

import (
	"testing"

	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/stretchr/testify/require"

	"github.com/ledgerwatch/erigon/cl/cltypes"
	"github.com/ledgerwatch/erigon/cl/cltypes/ssz_utils"
)

var testMetadata = &cltypes.Metadata{
	SeqNumber: 99,
	Attnets:   69,
}

var testPing = &cltypes.Ping{
	Id: 420,
}

var testSingleRoot = &cltypes.SingleRoot{
	Root: libcommon.HexToHash("96"),
}

var testLcRangeRequest = &cltypes.LightClientUpdatesByRangeRequest{
	Period: 69,
	Count:  666,
}

var testBlockRangeRequest = &cltypes.BeaconBlocksByRangeRequest{
	StartSlot: 999,
	Count:     666,
}

var testStatus = &cltypes.Status{
	FinalizedEpoch: 666,
	HeadSlot:       94,
	HeadRoot:       libcommon.HexToHash("a"),
	FinalizedRoot:  libcommon.HexToHash("bbba"),
}

var testHeader = &cltypes.BeaconBlockHeader{
	Slot:          2,
	ProposerIndex: 24,
	ParentRoot:    libcommon.HexToHash("a"),
	Root:          libcommon.HexToHash("d"),
	BodyRoot:      libcommon.HexToHash("ad"),
}

var testLcUpdate = &cltypes.LightClientUpdate{
	AttestedHeader: testHeader,
	NextSyncCommitee: &cltypes.SyncCommittee{
		PubKeys: make([][48]byte, 512),
	},
	NextSyncCommitteeBranch: make([]libcommon.Hash, 5),
	FinalizedHeader:         testHeader,
	FinalityBranch:          make([]libcommon.Hash, 6),
	SyncAggregate:           &cltypes.SyncAggregate{},
	SignatureSlot:           294,
}

var testLcUpdateFinality = &cltypes.LightClientFinalityUpdate{
	AttestedHeader:  testHeader,
	FinalizedHeader: testHeader,
	FinalityBranch:  make([]libcommon.Hash, 6),
	SyncAggregate:   &cltypes.SyncAggregate{},
	SignatureSlot:   294,
}

var testLcUpdateOptimistic = &cltypes.LightClientOptimisticUpdate{
	AttestedHeader: testHeader,
	SyncAggregate:  &cltypes.SyncAggregate{},
	SignatureSlot:  294,
}

var testLcBootstrap = &cltypes.LightClientBootstrap{
	Header: testHeader,
	CurrentSyncCommittee: &cltypes.SyncCommittee{
		PubKeys: make([][48]byte, 512),
	},
	CurrentSyncCommitteeBranch: make([]libcommon.Hash, 5),
}

func TestMarshalNetworkTypes(t *testing.T) {
	cases := []ssz_utils.EncodableSSZ{
		testMetadata,
		testPing,
		testSingleRoot,
		testLcRangeRequest,
		testBlockRangeRequest,
		testStatus,
		testLcUpdate,
		testLcUpdateFinality,
		testLcUpdateOptimistic,
		testLcBootstrap,
	}

	unmarshalDestinations := []ssz_utils.EncodableSSZ{
		&cltypes.Metadata{},
		&cltypes.Ping{},
		&cltypes.SingleRoot{},
		&cltypes.LightClientUpdatesByRangeRequest{},
		&cltypes.BeaconBlocksByRangeRequest{},
		&cltypes.Status{},
		&cltypes.LightClientUpdate{},
		&cltypes.LightClientFinalityUpdate{},
		&cltypes.LightClientOptimisticUpdate{},
		&cltypes.LightClientBootstrap{},
	}
	for i, tc := range cases {
		marshalledBytes, err := tc.EncodeSSZ(nil)
		require.NoError(t, err)
		require.Equal(t, len(marshalledBytes), tc.EncodingSizeSSZ())
		require.NoError(t, unmarshalDestinations[i].DecodeSSZ(marshalledBytes))
		require.Equal(t, tc, unmarshalDestinations[i])
	}
}
