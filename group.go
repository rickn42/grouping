package grouping

import (
	"fmt"
	"sort"
)

type Group struct {
	ScorePerMet    float64
	BoringScoreSum float64
	MemberLimit    int
	Members        []*Person
}

func NewGroup(scorePerMet float64, memberLimit int) *Group {
	return &Group{ScorePerMet: scorePerMet, MemberLimit: memberLimit}
}

func (g *Group) String() string {
	sort.Slice(g.Members, func(i, j int) bool {
		return g.Members[i].Id < g.Members[j].Id
	})

	return fmt.Sprintf("%v", g.Members)
}

func (g *Group) IsFull() bool {
	return len(g.Members) >= g.MemberLimit
}

func (g *Group) AddMember(p *Person) {
	if g.IsFull() {
		panic("member is full!")
	}

	// update current group boring score
	g.BoringScoreSum += p.CalcBoringScore(g.Members)

	// update boring score each other
	for _, member := range g.Members {
		member.AddScore(p, g.ScorePerMet)
		p.AddScore(member, g.ScorePerMet)
	}

	g.Members = append(g.Members, p)
}
