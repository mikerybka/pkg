package react

import (
	"bytes"
	"encoding/json"
	"io"
)

type Element struct {
	Type  string
	Props map[string]string
}

func (el *Element) Render(w io.Writer, indent int) (int, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("<")
	buf.WriteString(el.Type)
	for k, v := range el.Props {
		if k == "children" {
			continue
		}
		buf.WriteString(" ")
		buf.WriteString(k)
		buf.WriteString("='")
		buf.WriteString(v)
		buf.WriteString("'")
	}
	buf.WriteString(">")
	children := []Element{}
	json.Unmarshal([]byte(el.Props["children"]), &children)
	for _, ch := range children {
		buf.WriteString("\n")
		for i := 0; i < indent+1; i++ {
			buf.WriteString("\t")
		}
		ch.Render(buf, indent+1)
	}
	buf.WriteString("</")
	buf.WriteString(el.Type)
	buf.WriteString(">")
	return w.Write(buf.Bytes())
}
