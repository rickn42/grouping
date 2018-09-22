package main

import (
	"fmt"
	"sort"
)

type Group struct {
	ScorePerMet    int
	BoringScoreSum int
	MemberLimt     int
	Members        []*Person
}

func NewGroup(scorePerMet int, memberLimit int) *Group {
	return &Group{ScorePerMet: scorePerMet, MemberLimt: memberLimit}
}

func (g *Group) String() string {
	sort.Slice(g.Members, func(i, j int) bool {
		return g.Members[i].Id < g.Members[j].Id
	})

	return fmt.Sprintf("%v(%02d)", g.Members, g.BoringScoreSum)
}

func (g *Group) IsFull() bool {
	return len(g.Members) >= g.MemberLimt
}

func (g *Group) AddMember(p *Person) {
	if g.IsFull() {
		panic("member is full!")
	}
	g.BoringScoreSum += p.CalcBoringScore(g.Members)
	for _, member := range g.Members {
		member.AddScore(p, g.ScorePerMet)
		p.AddScore(member, g.ScorePerMet)
	}
	g.Members = append(g.Members, p)
}
