// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/schema/db/ent/entgen/usermodel"
)

// UserModel is the model entity for the UserModel schema.
type UserModel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// Email holds the value of the "email" field.
	Email        string `json:"email,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usermodel.FieldID:
			values[i] = new(sql.NullInt64)
		case usermodel.FieldName, usermodel.FieldUserName, usermodel.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserModel fields.
func (um *UserModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usermodel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			um.ID = int(value.Int64)
		case usermodel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				um.Name = value.String
			}
		case usermodel.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				um.UserName = value.String
			}
		case usermodel.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				um.Email = value.String
			}
		default:
			um.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserModel.
// This includes values selected through modifiers, order, etc.
func (um *UserModel) Value(name string) (ent.Value, error) {
	return um.selectValues.Get(name)
}

// Update returns a builder for updating this UserModel.
// Note that you need to call UserModel.Unwrap() before calling this method if this UserModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (um *UserModel) Update() *UserModelUpdateOne {
	return NewUserModelClient(um.config).UpdateOne(um)
}

// Unwrap unwraps the UserModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (um *UserModel) Unwrap() *UserModel {
	_tx, ok := um.config.driver.(*txDriver)
	if !ok {
		panic("entgen: UserModel is not a transactional entity")
	}
	um.config.driver = _tx.drv
	return um
}

// String implements the fmt.Stringer.
func (um *UserModel) String() string {
	var builder strings.Builder
	builder.WriteString("UserModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", um.ID))
	builder.WriteString("name=")
	builder.WriteString(um.Name)
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(um.UserName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(um.Email)
	builder.WriteByte(')')
	return builder.String()
}

// UserModels is a parsable slice of UserModel.
type UserModels []*UserModel