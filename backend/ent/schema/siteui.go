package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Siteui holds the schema definition for the Siteui entity.
type Siteui struct {
	ent.Schema
}

// Indexes of the Siteui.
func (Siteui) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").Unique(),
	}
}

// Mixin of the Siteui.
func (Siteui) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Siteui.
func (Siteui) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().StructTag(`json:"id"`),
		field.UUID("user_id", uuid.UUID{}).Unique().StructTag(`json:"userId"`),
		field.String("sitename").MaxLen(32).StructTag(`json:"sitename"`),
		field.String("homepageImgUrl").MaxLen(512).Default("").StructTag(`json:"homepageImgUrl"`),
		field.String("homepageText").MaxLen(255).StructTag(`json:"homepageText"`),
		field.String("homepageTextColor").MaxLen(16).StructTag(`json:"homepageTextColor"`),
	}
}

// Edges of the Siteui.
func (Siteui) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("siteui").
			Unique().
			Field("user_id").
			Required(),
	}

}

// Annotations of the Siteui.
func (Siteui) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// not marshal edges
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
