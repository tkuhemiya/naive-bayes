# naive bayes implementation

A small Go experiment that trains a naive Bayes spam detector on whatever text you feed it. It keeps counts of every token seen in spam and ham files, applies Laplace smoothing, and then predicts whether a new string is spam or not.

## setup

The program expects two folders:

- `data/ham/` – place `.txt` files that should be treated as ham (legitimate) examples.
- `data/spam/` – place `.txt` files that are spam examples.

Each `.txt` file can contain whatever text you want the model to learn from. The training loop tokenizes the files and builds counts for each class.

## running

```sh
go run .
```

That will load the files, train a model with Laplace smoothing (alpha = 1.0 by default), and print each class count along with a sample prediction. You can swap out the hardcoded `test` string in `main.go` with your own text to play with predictions.

If you add new files or change the test string, rerun `go run main.go` to see updated results.
