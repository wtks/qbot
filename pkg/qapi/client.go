package qapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/dghubble/sling"
	"github.com/wtks/qbot/internal/apiclient"
	"net/http"
	"sync"
)

type WrappedClient struct {
	s *sling.Sling
	c *apiclient.ClientWithResponses

	stampToIdMapping map[string]string
	idToStampMapping map[string]string
	stampMappingLock sync.RWMutex

	userNameToIdMapping map[string]string
	idToUserNameMapping map[string]string
	userNameMappingLock sync.RWMutex
}

func NewWrappedClient(accessToken string) (*WrappedClient, error) {
	btp, _ := securityprovider.NewSecurityProviderBearerToken(accessToken)
	c, _ := apiclient.NewClientWithResponses(Endpoint+"/api/v3", apiclient.WithRequestEditorFn(btp.Intercept))
	wc := &WrappedClient{
		s:                   sling.New().Base(Endpoint).Set("Authorization", "Bearer "+accessToken),
		c:                   c,
		stampToIdMapping:    map[string]string{},
		idToStampMapping:    map[string]string{},
		userNameToIdMapping: map[string]string{},
		idToUserNameMapping: map[string]string{},
	}

	if err := wc.fetchStampsInfo(); err != nil {
		return nil, fmt.Errorf("failed to fetch stamps: %w", err)
	}
	if err := wc.fetchUsersInfo(); err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	return wc, nil
}

func (c *WrappedClient) NewReq() *sling.Sling {
	return c.s.New()
}

func (c *WrappedClient) PostMessage(channelID string, text string) (string, error) {
	r, err := c.c.PostMessageWithResponse(context.Background(), apiclient.ChannelIdInPath(channelID), apiclient.PostMessageJSONRequestBody{
		Content: text,
	})
	if err != nil {
		return "", err
	}
	if r.JSON201 == nil {
		return "", fmt.Errorf(r.Status())
	}
	return r.JSON201.Id, nil
}

func (c *WrappedClient) AddMessageStamp(messageID string, stamp string) error {
	r, err := c.c.AddMessageStampWithResponse(context.Background(), apiclient.MessageIdInPath(messageID), apiclient.StampIdInPath(c.GetStampIDByName(stamp)), apiclient.AddMessageStampJSONRequestBody{Count: 1})
	if err != nil {
		return err
	}
	if r.StatusCode() != http.StatusNoContent {
		return errors.New(r.Status())
	}
	return nil
}

func (c *WrappedClient) GetStampIDByName(name string) string {
	c.stampMappingLock.RLock()
	s, ok := c.stampToIdMapping[name]
	c.stampMappingLock.RUnlock()

	if !ok {
		_ = c.fetchStampsInfo()
		return c.stampToIdMapping[name]
	}
	return s
}

func (c *WrappedClient) GetStampNameByID(id string) string {
	c.stampMappingLock.RLock()
	s, ok := c.idToStampMapping[id]
	c.stampMappingLock.RUnlock()

	if !ok {
		_ = c.fetchStampsInfo()
		return c.idToStampMapping[id]
	}
	return s
}

func (c *WrappedClient) fetchStampsInfo() error {
	trueV := true
	stamps, err := c.c.GetStampsWithResponse(context.Background(), &apiclient.GetStampsParams{IncludeUnicode: &trueV})
	if err != nil {
		return err
	}
	c.stampMappingLock.Lock()
	defer c.stampMappingLock.Unlock()
	for _, stamp := range *stamps.JSON200 {
		c.stampToIdMapping[stamp.Name] = stamp.Id
		c.idToStampMapping[stamp.Id] = stamp.Name
	}
	return nil
}

func (c *WrappedClient) GetUserIDByName(name string) string {
	c.userNameMappingLock.RLock()
	s, ok := c.userNameToIdMapping[name]
	c.userNameMappingLock.RUnlock()

	if !ok {
		_ = c.fetchUsersInfo()
		return c.userNameToIdMapping[name]
	}
	return s
}

func (c *WrappedClient) GetUserNameByID(id string) string {
	c.userNameMappingLock.RLock()
	s, ok := c.idToUserNameMapping[id]
	c.userNameMappingLock.RUnlock()

	if !ok {
		_ = c.fetchUsersInfo()
		return c.idToUserNameMapping[id]
	}
	return s
}

func (c *WrappedClient) fetchUsersInfo() error {
	trueV := true
	users, err := c.c.GetUsersWithResponse(context.Background(), &apiclient.GetUsersParams{IncludeSuspended: &trueV})
	if err != nil {
		return err
	}
	c.userNameMappingLock.Lock()
	defer c.userNameMappingLock.Unlock()
	for _, user := range *users.JSON200 {
		c.userNameToIdMapping[user.Name] = user.Id
		c.idToUserNameMapping[user.Id] = user.Name
	}
	return nil
}
