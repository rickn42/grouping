package main

import (
	"fmt"
	"math"

	"github.com/rickn42/study/group/grouping"
)

type GroupingFunc func(ps []*grouping.Person, grpOpt grouping.Option) []*grouping.Group

func evaluation(groupingFunc GroupingFunc, evalOpt EvaluationOption) {

	// statistics for record meeting count each person
	statistics := make([][]int, evalOpt.PersonCnt)
	for i := range statistics {
		statistics[i] = make([]int, evalOpt.PersonCnt)
	}

	persons := make([]*grouping.Person, evalOpt.PersonCnt)
	for i := range persons {
		persons[i] = grouping.NewPerson(i)
	}

	groupingOpt := grouping.Option{
		MemberCnt:    evalOpt.MemberOfGroup,
		BoringAmount: evalOpt.IdealMeetTurn(),
	}

	// repeat grouping
	//fmt.Println("repeat grouping...")
	var maxGroupScore float64
	var cntGroupScore int
	var sumGroupScore float64
	for i := 0; i < evalOpt.TurnCnt; i++ {
		groups := groupingFunc(persons, groupingOpt)

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
		grouping.Alleviate(persons, evalOpt.AlleviateAmount)
	}

	// print result statistics
	var maxScore, minScore, sumScore, avgScore float64
	idealMeetCnt := evalOpt.IdealMeetCnt()
	minScore = math.MaxInt32
	for i, metCnts := range statistics {
		var score float64
		for j, cnt := range metCnts {
			if delta := idealMeetCnt - float64(cnt); i != j {
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
	fmt.Println(evalOpt)
	avgScore = sumScore / float64(len(statistics))
	fmt.Printf("statistic score: min %.1f, max %.1f, avg %.1f\n", minScore, maxScore, avgScore)
	avgGroupScore := sumGroupScore / float64(cntGroupScore)
	fmt.Printf("group boring score: max %.1f, avg %.1f", maxGroupScore, avgGroupScore)
}
