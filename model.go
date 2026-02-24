package main

import "math"

type Model struct {
	hamCounts  map[string]int
	spamCounts map[string]int

	hamTotal  int
	spamTotal int

	hamDocCount  int
	spamDocCount int

	vocab map[string]struct{}

	alpha float64
}

func TrainNaiveBayes(hams, spams []string, alpha float64) Model {
	model := Model{
		hamCounts:    map[string]int{},
		spamCounts:   map[string]int{},
		vocab:        map[string]struct{}{},
		alpha:        alpha,
		hamDocCount:  len(hams),
		spamDocCount: len(spams),
	}

	for _, doc := range hams {
		for _, tok := range Tokenize(doc) {
			model.hamCounts[tok]++
			model.hamTotal++
			model.vocab[tok] = struct{}{}
		}
	}
	for _, doc := range spams {
		for _, tok := range Tokenize(doc) {
			model.spamCounts[tok]++
			model.spamTotal++
			model.vocab[tok] = struct{}{}
		}
	}

	return model
}

func (m Model) hamLiklihood (tok string) float64 {
	V := float64(len(m.vocab))
	num := float64(m.hamCounts[tok]) + m.alpha
	den := float64(m.hamTotal) + m.alpha * V
	return math.Log(num / den)
}

func (m Model) spamLiklihood (tok string) float64 {
	V := float64(len(m.vocab))
	num := float64(m.spamCounts[tok]) + m.alpha
	den := float64(m.spamTotal) + m.alpha * V
	return math.Log(num / den)
}

func (m Model) Predict(doc string) (label string, logHam float64, logSpam float64) {
	docToks := Tokenize(doc)

	totalDocs := float64(m.hamDocCount + m.spamDocCount)
	logHam = math.Log(float64(m.hamDocCount) / totalDocs)
	logSpam = math.Log(float64(m.spamDocCount) / totalDocs)

	docCount := map[string]int{}
	for _, tok := range docToks {
		docCount[tok]++
	}

	for tok, count := range docCount {
		if _, ok := m.vocab[tok]; !ok {
			// ignore new vocab
			continue
		}

		cc := float64(count)
		logHam += cc * m.hamLiklihood(tok)
		logSpam += cc * m.spamLiklihood(tok)

	}

	if logSpam > logHam {
		return "SPAM", logHam, logSpam
	}
	return "HAM", logHam, logSpam
}
