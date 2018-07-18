package handler
//
//import (
//	"context"
//	"go.uber.org/zap"
//	story "story/grpc/proto/story"
//	storyModel "story/page/model/story"
//	"story/library/logger"
//	"story/library"
//	"github.com/grpc/grpc-go/status"
//	"google.golang.org/grpc/codes"
//	"fmt"
//	"errors"
//	"time"
//)
//
//type Account struct{}
//
//// Call is a single request handler called via client.Call or the generated client code
//func (e *Account) InsertAccountInfo(ctx context.Context, req *story.Request, rsp *story.ResponseSafe) (err error) {
//	logger.ZapInfo.Info("Received Account.InsertAccountInfo request", zap.String("request", req.String()))
//	defer func() {
//		if err := recover(); err != nil {
//			rsp.Code = library.InternalError
//			isError, ok := err.(error)
//			if ok {
//				rsp.Message = isError.Error()
//				err = status.Error(codes.Aborted, isError.Error())
//			} else {
//				rsp.Message = library.CodeString(library.InternalError)
//				err = fmt.Sprintf("%v", err)
//			}
//		}
//	}()
//	storyService := accModel.LoadAccountService()
//	if storyService == nil {
//		panic(errors.New("load story service fail"))
//	}
//	if library.IsEmpty(req.Info.Password) {
//		panic(errors.New("password invalid"))
//	}
//	userProfile := new(accModel.UserProfile)
//	userProfile.Openid = req.Info.OpenId
//	userProfile.Passid = req.Info.PassId
//	userProfile.Email = req.Info.Email
//	userProfile.Avatar = req.Info.Avatar
//	userProfile.Password = library.EncodeMd5(req.Info.Password)
//	userProfile.Update_time = time.Now().Unix()
//	userProfile.Phone = req.Info.Phone
//	userProfile.Nick_name = req.Info.NickName
//
//	inserErr := storyService.InsertNewUser(userProfile)
//	if inserErr != nil {
//		panic(inserErr)
//	}
//	safeAccount := story.SafeAccount{
//		OpenId:   userProfile.Openid,
//		PassId:   userProfile.Passid,
//		Phone:    userProfile.Phone,
//		Ext:      userProfile.Ext,
//		Avatar:   userProfile.Avatar,
//		NickName: userProfile.Nick_name,
//		Email:    userProfile.Email,
//	}
//	rsp.Data = &safeAccount
//	rsp.Code = library.CodeSucc
//	rsp.Message = library.CodeString(library.CodeSucc)
//	return status.Error(codes.OK, "success")
//}
