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
	GPA                 float64
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

		gpa, _ := strconv.ParseFloat(fields[2], 64)
		a := new(applicant)

		a.FirstName = fields[0]
		a.LastName = fields[1]
		a.GPA = gpa
		a.Departments = fields[3:]

		applicants = append(applicants, a)
	}
	return applicants
}

func sortApplicantsByGPA(applicants []*applicant) []*applicant {

	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].GPA != applicants[j].GPA {
			return applicants[i].GPA > applicants[j].GPA
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
		departments[d] = sortApplicantsByGPA(departments[d])
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
	
	for _, as := range accepted {
		sortApplicantsByGPA(as)
	}

	return accepted
}

func printDepartment(departments map[string][]*applicant, names []string) {
	for _, n := range names {
		fmt.Println(n, len(departments[n]))
		for _, a := range departments[n] {
			if !a.Accepted {
				continue
			}

			fmt.Printf("%s %s %.2f\n", a.FirstName, a.LastName, a.GPA)
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
