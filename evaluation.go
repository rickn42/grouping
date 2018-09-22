package main

import (
	"fmt"
	"math"
)

func evaluation(personCnt int, memberCnt int, turnCnt int) {

	var idealMetCnt = float64(turnCnt*(memberCnt-1)) / float64(personCnt-1)
	var alleviateScore = 1.5

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
	var maxGroupScore float64
	var cntGroupScore int
	var sumGroupScore float64
	for i := 0; i < turnCnt; i++ {
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

			if g.BoringScoreSum > maxGroupScore {
				maxGroupScore = g.BoringScoreSum
			}
			sumGroupScore += g.BoringScoreSum
			cntGroupScore += 1
		}
		AlleviateBoringScore(persons, alleviateScore)
	}

	// print result statistics
	var maxScore, minScore, sumScore, avgScore float64
	minScore = math.MaxInt32
	for i, metCnts := range statistics {
		var score float64
		for j, cnt := range metCnts {
			if delta := idealMetCnt - float64(cnt); i != j {
				score += delta * delta
			}
		}
		fmt.Printf("person(%d) score %1.f\n%v\n", i, score, metCnts)

		if score > maxScore {
			maxScore = score
		}
		if score < minScore {
			minScore = score
		}
		sumScore += score
	}

	// print result
	fmt.Printf("grouping with person %d, member %d, generation %d, ideal met count %.1f\n", personCnt, memberCnt, turnCnt, idealMetCnt)
	avgScore = sumScore / float64(len(statistics))
	fmt.Printf("statistic score: min %.1f, max %.1f, avg %.1f\n", minScore, maxScore, avgScore)
	avgGroupScore := sumGroupScore / float64(cntGroupScore)
	fmt.Printf("group boring score: max %.1f, avg %.1f", maxGroupScore, avgGroupScore)
}
