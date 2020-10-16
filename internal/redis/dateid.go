package redis

import (
	"time"

	"github.com/ntec-io/Nostradamus/internal/logger"
)

const TimeLayout string = "02-01-2006"

func (c Client) SetDateIDs(m map[time.Time]string) (err error) {
	for k, v := range m {
		err = c.SetDateID(k, v)
		if err != nil {
			return
		}
	}
	return
}

func (c Client) SetDateID(t time.Time, id string) (err error) {
	logger.Log().Debugf("Saving DateID: %s = %s", t.Format(TimeLayout), id)
	err = c.rdb.Set(c.ctx, "nostradamus.dateid."+t.Format(TimeLayout), id, 0).Err()
	return
}

func (c Client) GetDateID(t time.Time) (string, error) {
	logger.Log().Debugf("Getting DateID for date: %s", t.Format(TimeLayout))
	return c.rdb.Get(c.ctx, "nostradamus.dateid."+t.Format(TimeLayout)).Result()
}

func (c Client) GetLastDateID(t time.Time) (res string, err error) {
	if res, err = c.GetDateID(t); err != nil {
		return c.GetLastDateID(t.AddDate(0, 0, -1))
	}
	return
}
