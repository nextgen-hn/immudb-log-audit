package immudb

import (
	"testing"

	"github.com/codenotary/immudb-log-audit/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWriteReadTypeParser is a regression test for the argument-order bug fixed in
// commit 673ddf045d0ab1c8f9088fd811bf8085be63b8d6.
//
// createKV / createSQL were calling WriteTypeParser(typ, collection, parser) instead
// of WriteTypeParser(collection, typ, parser); this test verifies the correct order
// is preserved end-to-end.
func TestWriteReadTypeParser(t *testing.T) {
	immuCli, _, containerID := utils.RunImmudbContainer()
	defer utils.StopImmudbContainer(containerID)

	cfg := NewConfigs(immuCli)

	const (
		collection = "testcollection"
		wantType   = "kv"
		wantParser = "pgauditjsonlog"
	)

	err := cfg.WriteTypeParser(collection, wantType, wantParser)
	require.NoError(t, err)

	gotType, gotParser, err := cfg.ReadTypeParser(collection)
	require.NoError(t, err)

	assert.Equal(t, wantType, gotType, "type must round-trip correctly (swapped args would store the collection name here)")
	assert.Equal(t, wantParser, gotParser, "parser must round-trip correctly")
}
