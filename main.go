package main

import (
	"fmt"
	"math"
)

func main() {
	hams, err := FileToString("./data/ham/*.txt")
	if err != nil {
		panic(err)
	}
	spams, err := FileToString("./data/spam/*.txt")
	if err != nil {
		panic(err)
	}

	model := TrainNaiveBayes(hams, spams, 1.0)

	fmt.Println("Total ham Docs: ", model.hamDocCount)
	fmt.Println("Total spam Docs: ", model.spamDocCount)

	test := `HELLOHELLOHELLO`

	label, lHam, lSpam := model.Predict(test)

	fmt.Println("Prediction:", label)
	maxLog := math.Max(lHam, lSpam)
	pHam := math.Exp(lHam-maxLog) / (math.Exp(lHam-maxLog) + math.Exp(lSpam-maxLog))
	pSpam := 1 - pHam
	fmt.Printf("P(ham|x)=%.4f P(spam|x)=%.4f\n", pHam, pSpam)
}
