package main

import (
	"fmt"
	"testing"
)

func Test_isEndStart(t *testing.T) {
	tests := []struct {
		sentence string
		want     error
	}{
		{
			sentence: "hello world",
			want:     fmt.Errorf("last letter of hello is not the same as the first letter of world"),
		},
		{
			sentence: "The engine exceeds safe engine engineering guidelines, so owners salivated.",
			want:     nil,
		},
		{
			sentence: "The elephant tour reminds Steven; never ride elephants seated down.",
			want:     nil,
		},
		{
			sentence: "Brenda always surrenders so ordinary yellow will last to Orlando.",
			want:     nil,
		},
		{
			sentence: "White earplugs sometimes save everyone energy, Your Royal Lordship Prince.",
			want:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.sentence, func(t *testing.T) {
			got := isEndStart(tt.sentence)
			if got != nil && tt.want != nil {
				// Compare error messages instead of error instances
				if got.Error() != tt.want.Error() {
					t.Errorf("isEndStart() error = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				// One of them is nil and the other isn't
				t.Errorf("isEndStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
