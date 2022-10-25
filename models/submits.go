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

// Submit is an object representing the database table.
type Submit struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID      int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ProblemID   int       `boil:"problem_id" json:"problem_id" toml:"problem_id" yaml:"problem_id"`
	Status      int       `boil:"status" json:"status" toml:"status" yaml:"status"`
	SolutionRaw string    `boil:"solution_raw" json:"solution_raw" toml:"solution_raw" yaml:"solution_raw"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *submitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L submitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SubmitColumns = struct {
	ID          string
	UserID      string
	ProblemID   string
	Status      string
	SolutionRaw string
	CreatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	ProblemID:   "problem_id",
	Status:      "status",
	SolutionRaw: "solution_raw",
	CreatedAt:   "created_at",
}

var SubmitTableColumns = struct {
	ID          string
	UserID      string
	ProblemID   string
	Status      string
	SolutionRaw string
	CreatedAt   string
}{
	ID:          "submits.id",
	UserID:      "submits.user_id",
	ProblemID:   "submits.problem_id",
	Status:      "submits.status",
	SolutionRaw: "submits.solution_raw",
	CreatedAt:   "submits.created_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var SubmitWhere = struct {
	ID          whereHelperint
	UserID      whereHelperint
	ProblemID   whereHelperint
	Status      whereHelperint
	SolutionRaw whereHelperstring
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint{field: "\"submits\".\"id\""},
	UserID:      whereHelperint{field: "\"submits\".\"user_id\""},
	ProblemID:   whereHelperint{field: "\"submits\".\"problem_id\""},
	Status:      whereHelperint{field: "\"submits\".\"status\""},
	SolutionRaw: whereHelperstring{field: "\"submits\".\"solution_raw\""},
	CreatedAt:   whereHelpertime_Time{field: "\"submits\".\"created_at\""},
}

// SubmitRels is where relationship names are stored.
var SubmitRels = struct {
	Problem string
	User    string
}{
	Problem: "Problem",
	User:    "User",
}

// submitR is where relationships are stored.
type submitR struct {
	Problem *Problem `boil:"Problem" json:"Problem" toml:"Problem" yaml:"Problem"`
	User    *User    `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*submitR) NewStruct() *submitR {
	return &submitR{}
}

func (r *submitR) GetProblem() *Problem {
	if r == nil {
		return nil
	}
	return r.Problem
}

