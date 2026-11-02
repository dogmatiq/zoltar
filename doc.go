// Package zoltar is a static analyzer that predicts the values in a Go program
// without running it.
//
// Create a [Predictor] for a program with [New], predict a call to a function
// with [Predictor.Predict], then get values from the resulting [Prediction],
// such as with [Prediction.ValueOf]. The examples show each step.
//
// Unless otherwise noted, no type in this package is safe for concurrent use by
// multiple goroutines.
package zoltar
