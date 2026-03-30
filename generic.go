package main

func Add[T int32 | float64](a, b T) T { return a + b }
