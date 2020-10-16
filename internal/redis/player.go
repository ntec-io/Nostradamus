package redis

import (
	"encoding/json"
	"fmt"

	"github.com/ntec-io/Nostradamus/pkg/fifaindex"
)

type Player struct {
	ID            int
	FifaindexLink string
	Name          string
	Stats         fifaindex.Player
	DateID        string
}

func (c Client) SavePlayer(p Player) (err error) {
	keyBase := "nostradamus.player." + fmt.Sprint(p.ID) + "."

	err = c.rdb.Set(c.ctx, keyBase+"name", p.Name, 0).Err()
	if err != nil {
		return
	}

	err = c.rdb.Set(c.ctx, keyBase+"fifaindexlink", p.FifaindexLink, 0).Err()
	if err != nil {
		return
	}

	d, err := json.Marshal(p.Stats)
	if err != nil {
		return
	}
	err = c.rdb.Set(c.ctx, keyBase+"stats."+p.DateID, string(d), 0).Err()

	return
}
