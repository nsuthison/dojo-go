package solutions

import (
	. "github.com/nsuthison/dojo-go/problems/leetcode/easy/employee-importance/models"
)

// Question: https://leetcode.com/problems/employee-importance/
func getImportance(employees []*Employee, id int) int {
	emplyeesMap := createEmployeeMapFrom(employees)

	toFind, _ := emplyeesMap[id]
	sumImportancePointsMemoize := make(map[int]int)

	return getSumImportanceFrom(toFind, emplyeesMap, sumImportancePointsMemoize)
}

func getSumImportanceFrom(employee Employee, employeesMap map[int]Employee, sumImportancePointsMemoize map[int]int) int {

	sumPoints, isExist := sumImportancePointsMemoize[employee.Id]
	if !isExist {
		sumPoints = employee.Importance

		if len(employee.Subordinates) != 0 {
			for _, employeeId := range employee.Subordinates {
				employeeToFind, _ := employeesMap[employeeId]
				sumPoints += getSumImportanceFrom(employeeToFind, employeesMap, sumImportancePointsMemoize)
			}
		}

		sumImportancePointsMemoize[employee.Id] = sumPoints
	}

	return sumPoints
}

func createEmployeeMapFrom(employees []*Employee) (employeeMap map[int]Employee) {

	employeeMap = make(map[int]Employee)

	for _, employee := range employees {
		employeeMap[employee.Id] = *employee
	}

	return employeeMap
}
