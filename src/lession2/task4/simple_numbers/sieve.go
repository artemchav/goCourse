package simple_numbers

/*
* Author: Artem Chistkov
* License: JFF (Just For Fun) =)

* it's the simpliest and the longest way to get simple numbers.
* We just divide all next value by every known before simple value
* there are no properties of using Erathosthene's sieve.
* Stupid as wood.
*/

var (
	// we don't know anything about cache yet. so we do it at least in a singleton style )
	simpleNumbers []int
	// just to decrease iterations count if it was called twice at minimum
	lastCheckedValue int
)

// if we want to get specified count of simple numbers from the start
func GetSimplesByCount(n int) []int {

	l := len(simpleNumbers)
	switch {
	case l == 0: // we didn't create the list before
		return newSimplesByCount(n)
	case l < n: // we have created it but it's not enough
		return extendSimplesByCount(n)
	case l == n: // no changes needed
		return simpleNumbers
	case l > n: // we don't need so big list as we have calculated before
		return simpleNumbers[0:n]
	}

	return simpleNumbers
}

// creates new simpleNumbers with specified numbers count
func newSimplesByCount(n int) []int {

	simpleNumbers = make([]int, 0, n)
	simpleNumbers =  append(simpleNumbers, 2)

	return extendSimplesByCount(n)
}

func extendSimplesByCount(n int) []int {

	simple := simpleNumbers[len(simpleNumbers)-1]

	for i := simple; len(simpleNumbers) < n; i++ {
		lastCheckedValue = i
		if isSimple(i) {
			simpleNumbers = append(simpleNumbers, i)
		}
	}

	return simpleNumbers
}

/*
* creates new simpleNumbers and fills it by simple numbers between zero and
* value specified by
 */
func newSimplesByValueLimit(n int) []int {

	simpleNumbers = make([]int, 0)
	simple := 2

	for i := simple; i <= n; i++ {
		if isSimple(i) {
			simpleNumbers = append(simpleNumbers, i)
		}
	}

	return simpleNumbers
}

// If the value is dividing by any known simple number then it's not simple
func isSimple(n int) bool {

	for _, num := range simpleNumbers {
		if n%num == 0 {
			return false
		}
	}

	return true
}

func GetSimpleNumbers() []int     { return simpleNumbers }
