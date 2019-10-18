package objects

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func BuildAuthTestResponse() *AuthTestResponseBuilder {
	var b AuthTestResponseBuilder
	return &b
}

func (b *AuthTestResponseBuilder) URL(v string) *AuthTestResponseBuilder {
	b.url = v
	return b
}

func (b *AuthTestResponseBuilder) Team(v string) *AuthTestResponseBuilder {
	b.team = v
	return b
}

func (b *AuthTestResponseBuilder) User(v string) *AuthTestResponseBuilder {
	b.user = v
	return b
}

func (b *AuthTestResponseBuilder) TeamID(v string) *AuthTestResponseBuilder {
	b.teamId = v
	return b
}

func (b *AuthTestResponseBuilder) UserID(v string) *AuthTestResponseBuilder {
	b.userId = v
	return b
}

func (b *AuthTestResponseBuilder) Build() (*AuthTestResponse, error) {
	var v AuthTestResponse
	v.url = b.url
	v.team = b.team
	v.user = b.user
	v.teamId = b.teamId
	v.userId = b.userId
	return &v, nil
}

func (b *AuthTestResponseBuilder) MustBuild() *AuthTestResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during AuthTestResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *AuthTestResponse) URL() string {
	return b.url
}

func (b *AuthTestResponse) Team() string {
	return b.team
}

func (b *AuthTestResponse) User() string {
	return b.user
}

func (b *AuthTestResponse) TeamID() string {
	return b.teamId
}

func (b *AuthTestResponse) UserID() string {
	return b.userId
}

func (v *AuthTestResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		URL    string `json:"url"`
		Team   string `json:"team"`
		User   string `json:"user"`
		TeamID string `json:"team_id"`
		UserID string `json:"user_id"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildAuthTestResponse().
		URL(proxy.URL).
		Team(proxy.Team).
		User(proxy.User).
		TeamID(proxy.TeamID).
		UserID(proxy.UserID).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildChannelsHistoryResponse() *ChannelsHistoryResponseBuilder {
	var b ChannelsHistoryResponseBuilder
	return &b
}

func (b *ChannelsHistoryResponseBuilder) HasMore(v bool) *ChannelsHistoryResponseBuilder {
	b.hasMore = v
	return b
}

func (b *ChannelsHistoryResponseBuilder) Latest(v string) *ChannelsHistoryResponseBuilder {
	b.latest = v
	return b
}

func (b *ChannelsHistoryResponseBuilder) Messages(v ...*Message) *ChannelsHistoryResponseBuilder {
	b.messages = v
	return b
}

func (b *ChannelsHistoryResponseBuilder) Build() (*ChannelsHistoryResponse, error) {
	var v ChannelsHistoryResponse
	v.hasMore = b.hasMore
	v.latest = b.latest
	v.messages = b.messages
	return &v, nil
}

func (b *ChannelsHistoryResponseBuilder) MustBuild() *ChannelsHistoryResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during ChannelsHistoryResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *ChannelsHistoryResponse) HasMore() bool {
	return b.hasMore
}

func (b *ChannelsHistoryResponse) Latest() string {
	return b.latest
}

func (b *ChannelsHistoryResponse) Messages() MessageList {
	return b.messages
}

func (v *ChannelsHistoryResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		HasMore  bool        `json:"has_more"`
		Latest   string      `json:"latest"`
		Messages MessageList `json:"messages"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildChannelsHistoryResponse().
		HasMore(proxy.HasMore).
		Latest(proxy.Latest).
		Messages(proxy.Messages...).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildChatResponse() *ChatResponseBuilder {
	var b ChatResponseBuilder
	return &b
}

func (b *ChatResponseBuilder) Channel(v string) *ChatResponseBuilder {
	b.channel = v
	return b
}

func (b *ChatResponseBuilder) Timestamp(v string) *ChatResponseBuilder {
	b.ts = v
	return b
}

func (b *ChatResponseBuilder) Message(v interface{}) *ChatResponseBuilder {
	b.message = v
	return b
}

func (b *ChatResponseBuilder) Build() (*ChatResponse, error) {
	var v ChatResponse
	v.channel = b.channel
	v.ts = b.ts
	v.message = b.message
	return &v, nil
}

func (b *ChatResponseBuilder) MustBuild() *ChatResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during ChatResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *ChatResponse) Channel() string {
	return b.channel
}

