package iteration

//const refactorCount = 5

func repeat(word string, refactorCount int) string{
	var repeated string

	for i := 0; i < refactorCount; i++{
		repeated += word
	}
	return repeated
}