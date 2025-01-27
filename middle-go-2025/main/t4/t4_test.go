package t4

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask4(t *testing.T) {
	testCases := readTestData(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldStdin := os.Stdin
			oldStdout := os.Stdout
			t.Cleanup(func() {
				require.NoError(t, tc.inputFile.Close())
				os.Stdin = oldStdin
				os.Stdout = oldStdout
			})

			os.Stdin = tc.inputFile

			outFile, err := os.CreateTemp("", fmt.Sprintf("%s-out*", tc.name))
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, outFile.Close())
				require.NoError(t, os.Remove(outFile.Name()))
			})

			os.Stdout = outFile

			Task4()

			_, err = outFile.Seek(0, 0)
			require.NoError(t, err)

			var actual []map[string]any
			require.NoError(t, json.NewDecoder(outFile).Decode(&actual))

			require.Equal(t, tc.expected, actual)
		})
	}
}

type testCase struct {
	name      string
	inputFile *os.File
	expected  []map[string]any
}

func readTestData(t *testing.T) []testCase {
	files, err := os.ReadDir("testdata")
	require.NoError(t, err)

	testCasesMap := map[string]testCase{}
	for _, fileEntry := range files {
		require.False(t, fileEntry.IsDir())

		name := fileEntry.Name()
		var (
			tcName   = name
			isAnswer bool
		)
		if strings.HasSuffix(name, ".a") {
			tcName = strings.TrimSuffix(name, ".a")
			isAnswer = true
		}

		file, oErr := os.Open(filepath.Join("testdata", name))
		require.NoError(t, oErr)

		tc, ok := testCasesMap[tcName]
		if !ok {
			tc = testCase{
				name: tcName,
			}
		}

		if isAnswer {
			var expected []map[string]any
			require.NoError(t, json.NewDecoder(file).Decode(&expected))
			require.NoError(t, file.Close())
			tc.expected = expected
		} else {
			tc.inputFile = file
		}

		testCasesMap[tcName] = tc
	}

	testCases := make([]testCase, 0, len(testCasesMap))
	for _, tc := range testCasesMap {
		testCases = append(testCases, tc)
	}

	return testCases
}
