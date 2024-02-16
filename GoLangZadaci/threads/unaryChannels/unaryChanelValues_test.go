package unaryChannels

/*
Our Mailio server isn't able to boot up until it receives the signal that its databases are all online,
and it learns about them being online by waiting for tokens (empty structs) on a channel.

Complete the waitForDbs function.
It should block until it receives numDBs tokens on the dbChan channel.
Each time it reads a token the getDatabasesChannel goroutine will print a message to the console for you.

*/

//import (
//	"fmt"
//	"testing"
//)
//
//func Test(t *testing.T) {
//	type testCase struct {
//		numDBs int
//	}
//	tests := []testCase{
//		{3},
//		{4},
//	}
//	if withSubmit {
//		tests = append(tests, []testCase{
//			{0},
//			{13},
//		}...)
//	}
//
//	for _, test := range tests {
//		dbChan := getDatabasesChannel(test.numDBs)
//		waitForDbs(test.numDBs, dbChan)
//		if len(dbChan) != 0 {
//			t.Errorf(`Test Failed: (%v) ->
//expected length: %v
//actual length: %v
//=======================
//`,
//				test.numDBs,
//				0,
//				len(dbChan),
//			)
//		} else {
//			fmt.Printf(`Test Passed: (%v) ->
//expected length: %v
//acual length: %v
//=======================
//`,
//				test.numDBs,
//				0,
//				len(dbChan),
//			)
//		}
//	}
//}
//
//// withSubmit is set at compile time depending
//// on which button is used to run the tests
//var withSubmit = true
