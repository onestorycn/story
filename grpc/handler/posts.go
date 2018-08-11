package handler

import (
	"context"
	"go.uber.org/zap"
	posts "story/grpc/proto/posts"
	storyModel "story/page/model/story"
	"story/library/logger"
	"story/library"
	"github.com/grpc/grpc-go/status"
	"google.golang.org/grpc/codes"
	"fmt"
	"errors"
	"time"
	"strconv"
)

type Posts struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Posts) InsertPostsInfo(ctx context.Context, req *posts.Request, rsp *posts.ResponseSafe) (err error) {
	logger.ZapInfo.Info("Received Posts.InsertPostsInfo request", zap.String("request", req.String()))
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
	storyService := storyModel.LoadPostService()
	if storyService == nil {
		panic(errors.New("load story service fail"))
	}
	now := time.Now().Unix()
	postInfo := new(storyModel.Post)
	postInfo.Storyid = req.Info.StoryId
	postInfo.Passid = req.Info.PassId
	postInfo.Header = req.Info.Header
	postInfo.Content = req.Info.Content
	postInfo.Rel = req.Info.Rel
	postInfo.Uid = req.Info.Uid
	postInfo.Update_time = now
	postInfo.Create_time = now
	dateconv := library.TimeFormatter(now, "20060102")
	dateConvInt64, errDate  := strconv.ParseInt(dateconv, 10, 64)
	if errDate != nil {
		panic(errDate)
	}
	postInfo.Create_date = dateConvInt64

	inserErr := storyService.InsertNewPost(postInfo)
	if inserErr != nil {
		panic(inserErr)
	}
	safePosts := posts.SafePost {
		StoryId:    postInfo.Storyid,
		PassId:     postInfo.Passid,
		Header:     postInfo.Header,
		Ext:        postInfo.Ext,
		Content:    postInfo.Content,
		Rel:        postInfo.Rel,
		CreateTime: postInfo.Create_time,
		UpdateTime: postInfo.Update_time,
	}
	rsp.Data = &safePosts
	rsp.Code = library.CodeSucc
	rsp.Message = library.CodeString(library.CodeSucc)
	return status.Error(codes.OK, "success")
}


func (e *Posts) UpdatePostsInfo(ctx context.Context, req *posts.Request, rsp *posts.Response) (err error) {
	logger.ZapInfo.Info("Received Posts.InsertPostsInfo request", zap.String("request", req.String()))
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
	storyService := storyModel.LoadPostService()
	if storyService == nil {
		panic(errors.New("load story service fail"))
	}
	return nil
}

func (e *Posts) GetPostsList(ctx context.Context, req *posts.RequestQuery, rsp *posts.ResponseList) (err error) {
	logger.ZapInfo.Info("Received Posts.InsertPostsInfo request", zap.String("request", req.String()))
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
	storyService := storyModel.LoadPostService()
	if storyService == nil {
		panic(errors.New("load story service fail"))
	}

	list , errGet := storyService.GetPostListByConds(req.PassId, req.StoryId, req.StartTime, req.EndTime, req.Limit, req.Page, req.IsDesc)
	if errGet != nil{

	}
	outputList := make([]*posts.SafePost, 0)
	for _, val := range list {
		safePost := new(posts.SafePost)
		safePost.StoryId = val.Storyid
		safePost.PassId = val.Passid
		safePost.Content = val.Content
		safePost.Ext = val.Ext
		safePost.Rel = val.Rel
		safePost.Header = val.Header
		safePost.UpdateTime = val.Update_time
		safePost.CreateTime = val.Create_time
		safePost.CreateDate = val.Create_date
		outputList = append(outputList, safePost)
	}
	rsp.List = outputList
	return nil
}

func (e *Posts) GetPostById(ctx context.Context, req *posts.RequestQuerySingle, rsp *posts.ResponseSafe) (err error) {
	logger.ZapInfo.Info("Received Posts.InsertPostsInfo request", zap.String("request", req.String()))
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
	storyService := storyModel.LoadPostService()
	if storyService == nil {
		panic(errors.New("load story service fail"))
	}

	return nil
}
