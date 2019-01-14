package client

import "github.com/micro/go-micro/client"
import proto "cinema/api/proto/user"

// user 服务客户端
func UserClient() proto.UserService {
	return proto.NewUserService("com.cinema.srv.user", client.DefaultClient)
}
