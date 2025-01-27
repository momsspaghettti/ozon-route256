package t3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask3(t *testing.T) {
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

			Task3()

			_, err = outFile.Seek(0, 0)
			require.NoError(t, err)

			actual := readAnswer(t, outFile)
			require.Equal(t, tc.expected, actual)
		})
	}
}

type testCase struct {
	name      string
	inputFile *os.File
	expected  []int
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
			tc.expected = readAnswer(t, file)
			require.NoError(t, file.Close())
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

func readAnswer(t *testing.T, in io.Reader) []int {
	res := make([]int, 0)
	r := bufio.NewReader(in)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		line = line[:len(line)-1]
		if len(line) == 0 {
			break
		}

		num, err := strconv.Atoi(line)
		require.NoError(t, err)

		res = append(res, num)
	}

	return res
}
