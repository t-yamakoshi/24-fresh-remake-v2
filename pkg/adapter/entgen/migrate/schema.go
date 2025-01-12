// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserModelsColumns holds the columns for the "user_models" table.
	UserModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "bigint unsigned"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "cognito_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
		{Name: "birth_date", Type: field.TypeTime},
		{Name: "gender", Type: field.TypeEnum, Nullable: true, Enums: []string{"men", "women", "other"}},
		{Name: "self_introduction", Type: field.TypeString, Nullable: true},
		{Name: "profile_image", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString},
	}
	// UserModelsTable holds the schema information for the "user_models" table.
	UserModelsTable = &schema.Table{
		Name:       "user_models",
		Columns:    UserModelsColumns,
		PrimaryKey: []*schema.Column{UserModelsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserModelsTable,
	}
)

func init() {
}