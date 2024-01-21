/*
LeetCode 690: https://leetcode.com/problems/employee-importance/description/
*/

package leetcode

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func buildEmployeeGraph(employees []*Employee) map[int]*Employee {
	graph := make(map[int]*Employee)
	for _, employee := range employees {
		graph[employee.Id] = employee
	}

	return graph
}

func getEmployeeImportance(graph map[int]*Employee, employee *Employee) int {
	importance := employee.Importance
	for _, id := range employee.Subordinates {
		subordinate := graph[id]
		importance += getEmployeeImportance(graph, subordinate)
	}

	return importance
}

func getImportance(employees []*Employee, id int) int {
	graph := buildEmployeeGraph(employees)
	return getEmployeeImportance(graph, graph[id])
}
