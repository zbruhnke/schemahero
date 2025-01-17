package mysql

import (
	"testing"

	schemasv1alpha2 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha2"

	"github.com/stretchr/testify/assert"
)

func Test_AddForeignKeyStatement(t *testing.T) {
	tests := []struct {
		name              string
		tableName         string
		schemaForeignKey  *schemasv1alpha2.SQLTableForeignKey
		expectedStatement string
	}{
		{
			name:      "no name, one column, no on delete",
			tableName: "t2",
			schemaForeignKey: &schemasv1alpha2.SQLTableForeignKey{
				Columns: []string{
					"c2",
				},
				References: schemasv1alpha2.SQLTableForeignKeyReferences{
					Table: "t1",
					Columns: []string{
						"c1",
					},
				},
			},
			expectedStatement: `alter table t2 add constraint t2_c2_fkey foreign key (c2) references t1 (c1)`,
		},
		{
			name:      "named, one column, no on delete",
			tableName: "t2",
			schemaForeignKey: &schemasv1alpha2.SQLTableForeignKey{
				Name: "hi_i_am_a_fkey",
				Columns: []string{
					"c2",
				},
				References: schemasv1alpha2.SQLTableForeignKeyReferences{
					Table: "t1",
					Columns: []string{
						"c1",
					},
				},
			},
			expectedStatement: `alter table t2 add constraint hi_i_am_a_fkey foreign key (c2) references t1 (c1)`,
		},
		{
			name:      "no name, two columns, no on delete",
			tableName: "t2",
			schemaForeignKey: &schemasv1alpha2.SQLTableForeignKey{
				Columns: []string{
					"c2",
					"c22",
				},
				References: schemasv1alpha2.SQLTableForeignKeyReferences{
					Table: "t1",
					Columns: []string{
						"c1",
						"c11",
					},
				},
			},
			expectedStatement: `alter table t2 add constraint t2_c2_c22_fkey foreign key (c2, c22) references t1 (c1, c11)`,
		},
		{
			name:      "no name, one column, on delete cascade",
			tableName: "t2",
			schemaForeignKey: &schemasv1alpha2.SQLTableForeignKey{
				OnDelete: "cascade",
				Columns: []string{
					"c2",
				},
				References: schemasv1alpha2.SQLTableForeignKeyReferences{
					Table: "t1",
					Columns: []string{
						"c1",
					},
				},
			},
			expectedStatement: `alter table t2 add constraint t2_c2_fkey foreign key (c2) references t1 (c1) on delete cascade`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addForeignKeyStatement := AddForeignKeyStatement(test.tableName, test.schemaForeignKey)

			assert.Equal(t, test.expectedStatement, addForeignKeyStatement)
		})
	}
}
