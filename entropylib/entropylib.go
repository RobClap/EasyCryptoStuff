package entropylib

import (
	"io/ioutil"
	"math"
)

//This actually calculates the shannon entropy and returns also the slice
//used to count occurencies
func coreCalc(input []byte) (entropy float64, bytes []uint32) {
	totalCharsCounter := 0
	bytes = make([]uint32, 256)
	for _, value := range input {
		bytes[value]++
		totalCharsCounter++
	}
	for _, value := range bytes {
		fvalue := float64(value) / float64(totalCharsCounter)
		if fvalue > 0 {
			entropy = entropy - float64(fvalue)*math.Log2(float64(fvalue))
		}
	}
	return
}

//Returns a value between 1 and 8
func CalculateShannon(input string) (entropy float64) {
	entropy, _ = coreCalc([]byte(input))
	return
}

//Returns a value of entropy assuming the whole character set is being used
//The value si always between 0.0 and 1.0
func CalculateBalanced(input string) (entropy float64) {
	entropy, bytes := coreCalc([]byte(input))
	counter := 0
	for _, value := range bytes {
		if value > 0 {
			counter++
		}
	}
	if counter > 1 {
		entropy = entropy / math.Log2(float64(counter))
	}
	return
}

//Returns a value of entropy assuming the character set is as long as the
//given parameter.
//This function does not check if the assertion above is valid.
func CalculateOnCharset(input string, charsetLen int) (entropy float64) {
	entropy, _ = coreCalc([]byte(input))
	if charsetLen > 1 {
		entropy = entropy / math.Log2(float64(charsetLen))
	}
	return
}

func CalculateOfFile(filepath string) (entropy float64, err error) {
	content, err := ioutil.ReadFile(filepath)
	entropy, _ = coreCalc(content)
	if err != nil {
		return -1.0, err
	}
	return entropy, nil
}
