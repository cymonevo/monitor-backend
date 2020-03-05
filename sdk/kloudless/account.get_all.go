package kloudless

import (
	"context"
	"errors"

	"github.com/cymonevo/monitor-backend/handler"
	"github.com/cymonevo/monitor-backend/internal/log"
	"github.com/cymonevo/monitor-backend/sdk"
)

const getAllAccounts = "/accounts"

func (c *clientImpl) GetAccounts(ctx context.Context) (interface{}, error) {
	resp, err := c.client.Get(getAllAccounts, nil, c.headers(nil))
	if err != nil {
		log.ErrorDetail("GetAccounts", "error get accounts", err)
		return nil, err
	}
	if !sdk.IsSuccess(resp.StatusCode) {
		log.Warnf("GetAccounts", "status %d %s", resp.StatusCode, resp.Status)
		return nil, errors.New(resp.Status)
	}
	var data interface{}
	err = handler.ParseBody(resp.Body, &data)
	if err != nil {
		log.ErrorDetail("GetAccounts", "error unmarshal data", err)
		return nil, err
	}
	log.Infof("SDK", "success get accounts %+v", data)
	return data, nil
}
