package tests

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"zeus/pkg/service"
)

func TestPermService_CheckPermByUid(t *testing.T) {
	ps := service.PermService{}
	if ps.CheckPermByUid(6, "/admin/image/system", "lvyou-admin") {
		t.Log("success")
	} else {
		t.Log("no permission")
	}
}
func TestPermService_Transform(t *testing.T) {
	ps := service.PermService{}
	t.Log(ps.TransformPerm("/parking/user/list"))
}
