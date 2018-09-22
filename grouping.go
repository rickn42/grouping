package main

import (
	"math/rand"
	"sort"
)

func Grouping(ps []*Person, memberCount int, boringMount int) (groups []*Group) {
	// make empty groups
	for i := 0; i < len(ps)/memberCount; i++ {
		groups = append(groups, NewGroup(boringMount, memberCount))
	}

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

func AlleviateBoringScore(ps []*Person, alleviateScore int) {
	for _, p := range ps {
		p.AlleviateBoringScore(alleviateScore)
	}
}
