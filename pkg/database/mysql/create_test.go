package mysql

import (
	"testing"

	schemasv1alpha2 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateTableStatement(t *testing.T) {
	tests := []struct {
		name              string
		tableSchema       *schemasv1alpha2.SQLTableSchema
		tableName         string
		expectedStatement string
	}{
		{
			name: "simple",
			tableSchema: &schemasv1alpha2.SQLTableSchema{
				PrimaryKey: []string{
					"id",
				},
				Columns: []*schemasv1alpha2.SQLTableColumn{
					&schemasv1alpha2.SQLTableColumn{
						Name: "id",
						Type: "integer",
					},
				},
			},
			tableName:         "simple",
			expectedStatement: "create table `simple` (`id` int (11), primary key (`id`))",
		},
		{
			name: "composite primary key",
			tableSchema: &schemasv1alpha2.SQLTableSchema{
				PrimaryKey: []string{
					"one",
					"two",
				},
				Columns: []*schemasv1alpha2.SQLTableColumn{
					&schemasv1alpha2.SQLTableColumn{
						Name: "one",
						Type: "integer",
					},
					&schemasv1alpha2.SQLTableColumn{
						Name: "two",
						Type: "integer",
					},
					&schemasv1alpha2.SQLTableColumn{
						Name: "three",
						Type: "varchar(255)",
					},
				},
			},
			tableName:         "composite_primary_key",
			expectedStatement: "create table `composite_primary_key` (`one` int (11), `two` int (11), `three` varchar (255), primary key (`one`, `two`))",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)

			createTableStatement, err := CreateTableStatement(test.tableName, test.tableSchema)
			req.NoError(err)

			assert.Equal(t, test.expectedStatement, createTableStatement)
		})
	}
}
