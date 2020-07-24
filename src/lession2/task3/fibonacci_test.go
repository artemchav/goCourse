package main

import "testing"
import "github.com/stretchr/testify/assert"


func TestFibonacci (t *testing.T){
	var fibonacci = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	assert.Equal(t, fibonacci, Fibonacci(15), "Some shit happens")
}
