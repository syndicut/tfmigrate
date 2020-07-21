package tfexec

import (
	"context"
	"testing"
)

func TestTerraformCLIRun(t *testing.T) {
	cases := []struct {
		desc         string
		mockCommands []*mockCommand
		args         []string
		want         string
		ok           bool
	}{
		{
			desc: "run terraform version",
			mockCommands: []*mockCommand{
				{
					args:     []string{"terraform", "version"},
					stdout:   "Terraform v0.12.28\n",
					exitCode: 0,
				},
			},
			args: []string{"version"},
			want: "Terraform v0.12.28\n",
			ok:   true,
		},
		{
			desc: "failed to run terraform version",
			mockCommands: []*mockCommand{
				{
					args:     []string{"terraform", "version"},
					exitCode: 1,
				},
			},
			args: []string{"version"},
			want: "",
			ok:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			e := NewMockExecutor(tc.mockCommands)
			terraformCLI := NewTerraformCLI(e)
			got, _, err := terraformCLI.Run(context.Background(), tc.args...)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected to return an error, but no error, got = %s", got)
			}
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