func (b *ChatResponse) Timestamp() string {
	return b.ts
}

func (b *ChatResponse) Message() interface{} {
	return b.message
}

func (v *ChatResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		Channel   string      `json:"channel"`
		Timestamp string      `json:"ts"`
		Message   interface{} `json:"message"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildChatResponse().
		Channel(proxy.Channel).
		Timestamp(proxy.Timestamp).
		Message(proxy.Message).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildEphemeralResponse() *EphemeralResponseBuilder {
	var b EphemeralResponseBuilder
	return &b
}

func (b *EphemeralResponseBuilder) MessageTs(v string) *EphemeralResponseBuilder {
	b.messageTs = v
	return b
}

func (b *EphemeralResponseBuilder) Build() (*EphemeralResponse, error) {
	var v EphemeralResponse
	v.messageTs = b.messageTs
	return &v, nil
}

func (b *EphemeralResponseBuilder) MustBuild() *EphemeralResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during EphemeralResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *EphemeralResponse) MessageTs() string {
	return b.messageTs
}

func (v *EphemeralResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		MessageTs string `json:"message_ts"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildEphemeralResponse().
		MessageTs(proxy.MessageTs).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildGenericResponse() *GenericResponseBuilder {
	var b GenericResponseBuilder
	return &b
}

func (b *GenericResponseBuilder) OK(v bool) *GenericResponseBuilder {
	b.ok = v
	return b
}

func (b *GenericResponseBuilder) ReplyTo(v int) *GenericResponseBuilder {
	b.replyTo = v
	return b
}

func (b *GenericResponseBuilder) Error(v *ErrorResponse) *GenericResponseBuilder {
	b.error = v
	return b
}

func (b *GenericResponseBuilder) Timestamp(v string) *GenericResponseBuilder {
	b.ts = v
	return b
}

func (b *GenericResponseBuilder) Build() (*GenericResponse, error) {
	var v GenericResponse
	v.ok = b.ok
	v.replyTo = b.replyTo
	v.error = b.error
	v.ts = b.ts
	return &v, nil
}

func (b *GenericResponseBuilder) MustBuild() *GenericResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during GenericResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *GenericResponse) OK() bool {
	return b.ok
}

func (b *GenericResponse) ReplyTo() int {
	return b.replyTo
}

func (b *GenericResponse) Error() *ErrorResponse {
	return b.error
}

func (b *GenericResponse) Timestamp() string {
	return b.ts
}

