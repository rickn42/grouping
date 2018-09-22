package main

import "fmt"

type EvaluationOption struct {
	PersonCnt       int
	MemberOfGroup   int
	TurnCnt         int
	AlleviateAmount float64
}

func (opt EvaluationOption) String() string {
	return fmt.Sprintf("grouping with person %d, %d member of a group, turn %d, ideal meet count %.1f",
		opt.PersonCnt, opt.MemberOfGroup, opt.TurnCnt, opt.IdealMeetCnt())
}

func (opt EvaluationOption) IdealMeetCnt() float64 {
	return float64(opt.TurnCnt*(opt.MemberOfGroup-1)) / float64(opt.PersonCnt-1)
}

func (opt EvaluationOption) IdealMeetTurn() float64 {
	return float64(opt.TurnCnt) / opt.IdealMeetCnt()
}
