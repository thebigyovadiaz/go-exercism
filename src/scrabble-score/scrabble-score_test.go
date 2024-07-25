package scrabble_score

import "testing"

func TestScrabbleScore(t *testing.T) {
	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty word",
			args: args{word: ""},
			want: 0,
		},
		{
			name: "single letter",
			args: args{word: "a"},
			want: 1,
		},
		{
			name: "multiple letters",
			args: args{word: "cabbage"},
			want: 14,
		},
		{
			name: "mixed case",
			args: args{word: "caBbaGe"},
			want: 14,
		},
		{
			name: "all letters",
			args: args{word: "abcdefghijklmnopqrstuvwxyz"},
			want: 87,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScrabbleScore(tt.args.word); got != tt.want {
				t.Errorf("ScrabbleScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