func (v *GenericResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		OK        bool           `json:"ok"`
		ReplyTo   int            `json:"reply_to"`
		Error     *ErrorResponse `json:"error"`
		Timestamp string         `json:"ts"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildGenericResponse().
		OK(proxy.OK).
		ReplyTo(proxy.ReplyTo).
		Error(proxy.Error).
		Timestamp(proxy.Timestamp).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildOAuthAccessResponse() *OAuthAccessResponseBuilder {
	var b OAuthAccessResponseBuilder
	return &b
}

func (b *OAuthAccessResponseBuilder) AccessToken(v string) *OAuthAccessResponseBuilder {
	b.accessToken = v
	return b
}

func (b *OAuthAccessResponseBuilder) Scope(v string) *OAuthAccessResponseBuilder {
	b.scope = v
	return b
}

func (b *OAuthAccessResponseBuilder) Build() (*OAuthAccessResponse, error) {
	var v OAuthAccessResponse
	v.accessToken = b.accessToken
	v.scope = b.scope
	return &v, nil
}

func (b *OAuthAccessResponseBuilder) MustBuild() *OAuthAccessResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during OAuthAccessResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *OAuthAccessResponse) AccessToken() string {
	return b.accessToken
}

func (b *OAuthAccessResponse) Scope() string {
	return b.scope
}

func (v *OAuthAccessResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildOAuthAccessResponse().
		AccessToken(proxy.AccessToken).
		Scope(proxy.Scope).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildPermalinkResponse() *PermalinkResponseBuilder {
	var b PermalinkResponseBuilder
	return &b
}

func (b *PermalinkResponseBuilder) Channel(v string) *PermalinkResponseBuilder {
	b.channel = v
	return b
}

func (b *PermalinkResponseBuilder) Permalink(v string) *PermalinkResponseBuilder {
	b.permalink = v
	return b
}

func (b *PermalinkResponseBuilder) Build() (*PermalinkResponse, error) {
	var v PermalinkResponse
	v.channel = b.channel
	v.permalink = b.permalink
	return &v, nil
}

func (b *PermalinkResponseBuilder) MustBuild() *PermalinkResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during PermalinkResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *PermalinkResponse) Channel() string {
	return b.channel
}

func (b *PermalinkResponse) Permalink() string {
	return b.permalink
}

func (v *PermalinkResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		Channel   string `json:"channel"`
		Permalink string `json:"permalink"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildPermalinkResponse().
		Channel(proxy.Channel).
		Permalink(proxy.Permalink).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildRTMResponse() *RTMResponseBuilder {
	var b RTMResponseBuilder
	return &b
}

func (b *RTMResponseBuilder) URL(v string) *RTMResponseBuilder {
	b.url = v
	return b
}

func (b *RTMResponseBuilder) Self(v *UserDetails) *RTMResponseBuilder {
	b.self = v
	return b
}

func (b *RTMResponseBuilder) Team(v *Team) *RTMResponseBuilder {
	b.team = v
	return b
}

func (b *RTMResponseBuilder) Users(v ...*User) *RTMResponseBuilder {
	b.users = v
	return b
}

func (b *RTMResponseBuilder) Channels(v ...*Channel) *RTMResponseBuilder {
	b.channels = v
	return b
}

func (b *RTMResponseBuilder) Groups(v ...*Group) *RTMResponseBuilder {
	b.groups = v
	return b
}

func (b *RTMResponseBuilder) Bots(v ...*Bot) *RTMResponseBuilder {
	b.bots = v
	return b
}

func (b *RTMResponseBuilder) IMs(v ...*IM) *RTMResponseBuilder {
	b.ims = v
	return b
}

func (b *RTMResponseBuilder) Build() (*RTMResponse, error) {
	var v RTMResponse
	v.url = b.url
	v.self = b.self
	v.team = b.team
	v.users = b.users
	v.channels = b.channels
	v.groups = b.groups
	v.bots = b.bots
	v.ims = b.ims
	return &v, nil
}

func (b *RTMResponseBuilder) MustBuild() *RTMResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during RTMResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *RTMResponse) URL() string {
	return b.url
}

func (b *RTMResponse) Self() *UserDetails {
	return b.self
}

func (b *RTMResponse) Team() *Team {
	return b.team
}

func (b *RTMResponse) Users() []*User {
	return b.users
}

func (b *RTMResponse) Channels() []*Channel {
	return b.channels
}

func (b *RTMResponse) Groups() []*Group {
	return b.groups
}

func (b *RTMResponse) Bots() []*Bot {
	return b.bots
}

func (b *RTMResponse) IMs() []*IM {
	return b.ims
}

