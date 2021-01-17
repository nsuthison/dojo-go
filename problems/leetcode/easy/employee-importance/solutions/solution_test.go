package solutions

import (
	"testing"

	. "github.com/nsuthison/dojo-go/problems/leetcode/easy/employee-importance/models"
	. "github.com/smartystreets/goconvey/convey"
)

var getImportanceTestCases = []struct {
	employees      []*Employee
	id             int
	expectedResult int
}{
	{
		employees: []*Employee{
			{
				Id:           1,
				Importance:   5,
				Subordinates: []int{2, 3},
			},
			{
				Id:           2,
				Importance:   3,
				Subordinates: []int{},
			},
			{
				Id:           3,
				Importance:   3,
				Subordinates: []int{},
			},
		},
		id:             3,
		expectedResult: 11,
	},
}

func Test_getImportance(t *testing.T) {
	for _, testCase := range getImportanceTestCases {
		t.Run("getImportance test", func(t *testing.T) {
			Convey("Given employees and employee id to get importance", t, func() {
				employees := testCase.employees
				id := testCase.id

				Convey("When find importance", func() {
					importnacePoints := getImportance(employees, id)

					Convey("Then importancePoints should equal to summation of the employee and all subordinates", func() {
						So(importnacePoints, ShouldEqual, testCase.expectedResult)
					})
				})
			})
		})
	}
}
