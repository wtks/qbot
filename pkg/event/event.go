package event

import (
	"time"
)

const TagAdded = "TAG_ADDED"

type TagAddedPayload struct {
	EventTime time.Time `json:"eventTime"`
	TagID     string    `json:"tagId"`
	Tag       string    `json:"tag"`
}
type TagAddedFunc func(payload TagAddedPayload)
type TagAddedFuncInterface interface {
	HandleTagAdded(payload TagAddedPayload)
}

const TagRemoved = "TAG_REMOVED"

type TagRemovedPayload struct {
	EventTime time.Time `json:"eventTime"`
	TagID     string    `json:"tagId"`
	Tag       string    `json:"tag"`
}
type TagRemovedFunc func(payload TagRemovedPayload)
type TagRemovedFuncInterface interface {
	HandleTagRemoved(payload TagRemovedPayload)
}

const UserCreated = "USER_CREATED"

type UserCreatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	User      User      `json:"user"`
}
type UserCreatedFunc func(payload UserCreatedPayload)
type UserCreatedFuncInterface interface {
	HandleUserCreated(payload UserCreatedPayload)
}

const Ping = "PING"

type PingPayload struct {
	EventTime time.Time `json:"eventTime"`
}
type PingFunc func(payload PingPayload)
type PingFuncInterface interface {
	HandlePing(payload PingPayload)
}

const Joined = "JOINED"

type JoinedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Channel   Channel   `json:"channel"`
}
type JoinedFunc func(payload JoinedPayload)
type JoinedFuncInterface interface {
	HandleJoined(payload JoinedPayload)
}

const Left = "LEFT"

type LeftPayload struct {
	EventTime time.Time `json:"eventTime"`
	Channel   Channel   `json:"channel"`
}
type LeftFunc func(payload LeftPayload)
type LeftFuncInterface interface {
	HandleLeft(payload LeftPayload)
}

const StampCreated = "STAMP_CREATED"

type StampCreatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	FileID    string    `json:"fileId"`
	Creator   User      `json:"creator"`
}
type StampCreatedFunc func(payload StampCreatedPayload)
type StampCreatedFuncInterface interface {
	HandleStampCreated(payload StampCreatedPayload)
}

const ChannelCreated = "CHANNEL_CREATED"

type ChannelCreatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Channel   Channel   `json:"channel"`
}
type ChannelCreatedFunc func(payload ChannelCreatedPayload)
type ChannelCreatedFuncInterface interface {
	HandleChannelCreated(payload ChannelCreatedPayload)
}

const ChannelTopicChanged = "CHANNEL_TOPIC_CHANGED"

type ChannelTopicChangedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Channel   Channel   `json:"channel"`
	Topic     string    `json:"topic"`
	Updater   User      `json:"updater"`
}
type ChannelTopicChangedFunc func(payload ChannelTopicChangedPayload)
type ChannelTopicChangedFuncInterface interface {
	HandleChannelTopicChanged(payload ChannelTopicChangedPayload)
}

const MessageCreated = "MESSAGE_CREATED"

type MessageCreatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Message   Message   `json:"message"`
}
type MessageCreatedFunc func(payload MessageCreatedPayload)
type MessageCreatedFuncInterface interface {
	HandleMessageCreated(payload MessageCreatedPayload)
}

const MessageDeleted = "MESSAGE_DELETED"

type MessageDeletedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Message   struct {
		ID        string `json:"id"`
		ChannelID string `json:"channelId"`
	} `json:"message"`
}
type MessageDeletedFunc func(payload MessageDeletedPayload)
type MessageDeletedFuncInterface interface {
	HandleMessageDeleted(payload MessageDeletedPayload)
}

const MessageUpdated = "MESSAGE_UPDATED"

type MessageUpdatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Message   Message   `json:"message"`
}
type MessageUpdatedFunc func(payload MessageUpdatedPayload)
type MessageUpdatedFuncInterface interface {
	HandleMessageUpdated(payload MessageUpdatedPayload)
}

const DirectMessageCreated = "DIRECT_MESSAGE_CREATED"

type DirectMessageCreatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Message   Message   `json:"message"`
}
type DirectMessageCreatedFunc func(payload DirectMessageCreatedPayload)
type DirectMessageCreatedFuncInterface interface {
	HandleDirectMessageCreated(payload DirectMessageCreatedPayload)
}

const DirectMessageDeleted = "DIRECT_MESSAGE_DELETED"

type DirectMessageDeletedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Message   struct {
		ID        string `json:"id"`
		UserID    string `json:"userId"`
		ChannelID string `json:"channelId"`
	} `json:"message"`
}
type DirectMessageDeletedFunc func(payload DirectMessageDeletedPayload)
type DirectMessageDeletedFuncInterface interface {
	HandleDirectMessageDeleted(payload DirectMessageDeletedPayload)
}

const DirectMessageUpdated = "DIRECT_MESSAGE_UPDATED"

type DirectMessageUpdatedPayload struct {
	EventTime time.Time `json:"eventTime"`
	Message   Message   `json:"message"`
}
type DirectMessageUpdatedFunc func(payload DirectMessageUpdatedPayload)
type DirectMessageUpdatedFuncInterface interface {
	HandleDirectMessageUpdated(payload DirectMessageUpdatedPayload)
}

const BotMessageStampsUpdated = "BOT_MESSAGE_STAMPS_UPDATED"

type BotMessageStampsUpdatedPayload struct {
	EventTime time.Time      `json:"eventTime"`
	MessageID string         `json:"messageId"`
	Stamps    []MessageStamp `json:"stamps"`
}
type BotMessageStampsUpdatedFunc func(payload BotMessageStampsUpdatedPayload)
type BotMessageStampsUpdatedFuncInterface interface {
	HandleBotMessageStampsUpdated(payload BotMessageStampsUpdatedPayload)
}
