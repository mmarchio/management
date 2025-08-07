package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)


type ShallowStats struct {
	ShallowModel
	ID 				StatsID `json:"stats_id"`
	Start 			time.Time `json:"start"`
	End 			time.Time `json:"end"`
	Input 			string `json:"input"`
	Output 			string `json:"output"`
	Duration 		time.Duration `json:"duration"`
	FilesArrayModel []string `json:"files_array_model"`
	Status 			string `json:"status"`
}

func (c *ShallowStats) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowStats) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowStats) New(id *string) ShallowStats {
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

