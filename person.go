package main

import (
	"fmt"
)

type Person struct {
	Id     int
	Boring map[*Person]float64
}

func NewPerson(id int) *Person {
	return &Person{
		Id:     id,
		Boring: make(map[*Person]float64),
	}
}

func (r *Person) String() string {
	return fmt.Sprintf("%02d", r.Id)
}

func (r *Person) AddScore(p *Person, score float64) {
	n, _ := r.Boring[p]
	r.Boring[p] = n + score
}

func (r *Person) CalcBoringScore(ps []*Person) (score float64) {
	for _, p := range ps {
		if i, ok := r.Boring[p]; ok {
			score += float64(i * i)
		}
	}
	return score
}

func (r *Person) CurBoringScore() (score float64) {
	for _, n := range r.Boring {
		score += float64(n * n)
	}
	return score
}

func (r *Person) AlleviateBoringScore(n float64) {
	for p, score := range r.Boring {
		if newScore := score - n; newScore <= 0 {
			delete(r.Boring, p)
		} else {
			r.Boring[p] = newScore
		}
	}
}
