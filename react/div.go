package react

func Div(props map[string]any) *Element {
	return &Element{
		Type:  "div",
		Props: stringifyProps(props),
	}
}
