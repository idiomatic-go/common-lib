package urn

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

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
		cell := &Cell{}
		urn.Err = cell.Parse(exp)
		if urn.Err != nil {
			return
		}
		urn.QbeGrid = append(urn.QbeGrid, *cell)
	}
}

func Build(nid string, field string, criteria any) *URN {
	return BuildMulti(nid, Cell{Field: field, Criteria: criteria})
}

func BuildMulti(nid string, cells ...Cell) *URN {
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
