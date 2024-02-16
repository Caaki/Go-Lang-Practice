package closingChanels

//func countReports(numSentCh chan int) int {
//	numOfReports := 0
//
//	for {
//		elem, ok := <-numSentCh
//		if !ok {
//			break
//		} else {
//			numOfReports += elem
//		}
//
//	}
//	return numOfReports
//}

func countReports(numSentCh chan int) int {
	numOfReports := 0

	for elem := range numSentCh {
		numOfReports += elem
	}
	return numOfReports
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
