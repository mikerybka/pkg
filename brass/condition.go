package brass

type Condition struct {
	LHS *Expression `json:"lhs"`
	Op  string      `json:"op"`
	RHS *Expression `json:"rhs"`
}
