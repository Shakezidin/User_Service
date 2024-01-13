package handler

import (
	cpb "github.com/Shakezidin/pkg/user/pb"
	SVCinter "github.com/Shakezidin/pkg/user/service/interface"
)

type UserHandler struct {
	SVC SVCinter.UserSVCInter
	cpb.UserServer
}

func NewUserHandler(svc SVCinter.UserSVCInter) *UserHandler {
	return &UserHandler{
		SVC: svc,
	}
}
