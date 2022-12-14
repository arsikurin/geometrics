// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Problem is an object representing the database table.
type Problem struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string `boil:"description" json:"description" toml:"description" yaml:"description"`
	SolutionRaw string `boil:"solution_raw" json:"solution_raw" toml:"solution_raw" yaml:"solution_raw"`
	Toolbar     string `boil:"toolbar" json:"toolbar" toml:"toolbar" yaml:"toolbar"`

	R *problemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L problemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProblemColumns = struct {
	ID          string
	Name        string
	Description string
	SolutionRaw string
	Toolbar     string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	SolutionRaw: "solution_raw",
	Toolbar:     "toolbar",
}

var ProblemTableColumns = struct {
	ID          string
	Name        string
	Description string
	SolutionRaw string
	Toolbar     string
}{
	ID:          "problems.id",
	Name:        "problems.name",
	Description: "problems.description",
	SolutionRaw: "problems.solution_raw",
	Toolbar:     "problems.toolbar",
}

// Generated where

var ProblemWhere = struct {
	ID          whereHelperint
	Name        whereHelperstring
	Description whereHelperstring
	SolutionRaw whereHelperstring
	Toolbar     whereHelperstring
}{
	ID:          whereHelperint{field: "\"problems\".\"id\""},
	Name:        whereHelperstring{field: "\"problems\".\"name\""},
	Description: whereHelperstring{field: "\"problems\".\"description\""},
	SolutionRaw: whereHelperstring{field: "\"problems\".\"solution_raw\""},
	Toolbar:     whereHelperstring{field: "\"problems\".\"toolbar\""},
}

// ProblemRels is where relationship names are stored.
var ProblemRels = struct {
	CoursesProblems string
	Submits         string
}{
	CoursesProblems: "CoursesProblems",
	Submits:         "Submits",
}

// problemR is where relationships are stored.
type problemR struct {
	CoursesProblems CoursesProblemSlice `boil:"CoursesProblems" json:"CoursesProblems" toml:"CoursesProblems" yaml:"CoursesProblems"`
	Submits         SubmitSlice         `boil:"Submits" json:"Submits" toml:"Submits" yaml:"Submits"`
}

// NewStruct creates a new relationship struct
func (*problemR) NewStruct() *problemR {
	return &problemR{}
}

func (r *problemR) GetCoursesProblems() CoursesProblemSlice {
	if r == nil {
		return nil
	}
	return r.CoursesProblems
}

func (r *problemR) GetSubmits() SubmitSlice {
	if r == nil {
		return nil
	}
	return r.Submits
}

// problemL is where Load methods for each relationship are stored.
type problemL struct{}

var (
	problemAllColumns            = []string{"id", "name", "description", "solution_raw", "toolbar"}
	problemColumnsWithoutDefault = []string{"name", "description", "solution_raw"}
	problemColumnsWithDefault    = []string{"id", "toolbar"}
	problemPrimaryKeyColumns     = []string{"id"}
	problemGeneratedColumns      = []string{}
)

type (
	// ProblemSlice is an alias for a slice of pointers to Problem.
	// This should almost always be used instead of []Problem.
	ProblemSlice []*Problem
	// ProblemHook is the signature for custom Problem hook methods
	ProblemHook func(context.Context, boil.ContextExecutor, *Problem) error

	problemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	problemType                 = reflect.TypeOf(&Problem{})
	problemMapping              = queries.MakeStructMapping(problemType)
	problemPrimaryKeyMapping, _ = queries.BindMapping(problemType, problemMapping, problemPrimaryKeyColumns)
	problemInsertCacheMut       sync.RWMutex
	problemInsertCache          = make(map[string]insertCache)
	problemUpdateCacheMut       sync.RWMutex
	problemUpdateCache          = make(map[string]updateCache)
	problemUpsertCacheMut       sync.RWMutex
	problemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var problemAfterSelectHooks []ProblemHook

var problemBeforeInsertHooks []ProblemHook
var problemAfterInsertHooks []ProblemHook

var problemBeforeUpdateHooks []ProblemHook
var problemAfterUpdateHooks []ProblemHook

var problemBeforeDeleteHooks []ProblemHook
var problemAfterDeleteHooks []ProblemHook

var problemBeforeUpsertHooks []ProblemHook
var problemAfterUpsertHooks []ProblemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Problem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Problem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Problem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Problem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Problem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Problem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Problem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Problem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Problem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProblemHook registers your hook function for all future operations.
func AddProblemHook(hookPoint boil.HookPoint, problemHook ProblemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		problemAfterSelectHooks = append(problemAfterSelectHooks, problemHook)
	case boil.BeforeInsertHook:
		problemBeforeInsertHooks = append(problemBeforeInsertHooks, problemHook)
	case boil.AfterInsertHook:
		problemAfterInsertHooks = append(problemAfterInsertHooks, problemHook)
	case boil.BeforeUpdateHook:
		problemBeforeUpdateHooks = append(problemBeforeUpdateHooks, problemHook)
	case boil.AfterUpdateHook:
		problemAfterUpdateHooks = append(problemAfterUpdateHooks, problemHook)
	case boil.BeforeDeleteHook:
		problemBeforeDeleteHooks = append(problemBeforeDeleteHooks, problemHook)
	case boil.AfterDeleteHook:
		problemAfterDeleteHooks = append(problemAfterDeleteHooks, problemHook)
	case boil.BeforeUpsertHook:
		problemBeforeUpsertHooks = append(problemBeforeUpsertHooks, problemHook)
	case boil.AfterUpsertHook:
		problemAfterUpsertHooks = append(problemAfterUpsertHooks, problemHook)
	}
}