func (r *submitR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// submitL is where Load methods for each relationship are stored.
type submitL struct{}

var (
	submitAllColumns            = []string{"id", "user_id", "problem_id", "status", "solution_raw", "created_at"}
	submitColumnsWithoutDefault = []string{"user_id", "problem_id", "status", "solution_raw"}
	submitColumnsWithDefault    = []string{"id", "created_at"}
	submitPrimaryKeyColumns     = []string{"id"}
	submitGeneratedColumns      = []string{}
)

type (
	// SubmitSlice is an alias for a slice of pointers to Submit.
	// This should almost always be used instead of []Submit.
	SubmitSlice []*Submit
	// SubmitHook is the signature for custom Submit hook methods
	SubmitHook func(context.Context, boil.ContextExecutor, *Submit) error

	submitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	submitType                 = reflect.TypeOf(&Submit{})
	submitMapping              = queries.MakeStructMapping(submitType)
	submitPrimaryKeyMapping, _ = queries.BindMapping(submitType, submitMapping, submitPrimaryKeyColumns)
	submitInsertCacheMut       sync.RWMutex
	submitInsertCache          = make(map[string]insertCache)
	submitUpdateCacheMut       sync.RWMutex
	submitUpdateCache          = make(map[string]updateCache)
	submitUpsertCacheMut       sync.RWMutex
	submitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var submitAfterSelectHooks []SubmitHook

var submitBeforeInsertHooks []SubmitHook
var submitAfterInsertHooks []SubmitHook

var submitBeforeUpdateHooks []SubmitHook
var submitAfterUpdateHooks []SubmitHook

var submitBeforeDeleteHooks []SubmitHook
var submitAfterDeleteHooks []SubmitHook

var submitBeforeUpsertHooks []SubmitHook
var submitAfterUpsertHooks []SubmitHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Submit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Submit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Submit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Submit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Submit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Submit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Submit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Submit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Submit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range submitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSubmitHook registers your hook function for all future operations.
func AddSubmitHook(hookPoint boil.HookPoint, submitHook SubmitHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		submitAfterSelectHooks = append(submitAfterSelectHooks, submitHook)
	case boil.BeforeInsertHook:
		submitBeforeInsertHooks = append(submitBeforeInsertHooks, submitHook)
	case boil.AfterInsertHook:
		submitAfterInsertHooks = append(submitAfterInsertHooks, submitHook)
	case boil.BeforeUpdateHook:
		submitBeforeUpdateHooks = append(submitBeforeUpdateHooks, submitHook)
	case boil.AfterUpdateHook:
		submitAfterUpdateHooks = append(submitAfterUpdateHooks, submitHook)
	case boil.BeforeDeleteHook:
		submitBeforeDeleteHooks = append(submitBeforeDeleteHooks, submitHook)
	case boil.AfterDeleteHook:
		submitAfterDeleteHooks = append(submitAfterDeleteHooks, submitHook)
	case boil.BeforeUpsertHook:
		submitBeforeUpsertHooks = append(submitBeforeUpsertHooks, submitHook)
	case boil.AfterUpsertHook:
		submitAfterUpsertHooks = append(submitAfterUpsertHooks, submitHook)
	}
}

// OneG returns a single submit record from the query using the global executor.
func (q submitQuery) OneG(ctx context.Context) (*Submit, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single submit record from the query.
func (q submitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Submit, error) {
	o := &Submit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for submits")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Submit records from the query using the global executor.
func (q submitQuery) AllG(ctx context.Context) (SubmitSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Submit records from the query.
func (q submitQuery) All(ctx context.Context, exec boil.ContextExecutor) (SubmitSlice, error) {
	var o []*Submit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Submit slice")
	}

	if len(submitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Submit records in the query using the global executor
func (q submitQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Submit records in the query.
func (q submitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count submits rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q submitQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q submitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if submits exists")
	}

	return count > 0, nil
}

// Problem pointed to by the foreign key.
func (o *Submit) Problem(mods ...qm.QueryMod) problemQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProblemID),
	}

	queryMods = append(queryMods, mods...)

	return Problems(queryMods...)
}

// User pointed to by the foreign key.
func (o *Submit) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadProblem allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (submitL) LoadProblem(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubmit interface{}, mods queries.Applicator) error {
	var slice []*Submit
	var object *Submit

	if singular {
		var ok bool
		object, ok = maybeSubmit.(*Submit)
		if !ok {
			object = new(Submit)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSubmit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSubmit))
			}
		}
	} else {
		s, ok := maybeSubmit.(*[]*Submit)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSubmit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSubmit))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &submitR{}
		}
		args = append(args, object.ProblemID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &submitR{}
			}

			for _, a := range args {
				if a == obj.ProblemID {
					continue Outer
				}
			}

			args = append(args, obj.ProblemID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`problems`),
		qm.WhereIn(`problems.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Problem")
	}

	var resultSlice []*Problem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Problem")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for problems")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for problems")
	}

	if len(submitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Problem = foreign
		if foreign.R == nil {
			foreign.R = &problemR{}
		}
		foreign.R.Submits = append(foreign.R.Submits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProblemID == foreign.ID {
				local.R.Problem = foreign
				if foreign.R == nil {
					foreign.R = &problemR{}
				}
				foreign.R.Submits = append(foreign.R.Submits, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (submitL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubmit interface{}, mods queries.Applicator) error {
	var slice []*Submit
	var object *Submit

	if singular {
		var ok bool
		object, ok = maybeSubmit.(*Submit)
		if !ok {
			object = new(Submit)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSubmit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSubmit))
			}
		}
	} else {
		s, ok := maybeSubmit.(*[]*Submit)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSubmit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSubmit))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &submitR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &submitR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(submitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Submits = append(foreign.R.Submits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Submits = append(foreign.R.Submits, local)
				break
			}
		}
	}

	return nil
}

// SetProblemG of the submit to the related item.
// Sets o.R.Problem to related.
// Adds o to related.R.Submits.
// Uses the global database handle.
func (o *Submit) SetProblemG(ctx context.Context, insert bool, related *Problem) error {
	return o.SetProblem(ctx, boil.GetContextDB(), insert, related)
}

// SetProblem of the submit to the related item.
// Sets o.R.Problem to related.
// Adds o to related.R.Submits.
func (o *Submit) SetProblem(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Problem) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"submits\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"problem_id"}),
		strmangle.WhereClause("\"", "\"", 2, submitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProblemID = related.ID
	if o.R == nil {
		o.R = &submitR{
			Problem: related,
		}
	} else {
		o.R.Problem = related
	}

	if related.R == nil {
		related.R = &problemR{
			Submits: SubmitSlice{o},
		}
	} else {
		related.R.Submits = append(related.R.Submits, o)
	}

	return nil
}

// SetUserG of the submit to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Submits.
// Uses the global database handle.
func (o *Submit) SetUserG(ctx context.Context, insert bool, related *User) error {
	return o.SetUser(ctx, boil.GetContextDB(), insert, related)
}

// SetUser of the submit to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Submits.
func (o *Submit) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"submits\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, submitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &submitR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Submits: SubmitSlice{o},
		}
	} else {
		related.R.Submits = append(related.R.Submits, o)
	}

	return nil
}

// Submits retrieves all the records using an executor.
func Submits(mods ...qm.QueryMod) submitQuery {
	mods = append(mods, qm.From("\"submits\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"submits\".*"})
	}

	return submitQuery{q}
}

// FindSubmitG retrieves a single record by ID.
func FindSubmitG(ctx context.Context, iD int, selectCols ...string) (*Submit, error) {
	return FindSubmit(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindSubmit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSubmit(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Submit, error) {
	submitObj := &Submit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"submits\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, submitObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from submits")
	}

	if err = submitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return submitObj, err
	}

	return submitObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Submit) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Submit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no submits provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(submitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	submitInsertCacheMut.RLock()
	cache, cached := submitInsertCache[key]
	submitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			submitAllColumns,
			submitColumnsWithDefault,
			submitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(submitType, submitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(submitType, submitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"submits\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"submits\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into submits")
	}

	if !cached {
		submitInsertCacheMut.Lock()
		submitInsertCache[key] = cache
		submitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Submit record using the global executor.
// See Update for more documentation.
func (o *Submit) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Submit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Submit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	submitUpdateCacheMut.RLock()
	cache, cached := submitUpdateCache[key]
	submitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			submitAllColumns,
			submitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update submits, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"submits\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, submitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(submitType, submitMapping, append(wl, submitPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update submits row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for submits")
	}

	if !cached {
		submitUpdateCacheMut.Lock()
		submitUpdateCache[key] = cache
		submitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q submitQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q submitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for submits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for submits")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SubmitSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SubmitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), submitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"submits\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, submitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in submit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all submit")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Submit) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Submit) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no submits provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(submitColumnsWithDefault, o)

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

	submitUpsertCacheMut.RLock()
	cache, cached := submitUpsertCache[key]
	submitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			submitAllColumns,
			submitColumnsWithDefault,
			submitColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			submitAllColumns,
			submitPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert submits, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(submitPrimaryKeyColumns))
			copy(conflict, submitPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"submits\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(submitType, submitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(submitType, submitMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert submits")
	}

	if !cached {
		submitUpsertCacheMut.Lock()
		submitUpsertCache[key] = cache
		submitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Submit record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Submit) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Submit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Submit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Submit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), submitPrimaryKeyMapping)
	sql := "DELETE FROM \"submits\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from submits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for submits")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q submitQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q submitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no submitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from submits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for submits")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SubmitSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SubmitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(submitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), submitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"submits\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, submitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from submit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for submits")
	}

	if len(submitAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Submit) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Submit provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Submit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSubmit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubmitSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty SubmitSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubmitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SubmitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), submitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"submits\".* FROM \"submits\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, submitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SubmitSlice")
	}

	*o = slice

	return nil
}

// SubmitExistsG checks if the Submit row exists.
func SubmitExistsG(ctx context.Context, iD int) (bool, error) {
	return SubmitExists(ctx, boil.GetContextDB(), iD)
}

// SubmitExists checks if the Submit row exists.
func SubmitExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"submits\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if submits exists")
	}

	return exists, nil
}
