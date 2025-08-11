package types

import (
	"context"
	"encoding/json"
)

type NodeOrder struct {
	EmbedModel
	ID NodeOrderID 			`form:"id" json:"id"`
	WorkflowID WorkflowID 	`form:"workflow_id" json:"workflow_id"`
	NodeID string 			`form:"node_id" json:"node_id"`
	NodeType string 		`form:"node_type" json:"node_type"`
	Order int 				`form:"order" json:"order"`
}

func (c NodeOrder) IsNil() bool {
	if c.ID.IsNil() && c.ID.IsNil() && c.WorkflowID.IsNil() && c.NodeID == "" && c.NodeType == ""{
		return true
	}
	return false
}

func (c NodeOrder) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowNodeOrder{}
	sm.ShallowModel = sm.ShallowModel.FromEmbedModel(c.EmbedModel)
	sm.ID = c.ID
	return sms
}

func (c *NodeOrder) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c NodeOrder) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
