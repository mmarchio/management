package types

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowNodeOrder struct {
	ShallowModel
	WorkflowID WorkflowID `form:"workflow_id" json:"workflow_id"`
	NodeID string `form:"node_id" json:"node_id"`
	NodeType string `form:"node_type" json:"node_type"`
	Order int `form:"order" json:"order"`
}

func (c ShallowNodeOrder) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowNodeOrder) Expand(e echo.Context) (*NodeOrder, error) {
	r := NodeOrder{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err).Log()
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		return &r, nil
	}
	r.EmbedModel = r.EmbedModel.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.WorkflowID = c.WorkflowID
	r.NodeID = c.NodeID
	r.NodeType = c.NodeType
	r.Order = c.Order
	return &r, nil
}

func (c *ShallowNodeOrder) Unmarshal(e echo.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowNodeOrder) Marshal(e echo.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowNodeOrder) IsShallowModel() bool {
	return true
}
