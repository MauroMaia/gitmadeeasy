package gitcmd

import (
	"encoding/json"
	"testing"
)


func TestListCommitIDs(t *testing.T) {

	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{"Success test case",true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ListCommitIDs()
			
			vaultJsonBytes, err := json.Marshal(got)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Got: %s\n", string(vaultJsonBytes))

			if len(got) > 0 != tt.want {
				t.Errorf("ListCommitIDs() got = %d, want > 0\n", len(got))
			}
		})
	}
}
