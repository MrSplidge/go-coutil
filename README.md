# About

Provides a work pool with a limited number of goroutines (threads) for processing an input slice of work items. It is useful where a large number of inputs should be processed in parallel but the outputs should be processed serially.

# Example

See ```workpool_test.go``` for examples.

# Licence

Please see the included BSD 3-clause LICENSE file.