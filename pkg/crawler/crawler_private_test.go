package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseObjectName(t *testing.T) {
	tests := []struct {
		name           string
		objName        string
		wantGroupID    string
		wantArtifactID string
		wantVersion    string
	}{
		{
			name:           "happy path",
			objName:        "maven2/org/json/json/20070829/json-20070829.jar.sha1",
			wantGroupID:    "org.json",
			wantArtifactID: "json",
			wantVersion:    "20070829",
		},
		{
			name:           "sad path (no groupID)",
			objName:        "maven2/org/20070829/org-20070829.jar.sha1",
			wantGroupID:    "",
			wantArtifactID: "",
			wantVersion:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroupID, gotArtifactID, gotVersion := parseObjectName(tt.objName)
			assert.Equal(t, tt.wantGroupID, gotGroupID)
			assert.Equal(t, tt.wantArtifactID, gotArtifactID)
			assert.Equal(t, tt.wantVersion, gotVersion)
		})
	}
}
