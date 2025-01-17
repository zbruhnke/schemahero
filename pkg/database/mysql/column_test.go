package mysql

import (
	"testing"

	schemasv1alpha2 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha2"
	"github.com/schemahero/schemahero/pkg/database/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mysqlColumnAsInsert(t *testing.T) {
	default11 := "11"
	tests := []struct {
		name              string
		column            *schemasv1alpha2.SQLTableColumn
		expectedStatement string
	}{
		{
			name: "simple",
			column: &schemasv1alpha2.SQLTableColumn{
				Name: "c",
				Type: "integer",
				Constraints: &schemasv1alpha2.SQLTableColumnConstraints{
					NotNull: &trueValue,
				},
				Default: &default11,
			},
			expectedStatement: "`c` int (11) not null default '11'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)

			generatedStatement, err := mysqlColumnAsInsert(test.column)
			req.NoError(err)
			assert.Equal(t, test.expectedStatement, generatedStatement)
		})
	}
}

func Test_InsertColumnStatement(t *testing.T) {
	tests := []struct {
		name              string
		tableName         string
		desiredColumn     *schemasv1alpha2.SQLTableColumn
		expectedStatement string
	}{
		{
			name:      "add column",
			tableName: "t",
			desiredColumn: &schemasv1alpha2.SQLTableColumn{
				Name: "a",
				Type: "integer",
			},
			expectedStatement: "alter table `t` add column `a` int (11)",
		},
		{
			name:      "add not null column",
			tableName: "t",
			desiredColumn: &schemasv1alpha2.SQLTableColumn{
				Name: "a",
				Type: "integer",
				Constraints: &schemasv1alpha2.SQLTableColumnConstraints{
					NotNull: &trueValue,
				},
			},
			expectedStatement: "alter table `t` add column `a` int (11) not null",
		},
		{
			name:      "add null column",
			tableName: "t",
			desiredColumn: &schemasv1alpha2.SQLTableColumn{
				Name: "a",
				Type: "integer",
				Constraints: &schemasv1alpha2.SQLTableColumnConstraints{
					NotNull: &falseValue,
				},
			},
			expectedStatement: "alter table `t` add column `a` int (11) null",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)

			generatedStatement, err := InsertColumnStatement(test.tableName, test.desiredColumn)
			req.NoError(err)
			assert.Equal(t, test.expectedStatement, generatedStatement)
		})
	}
}

func Test_schemaColumnToMysqlColumn(t *testing.T) {
	tests := []struct {
		name           string
		schemaColumn   *schemasv1alpha2.SQLTableColumn
		expectedColumn *types.Column
	}{
		{
			name: "varchar (10)",
			schemaColumn: &schemasv1alpha2.SQLTableColumn{
				Name: "vc",
				Type: "varchar (10)",
			},
			expectedColumn: &types.Column{
				Name:          "vc",
				DataType:      "varchar (10)",
				ColumnDefault: nil,
			},
		},
		{
			name: "bool",
			schemaColumn: &schemasv1alpha2.SQLTableColumn{
				Name: "b",
				Type: "bool",
			},
			expectedColumn: &types.Column{
				Name:          "b",
				DataType:      "tinyint (1)",
				ColumnDefault: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)

			column, err := schemaColumnToColumn(test.schemaColumn)
			req.NoError(err)
			assert.Equal(t, test.expectedColumn, column)
		})
	}

}
