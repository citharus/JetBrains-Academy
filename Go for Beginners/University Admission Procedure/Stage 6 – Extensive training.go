package main

import (
	"bufio"
	"fmt"
	"log"
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

func getApplicants() ([]*applicant, error) {
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
		
		scores, err := calculateScores(fields[2:6])
		if err != nil {return nil, err}
		a.Scores = scores

		applicants = append(applicants, a)
	}
	return applicants, nil
}

func calculateScores(rawScores []string) (map[string]float64, error) {
	var scoreList []float64
	for _, s := range rawScores {
		score, err := strconv.ParseFloat(s, 64)
		if err != nil {return nil, err}
		scoreList = append(scoreList, score)
	}

	scores := make(map[string]float64)
	scores["Physics"] = (scoreList[0] + scoreList[2]) / 2
	scores["Engineering"] = (scoreList[2] + scoreList[3]) / 2
	scores["Biotech"] = (scoreList[1] + scoreList[0]) / 2
	scores["Chemistry"] = scoreList[1]
	scores["Mathematics"] = scoreList[2]

	return scores, nil
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

func assignApplicantsToDepartments(applicants []*applicant, maxApplicants int) map[string][]*applicant {
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

func saveDepartments(departments map[string][]*applicant) error {
	for d, as := range departments {
		file, _ := os.Create(fmt.Sprintf("%s.txt", strings.ToLower(d)))
		for _, a := range as {
			_, err := fmt.Fprintf(file, "%s %s %.1f\n", a.FirstName, a.LastName, a.Scores[d])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	var n int
	fmt.Scan(&n)

	applicants, err := getApplicants()
	if err != nil {log.Fatal(err)}

	departments := assignApplicantsToDepartments(applicants, n)

	err = saveDepartments(departments)
	if err != nil {log.Fatal(err)}
}
