package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Submission holds the schema definition for the Submission entity.
type Submission struct {
	ent.Schema
}

// Fields of the Submission.
func (Submission) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("pending", "compiling", "judging", "finished").Default("pending"),
		field.Enum("verdict").Values("OK", "TLE", "MLE", "ILE", "WA", "CE", "RE", "PE", "CRASHED", "OTHER").Default("OK"),
		field.Int("test_count").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Submission.
func (Submission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("problem", Problem.Type).Ref("submissions").Unique().Required(),
	}
}
