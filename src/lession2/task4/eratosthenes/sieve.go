package eratosthenes

/*
* Author: Artem Chistkov
* License: JFF (Just For Fun) =)
*/

const (
	initialSieveLength = 1000
	extCoefficient     = 1.3
)

var (
	// map of indexes of last zeroed position in sieve for every simple value
	// needed to make it available to extend the sieve if needed
	// key is a simple number
	// value is a last index in sieve where we zeroed value with step by the value
	lastZeroedPositions map[int]int

	//to decrease unnecessary steps
	lastSimpleValue = 2

	//current Erathothenes sieve
	sieve []int

	//sorted list of known simple numbers
	simpleNumbers []int
)

func init() {
	lastZeroedPositions = make(map[int]int)

	// we don't want to recreate this slice on every appending to capacity 2,4,8,16 etc.
	simpleNumbers = make([]int, 0, initialSieveLength)

	// and will do small initial sieve
	newSieveByValue(initialSieveLength)
	saveSimpleNumbers()
}

func newSieveByValue(n int) {

	sieve = make([]int, n+1, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = i
	}
	cleanSieve(lastSimpleValue)
}

//recursievely clean sieve by zeroing all non simple values
func cleanSieve(n int) {

	simpleNumbers = append(simpleNumbers, n)
	for pos := n + n; pos < len(sieve); pos = pos + n {
		sieve[pos] = 0
		lastZeroedPositions[n] = pos
	}

	// when done we can find next non zero value after the value and it will be simple
	// so start new iteration of cleaning by next value
	if n = findNextSimple(n); n > 0 {
		cleanSieve(n)
	}
}

// find next simple by iterating one by one sieve's value from the start value
// until we find non zero value
func findNextSimple(n int) int {
	for i := n + 1; i < len(sieve); i++ {
		if sieve[i] > 0 {
			return i
		}
	}
	return 0
}

func extendSieve(n int) {
	for i := len(sieve); len(sieve) < n; i++ {
		sieve = append(sieve, i)
	}

	cleanByKnownPositions()                     //we remember all last zeroed indexes for every simple value
	cleanSieve(findNextSimple(lastSimpleValue)) // clean new sieve starting from the next simple value
	saveSimpleNumbers()
}

func cleanByKnownPositions() {
	for num, lastIdx := range lastZeroedPositions {
		for pos := lastIdx; pos < len(sieve); pos = pos + num {
			sieve[pos] = 0
		}
	}
}

func saveSimpleNumbers() {
	simpleNumbers = make([]int, 0, int(float64(len(sieve))*extCoefficient))
	for _, num := range sieve {
		if num > 0 {
			simpleNumbers = append(simpleNumbers, num)
		}
	}
}

// Public functions
func GetSieveByValue(n int) []int {
	if len(sieve) < n {
		extendSieve(n)
	}
	return sieve[:n]
}

func GetAllKnownSimples() []int { return simpleNumbers }

func GetSimplesByValue(n int) []int {
	for n > simpleNumbers[len(simpleNumbers)-1] {
		extendSieve(int(float64(len(sieve)) * extCoefficient))
	}

	//FIXME: find better way to do it =) I don't like this.
	i := len(simpleNumbers)-1
	for n < simpleNumbers[i] { i-- }
	return simpleNumbers[:i+1]
}

func GetSimplesByCount(n int) []int {
	if len(simpleNumbers) < n {
		for len(simpleNumbers) < n {
			extendSieve(int(float64(len(sieve)) * extCoefficient))
		}
	}
	return simpleNumbers[:n]
}
