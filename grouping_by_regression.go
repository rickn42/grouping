package grouping

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func GroupingByRegression(ps []*Person, memberCount int, boringMount float64) (groups []*Group) {
	groupCnt := int(math.Ceil(float64(len(ps)) / float64(memberCount)))

	// make empty groups
	for i := 0; i < groupCnt; i++ {
		groups = append(groups, NewGroup(boringMount, memberCount))
	}

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})

	for _, p := range ps {

		// sorting groups by boring score
		sort.Slice(groups, func(i, j int) bool {
			if p.CalcBoringScore(groups[i].Members)-p.CalcBoringScore(groups[j].Members) < 0 {
				return true
			} else if len(groups[i].Members) < len(groups[j].Members) {
				return true
			}
			return false
		})

		// add person to group
		for _, g := range groups {
			if g.IsFull() {
				continue
			} else {
				g.AddMember(p)
				break
			}
		}
	}

	return groups
}
