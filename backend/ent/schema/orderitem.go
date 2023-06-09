package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.UUID("order_id", uuid.UUID{}).StructTag(`json:"orderId"`),
		field.UUID("product_id", uuid.UUID{}).StructTag(`json:"productId"`),
		field.String("purchased_name").MaxLen(255).StructTag(`json:"purchasedName"`),
		field.Float("purchased_price").Min(0.0).StructTag(`json:"purchasedPrice"`),
		field.Int("quantity").Min(1).StructTag(`json:"quantity"`),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Order.Type).
			Ref("orderitems").
			Unique().
			Field("order_id").
			Required(),
	}
}

// Annotations of the OrderItem.
func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// not marshal edges
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
