package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname").MaxLen(20),
		field.String("lastname").MaxLen(20),
		field.String("email").MaxLen(20).Unique(),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
