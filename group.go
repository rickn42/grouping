package main

import (
	"fmt"
	"sort"
)

type Group struct {
	ScorePerMet    int
	BoringScoreSum int
	Members        []*Person
}

func NewGroup(scorePerMet int) *Group {
	return &Group{ScorePerMet: scorePerMet}
}

func (g *Group) String() string {
	sort.Slice(g.Members, func(i, j int) bool {
		return g.Members[i].Id < g.Members[j].Id
	})

	return fmt.Sprintf("%v(%02d)", g.Members, g.BoringScoreSum)
}

func (g *Group) IsFull(limit int) bool {
	return len(g.Members) >= limit
}

func (g *Group) AddMember(p *Person) {
	g.BoringScoreSum += p.CalcBoringScore(g.Members)
	for _, member := range g.Members {
		member.AddScore(p, g.ScorePerMet)
		p.AddScore(member, g.ScorePerMet)
	}
	g.Members = append(g.Members, p)
}