func (v *RTMResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		URL      string       `json:"url"`
		Self     *UserDetails `json:"self"`
		Team     *Team        `json:"team"`
		Users    []*User      `json:"users"`
		Channels []*Channel   `json:"channels"`
		Groups   []*Group     `json:"groups"`
		Bots     []*Bot       `json:"bots"`
		IMs      []*IM        `json:"ims"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildRTMResponse().
		URL(proxy.URL).
		Self(proxy.Self).
		Team(proxy.Team).
		Users(proxy.Users...).
		Channels(proxy.Channels...).
		Groups(proxy.Groups...).
		Bots(proxy.Bots...).
		IMs(proxy.IMs...).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildReactionsGetResponse() *ReactionsGetResponseBuilder {
	var b ReactionsGetResponseBuilder
	return &b
}

func (b *ReactionsGetResponseBuilder) Channel(v string) *ReactionsGetResponseBuilder {
	b.channel = v
	return b
}

func (b *ReactionsGetResponseBuilder) Comment(v string) *ReactionsGetResponseBuilder {
	b.comment = v
	return b
}

func (b *ReactionsGetResponseBuilder) File(v *File) *ReactionsGetResponseBuilder {
	b.file = v
	return b
}

func (b *ReactionsGetResponseBuilder) Message(v *Message) *ReactionsGetResponseBuilder {
	b.message = v
	return b
}

func (b *ReactionsGetResponseBuilder) Type(v string) *ReactionsGetResponseBuilder {
	b.typ = v
	return b
}

func (b *ReactionsGetResponseBuilder) Build() (*ReactionsGetResponse, error) {
	var v ReactionsGetResponse
	v.channel = b.channel
	v.comment = b.comment
	v.file = b.file
	v.message = b.message
	v.typ = b.typ
	return &v, nil
}

func (b *ReactionsGetResponseBuilder) MustBuild() *ReactionsGetResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during ReactionsGetResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *ReactionsGetResponse) Channel() string {
	return b.channel
}

func (b *ReactionsGetResponse) Comment() string {
	return b.comment
}

func (b *ReactionsGetResponse) File() *File {
	return b.file
}

func (b *ReactionsGetResponse) Message() *Message {
	return b.message
}

func (b *ReactionsGetResponse) Type() string {
	return b.typ
}

func (v *ReactionsGetResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		Channel string   `json:"channel"`
		Comment string   `json:"comment"`
		File    *File    `json:"file"`
		Message *Message `json:"message"`
		Type    string   `json:"typ"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildReactionsGetResponse().
		Channel(proxy.Channel).
		Comment(proxy.Comment).
		File(proxy.File).
		Message(proxy.Message).
		Type(proxy.Type).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}

func BuildReactionsListResponse() *ReactionsListResponseBuilder {
	var b ReactionsListResponseBuilder
	return &b
}

func (b *ReactionsListResponseBuilder) Items(v ...*ReactionsGetResponse) *ReactionsListResponseBuilder {
	b.items = v
	return b
}

func (b *ReactionsListResponseBuilder) Paging(v *Paging) *ReactionsListResponseBuilder {
	b.paging = v
	return b
}

func (b *ReactionsListResponseBuilder) Build() (*ReactionsListResponse, error) {
	var v ReactionsListResponse
	v.items = b.items
	v.paging = b.paging
	return &v, nil
}

func (b *ReactionsListResponseBuilder) MustBuild() *ReactionsListResponse {
	v, err := b.Build()
	if err != nil {
		panic("error during ReactionsListResponse.MustBuild: " + err.Error())
	}
	return v
}

func (b *ReactionsListResponse) Items() ReactionsGetResponseList {
	return b.items
}

func (b *ReactionsListResponse) Paging() *Paging {
	return b.paging
}

func (v *ReactionsListResponse) UnmarshalJSON(data []byte) error {
	var proxy struct {
		Items  ReactionsGetResponseList `json:"items"`
		Paging *Paging                  `json:"paging"`
	}
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	x, err := BuildReactionsListResponse().
		Items(proxy.Items...).
		Paging(proxy.Paging).
		Build()
	if err != nil {
		return errors.Wrap(err, `failed to build object from JSON`)
	}
	*v = *x
	return nil
}
