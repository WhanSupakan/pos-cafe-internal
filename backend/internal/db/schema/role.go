package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").NotEmpty().Unique(),
		field.String("description").NotEmpty(),
		field.Time("created_at").Default(time.Now()).SchemaType(map[string]string{
			dialect.Postgres: "timestamptz",
		}),
		field.Time("updated_at").Default(time.Now()).SchemaType(map[string]string{
			dialect.Postgres: "timestamptz",
		}),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
