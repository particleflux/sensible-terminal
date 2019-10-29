// Package sensible_terminal auto-detects installed terminal emulators
package sensible_terminal

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func resetEnv(vars []string) {
	for _, envString := range vars {
		env := strings.SplitN(envString, "=", 2)
		if err := os.Setenv(env[0], env[1]); err != nil {
			panic(err)
		}
	}
}

func TestAll(t *testing.T) {
	defer resetEnv(os.Environ())

	tests := []struct {
		name    string
		testDir string
		want    []string
		wantErr bool
	}{
		{
			name:    "None found",
			testDir: "testdata/empty",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "One found",
			testDir: "testdata/single",
			want:    []string{"xterm"},
			wantErr: false,
		},
		{
			name:    "Multiple found",
			testDir: "testdata/multiple",
			want:    []string{"urxvt", "rxvt"},
			wantErr: false,
		},
		{
			name:    "Non-executables",
			testDir: "testdata/non-executable",
			want:    []string{"mate-terminal"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Setenv("PATH", tt.testDir); err != nil {
				panic(err)
			}

			got, err := All()
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	defer resetEnv(os.Environ())

	tests := []struct {
		name    string
		testDir string
		want    string
		wantErr bool
	}{
		{
			name:    "None found",
			testDir: "testdata/empty",
			want:    "",
			wantErr: true,
		},
		{
			name:    "One found",
			testDir: "testdata/single",
			want:    "xterm",
			wantErr: false,
		},
		{
			name:    "Multiple found",
			testDir: "testdata/multiple",
			want:    "urxvt",
			wantErr: false,
		},
		{
			name:    "Non-executables",
			testDir: "testdata/non-executable",
			want:    "mate-terminal",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Setenv("PATH", tt.testDir); err != nil {
				panic(err)
			}

			got, err := First()
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}
