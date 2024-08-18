// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/syoi-org/judge.syoi.org/ent/problem"
	"github.com/syoi-org/judge.syoi.org/ent/submission"
)

// SubmissionCreate is the builder for creating a Submission entity.
type SubmissionCreate struct {
	config
	mutation *SubmissionMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (sc *SubmissionCreate) SetStatus(s submission.Status) *SubmissionCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableStatus(s *submission.Status) *SubmissionCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetVerdict sets the "verdict" field.
func (sc *SubmissionCreate) SetVerdict(s submission.Verdict) *SubmissionCreate {
	sc.mutation.SetVerdict(s)
	return sc
}

// SetNillableVerdict sets the "verdict" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableVerdict(s *submission.Verdict) *SubmissionCreate {
	if s != nil {
		sc.SetVerdict(*s)
	}
	return sc
}

// SetTestCount sets the "test_count" field.
func (sc *SubmissionCreate) SetTestCount(i int) *SubmissionCreate {
	sc.mutation.SetTestCount(i)
	return sc
}

// SetNillableTestCount sets the "test_count" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableTestCount(i *int) *SubmissionCreate {
	if i != nil {
		sc.SetTestCount(*i)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubmissionCreate) SetCreatedAt(t time.Time) *SubmissionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableCreatedAt(t *time.Time) *SubmissionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubmissionCreate) SetUpdatedAt(t time.Time) *SubmissionCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableUpdatedAt(t *time.Time) *SubmissionCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetProblemID sets the "problem" edge to the Problem entity by ID.
func (sc *SubmissionCreate) SetProblemID(id int) *SubmissionCreate {
	sc.mutation.SetProblemID(id)
	return sc
}

// SetProblem sets the "problem" edge to the Problem entity.
func (sc *SubmissionCreate) SetProblem(p *Problem) *SubmissionCreate {
	return sc.SetProblemID(p.ID)
}

// Mutation returns the SubmissionMutation object of the builder.
func (sc *SubmissionCreate) Mutation() *SubmissionMutation {
	return sc.mutation
}

// Save creates the Submission in the database.
func (sc *SubmissionCreate) Save(ctx context.Context) (*Submission, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubmissionCreate) SaveX(ctx context.Context) *Submission {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubmissionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubmissionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubmissionCreate) defaults() {
	if _, ok := sc.mutation.Status(); !ok {
		v := submission.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Verdict(); !ok {
		v := submission.DefaultVerdict
		sc.mutation.SetVerdict(v)
	}
	if _, ok := sc.mutation.TestCount(); !ok {
		v := submission.DefaultTestCount
		sc.mutation.SetTestCount(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := submission.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := submission.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubmissionCreate) check() error {
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Submission.status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := submission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Submission.status": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Verdict(); !ok {
		return &ValidationError{Name: "verdict", err: errors.New(`ent: missing required field "Submission.verdict"`)}
	}
	if v, ok := sc.mutation.Verdict(); ok {
		if err := submission.VerdictValidator(v); err != nil {
			return &ValidationError{Name: "verdict", err: fmt.Errorf(`ent: validator failed for field "Submission.verdict": %w`, err)}
		}
	}
	if _, ok := sc.mutation.TestCount(); !ok {
		return &ValidationError{Name: "test_count", err: errors.New(`ent: missing required field "Submission.test_count"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Submission.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Submission.updated_at"`)}
	}
	if len(sc.mutation.ProblemIDs()) == 0 {
		return &ValidationError{Name: "problem", err: errors.New(`ent: missing required edge "Submission.problem"`)}
	}
	return nil
}

func (sc *SubmissionCreate) sqlSave(ctx context.Context) (*Submission, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubmissionCreate) createSpec() (*Submission, *sqlgraph.CreateSpec) {
	var (
		_node = &Submission{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(submission.Table, sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(submission.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Verdict(); ok {
		_spec.SetField(submission.FieldVerdict, field.TypeEnum, value)
		_node.Verdict = value
	}
	if value, ok := sc.mutation.TestCount(); ok {
		_spec.SetField(submission.FieldTestCount, field.TypeInt, value)
		_node.TestCount = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(submission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(submission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := sc.mutation.ProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submission.ProblemTable,
			Columns: []string{submission.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.problem_submissions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubmissionCreateBulk is the builder for creating many Submission entities in bulk.
type SubmissionCreateBulk struct {
	config
	err      error
	builders []*SubmissionCreate
}

// Save creates the Submission entities in the database.
func (scb *SubmissionCreateBulk) Save(ctx context.Context) ([]*Submission, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Submission, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubmissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubmissionCreateBulk) SaveX(ctx context.Context) []*Submission {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubmissionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubmissionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
