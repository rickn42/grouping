package main

import (
	"fmt"
	"math"

	"github.com/rickn42/grouping"
)

type GroupingFunc func(ps []*grouping.Person, grpOpt grouping.Option) []*grouping.Group

func evaluation(groupingFunc GroupingFunc, evalOpt EvaluationOption) {

	// statistics for record meeting each other.
	// First index is a person number.
	// Second index is a person who met.
	// Third index is turn numbers of meeting.
	// ex) statistics[i][j] = [1,4,8,12] means that person(i) met person(j) 4 times (turn number 1, 4, 8, 12)
	statistics := make([][][]int, evalOpt.PersonCnt)
	for i := range statistics {
		statistics[i] = make([][]int, evalOpt.PersonCnt)
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
						statistics[m.Id][m2.Id] = append(statistics[m.Id][m2.Id], i)
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

	// cale result
	var maxScore, minScore, sumScore, avgScore float64
	turnIntervalCost := make([]float64, evalOpt.PersonCnt)
	idealMeetTurn := evalOpt.IdealMeetTurn()
	idealMeetCnt := evalOpt.IdealMeetCnt()
	minScore = math.MaxInt32
	for i, j := range statistics {
		var score float64
		for j, turns := range j {
			if delta := idealMeetCnt - float64(len(turns)); i != j {
				score += delta * delta
			}
			for idx, turn := range turns {
				if idx != 0 {
					d := idealMeetTurn - float64(turn - turns[idx-1])
					turnIntervalCost[i] += d * d
				}
			}
		}
		meetCnt := make([]int, evalOpt.PersonCnt)
		for who, turns := range j {
			meetCnt[who] = len(turns)
		}
		if evalOpt.Verbose {
			fmt.Printf("person(%d) score %1.f\nmeet count %v\n", i, score, meetCnt)
		}

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
	fmt.Println("* statistics score")
	fmt.Printf("min %.1f, max %.1f, avg %.1f\n", minScore, maxScore, avgScore)
	avgGroupScore := sumGroupScore / float64(cntGroupScore)
	fmt.Println("* group boring cost")
	fmt.Printf("max %.1f, avg %.1f\n", maxGroupScore, avgGroupScore)

	fmt.Println("* meet turn interval cost")
	var minTurnCost, maxTurnCost, sumTurnCost float64
	minTurnCost = math.MaxInt32
	for _, cost := range turnIntervalCost {
		//fmt.Printf("person(%d) %01.f\n", i, cost)
		if cost < minTurnCost {
			minTurnCost = cost
		}
		if cost > maxTurnCost {
			maxTurnCost = cost
		}
		sumTurnCost += cost
	}
	fmt.Printf("min %.1f max %.1f avg %1.f\n", minTurnCost, maxTurnCost, sumTurnCost/float64(len(turnIntervalCost)))
}
