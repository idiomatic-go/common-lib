package urn

import (
	"errors"
	"fmt"
	"strings"
)

const (
	QbeNid = "qbe"
	//ContentLocation = "content-location"
	//Embedded        = "embedded"
	//EmbeddedContent = ContentLocation + "=" + Embedded
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

type URN struct {
	Nid      string
	Nss      string
	RawQuery string
	QbeGrid  []Cell
	Err      error
}

func (u *URN) String() string {
	var sb strings.Builder
	sb.WriteString(u.Nid)
	sb.WriteString(":")
	sb.WriteString(u.Nss)
	return sb.String()
}

func (u *URN) Cell(field string) Cell {
	for i, cell := range u.QbeGrid {
		if cell.Field == field {
			return u.QbeGrid[i]
		}
	}
	return Cell{}
}

/*
func (u *URN) IsEmbeddedContent() bool {
	if u.RawQuery == "" {
		return false
	}
	//if list, ok := u.Values[ContentLocation]; ok {
	return strings.Index(u.RawQuery, EmbeddedContent) != -1
}

*/
