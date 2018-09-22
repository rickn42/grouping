package main

import (
	"fmt"
	"math"
)

func evaluation(personCnt int, memberCnt int) {

	generationCnt := personCnt * 10
	idealMetCnt := (generationCnt * (memberCnt - 1)) / (personCnt - 1)
	alleviateScore := 2

	fmt.Printf("grouping with person %d, member %d, generation %d, ideal met count %d\n", personCnt, memberCnt, generationCnt, idealMetCnt)

	// statistics for record meeting count each person
	statistics := make([][]int, personCnt)
	for i := range statistics {
		statistics[i] = make([]int, personCnt)
	}

	persons := make([]*Person, personCnt)
	for i := range persons {
		persons[i] = NewPerson(i)
	}

	// repeat grouping
	fmt.Println("repeat grouping...")
	for i := 0; i < generationCnt; i++ {
		groups := Grouping(persons, memberCnt, idealMetCnt)

		// record for statistics
		for _, g := range groups {
			for _, m := range g.Members {
				for _, m2 := range g.Members {
					if m.Id != m2.Id {
						statistics[m.Id][m2.Id] += 1
					}
				}
			}
		}
		AlleviateBoringScore(persons, alleviateScore)
	}

	// print result statistics
	var maxScore, minScore, sumScore, avgScore int
	minScore = math.MaxInt32
	for i, metCnts := range statistics {
		var score int
		for j, cnt := range metCnts {
			if delta := idealMetCnt - cnt; i != j {
				score += delta * delta
			}
		}
		fmt.Printf("person(%d) score %d\n%v\n", i, score, metCnts)

		if score > maxScore {
			maxScore = score
		}
		if score < minScore {
			minScore = score
		}
		sumScore += score
	}
	avgScore = sumScore / len(statistics)
	fmt.Printf("score: min %d, max %d, avg %d\n", minScore, maxScore, avgScore)
}
