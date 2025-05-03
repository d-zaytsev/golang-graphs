package algo

import (
	"testing"
)

func TestDinicAlgo(t *testing.T) {
	tests := []struct {
		name         string
		builder      func() (*NetworkTaskData, error)
		expectedFlow float64
	}{
		{"Dinic Test 1", buildTaskData1, 10},
		{"Dinic Test 2", buildTaskData2, 10},
		{"Dinic Test 3", buildTaskData3, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testData, err := tt.builder()
			if err != nil {
				t.Fatal(err)
			}

			res, err := testData.Dinic()
			if err != nil {
				t.Fatal(err)
			}

			if res != tt.expectedFlow {
				t.Errorf("Incorrect result path: got %v, want %v", res, tt.expectedFlow)
			}
		})
	}
}
