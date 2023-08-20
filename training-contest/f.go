package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct {
	coordinate int
	isStart    bool
}

func getPoint(h, m, s int, isStart bool) (valid bool, p point) {
	if h < 0 || h > 23 {
		return false, point{}
	}

	if m < 0 || m > 59 {
		return false, point{}
	}

	if s < 0 || s > 59 {
		return false, point{}
	}

	return true, point{
		coordinate: h*60*60 + m*60 + s,
		isStart:    isStart,
	}
}

func addIntervalPoints(h1, m1, s1, h2, m2, s2 int, pointsArr []point) (valid bool, resPointsArr []point) {
	valid, startPoint := getPoint(h1, m1, s1, true)
	if !valid {
		return false, pointsArr
	}

	valid, endPoint := getPoint(h2, m2, s2, false)
	if !valid {
		return false, pointsArr
	}

	if endPoint.coordinate < startPoint.coordinate {
		return false, pointsArr
	}

	return true, append(pointsArr, startPoint, endPoint)
}

func checkIntervalsPoints(pointsArr []point) bool {
	sort.Slice(pointsArr, func(i, j int) bool {
		if pointsArr[i].coordinate == pointsArr[j].coordinate {
			if pointsArr[i].isStart {
				return true
			}
			return false
		}
		return pointsArr[i].coordinate < pointsArr[j].coordinate
	})

	startsCount := 0
	for _, p := range pointsArr {
		if p.isStart {
			startsCount++
		} else {
			startsCount--
		}

		if startsCount > 1 {
			return false
		}
	}

	return true
}

func fFindAnswerForSet(in *bufio.Reader, out *bufio.Writer, pointsArr []point) {
	pointsArr = pointsArr[:0]

	var n int
	_, _ = fmt.Fscanf(in, "%d\n", &n)

	res := true
	// HH:MM:SS-HH:MM:SS
	inputFormat := "%d:%d:%d-%d:%d:%d\n"
	var h1, m1, s1, h2, m2, s2 int
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanf(in, inputFormat, &h1, &m1, &s1, &h2, &m2, &s2)
		if !res {
			continue
		}

		res, pointsArr = addIntervalPoints(h1, m1, s1, h2, m2, s2, pointsArr)
	}

	res = res && checkIntervalsPoints(pointsArr)

	if res {
		_, _ = fmt.Fprintln(out, "YES")
	} else {
		_, _ = fmt.Fprintln(out, "NO")
	}
}

func F() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCount int
	_, _ = fmt.Fscanf(in, "%d\n", &setsCount)

	pointsArr := make([]point, 0, 40000)
	for i := 0; i < setsCount; i++ {
		fFindAnswerForSet(in, out, pointsArr)
	}
}

func main() {
	F()
}
