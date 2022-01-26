package main

import "fmt"

var output []int
var inputIndexOfOutput int = 0

func main() {
	//inputList := []string{"a", "b", "a", "b", "c", "b", "a", "c", "a", "d", "e", "f", "e", "g", "d", "e", "h", "i", "j", "h", "k", "l", "i", "j"}
	inputList := []string{"z", "w", "h", "c", "b", "z", "m", "c", "h", "f", "h", "m", "i"} //This is the hardest and the most difficult array you can have
	//inputList := []string{"a", "b", "c", "a"}
	//inputList := []string{"a", "b", "c"}
	//inputList := []string{"z", "w", "c", "b", "z", "c", "h", "f", "i", "h", "i"}

	//inputList := []string{"z","w","c","b","z","c","h","f","i","h","i"}
	getLastOccuranceOf(0, inputList)

	if len(output) == 1 {
		output[0] = 1
	}

	fmt.Println("\n", output)
}

func getLastOccuranceOf(indexToStartSearchFrom int, inputList []string) {

	lastOccurance := 0

	for i := indexToStartSearchFrom; i < len(inputList); i++ {
		if inputList[i] == inputList[indexToStartSearchFrom] {
			lastOccurance = i
		}
	}

	isThereAnyOtherOccurance := isThereAnotherOccuranceOfLettersBetweenIndexToSearchAndLastOccuranceThatIsAppearedAfterLastOccurance(indexToStartSearchFrom, lastOccurance, inputList, -1)

	if isThereAnyOtherOccurance != -1 {
		fmt.Println(isThereAnyOtherOccurance)
		output = append(output, isThereAnyOtherOccurance+1-indexToStartSearchFrom)
		inputIndexOfOutput++

		if isThereAnyOtherOccurance+1 < len(inputList) {
			getLastOccuranceOf(isThereAnyOtherOccurance+1, inputList)
		}

	} else {
		output = append(output, lastOccurance+1-indexToStartSearchFrom)
		inputIndexOfOutput++

		if lastOccurance+1 < len(inputList) {

			getLastOccuranceOf(lastOccurance+1, inputList)
		}
	}

}

func isThereAnotherOccuranceOfLettersBetweenIndexToSearchAndLastOccuranceThatIsAppearedAfterLastOccurance(indexToStartSearchFrom, lastOccurance int, inputList []string, lastIndexOfOccuranceFromPreviousForLoop int) int {
	indexToReturn := -1
	for i := indexToStartSearchFrom; i < lastOccurance; i++ {
		for j := lastOccurance + 1; j < len(inputList); j++ {
			if inputList[i] == inputList[j] {
				if j > indexToReturn {
					indexToReturn = j
				}
				//fmt.Println("'", inputList[i], "' at index ", i, " is found again at index ", j, " after last occurance of '", inputList[lastOccurance], "' at index ", lastOccurance)
			}
		}
	}

	if indexToReturn != -1 {
		return isThereAnotherOccuranceOfLettersBetweenIndexToSearchAndLastOccuranceThatIsAppearedAfterLastOccurance(lastOccurance, indexToReturn, inputList, indexToReturn)
	}
	return lastIndexOfOccuranceFromPreviousForLoop
}

func Debug(...string) {

}