// OneG returns a single problem record from the query using the global executor.
func (q problemQuery) OneG(ctx context.Context) (*Problem, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single problem record from the query.
func (q problemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Problem, error) {
	o := &Problem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for problems")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Problem records from the query using the global executor.
func (q problemQuery) AllG(ctx context.Context) (ProblemSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Problem records from the query.
func (q problemQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProblemSlice, error) {
	var o []*Problem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Problem slice")
	}

	if len(problemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Problem records in the query using the global executor
func (q problemQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Problem records in the query.
func (q problemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count problems rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q problemQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q problemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if problems exists")
	}

	return count > 0, nil
}

// CoursesProblems retrieves all the courses_problem's CoursesProblems with an executor.
func (o *Problem) CoursesProblems(mods ...qm.QueryMod) coursesProblemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"courses_problems\".\"problem_id\"=?", o.ID),
	)

	return CoursesProblems(queryMods...)
}

// Submits retrieves all the submit's Submits with an executor.
func (o *Problem) Submits(mods ...qm.QueryMod) submitQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"submits\".\"problem_id\"=?", o.ID),
	)

	return Submits(queryMods...)
}

// LoadCoursesProblems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (problemL) LoadCoursesProblems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProblem interface{}, mods queries.Applicator) error {
	var slice []*Problem
	var object *Problem

	if singular {
		var ok bool
		object, ok = maybeProblem.(*Problem)
		if !ok {
			object = new(Problem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProblem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProblem))
			}
		}
	} else {
		s, ok := maybeProblem.(*[]*Problem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProblem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProblem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &problemR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &problemR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`courses_problems`),
		qm.WhereIn(`courses_problems.problem_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load courses_problems")
	}

	var resultSlice []*CoursesProblem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice courses_problems")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on courses_problems")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for courses_problems")
	}

	if len(coursesProblemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CoursesProblems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &coursesProblemR{}
			}
			foreign.R.Problem = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProblemID {
				local.R.CoursesProblems = append(local.R.CoursesProblems, foreign)
				if foreign.R == nil {
					foreign.R = &coursesProblemR{}
				}
				foreign.R.Problem = local
				break
			}
		}
	}

	return nil
}

// LoadSubmits allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (problemL) LoadSubmits(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProblem interface{}, mods queries.Applicator) error {
	var slice []*Problem
	var object *Problem

	if singular {
		var ok bool
		object, ok = maybeProblem.(*Problem)
		if !ok {
			object = new(Problem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProblem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProblem))
			}
		}
	} else {
		s, ok := maybeProblem.(*[]*Problem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProblem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProblem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &problemR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &problemR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`submits`),
		qm.WhereIn(`submits.problem_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load submits")
	}

	var resultSlice []*Submit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice submits")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on submits")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for submits")
	}

	if len(submitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Submits = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &submitR{}
			}
			foreign.R.Problem = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProblemID {
				local.R.Submits = append(local.R.Submits, foreign)
				if foreign.R == nil {
					foreign.R = &submitR{}
				}
				foreign.R.Problem = local
				break
			}
		}
	}

	return nil
}

// AddCoursesProblemsG adds the given related objects to the existing relationships
// of the problem, optionally inserting them as new records.
// Appends related to o.R.CoursesProblems.
// Sets related.R.Problem appropriately.
// Uses the global database handle.
func (o *Problem) AddCoursesProblemsG(ctx context.Context, insert bool, related ...*CoursesProblem) error {
	return o.AddCoursesProblems(ctx, boil.GetContextDB(), insert, related...)
}

