package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cuijxin/mysql-operator/pkg/cluster"
)

func TestGetReplicationGroupSeeds(t *testing.T) {
	testCases := []struct {
		seeds    string
		pod      *cluster.Instance
		expected []string
	}{
		{
			seeds:    "server-1-0:1234,server-1-1:1234",
			pod:      cluster.NewInstance("", "", "server-1", 0, -1, false),
			expected: []string{"server-1-0:1234", "server-1-1:1234"},
		}, {
			seeds:    "server-1-1:1234,server-1-0:1234",
			pod:      cluster.NewInstance("", "", "server-1", 0, -1, false),
			expected: []string{"server-1-0:1234", "server-1-1:1234"},
		}, {
			seeds:    "server-1-0:1234,server-1-1:1234",
			pod:      cluster.NewInstance("", "", "server-2", 0, -1, false),
			expected: []string{"server-1-0:1234", "server-1-1:1234"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.seeds, func(t *testing.T) {
			output, err := getReplicationGroupSeeds(tt.seeds, tt.pod)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, output)
		})
	}
}
