package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const threshold = 60.0

type applicant struct {
	GPA                 float64
	FirstName, LastName string
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var applicants []applicant

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")

		s := strings.Fields(line)
		gpa, _ := strconv.ParseFloat(s[2], 64)
		applicants = append(applicants, applicant{FirstName: s[0], LastName: s[1], GPA: gpa})
	}

	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].GPA != applicants[j].GPA {
			return applicants[i].GPA > applicants[j].GPA
		} else if applicants[i].FirstName != applicants[j].FirstName {
			return applicants[i].FirstName < applicants[j].FirstName
		}
		return applicants[i].LastName < applicants[j].LastName
	})

	fmt.Println("Successful applicants:")
	for i := 0; i < m; i++ {
		fmt.Println(applicants[i].FirstName, applicants[i].LastName)
	}
}