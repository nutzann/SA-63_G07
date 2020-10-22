// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AreasColumns holds the columns for the "areas" table.
	AreasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "disease_area", Type: field.TypeInt, Nullable: true},
		{Name: "employee_area", Type: field.TypeInt, Nullable: true},
		{Name: "level_area", Type: field.TypeInt, Nullable: true},
		{Name: "statistic_area", Type: field.TypeInt, Nullable: true},
	}
	// AreasTable holds the schema information for the "areas" table.
	AreasTable = &schema.Table{
		Name:       "areas",
		Columns:    AreasColumns,
		PrimaryKey: []*schema.Column{AreasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "areas_diseases_area",
				Columns: []*schema.Column{AreasColumns[2]},

				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "areas_employees_area",
				Columns: []*schema.Column{AreasColumns[3]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "areas_levels_area",
				Columns: []*schema.Column{AreasColumns[4]},

				RefColumns: []*schema.Column{LevelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "areas_statistics_area",
				Columns: []*schema.Column{AreasColumns[5]},

				RefColumns: []*schema.Column{StatisticsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:        "diseases",
		Columns:     DiseasesColumns,
		PrimaryKey:  []*schema.Column{DiseasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:        "employees",
		Columns:     EmployeesColumns,
		PrimaryKey:  []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// LevelsColumns holds the columns for the "levels" table.
	LevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// LevelsTable holds the schema information for the "levels" table.
	LevelsTable = &schema.Table{
		Name:        "levels",
		Columns:     LevelsColumns,
		PrimaryKey:  []*schema.Column{LevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatisticsColumns holds the columns for the "statistics" table.
	StatisticsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// StatisticsTable holds the schema information for the "statistics" table.
	StatisticsTable = &schema.Table{
		Name:        "statistics",
		Columns:     StatisticsColumns,
		PrimaryKey:  []*schema.Column{StatisticsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AreasTable,
		DiseasesTable,
		EmployeesTable,
		LevelsTable,
		StatisticsTable,
	}
)

func init() {
	AreasTable.ForeignKeys[0].RefTable = DiseasesTable
	AreasTable.ForeignKeys[1].RefTable = EmployeesTable
	AreasTable.ForeignKeys[2].RefTable = LevelsTable
	AreasTable.ForeignKeys[3].RefTable = StatisticsTable
}
