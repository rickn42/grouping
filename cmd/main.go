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
		PersonCnt:       50,
		MemberOfGroup:   3,
		TurnCnt:         90,
		AlleviateAmount: 0.0,
		Verbose: 		 true,
	}

	fmt.Println("--- grouping by random shuffle ---")
	evaluation(grouping.GroupingByRandom, opt)

	fmt.Println()
	fmt.Println("--- grouping by regression ---")
	evaluation(grouping.GroupingByRegression, opt)
}
