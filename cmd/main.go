package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/rickn42/study/group/grouping"
)

func main() {
	evaluationTest()
}

func evaluationTest() {
	personCnt := 50
	groupMemberCnt := 4
	turnCnt := 50

	fmt.Println("grouping by Grouping Function")
	evaluation(personCnt, groupMemberCnt, turnCnt, grouping.Grouping)

	fmt.Println("\n")
	fmt.Println("grouping by Only Random")
	evaluation(personCnt, groupMemberCnt, turnCnt, GroupingOnlyRandom)
}

func GroupingOnlyRandom(ps []*grouping.Person, memberCnt int, boringAmount float64) []*grouping.Group {
	groups := make([]*grouping.Group, int(math.Ceil(float64(len(ps))/float64(memberCnt))))
	for i := range groups {
		groups[i] = grouping.NewGroup(boringAmount, memberCnt)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})

	var groupIdx int
	for _, p := range ps {
		groupIdx %= len(groups)
		groups[groupIdx].AddMember(p)
		groupIdx++
	}

	return groups
}
