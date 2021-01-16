package solutions

import (
	. "github.com/nsuthison/dojo-go/problems/leetcode/easy/employee-importance/models"
)

// Question: https://leetcode.com/problems/employee-importance/
func getImportance(employees []*Employee, id int) int {
	return 0
}

func createEmployeeMapFrom(employees []Employee) (employeeMap map[int]Employee) {

	employeeMap = make(map[int]Employee)

	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}

	return employeeMap
}

// func getAllSurbodinatesFrom(employeeId int, employees []*Employee) []int {

// }
