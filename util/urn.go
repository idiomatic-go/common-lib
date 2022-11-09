package util

import (
	"errors"
	"fmt"
	"strings"
)

func (u URN) String() string {
	var sb strings.Builder
	sb.WriteString(u.Nid)
	sb.WriteString(":")
	for i, cell := range u.Grid {
		if i > 0 {
			u.Nss += ","
		}
		u.Nss += cell.String()
	}
	sb.WriteString(u.Nss)
	return sb.String()
}

func (c Cell) String() string {
	return fmt.Sprintf("%v=%v", c.Field, c.Criteria)
}

/*

func (c *Cell) Parse(exp string) error {
	tokens := strings.Split(exp, "=")
	if len(tokens) < 2 {
		return errors.New(fmt.Sprintf("invalid QBE expression, missing token : %v", exp))
	}
	c.Field = tokens[0]
	c.Criteria = tokens[1]
	return nil
}



func Parse(urn string) *URN {
	if urn == "" {
		return &URN{Nid: "", Nss: "", Grid: nil, Err: errors.New("invalid URN, urn is empty")}
	}
	url0, err := url.Parse(strings.TrimPrefix(urn, "urn:"))
	if err != nil {
		return &URN{Nid: "", Nss: "", Grid: nil, Err: err}
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
		urn.Grid = append(urn.Grid, *cell)
	}
}


*/

func NewURN(nid string, cells ...Cell) URN {
	u := URN{Nid: QbeNid}
	if nid != "" {
		u.Nid = nid
	}
	/*
		if u.Nid == "" {
			u.Err = errors.New("invalid URN, Nid is empty")
			return u
		}
		if len(cells) == 0 {
			u.Err = errors.New("invalid URN, cells are empty")
			return u
		}

	*/
	for _, cell := range cells {
		//if cell.Field == "" {
		//	u.Err = errors.New("invalid URN, cell field is empty")
		//	return u
		//}
		//if i > 0 {
		//	u.Nss += ","
		//	}
		//	u.Nss += cell.String()
		u.Grid = append(u.Grid, cell)
	}
	return u
}

func ValidateURN(urn URN) error {
	if urn.Nid == "" {
		return errors.New("invalid URN, Nid is empty")
	}
	if len(urn.Grid) == 0 {
		return errors.New("invalid URN, Cells are empty")
	}
	for _, c := range urn.Grid {
		if c.Field == "" {
			return errors.New("invalid URN, cell field is empty")
		}
	}
	return nil
}
