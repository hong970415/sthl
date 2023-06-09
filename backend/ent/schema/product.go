package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.UUID("user_id", uuid.UUID{}).StructTag(`json:"userId"`),
		field.String("name").MaxLen(255).StructTag(`json:"name"`),
		field.Float("price").Min(1).Max(10000000.0).StructTag(`json:"price"`),
		field.Int32("quantity").Min(0).Max(10000000).StructTag(`json:"quantity"`),
		field.String("description").MaxLen(255).StructTag(`json:"description"`),
		field.String("status").MaxLen(255).StructTag(`json:"status"`),
		field.Bool("is_archived").Default(false).StructTag(`json:"isArchived"`),
		field.String("img_url").MaxLen(512).Default("").StructTag(`json:"imgUrl"`),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("products").
			Unique().
			Field("user_id").
			Required(),
	}
}

// Annotations of the Product.
func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// not marshal edges
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
