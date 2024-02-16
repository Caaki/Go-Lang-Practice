package tickersAndDefault

/*
ASSIGNMENT
Like all good back-end engineers, we frequently save backup snapshots of the Mailio database.

Complete the saveBackups function.

It should read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds.
After all, we want to show in the logs that the snapshot service is running.

*/

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestSaveBackups(t *testing.T) {
	expectedLogs := []string{
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"All backups saved!",
	}

	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	actualLogs := []string{}
	for actualLog := range logChan {
		fmt.Println(actualLog)
		actualLogs = append(actualLogs, actualLog)
	}

	if !slices.Equal(expectedLogs, actualLogs) {
		t.Errorf(`Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
	} else {
		fmt.Printf(`Test Passed:
expected:
%v
actual:
%v`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}
