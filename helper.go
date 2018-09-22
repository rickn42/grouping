package grouping

func AlleviateBoringScore(ps []*Person, alleviateScore float64) {
	for _, p := range ps {
		p.AlleviateBoringScore(alleviateScore)
	}
}
