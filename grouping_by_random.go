package grouping

import (
	"math"
	"math/rand"
	"time"
)

func GroupingByRandom(ps []*Person, opt Option) []*Group {
	groups := make([]*Group, int(math.Ceil(float64(len(ps))/float64(opt.MemberCnt))))
	for i := range groups {
		groups[i] = NewGroup(opt.BoringAmount, opt.MemberCnt)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})

	var groupIdx int
	for _, p := range ps {
		groupIdx %= len(groups)
		groups[groupIdx].AddMember(p)
		groupIdx++
	}

	return groups
}
