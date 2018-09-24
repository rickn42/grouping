package main

import (
	"fmt"

	"github.com/rickn42/grouping"
)

func main() {
	evaluationTest()
}

func evaluationTest() {
	opt := EvaluationOption{
		PersonCnt:       40,
		MemberOfGroup:   4,
		TurnCnt:         60,
		AlleviateAmount: 0.0,
		Verbose: 		 true,
		ShowTurn: 		 true,
	}

	fmt.Println("--- grouping by random shuffle ---")
	evaluation(grouping.GroupingByRandom, opt)

	fmt.Println()
	fmt.Println("--- grouping by regression ---")
	evaluation(grouping.GroupingByRegression, opt)
}
