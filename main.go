package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	totalMemberCount := 20
	groupMemberCount := 4
	rotateCount := totalMemberCount / (groupMemberCount - 1)
	alleviateScore := 2
	fmt.Println(totalMemberCount, groupMemberCount, rotateCount, alleviateScore)

	// generate persons
	var members []*Person
	for i := 0; i < totalMemberCount; i++ {
		members = append(members, NewPerson(i))
	}

	turn := 100
	for i := 0; i < turn; i++ {
		groups := Grouping(members, groupMemberCount, rotateCount)
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].BoringScoreSum > groups[j].BoringScoreSum
		})
		fmt.Println(groups)

		for _, m := range members {
			m.AlleviateBoringScore(alleviateScore)
		}
	}

	fmt.Println(members[0].Boring)
}

func Grouping(ps []*Person, memberCount int, boringScore int) (groups []*Group) {
	// make empty groups
	for i := 0; i < len(ps)/memberCount; i++ {
		groups = append(groups, NewGroup(boringScore))
	}

	// boring person first!
	sort.Slice(ps, func(i, j int) bool {
		scoreI := ps[i].CurBoringScore()
		scoreJ := ps[j].CurBoringScore()
		// a little random waggle
		return scoreI+rand.Intn(int(float64(scoreI)*0.1)+1) < scoreJ+rand.Intn(int(float64(scoreJ)*0.1)+1)
	})

	for _, p := range ps {

		// sorting groups by boring score
		sort.Slice(groups, func(i, j int) bool {
			ret := p.CalcBoringScore(groups[i].Members) - p.CalcBoringScore(groups[j].Members)
			if ret == 0 {
				return len(groups[i].Members) < len(groups[j].Members)
			} else {
				return ret < 0
			}
		})

		// add person to group
		for _, g := range groups {
			if g.IsFull(memberCount) {
				continue
			} else {
				g.AddMember(p)
				break
			}
		}
	}

	return groups
}
