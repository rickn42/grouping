package main

import (
	"fmt"

	"github.com/rickn42/study/group/grouping"
)

func main() {
	evaluationTest()
}

func evaluationTest() {
	opt := EvaluationOption{
		PersonCnt:       50,
		MemberOfGroup:   4,
		TurnCnt:         50,
		AlleviateAmount: 0.0,
	}

	fmt.Println("grouping by regression")
	evaluation(grouping.GroupingByRegression, opt)

	fmt.Println("")
	fmt.Println("grouping by random shuffle")
	evaluation(grouping.GroupingByRandom, opt)
}
