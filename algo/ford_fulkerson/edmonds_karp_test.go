package algo

import (
	"testing"
)

func TestEdmondsKarp(t *testing.T) {
	tests := buildTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testData, err := tt.builder()
			if err != nil {
				t.Fatal(err)
			}

			res, err := testData.EdmondsKarp()
			if err != nil {
				t.Fatal(err)
			}

			if res != tt.expectedFlow {
				t.Errorf("Incorrect result path: got %v, want %v", res, tt.expectedFlow)
			}
		})
	}
}
