package main

import (
	"fmt"
	"math"

	"github.com/rickn42/grouping"
)

type GroupingFunc func(ps []*grouping.Person, grpOpt grouping.Option) []*grouping.Group

func evaluation(groupingFunc GroupingFunc, evalOpt EvaluationOption) {

	// meetCounter for record meeting each other.
	// First index is a person number.
	// Second index is a person who met.
	// Third index is turn numbers of meeting.
	// ex) meetCounter[i][j] = [1,4,8,12] means that person(i) met person(j) 4 times (turn number 1, 4, 8, 12)
	meetCounter := make([][][]int, evalOpt.PersonCnt)
	for i := range meetCounter {
		meetCounter[i] = make([][]int, evalOpt.PersonCnt)
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
	if evalOpt.ShowTurn { fmt.Println("repeat grouping...") }
	var maxGroupScore float64
	var cntGroupScore int
	var sumGroupScore float64
	for turnNum := 0; turnNum < evalOpt.TurnCnt; turnNum++ {
		groups := groupingFunc(persons, groupingOpt)

		if evalOpt.ShowTurn {
			fmt.Printf("turn %03d: %v\n", turnNum, groups)
		}

		// record for meetCounter
		for _, g := range groups {
			for _, m := range g.Members {
				for _, m2 := range g.Members {
					if m.Id != m2.Id {
						meetCounter[m.Id][m2.Id] = append(meetCounter[m.Id][m2.Id], turnNum)
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
	for me, other := range meetCounter {

		var score float64
		for i, turns := range other {
			// update meet count score
			if delta := idealMeetCnt - float64(len(turns)); me != i {
				score += delta * delta
			}
			// update meet interval score
			for idx, turn := range turns {
				if idx != 0 {
					d := idealMeetTurn - float64(turn - turns[idx-1])
					turnIntervalCost[me] += d * d
				}
			}
		}

		// print meet count
		meetCnt := make([]int, evalOpt.PersonCnt)
		for who, turns := range other {
			meetCnt[who] = len(turns)
		}
		if evalOpt.Verbose {
			fmt.Printf("person(%d) score %1.f\nmeet count %v\n", me, score, meetCnt)
		}

		// update min, max score
		if score > maxScore {
			maxScore = score
		}
		if score < minScore {
			minScore = score
		}
		sumScore += score
	}

	// print result
	fmt.Println()
	fmt.Println(evalOpt)
	avgScore = sumScore / float64(len(meetCounter))

	fmt.Println()
	fmt.Println("* difference score between ideal and result meet count")
	fmt.Printf("min %.1f, max %.1f, avg %.1f\n", minScore, maxScore, avgScore)

	//fmt.Println()
	//avgGroupScore := sumGroupScore / float64(cntGroupScore)
	//fmt.Println("* group boring cost")
	//fmt.Printf("max %.1f, avg %.1f\n", maxGroupScore, avgGroupScore)

	fmt.Println()
	fmt.Println("* difference score between ideal and result meet interval")
	var minTurnCost, maxTurnCost, sumTurnCost float64
	minTurnCost = math.MaxInt32
	if evalOpt.Verbose {
		fmt.Print("[")
	}
	for _, cost := range turnIntervalCost {
		if evalOpt.Verbose {
			fmt.Printf("%01.f, ",cost)
		}
		if cost < minTurnCost {
			minTurnCost = cost
		}
		if cost > maxTurnCost {
			maxTurnCost = cost
		}
		sumTurnCost += cost
	}
	if evalOpt.Verbose {
		fmt.Println("]")
	}
	fmt.Printf("min %.1f max %.1f avg %1.f\n", minTurnCost, maxTurnCost, sumTurnCost/float64(len(turnIntervalCost)))
}
