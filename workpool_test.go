package coutil

import (
	"reflect"
	"slices"
	"testing"
)

func TestWorkPool(t *testing.T) {
	// Test with a single work item and worker count of 1
	t.Run("SingleWorkItemAndWorkerCount1", func(t *testing.T) {
		input := []int{1}
		output := []int{}
		expectedOutput := []int{10}

		WorkPool(
			1,
			input,
			func(workItem int) int {
				return workItem * 10
			},
			func(resultItem int) {
				output = append(output, resultItem)
			})

		//Result order is not guaranteed
		slices.Sort(output)
		slices.Sort(expectedOutput)

		if !reflect.DeepEqual(output, expectedOutput) {
			t.Errorf("Expected %v, got %v", expectedOutput, output)
		}
	})

	// Test with multiple work items and worker count equal to the number of work items
	t.Run("MultipleWorkItemsAndWorkerCountEqual", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		output := []string{}
		expectedOutput := []string{"a!", "b!", "c!", "d!"}

		WorkPool(
			len(input),
			input,
			func(workItem string) string {
				return workItem + "!"
			},
			func(resultItem string) {
				output = append(output, resultItem)
			})

		//Result order is not guaranteed
		slices.Sort(output)
		slices.Sort(expectedOutput)

		if !reflect.DeepEqual(output, expectedOutput) {
			t.Errorf("Expected %v, got %v", expectedOutput, output)
		}
	})

	// Test with multiple work items and worker count larger than the number of work items
	t.Run("MultipleWorkItemsAndLargerWorkerCount", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		output := []string{}
		expectedOutput := []string{"a!", "b!", "c!", "d!"}

		WorkPool(
			8,
			input,
			func(workItem string) string {
				return workItem + "!"
			},
			func(resultItem string) {
				output = append(output, resultItem)
			})

		//Result order is not guaranteed
		slices.Sort(output)
		slices.Sort(expectedOutput)

		if !reflect.DeepEqual(output, expectedOutput) {
			t.Errorf("Expected %v, got %v", expectedOutput, output)
		}
	})

	// Test with a large number of work items and a small worker count
	t.Run("LargeNumberOfWorkItemsAndSmallWorkerCount", func(t *testing.T) {
		input := []int{}
		output := []int{}
		expectedOutput := []int{}

		for i := 0; i < 100000; i++ {
			input = append(input, i)
			expectedOutput = append(expectedOutput, i/5)
		}

		WorkPool(
			3,
			input,
			func(workItem int) int {
				return workItem / 5
			},
			func(resultItem int) {
				output = append(output, resultItem)
			})

		//Result order is not guaranteed
		slices.Sort(output)
		slices.Sort(expectedOutput)

		if !reflect.DeepEqual(output, expectedOutput) {
			t.Errorf("Expected %v..., got %v...", expectedOutput[:100], output[:100])
		}
	})
}
