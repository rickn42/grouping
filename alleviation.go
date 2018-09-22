package grouping

func Alleviate(ps []*Person, alleviateScore float64) {
	for _, p := range ps {
		p.AlleviateBoringScore(alleviateScore)
	}
}
