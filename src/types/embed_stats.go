package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
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

func (c *Stats) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Stats) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c Stats) New(id string) Stats {
	if id != "" {
		c.ID = StatsID(id)
	} else {
		c.ID = StatsID(uuid.NewString())
	}
	c.Status = "queued"
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	c.ContentType = "contextitem"
	return c
}

