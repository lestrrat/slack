package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go (generateServiceDetailsFile). DO NOT EDIT!

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat-go/slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = json.Marshal
var _ = objects.EpochTime(0)

// DialogOpenCall is created by DialogService.Open method call
type DialogOpenCall struct {
	service    *DialogService
	dialog     *objects.Dialog
	trigger_id string
}

// Open creates a DialogOpenCall object in preparation for accessing the dialog.open endpoint
func (s *DialogService) Open(dialog *objects.Dialog, trigger_id string) *DialogOpenCall {
	var call DialogOpenCall
	call.service = s
	call.dialog = dialog
	call.trigger_id = trigger_id
	return &call
}

// ValidateArgs checks that all required fields are set in the DialogOpenCall object
func (c *DialogOpenCall) ValidateArgs() error {
	if c.dialog == nil {
		return errors.New(`required field dialog not initialized`)
	}
	if len(c.trigger_id) <= 0 {
		return errors.New(`required field trigger_id not initialized`)
	}
	return nil
}

// Values returns the DialogOpenCall object as url.Values
func (c *DialogOpenCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	dialogEncoded, err := json.Marshal(c.dialog)
	if err != nil {
		return nil, errors.Wrap(err, `failed to encode field`)
	}
	v.Set("dialog", string(dialogEncoded))

	v.Set("trigger_id", c.trigger_id)
	return v, nil
}

// Do executes the call to access dialog.open endpoint
func (c *DialogOpenCall) Do(ctx context.Context) (*objects.DialogResponse, error) {
	const endpoint = "dialog.open"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.DialogResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to dialog.open`)
	}
	if !res.OK() {
		return nil, errors.New(res.Error().String())
	}

	return res.DialogResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *DialogOpenCall) FromValues(v url.Values) error {
	var tmp DialogOpenCall
	if raw := strings.TrimSpace(v.Get("dialog")); len(raw) > 0 {
		if err := json.Unmarshal([]byte(raw), &tmp.dialog); err != nil {
			return errors.Wrap(err, `failed to decode value "dialog"`)
		}
	}
	if raw := strings.TrimSpace(v.Get("trigger_id")); len(raw) > 0 {
		tmp.trigger_id = raw
	}
	*c = tmp
	return nil
}
