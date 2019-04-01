// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package beegoormadapter

import (
	"errors"
	"runtime"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/lib/pq"
)

type CasbinRule struct {
	Id    int
	PType string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

func init() {
	orm.RegisterModel(new(CasbinRule))

	orm.RegisterDriver("mysql", orm.DRMySQL)
}

// Adapter represents the Xorm adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbSpecified    bool
	o              orm.Ormer
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
}

// NewAdapter is the constructor for Adapter.
// dbSpecified is an optional bool parameter. The default value is false.
// It's up to whether you have specified an existing DB in dataSourceName.
// If dbSpecified == true, you need to make sure the DB in dataSourceName exists.
// If dbSpecified == false, the adapter will automatically create a DB named "casbin".
func NewAdapter(driverName string, dataSourceName string, dbSpecified ...bool) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName

	if len(dbSpecified) == 0 {
		a.dbSpecified = false
	} else if len(dbSpecified) == 1 {
		a.dbSpecified = dbSpecified[0]
	} else {
		panic(errors.New("invalid parameter: dbSpecified"))
	}

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) registerDataBase(aliasName, driverName, dataSource string, params ...int) error {
	err := orm.RegisterDataBase(aliasName, driverName, dataSource, params...)
	if err != nil && strings.HasSuffix(err.Error(), "already registered, cannot reuse") {
		return nil
	}
	return err
}

func (a *Adapter) createDatabase() error {
	var err error
	var o orm.Ormer
	if a.driverName == "postgres" {
		err = a.registerDataBase("create_casbin", a.driverName, a.dataSourceName + " dbname=postgres")
	} else {
		err = a.registerDataBase("create_casbin", a.driverName, a.dataSourceName)
	}
	if err != nil {
		return err
	}
	o = orm.NewOrm()
	if a.driverName == "postgres" {
		if 		_, err = o.Raw("CREATE DATABASE casbin").Exec(); err != nil {
			// 42P04 is	duplicate_database
			if err.(*pq.Error).Code == "42P04" {
				return nil
			}
		}
	} else {
		_, err = o.Raw("CREATE DATABASE IF NOT EXISTS casbin").Exec()
	}
	return err
}

func (a *Adapter) open() {
	var err error

	err = a.registerDataBase("default", a.driverName, a.dataSourceName)
	if err != nil {
		panic(err)
	}

	if a.dbSpecified {
		err = a.registerDataBase("casbin", a.driverName, a.dataSourceName)
		if err != nil {
			panic(err)
		}
	} else {
		if err = a.createDatabase(); err != nil {
			panic(err)
		}

		if a.driverName == "postgres" {
			err = a.registerDataBase("casbin", a.driverName, a.dataSourceName + " dbname=casbin")
		} else {
			err = a.registerDataBase("casbin", a.driverName, a.dataSourceName + "casbin")
		}
		if err != nil {
			panic(err)
		}
	}

	a.o = orm.NewOrm()
	a.o.Using("casbin")

	a.createTable()
}

func (a *Adapter) close() {
	a.o = nil
}

func (a *Adapter) createTable() {
	//err := orm.RunSyncdb("casbin", false, true)
	//if err != nil {
	//	panic(err)
	//}
}

func (a *Adapter) dropTable() {
	//err := orm.RunSyncdb("casbin", true, true)
	//if err != nil {
	//	panic(err)
	//}
	o := orm.NewOrm()
	o.Raw("truncate table casbin_rule").Exec()
}

func loadPolicyLine(line CasbinRule, model model.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var lines []CasbinRule
	_, err := a.o.QueryTable("casbin_rule").All(&lines)
	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	return nil
}

func savePolicyLine(ptype string, rule []string) CasbinRule {
	line := CasbinRule{}

	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return line
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	a.dropTable()
	//a.createTable()

	var lines []CasbinRule

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	_, err := a.o.InsertMulti(len(lines), lines)
	return err
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.o.Insert(&line)
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.o.Delete(&line, "p_type", "v0", "v1", "v2", "v3", "v4", "v5")
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := CasbinRule{}

	line.PType = ptype
	filter := []string{}
	filter = append(filter, "p_type")
	if fieldIndex <= 0 && 0 < fieldIndex + len(fieldValues) {
		line.V0 = fieldValues[0 - fieldIndex]
		if line.V0 != "" {
			filter = append(filter, "v0")
		}
	}
	if fieldIndex <= 1 && 1 < fieldIndex + len(fieldValues) {
		line.V1 = fieldValues[1 - fieldIndex]
		if line.V1 != "" {
			filter = append(filter, "v1")
		}
	}
	if fieldIndex <= 2 && 2 < fieldIndex + len(fieldValues) {
		line.V2 = fieldValues[2 - fieldIndex]
		if line.V2 != "" {
			filter = append(filter, "v2")
		}
	}
	if fieldIndex <= 3 && 3 < fieldIndex + len(fieldValues) {
		line.V3 = fieldValues[3 - fieldIndex]
		if line.V3 != "" {
			filter = append(filter, "v3")
		}
	}
	if fieldIndex <= 4 && 4 < fieldIndex + len(fieldValues) {
		line.V4 = fieldValues[4 - fieldIndex]
		if line.V4 != "" {
			filter = append(filter, "v4")
		}
	}
	if fieldIndex <= 5 && 5 < fieldIndex + len(fieldValues) {
		line.V5 = fieldValues[5 - fieldIndex]
		if line.V5 != "" {
			filter = append(filter, "v5")
		}
	}

	_, err := a.o.Delete(&line, filter...)
	return err
}
