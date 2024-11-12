package react

import "encoding/json"

func CreateElement(typ Component, props map[string]any, children ...*Element) *Element {
	ch, err := json.Marshal(children)
	if err != nil {
		panic(err)
	}
	props["children"] = string(ch)
	return typ(props)
}
