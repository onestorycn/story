package handler

import (
	"context"
	"go.uber.org/zap"
	story "story/grpc/proto/story"
	"story/library/logger"
	"story/library"
	"github.com/grpc/grpc-go/status"
	"google.golang.org/grpc/codes"
	"fmt"
)

type Story struct{}

func (e *Story) CreateStory(ctx context.Context, req *story.StoryInfo, rsp *story.Response) (err error) {
	logger.ZapInfo.Info("Received Story.CreateStory request", zap.String("request", req.String()))
	defer func() {
		if err := recover(); err != nil {
			rsp.Code = library.InternalError
			isError, ok := err.(error)
			if ok {
				rsp.Message = isError.Error()
				err = status.Error(codes.Aborted, isError.Error())
			} else {
				rsp.Message = library.CodeString(library.InternalError)
				err = fmt.Sprintf("%v", err)
			}
		}
	}()
	rsp.Code = library.CodeSucc
	rsp.Message = library.CodeString(library.CodeSucc)
	return status.Error(codes.OK, "success")
}


func (e *Story) GetStoryList(ctx context.Context, req *story.RequestStoryId, rsp *story.ResponseStoryList) (err error) {
	logger.ZapInfo.Info("Received Story.CreateStory request", zap.String("request", req.String()))
	defer func() {
		if err := recover(); err != nil {
			rsp.Code = library.InternalError
			isError, ok := err.(error)
			if ok {
				rsp.Message = isError.Error()
				err = status.Error(codes.Aborted, isError.Error())
			} else {
				rsp.Message = library.CodeString(library.InternalError)
				err = fmt.Sprintf("%v", err)
			}
		}
	}()
	rsp.Code = library.CodeSucc
	rsp.Message = library.CodeString(library.CodeSucc)
	return status.Error(codes.OK, "success")
}



func (e *Story) UpdateStory(ctx context.Context, req *story.StoryInfo, rsp *story.Response) (err error) {
	logger.ZapInfo.Info("Received Story.CreateStory request", zap.String("request", req.String()))
	defer func() {
		if err := recover(); err != nil {
			rsp.Code = library.InternalError
			isError, ok := err.(error)
			if ok {
				rsp.Message = isError.Error()
				err = status.Error(codes.Aborted, isError.Error())
			} else {
				rsp.Message = library.CodeString(library.InternalError)
				err = fmt.Sprintf("%v", err)
			}
		}
	}()
	rsp.Code = library.CodeSucc
	rsp.Message = library.CodeString(library.CodeSucc)
	return status.Error(codes.OK, "success")
}