// AddCoursesProblems adds the given related objects to the existing relationships
// of the problem, optionally inserting them as new records.
// Appends related to o.R.CoursesProblems.
// Sets related.R.Problem appropriately.
func (o *Problem) AddCoursesProblems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CoursesProblem) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProblemID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"courses_problems\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"problem_id"}),
				strmangle.WhereClause("\"", "\"", 2, coursesProblemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProblemID = o.ID
		}
	}

	if o.R == nil {
		o.R = &problemR{
			CoursesProblems: related,
		}
	} else {
		o.R.CoursesProblems = append(o.R.CoursesProblems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &coursesProblemR{
				Problem: o,
			}
		} else {
			rel.R.Problem = o
		}
	}
	return nil
}

// AddSubmitsG adds the given related objects to the existing relationships
// of the problem, optionally inserting them as new records.
// Appends related to o.R.Submits.
// Sets related.R.Problem appropriately.
// Uses the global database handle.
func (o *Problem) AddSubmitsG(ctx context.Context, insert bool, related ...*Submit) error {
	return o.AddSubmits(ctx, boil.GetContextDB(), insert, related...)
}

// AddSubmits adds the given related objects to the existing relationships
// of the problem, optionally inserting them as new records.
// Appends related to o.R.Submits.
// Sets related.R.Problem appropriately.
func (o *Problem) AddSubmits(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Submit) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProblemID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"submits\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"problem_id"}),
				strmangle.WhereClause("\"", "\"", 2, submitPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProblemID = o.ID
		}
	}

	if o.R == nil {
		o.R = &problemR{
			Submits: related,
		}
	} else {
		o.R.Submits = append(o.R.Submits, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &submitR{
				Problem: o,
			}
		} else {
			rel.R.Problem = o
		}
	}
	return nil
}

// Problems retrieves all the records using an executor.
func Problems(mods ...qm.QueryMod) problemQuery {
	mods = append(mods, qm.From("\"problems\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"problems\".*"})
	}

	return problemQuery{q}
}

// FindProblemG retrieves a single record by ID.
func FindProblemG(ctx context.Context, iD int, selectCols ...string) (*Problem, error) {
	return FindProblem(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProblem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProblem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Problem, error) {
	problemObj := &Problem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"problems\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, problemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from problems")
	}

	if err = problemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return problemObj, err
	}

	return problemObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Problem) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Problem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no problems provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(problemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	problemInsertCacheMut.RLock()
	cache, cached := problemInsertCache[key]
	problemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			problemAllColumns,
			problemColumnsWithDefault,
			problemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(problemType, problemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(problemType, problemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"problems\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"problems\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into problems")
	}

	if !cached {
		problemInsertCacheMut.Lock()
		problemInsertCache[key] = cache
		problemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Problem record using the global executor.
// See Update for more documentation.
func (o *Problem) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Problem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Problem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	problemUpdateCacheMut.RLock()
	cache, cached := problemUpdateCache[key]
	problemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			problemAllColumns,
			problemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update problems, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"problems\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, problemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(problemType, problemMapping, append(wl, problemPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update problems row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for problems")
	}

	if !cached {
		problemUpdateCacheMut.Lock()
		problemUpdateCache[key] = cache
		problemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q problemQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q problemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for problems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for problems")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProblemSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProblemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"problems\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, problemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in problem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all problem")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Problem) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Problem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no problems provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(problemColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	problemUpsertCacheMut.RLock()
	cache, cached := problemUpsertCache[key]
	problemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			problemAllColumns,
			problemColumnsWithDefault,
			problemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			problemAllColumns,
			problemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert problems, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(problemPrimaryKeyColumns))
			copy(conflict, problemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"problems\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(problemType, problemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(problemType, problemMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert problems")
	}

	if !cached {
		problemUpsertCacheMut.Lock()
		problemUpsertCache[key] = cache
		problemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Problem record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Problem) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Problem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Problem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Problem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), problemPrimaryKeyMapping)
	sql := "DELETE FROM \"problems\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from problems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for problems")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q problemQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q problemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no problemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from problems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for problems")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProblemSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProblemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(problemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"problems\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, problemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from problem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for problems")
	}

	if len(problemAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Problem) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Problem provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Problem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProblem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProblemSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ProblemSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProblemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProblemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"problems\".* FROM \"problems\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, problemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProblemSlice")
	}

	*o = slice

	return nil
}

// ProblemExistsG checks if the Problem row exists.
func ProblemExistsG(ctx context.Context, iD int) (bool, error) {
	return ProblemExists(ctx, boil.GetContextDB(), iD)
}

// ProblemExists checks if the Problem row exists.
func ProblemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"problems\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if problems exists")
	}

	return exists, nil
}
