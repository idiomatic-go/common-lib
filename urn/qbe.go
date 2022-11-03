package urn

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func ParseQbe(urn string) *QbeURN {
	if urn == "" {
		return &QbeURN{Nid: "", Nss: "", Grid: nil, Err: errors.New("invalid QbeURN, urn is empty")}
	}
	url, err := url.Parse(strings.TrimPrefix(urn, "urn:"))
	if err != nil {
		return &QbeURN{Nid: "", Nss: "", Grid: nil, Err: err}
	}
	if url.Scheme != QbeNid {
		return &QbeURN{Nid: "", Nss: "", Grid: nil, Err: errors.New(fmt.Sprintf("invalid QbeURN Nid : %v", url.Scheme))}
	}
	u := QbeURN{Nid: url.Scheme, Nss: url.Opaque}
	if u.Nid == "" || u.Nss == "" {
		u.Err = errors.New(fmt.Sprintf("invalid QbeURNn, Nid or Nss is empty : %v", u))
	} else {
		parseQbeGrid(&u)
	}
	if url.RawQuery != "" {
		u.RawQuery = url.RawQuery
	}
	return &u
}

func parseQbeGrid(urn *QbeURN) {
	cells := strings.Split(urn.Nss, ",")
	for _, exp := range cells {
		cell := &Cell{}
		urn.Err = cell.Parse(exp)
		if urn.Err != nil {
			return
		}
		urn.Grid = append(urn.Grid, *cell)
	}
}

func BuildQbe(embeddedContent bool, field string, criteria any) *QbeURN {
	return BuildQbeMulti(embeddedContent, Cell{Field: field, Criteria: criteria})
}

func BuildQbeMulti(embeddedContent bool, cells ...Cell) *QbeURN {
	u := QbeURN{Nid: QbeNid}
	if u.Nid == "" {
		u.Err = errors.New("invalid QbeURN, Nid is empty")
		return &u
	}
	if len(cells) == 0 {
		u.Err = errors.New("invalid QbeURN, cells are empty")
		return &u
	}
	if embeddedContent {
		u.RawQuery = EmbeddedContent
	}
	for i, cell := range cells {
		if cell.Field == "" {
			u.Err = errors.New("invalid QbeURN, cell field is empty")
			return &u
		}
		if i > 0 {
			u.Nss += ","
		}
		u.Nss += cell.String()
		u.Grid = append(u.Grid, cell)
	}
	if embeddedContent {
		u.Nss += fmt.Sprintf("?%v", EmbeddedContent)
	}
	return &u
}
