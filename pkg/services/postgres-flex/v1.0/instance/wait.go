package instance

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// WaitForCreateOrUpdate will wait for instance create/update to complete
// returned interface is of *InstanceSingleInstance
func (c *ClientWithResponses[K]) WaitForCreateOrUpdate(ctx context.Context, projectID, instanceID string) *wait.Handler {
	// artifical wait for instance to change from status ready to updating
	time.Sleep(5 * time.Second)
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.Error != nil {
			return nil, false, err
		}
		if s.JSON200 == nil || s.JSON200.Item == nil {
			return nil, false, errors.New("bad response")
		}
		if *s.JSON200.Item.Status == STATUS_READY {
			return s.JSON200.Item, true, nil
		}
		if *s.JSON200.Item.Status == STATUS_FAILED {
			return s.JSON200.Item, false, errors.New("received status FAILED from server")
		}
		return s.JSON200.Item, false, nil
	})
}

// WaitForDelete will wait for instance deletion
// returned value for deletion wait will always be nil
func (c *ClientWithResponses[K]) WaitForDelete(ctx context.Context, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if res.StatusCode() == http.StatusNotFound {
			return nil, true, nil
		}
		if res.Error != nil {
			return nil, false, res.Error
		}
		return nil, false, nil
	})
}