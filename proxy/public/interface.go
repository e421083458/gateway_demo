package public

import (
	"context"
	"net"
)

type TCPHandler interface {
	ServeTCP(ctx context.Context, conn net.Conn)
}
