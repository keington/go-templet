package net

import (
	"context"
	"errors"
	"google.golang.org/grpc/peer"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/26 20:32
 * @file: ip.go
 * @description: ip
 */

// GetClientIP 检查上下文以检索客户机的ip地址
func GetClientIP(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("couldn't parse client IP address")
	}
	return p.Addr.String(), nil
}
