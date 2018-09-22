package main

import (
	"fmt"

	"github.com/rickn42/study/group/grouping"
)

func main() {
	evaluationTest()
}

func evaluationTest() {
	personCnt := 50
	groupMemberCnt := 4
	turnCnt := 50

	fmt.Println("grouping by regression")
	evaluation(personCnt, groupMemberCnt, turnCnt, grouping.GroupingByRegression)

	fmt.Println("")
	fmt.Println("grouping by random shuffle")
	evaluation(personCnt, groupMemberCnt, turnCnt, grouping.GroupingByRandom)
}
