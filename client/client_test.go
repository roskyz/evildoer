package client

import (
	"context"
	"evildoer/protogen"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	t.Cleanup(cancel)
	client, err := NewGroupClient(ctx, "server.web.local:7070", nil)
	assert.Nil(t, err)

	ctx = context.Background()
	listModeResponse, err := client.ListMode(ctx, new(emptypb.Empty))
	assert.Nil(t, err)
	t.Log(listModeResponse)

	createResponse, err := client.CreateGroup(ctx, &protogen.Group{
		Key:         "test-auto-group",
		Name:        "test-auto-name",
		Description: "test if project's working right",
		State:       protogen.GroupState_GROUP_STATE_NORMAL,
		Creator:     "rex",
		AccessMode:  []protogen.AccessMode{protogen.AccessMode_ACCESS_MODE_API},
	})
	assert.Nil(t, err)
	t.Log(createResponse.IdResp)

	getResponse, err := client.GetGroup(ctx, &protogen.IDReqOrResponse{Id: []byte("647f48397cbf6a04dad83bc3")})
	assert.Nil(t, err)
	t.Log(getResponse.GetGroup())

	group := getResponse.GetGroup()
	group.Name = "change-name"
	t.Log(client.UpdateGroup(ctx, group))
}
