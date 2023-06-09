package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Imageinfo holds the schema definition for the Imageinfo entity.
type Imageinfo struct {
	ent.Schema
}

// Indexes of the Imageinfo.
func (Imageinfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "img_name").Unique(),
	}
}

// Mixin of the Imageinfo.
func (Imageinfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Imageinfo.
func (Imageinfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).StructTag(`json:"userId"`),
		field.String("img_url").MaxLen(512).StructTag(`json:"imgUrl"`),
		field.String("img_name").MaxLen(128).StructTag(`json:"imgName"`),
		field.Int64("img_size").NonNegative().StructTag(`json:"imgSize"`),
		field.String("img_s3_id_key").Unique().MaxLen(1024).StructTag(`json:"imgS3IdKey"`),
	}
}

// Edges of the Imageinfo.
func (Imageinfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("imagesinfo").
			Unique().
			Field("user_id").
			Required(),
	}
}

// Annotations of the OrderItem.
func (Imageinfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// not marshal edges
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
