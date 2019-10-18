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

// RTMStartCall is created by RTMService.Start method call
type RTMStartCall struct {
	service *RTMService
}

// Start creates a RTMStartCall object in preparation for accessing the rtm.start endpoint
func (s *RTMService) Start() *RTMStartCall {
	var call RTMStartCall
	call.service = s
	return &call
}

// ValidateArgs checks that all required fields are set in the RTMStartCall object
func (c *RTMStartCall) ValidateArgs() error {
	return nil
}

// Values returns the RTMStartCall object as url.Values
func (c *RTMStartCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

type RTMStartCallResponse struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}

func (r *RTMStartCallResponse) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal RTMStartCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *RTMStartCallResponse) payload() (*objects.RTMResponse, error) {
	var res0 objects.RTMResponse
	if err := json.Unmarshal(r.Payload0, &res0); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.RTMResponse from response`)
	}
	return &res0, nil
}

// Do executes the call to access rtm.start endpoint
func (c *RTMStartCall) Do(ctx context.Context) (*objects.RTMResponse, error) {
	const endpoint = "rtm.start"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res RTMStartCallResponse
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to rtm.start`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to rtm.start`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *RTMStartCall) FromValues(v url.Values) error {
	var tmp RTMStartCall
	*c = tmp
	return nil
}
