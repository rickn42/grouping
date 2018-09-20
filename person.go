package main

import (
	"fmt"
)

type Person struct {
	Id     int
	Boring map[*Person]int
}

func NewPerson(id int) *Person {
	return &Person{
		Id:     id,
		Boring: make(map[*Person]int),
	}
}

func (r *Person) String() string {
	return fmt.Sprintf("%02d", r.Id)
}

func (r *Person) AddScore(p *Person, score int) {
	n, _ := r.Boring[p]
	r.Boring[p] = n + score
}

func (r *Person) CalcBoringScore(ps []*Person) (score int) {
	for _, p := range ps {
		if i, ok := r.Boring[p]; ok {
			score += i * i
		}
	}
	return score
}

func (r *Person) CurBoringScore() (score int) {
	for _, n := range r.Boring {
		score += n * n
	}
	return score
}

func (r *Person) AlleviateBoringScore(n int) {
	for p, score := range r.Boring {
		if newScore := score - n; newScore <= 0 {
			delete(r.Boring, p)
		} else {
			r.Boring[p] = newScore
		}
	}
}
