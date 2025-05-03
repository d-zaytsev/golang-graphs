package algo

import (
	"testing"
)

func TestFordFulkerson(t *testing.T) {
	tests := []struct {
		name         string
		builder      func() (*NetworkTaskData, error)
		expectedFlow float64
	}{
		{"Test 1", buildTaskData1, 10},
		{"Test 2", buildTaskData2, 10},
		{"Test 3", buildTaskData3, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testData, err := tt.builder()
			if err != nil {
				t.Fatal(err)
			}

			res, err := testData.FordFulkerson()
			if err != nil {
				t.Fatal(err)
			}

			if res != tt.expectedFlow {
				t.Errorf("Incorrect result path: got %v, want %v", res, tt.expectedFlow)
			}
		})
	}
}
