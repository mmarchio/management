package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)


type Stats struct {
	EmbedModel
	ID 				StatsID `json:"stats_id"`
	Start 			time.Time `json:"start"`
	End 			time.Time `json:"end"`
	Input 			string `json:"input"`
	Output 			string `json:"output"`
	Duration 		time.Duration `json:"duration"`
	FilesArrayModel []File `json:"files_array_model"`
	Status 			string `json:"status"`
}

func (c Stats) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowStats{}
	sm.ShallowModel = sm.ShallowModel.FromEmbedModel(c.EmbedModel)
	sm.ID = c.ID
	sm.Start = c.Start
	sm.End = c.End
	sm.Input = c.Input
	sm.Output = c.Output
	sm.Duration = c.Duration
	sm.Status = c.Status
	sm.FilesArrayModel = make([]string, 0)
	for _, id := range c.FilesArrayModel {
		sm.FilesArrayModel = append(sm.FilesArrayModel, id.EmbedModel.ID)
		sms = append(sms, id.Pack()...)
	}
	sms = append(sms, sm)
	return sms
}

func (c *Stats) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Stats) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c Stats) New(id *string) Stats {
	if id != nil {
		c.ID = StatsID(*id)
	} else {
		c.ID = StatsID(uuid.NewString())
	}
	c.Status = "queued"
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	c.ContentType = "contextitem"
	return c
}

func (c Stats) Get(e echo.Context) (*Stats, error) {
	input := Content{}
	input.Model.ContentType = "stats"
	input.Model.ID = c.EmbedModel.ID
	input.ID = c.ID.String()
	output, err := input.Get(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	stats := c
	if err := json.Unmarshal([]byte(output.Content), &stats); err != nil {
		return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
	}
	return &stats, nil
}