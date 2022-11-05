package util

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func (c *QbeCell) Parse(exp string) error {
	tokens := strings.Split(exp, "=")
	if len(tokens) < 2 {
		return errors.New(fmt.Sprintf("invalid QBE expression, missing token : %v", exp))
	}
	c.Field = tokens[0]
	c.Criteria = tokens[1]
	return nil
}

func (c *QbeCell) String() string {
	return fmt.Sprintf("%v=%v", c.Field, c.Criteria)
}

func Parse(urn string) *URN {
	if urn == "" {
		return &URN{Nid: "", Nss: "", QbeGrid: nil, Err: errors.New("invalid URN, urn is empty")}
	}
	url0, err := url.Parse(strings.TrimPrefix(urn, "urn:"))
	if err != nil {
		return &URN{Nid: "", Nss: "", QbeGrid: nil, Err: err}
	}
	u := URN{Nid: url0.Scheme, Nss: url0.Opaque, RawQuery: url0.RawQuery}
	if u.Nid == "" || u.Nss == "" {
		u.Err = errors.New(fmt.Sprintf("invalid URN, Nid or Nss is empty : %v", u))
	} else {
		parseQbeGrid(&u)
	}
	return &u
}

func parseQbeGrid(urn *URN) {
	cells := strings.Split(urn.Nss, ",")
	for _, exp := range cells {
		cell := &QbeCell{}
		urn.Err = cell.Parse(exp)
		if urn.Err != nil {
			return
		}
		urn.QbeGrid = append(urn.QbeGrid, *cell)
	}
}

func Build(nid string, field string, criteria any) *URN {
	return BuildMulti(nid, QbeCell{Field: field, Criteria: criteria})
}

func BuildMulti(nid string, cells ...QbeCell) *URN {
	u := URN{Nid: QbeNid}
	if nid != "" {
		u.Nid = nid
	}
	if u.Nid == "" {
		u.Err = errors.New("invalid URN, Nid is empty")
		return &u
	}
	if len(cells) == 0 {
		u.Err = errors.New("invalid URN, cells are empty")
		return &u
	}
	for i, cell := range cells {
		if cell.Field == "" {
			u.Err = errors.New("invalid URN, cell field is empty")
			return &u
		}
		if i > 0 {
			u.Nss += ","
		}
		u.Nss += cell.String()
		u.QbeGrid = append(u.QbeGrid, cell)
	}
	return &u
}
