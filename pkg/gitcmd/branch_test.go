package gitcmd

import (
	"encoding/json"
	"testing"
)

func TestListBranches(t *testing.T) {

	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{"Success test case", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, _ := ListBranches()

			vaultJsonBytes, err := json.Marshal(got)
			if err != nil {
				t.Fatalf("Failed to parse json: %s \n", err)
			}

			t.Logf("Got: %s\n", string(vaultJsonBytes))

			if len(got) > 0 != tt.want {
				t.Errorf("ListBranches() got = %d, want > 0\n", len(got))
			}
		})
	}
}
