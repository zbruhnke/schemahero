package mysql

import (
	"fmt"
	"strings"

	schemasv1alpha2 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha2"
	"github.com/schemahero/schemahero/pkg/database/types"
)

func RemoveForeignKeyStatement(tableName string, foreignKey *types.ForeignKey) string {
	return fmt.Sprintf("alter table %s drop constraint %s", tableName, foreignKey.Name)
}

func AddForeignKeyStatement(tableName string, schemaForeignKey *schemasv1alpha2.SQLTableForeignKey) string {
	return fmt.Sprintf("alter table %s add %s", tableName, foreignKeyConstraintClause(tableName, schemaForeignKey))
}

func foreignKeyConstraintClause(tableName string, schemaForeignKey *schemasv1alpha2.SQLTableForeignKey) string {
	onDelete := ""
	if schemaForeignKey.OnDelete != "" {
		onDelete = fmt.Sprintf(" on delete %s", schemaForeignKey.OnDelete)
	}

	return fmt.Sprintf("constraint %s foreign key (%s) references %s (%s)%s",
		types.GenerateFKName(tableName, schemaForeignKey),
		strings.Join(schemaForeignKey.Columns, ", "),
		schemaForeignKey.References.Table,
		strings.Join(schemaForeignKey.References.Columns, ", "),
		onDelete)
}
