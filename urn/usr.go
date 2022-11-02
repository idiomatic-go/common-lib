package urn

import (
	"errors"
	"fmt"
	"strings"
)

const (
	QbeNid     = "qbe"
	SchemeName = "scheme"
)

type Cell struct {
	Field    string
	Criteria any
}

func (c *Cell) Parse(exp string) error {
	tokens := strings.Split(exp, "=")
	if len(tokens) < 2 {
		return errors.New(fmt.Sprintf("invalid QBE expression, missing token : %v", exp))
	}
	c.Field = tokens[0]
	c.Criteria = tokens[1]
	return nil
}

func (c *Cell) String() string {
	return fmt.Sprintf("%v=%v", c.Field, c.Criteria)
}

type QbeURN struct {
	Nid  string
	Nss  string
	Grid []Cell
	Err  error
}

func (u *QbeURN) String() string {
	var sb strings.Builder
	sb.WriteString(u.Nid)
	sb.WriteString(":")
	sb.WriteString(u.Nss)
	return sb.String()
}

func (u *QbeURN) Cell(field string) Cell {
	for i, cell := range u.Grid {
		if cell.Field == field {
			return u.Grid[i]
		}
	}
	return Cell{}
}

func (u *QbeURN) Scheme() string {
	c := u.Cell(SchemeName)
	if c.Criteria == nil {
		return ""
	}
	if s, ok := c.Criteria.(string); ok {
		return s
	}
	return ""
}
