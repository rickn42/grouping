package main

import "github.com/rickn42/study/group/grouping"

func main() {
	evaluationTest()
}

func evaluationTest() {
	personCnt := 50
	groupMemberCnt := 5
	turnCnt := 50
	evaluation(personCnt, groupMemberCnt, turnCnt, grouping.Grouping)
}
