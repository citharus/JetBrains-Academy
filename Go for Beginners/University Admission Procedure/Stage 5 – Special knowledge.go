package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type applicant struct {
	FirstName, LastName string
	Scores              map[string]float64
	Departments         []string
	Accepted            bool
}

func getApplicants() []*applicant {
	var applicants []*applicant

	file, _ := os.Open("applicants.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		a := new(applicant)

		a.FirstName = fields[0]
		a.LastName = fields[1]
		a.Departments = fields[6:]

		a.Scores = make(map[string]float64, 4)
		for i, s := range fields[2:6] {
			score, _ := strconv.ParseFloat(s, 64)

			switch i {
			case 0:
				a.Scores["Physics"] = score
			case 1:
				a.Scores["Chemistry"] = score
				a.Scores["Biotech"] = score
			case 2:
				a.Scores["Mathematics"] = score
			case 3:
				a.Scores["Engineering"] = score
			}
		}

		applicants = append(applicants, a)
	}
	return applicants
}

func sortApplicantsByDepartmentScore(applicants []*applicant, department string) []*applicant {
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].Scores[department] != applicants[j].Scores[department] {
			return applicants[i].Scores[department] > applicants[j].Scores[department]
		}
		return (applicants[i].FirstName + applicants[i].LastName) < (applicants[j].FirstName + applicants[j].LastName)
	})

	return applicants
}

func groupApplicantsByDepartment(applicants []*applicant, round int) map[string][]*applicant {
	departments := make(map[string][]*applicant)

	for _, a := range applicants {
		departments[a.Departments[round-1]] = append(departments[a.Departments[round-1]], a)
	}

	for d := range departments {
		departments[d] = sortApplicantsByDepartmentScore(departments[d], d)
	}

	return departments
}

func assignApplicants(applicants []*applicant, maxApplicants int) map[string][]*applicant {
	accepted := make(map[string][]*applicant)

	admission := func(round int) {
		for d, as := range groupApplicantsByDepartment(applicants, round) {
			for _, a := range as {
				if a.Accepted {
					continue
				}

				if len(accepted[d]) < maxApplicants {
					accepted[d] = append(accepted[d], a)
					a.Accepted = true
				}
			}
		}
	}

	admission(1)
	admission(2)
	admission(3)

	for d, as := range accepted {
		sortApplicantsByDepartmentScore(as, d)
	}

	return accepted
}

func printDepartment(departments map[string][]*applicant, names []string) {
	for _, n := range names {
		fmt.Println(n)
		for _, a := range departments[n] {
			fmt.Printf("%s %s %.1f\n", a.FirstName, a.LastName, a.Scores[n])
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	applicants := getApplicants()

	departments := assignApplicants(applicants, n)

	printDepartment(departments, []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"})
}
