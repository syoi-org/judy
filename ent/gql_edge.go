// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (j *Judge) Problems(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int,
) (*ProblemConnection, error) {
	opts := []ProblemPaginateOption{}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := j.Edges.totalCount[0][alias]
	if nodes, err := j.NamedProblems(alias); err == nil || hasTotalCount {
		pager, err := newProblemPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ProblemConnection{Edges: []*ProblemEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return j.QueryProblems().Paginate(ctx, after, first, before, last, opts...)
}

func (pr *Problem) Submissions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int,
) (*SubmissionConnection, error) {
	opts := []SubmissionPaginateOption{}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := pr.Edges.totalCount[0][alias]
	if nodes, err := pr.NamedSubmissions(alias); err == nil || hasTotalCount {
		pager, err := newSubmissionPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &SubmissionConnection{Edges: []*SubmissionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return pr.QuerySubmissions().Paginate(ctx, after, first, before, last, opts...)
}

func (pr *Problem) Judge(ctx context.Context) (*Judge, error) {
	result, err := pr.Edges.JudgeOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryJudge().Only(ctx)
	}
	return result, err
}

func (s *Submission) Problem(ctx context.Context) (*Problem, error) {
	result, err := s.Edges.ProblemOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryProblem().Only(ctx)
	}
	return result, err
}
