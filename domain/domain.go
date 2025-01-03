package domain

import "maps"

// Add more clean types for age sex
type Criminal struct {
	Name             string
	Age              int
	Gender           string
	CrimeDate        string
	CrimeType        string
	CrimeDescription string
	Status           string
}

type Criminals []Criminal

// Средний возраст по выборке
func getAvargeAge(criminals Criminals) (avargeAge int) {
	for _, c := range criminals {
		avargeAge += int(c.Age)
	}
	avargeAge /= len(criminals)
	return
}

// Наиболее часто встречающееся преступление
func getMostCommonCrime(criminals Criminals) (mostCommonCrime string) {
	if len(criminals) == 0 {
		panic("getMostCommonCrime get empty criminals")
	}
	crimesMap := make(map[string]int, len(criminals))
	for _, c := range criminals {
		if _, ok := crimesMap[c.CrimeType]; ok {
			crimesMap[c.CrimeType] += 1
		} else {
			crimesMap[c.CrimeType] = 1
		}
	}
	maxValue := 0
	for key := range maps.Keys(crimesMap) {
		if maxValue == 0 || maxValue < crimesMap[key] {
			maxValue = crimesMap[key]
			mostCommonCrime = key
		}
	}
	return
}

// Наиболее часто встречающееся пол преступника
func getMostCommonGender(criminals Criminals) (mostCommonGender string) {
	if len(criminals) == 0 {
		panic("getMostCommonGender get empty criminals")
	}
	crimesMap := make(map[string]int, len(criminals))
	for _, c := range criminals {
		if _, ok := crimesMap[c.Gender]; ok {
			crimesMap[c.Gender] += 1
		} else {
			crimesMap[c.Gender] = 1
		}
	}
	maxValue := 0
	for key := range maps.Keys(crimesMap) {
		if maxValue == 0 || maxValue < crimesMap[key] {
			maxValue = crimesMap[key]
			mostCommonGender = key
		}
	}
	return
}
