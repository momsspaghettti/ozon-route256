package t3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Dir struct {
	Files   []string `json:"files,omitempty"`
	Folders []*Dir   `json:"folders,omitempty"`
}

func Task3() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = out.Flush()
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		panic(fmt.Errorf("failed to read t: %w", err))
	}

	for i := range t {
		count, err := getInfectedFilesCount(in)
		if err != nil {
			panic(fmt.Errorf("failed to process #%d data set: %w", i, err))
		}

		if _, err = fmt.Fprintln(out, count); err != nil {
			panic(fmt.Errorf("failed to write #%d answer '%d': %w", i, count, err))
		}
	}
}

func getInfectedFilesCount(in *bufio.Reader) (int, error) {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return 0, fmt.Errorf("failed to read n: %w", err)
	}
	_, err := in.ReadBytes('\n')
	if err != nil {
		return 0, fmt.Errorf("failed to read N line new line: %w", err)
	}

	dTree, err := readDirTree(in, n)
	if err != nil {
		return 0, fmt.Errorf("failed to read dir tree: %w", err)
	}

	return countInfectedFiles(dTree, false), nil
}

func countInfectedFiles(dir *Dir, isDirInfected bool) int {
	hasInfectedFiles := false
	if isDirInfected {
		hasInfectedFiles = true
	} else {
		for _, fileName := range dir.Files {
			if strings.HasSuffix(fileName, ".hack") {
				hasInfectedFiles = true
				break
			}
		}
	}

	isDirInfected = isDirInfected || hasInfectedFiles

	infectedFilesCount := 0
	if isDirInfected {
		infectedFilesCount = len(dir.Files)
	}

	for _, cDir := range dir.Folders {
		infectedFilesCount += countInfectedFiles(cDir, isDirInfected)
	}

	return infectedFilesCount
}

func readDirTree(in *bufio.Reader, linesCount int) (*Dir, error) {
	r, w := io.Pipe()
	go func() {
		defer func() {
			_ = w.Close()
		}()

		for i := range linesCount {
			lineBytes, err := in.ReadBytes('\n')
			if err != nil {
				panic(fmt.Errorf("failed to read line #%d: %w", i, err))
			}
			if _, err = w.Write(lineBytes); err != nil {
				panic(fmt.Errorf("failed to write line #%d to pipe: %w", i, err))
			}
		}
	}()

	var d Dir
	if err := json.NewDecoder(r).Decode(&d); err != nil {
		return nil, fmt.Errorf("failed to decode dir: %w", err)
	}

	return &d, nil
}
