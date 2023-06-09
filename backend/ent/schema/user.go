package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.String("email").Unique().MaxLen(255).StructTag(`json:"email"`),
		field.String("hashed_pw").MaxLen(255).Sensitive(),
		field.Bool("email_verified").Default(false).StructTag(`json:"emailVerified"`),
		field.Bool("is_archived").Default(false).StructTag(`json:"isArchived"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("orders", Order.Type),
		edge.To("siteui", Siteui.Type).Unique(),
		edge.To("imagesinfo", Imageinfo.Type),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// not marshal edges
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
