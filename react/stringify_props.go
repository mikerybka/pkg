package react

import "encoding/json"

func stringifyProps(props map[string]any) map[string]string {
	res := map[string]string{}

	for k, v := range props {
		b, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		res[k] = string(b)
	}

	return res
}
