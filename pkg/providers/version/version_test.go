package version

import "testing"

func TestValidateK8sVersion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		version string
		wantErr bool
	}{
		{
			name:    "accepts min supported version",
			version: MinK8sVersion,
		},
		{
			name:    "accepts max supported version",
			version: MaxK8sVersion,
		},
		{
			name:    "accepts ack patch version suffix",
			version: "1.35.2-aliyun.1",
		},
		{
			name:    "rejects version below minimum",
			version: "1.27.9",
			wantErr: true,
		},
		{
			name:    "rejects version above maximum",
			version: "1.35.3",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateK8sVersion(tt.version)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error for version %q", tt.version)
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error for version %q, got %v", tt.version, err)
			}
		})
	}
}
