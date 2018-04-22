package v1alpha1

type PlanID string

type Plan struct {
	ID       PlanID `json:"id"`
	Attempts []int  `json:"attempts,omitempty"`

	Aggregate *AggregatePlan `json:"aggregate,omitempty"`
}

type AggregatePlan []Plan
