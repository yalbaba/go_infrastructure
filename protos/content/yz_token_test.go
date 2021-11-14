/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/16
   Description :
-------------------------------------------------
*/

package content

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestGetYouZanAccessToken(t *testing.T) {
	conn, err := grpc.Dial("47.108.132.82:9091", grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()
	c := NewContentServiceClient(conn)
	resp, err := c.GetYouZanAccessToken(context.Background(), &GetYouZanAccessTokenReq{})
	require.NoError(t, err)
	t.Log(resp.Token)
}
