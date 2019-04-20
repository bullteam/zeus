package tests

import "testing"

func TestDeptService_CheckDepartmentHasPeopleInside(t *testing.T) {
	if deptService.CheckIfPeopleInside(1) {
		t.Log("Folks in department")
	} else {
		t.Log("None belong to specific department")
	}
}
