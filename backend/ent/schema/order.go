package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.UUID("user_id", uuid.UUID{}).StructTag(`json:"userId"`),
		field.Float("discount").Min(0.0).Max(1.0).StructTag(`json:"discount"`),
		field.Float("total_amount").Min(0.0).StructTag(`json:"totalAmount"`),
		field.String("remark").MaxLen(255).StructTag(`json:"remark"`),
		field.String("status").MaxLen(255).StructTag(`json:"status"`),
		field.String("payment_status").MaxLen(255).StructTag(`json:"paymentStatus"`),
		field.String("payment_method").MaxLen(255).StructTag(`json:"paymentMethod"`),
		field.String("delivery_status").MaxLen(255).StructTag(`json:"deliveryStatus"`),
		field.String("shipping_address").MaxLen(512).StructTag(`json:"shippingAddress"`),
		field.String("tracking_number").MaxLen(255).StructTag(`json:"trackingNumber"`),
		field.Bool("is_archived").Default(false).StructTag(`json:"isArchived"`),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("orders").
			Unique().
			Field("user_id").
			Required(),
		edge.To("orderitems", OrderItem.Type),
	}
}

// Annotations of the OrderItem.
func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// not marshal edges
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
