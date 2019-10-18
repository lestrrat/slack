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

// OAuthAccessCall is created by OAuthService.Access method call
type OAuthAccessCall struct {
	service      *OAuthService
	clientID     string
	clientSecret string
	code         string
	redirectURI  string
}

// Access creates a OAuthAccessCall object in preparation for accessing the oauth.access endpoint
func (s *OAuthService) Access(clientID string, clientSecret string, code string) *OAuthAccessCall {
	var call OAuthAccessCall
	call.service = s
	call.clientID = clientID
	call.clientSecret = clientSecret
	call.code = code
	return &call
}

// RedirectURI sets the value for optional redirectURI parameter
func (c *OAuthAccessCall) RedirectURI(redirectURI string) *OAuthAccessCall {
	c.redirectURI = redirectURI
	return c
}

// ValidateArgs checks that all required fields are set in the OAuthAccessCall object
func (c *OAuthAccessCall) ValidateArgs() error {
	if len(c.clientID) <= 0 {
		return errors.New(`required field clientID not initialized`)
	}
	if len(c.clientSecret) <= 0 {
		return errors.New(`required field clientSecret not initialized`)
	}
	if len(c.code) <= 0 {
		return errors.New(`required field code not initialized`)
	}
	return nil
}

// Values returns the OAuthAccessCall object as url.Values
func (c *OAuthAccessCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}

	v.Set("client_id", c.clientID)

	v.Set("client_secret", c.clientSecret)

	v.Set("code", c.code)

	if len(c.redirectURI) > 0 {
		v.Set("redirect_uri", c.redirectURI)
	}
	return v, nil
}

type OAuthAccessCallResponse struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}

func (r *OAuthAccessCallResponse) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal OAuthAccessCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *OAuthAccessCallResponse) payload() (*objects.OAuthAccessResponse, error) {
	var res0 objects.OAuthAccessResponse
	if err := json.Unmarshal(r.Payload0, &res0); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.OAuthAccessResponse from response`)
	}
	return &res0, nil
}

// Do executes the call to access oauth.access endpoint
func (c *OAuthAccessCall) Do(ctx context.Context) (*objects.OAuthAccessResponse, error) {
	const endpoint = "oauth.access"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res OAuthAccessCallResponse
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to oauth.access`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to oauth.access`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *OAuthAccessCall) FromValues(v url.Values) error {
	var tmp OAuthAccessCall
	if raw := strings.TrimSpace(v.Get("client_id")); len(raw) > 0 {
		tmp.clientID = raw
	}
	if raw := strings.TrimSpace(v.Get("client_secret")); len(raw) > 0 {
		tmp.clientSecret = raw
	}
	if raw := strings.TrimSpace(v.Get("code")); len(raw) > 0 {
		tmp.code = raw
	}
	if raw := strings.TrimSpace(v.Get("redirect_uri")); len(raw) > 0 {
		tmp.redirectURI = raw
	}
	*c = tmp
	return nil
}
