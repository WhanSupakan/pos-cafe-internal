package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type EnumStatus string

const (
	StatusActive    EnumStatus = "ACTIVE"
	StatusSuspended EnumStatus = "SUSPENDED"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("code").NotEmpty().Unique(),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("username").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.UUID("role_id", uuid.UUID{}),
		field.String("status").NotEmpty().Default(string(StatusActive)).Validate(validateStatus),
		field.Time("created_at").Default(time.Now()).SchemaType(map[string]string{
			dialect.Postgres: "timestamptz",
		}),
		field.Time("updated_at").Default(time.Now()).SchemaType(map[string]string{
			dialect.Postgres: "timestamptz",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Validator for the User status.
func validateStatus(s string) error {
	switch s {
	case string(StatusActive):
		return nil
	case string(StatusSuspended):
		return nil
	default:
		return fmt.Errorf("invalid status: %s", s)
	}
}
