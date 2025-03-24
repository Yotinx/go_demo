package main

import (
	"code.byted.org/gopkg/logs"
	"context"
	rpc "github.com/Yotinx/go_demo/kitex_gen/toutiao/demo/rpc"
)

// GetItem implements the ItemServiceImpl interface.

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *rpc.GetItemRequest) (resp *rpc.GetItemResponse, err error) {
	// 输出日志
	logs.CtxDebug(ctx, "GetItem called")
	// 填写返回对象内容
	// 需要注意 这里是rpc.xxx因为我的路径是.../rpc引入的是包名为rpc的包，不能直接贴过去用
	resp = rpc.NewGetItemResponse()
	resp.Item = rpc.NewItem()
	resp.Item.Id = 1024
	resp.Item.Title = "Hello KiteX!"
	resp.Item.Content = "KiteX is the best framework!"
	return
}
