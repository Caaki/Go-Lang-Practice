package mutex

/*
ASSIGNMENT
We send emails across many different goroutines at Mailio. To keep track of how many we've sent to a given email address,
we use an in-memory map.

Our safeCounter struct is unsafe! Update the inc() and val() methods so that they utilize the safeCounter's
mutex to ensure that the map is not accessed by multiple goroutines at the same time.

NOTE: WASM IS SINGLE-THREADED
Now, it's worth pointing out that our execution engine on Boot.dev uses web assembly to run the code you write in your browser.
Web assembly is single-threaded,
which awkwardly means that maps are thread-safe in web assembly.
I've simulated a multithreaded environment with the slowIncrement function.

In reality, any Go code you write may or may not run on a single-core machine,
so it's always best to write your code so that it is safe no matter which hardware it runs on.





*/

import (
	"fmt"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		email string
		count int
	}
	var tests = []testCase{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"lila@bates.com", 31},
			{"sam@bates.com", 453},
		}...)
	}

	for _, test := range tests {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.RWMutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		if output := sc.val(test.email); output != test.count {
			t.Errorf(`Test Failed:
email: %v
count : %v
->
expected count: %v
actual count: %v
=======================
`, test.email, test.count, test.count, output)
		} else {
			fmt.Printf(`Test Passed:
email: %v
count : %v
->
expected count: %v
actual count: %v
=======================
			`, test.email, test.count, test.count, output)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
