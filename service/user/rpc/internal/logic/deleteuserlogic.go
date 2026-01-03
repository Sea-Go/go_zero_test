package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *__.DeleteUserReq) (*__.DeleteUserResp, error) {
	// todo: add your logic here and delete this line

	return &__.DeleteUserResp{}, nil
}
