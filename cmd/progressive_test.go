package cmd

import "testing"

func TestProgressiveRound(t *testing.T) {
	tests := []struct{
		name string
		f float32
		want int
	}{
		{
			name: "It rounds 0 to 0",
			f: 0,
			want: 0,
		},
		{
			name: "It rounds 0.1 to 0",
			f: 0.1,
			want: 0,
		},
		{
			name: "It rounds -0.1 to 0",
			f: -0.1,
			want: 0,
		},
		{
			name: "It rounds up for positives",
			f: 2.6,
			want: 3,
		},
		{
			name: "It rounds down for negatives",
			f: -2.6,
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := progressiveRound(tt.f); got != tt.want {
				t.Fatalf(`progressiveRound(%f) = %d, want %d`, tt.f, got, tt.want)
			}
		})
	}
}

func TestProgressiveBlend(t *testing.T) {
	tests := []struct{
		name string
		start int
		end int
		iteration int
		endPoint float32
		want int
	}{
		{
			name: "It is the start value at iteration 0",
			start: 100,
			end: 200,
			iteration: 0,
			endPoint: 1000,
			want: 100,
		},
		{
			name: "It is halfway between the start and end at 50 percent progress",
			start: 100,
			end: 200,
			iteration: 500,
			endPoint: 1000,
			want: 150,
		},
		{
			name: "It is the end value at 100 percent progress",
			start: 100,
			end: 200,
			iteration: 1000,
			endPoint: 1000,
			want: 200,
		},
		{
			name: "It can blend down",
			start: 200,
			end: 100,
			iteration: 750,
			endPoint: 1000,
			want: 125,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := progressiveBlend(tt.start, tt.end, tt.iteration, tt.endPoint); got != tt.want {
				t.Fatalf(`got = %v, want %v`, got, tt.want)
			}
		})
	}
}
