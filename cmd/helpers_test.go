package cmd

import "testing"

func TestAssertMinMax(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		max     int
		wantErr bool
	}{
		{
			name:    "min is less than max",
			min:     20,
			max:     30,
			wantErr: false,
		},
		{
			name:    "min is equal to max",
			min:     20,
			max:     20,
			wantErr: false,
		},
		{
			name:    "min is greater than max",
			min:     30,
			max:     20,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want string
			if tt.wantErr {
				want = "non-nil"
			} else {
				want = "nil"
			}
			if err := assertMinMax(tt.min, tt.max); (err != nil) != tt.wantErr {
				t.Fatalf(`err = %v, want %s`, err, want)
			}
		})
	}
}
