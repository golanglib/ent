// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/note"
	"entgo.io/ent/entc/integration/customid/ent/schema"
)

// Note is the model entity for the Note schema.
type Note struct {
	config `json:"-"`
	// ID of the ent.
	ID schema.NoteID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NoteQuery when eager-loading is set.
	Edges         NoteEdges `json:"edges"`
	note_children *schema.NoteID
}

// NoteEdges holds the relations/edges for other nodes in the graph.
type NoteEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Note `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Note `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NoteEdges) ParentOrErr() (*Note, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: note.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e NoteEdges) ChildrenOrErr() ([]*Note, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Note) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case note.FieldID, note.FieldText:
			values[i] = new(sql.NullString)
		case note.ForeignKeys[0]: // note_children
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Note", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Note fields.
func (n *Note) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case note.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				n.ID = schema.NoteID(value.String)
			}
		case note.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				n.Text = value.String
			}
		case note.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note_children", values[i])
			} else if value.Valid {
				n.note_children = new(schema.NoteID)
				*n.note_children = schema.NoteID(value.String)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the Note entity.
func (n *Note) QueryParent() *NoteQuery {
	return (&NoteClient{config: n.config}).QueryParent(n)
}

// QueryChildren queries the "children" edge of the Note entity.
func (n *Note) QueryChildren() *NoteQuery {
	return (&NoteClient{config: n.config}).QueryChildren(n)
}

// Update returns a builder for updating this Note.
// Note that you need to call Note.Unwrap() before calling this method if this Note
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Note) Update() *NoteUpdateOne {
	return (&NoteClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Note entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Note) Unwrap() *Note {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Note is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Note) String() string {
	var builder strings.Builder
	builder.WriteString("Note(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", text=")
	builder.WriteString(n.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Notes is a parsable slice of Note.
type Notes []*Note

func (n Notes) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}