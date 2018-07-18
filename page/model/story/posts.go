package story

import (
	"sync"
	"story/page/service/mysql"
	"story/library/config"
	"story/library"
	"errors"
	"time"
)

var (
	tableName  = "posts"
	once       sync.Once
	serviceMap = make(map[string]*PostService, 0)
	lock       = &sync.Mutex{}
)

type Post struct {
	Id          int64  `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id"`
	Storyid     int64  `gorm:"not null" json:"storyid"`
	Passid      string `gorm:"not null" json:"passid"`
	Uid         int64  `gorm:"not null" json:"uid"`
	Header      string `gorm:"not null" json:"header"`
	Rel         string `gorm:"not null" json:"rel"`
	Content     string `gorm:"not null" json:"content"`
	Ext         string `gorm:"not null" json:"ext"`
	Update_time int64  `gorm:"not null" json:"update_time"`
	Create_time int64  `gorm:"not null" json:"create_time"`
	Create_date int64  `gorm:"not null" json:"create_date"`
}

func (post *Post) TableName() string {
	return tableName
}

type PostService struct {
	DbInstance *mysql.MysqlDbInfo
	env        string
}

func LoadPostService() (postService *PostService) {
	var env string
	envGet := config.GetServiceEnv("env")
	if envGet == nil {
		env = "prod"
	} else {
		env = envGet.(string)
	}
	if _, ok := serviceMap[env]; ok {
		return serviceMap[env]
	}
	lock.Lock()
	defer lock.Unlock()
	if _, ok := serviceMap[env]; !ok {
		mysqlInstance := mysql.LoadMysqlConn(env)
		if mysqlInstance == nil || mysqlInstance.Conn == nil {
			return nil
		} else {
			sockService := new(PostService)
			sockService.DbInstance = mysqlInstance
			sockService.env = env
			serviceMap[env] = sockService
		}
	}
	return serviceMap[env]
}

func (postService *PostService) InsertNewPost(post *Post) error {
	conn := postService.DbInstance.CheckAndReturnConn()
	res := conn.Create(&post)
	if res.Error != nil {
		return res.Error
	}
	return nil
}


func (postService *PostService) UpdatePosts(Id int64, PassId, Header, Rel, Content, Ext string) (int64, error) {
	var updates = make(map[string]interface{}, 0)

	if Id == 0 || len(PassId) < 1{
		return 0, errors.New("update not valid")
	}
	if resBool := library.IsEmpty(Header); !resBool {
		updates["header"] = Header
	}
	if resBool := library.IsEmpty(Rel); !resBool {
		updates["rel"] = Rel
	}
	if resBool := library.IsEmpty(Content); !resBool {
		updates["content"] = Content
	}
	if resBool := library.IsEmpty(Ext); !resBool {
		updates["ext"] = Ext
	}

	if len(updates) == 0 {
		return 0, errors.New("no field to update")
	}

	updates["update_time"] = time.Now().Unix()

	conn := postService.DbInstance.CheckAndReturnConn()
	if conn == nil {
		return 0, errors.New("get db connection fail")
	}
	user := new(Post)
	upRes := conn.Model(&user).Where("id = ? AND passid = ?", Id, PassId).Updates(updates)
	if upRes.Error != nil {
		return 0, upRes.Error
	}
	return upRes.RowsAffected, nil
}

func (postService *PostService) GetPostListByConds(PassId string, StoryId, startTime, endTime, limit, page int64, isDesc bool) ([]*Post, error) {
	var post = make([]*Post, 0)
	if limit == 0 {
		limit = 10
	}

	if page < 1 {
		page = 1
	}

	conn := postService.DbInstance.CheckAndReturnConn()
	if conn == nil {
		return nil , errors.New("get db connection fail")
	}

	var startTimeUse int64
	var endTimeUse int64

	if endTime == 0 {
		endTimeUse = time.Now().Unix()
	}else{
		endTimeUse = endTime
	}
	startTimeUse= startTime
	offset := (page-1) * limit
	orderby := "id asc"
	if isDesc {
		orderby = "id desc"
	}
	res := conn.Where("passid = ? AND storyid = ? AND create_time >= ? AND create_time <= ?", PassId, StoryId, startTimeUse, endTimeUse).Offset(offset).Limit(limit).Order(orderby).Find(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}

func (postService *PostService) GetPostListByCondsCount(PassId string, StoryId, startTime, endTime, limit, page int64, isDesc bool) ([]*Post, error) {
	return nil, nil
}

func (postService *PostService) GetPostById(Id int64, Passid string) (*Post, error) {
	post := new(Post)
	conn := postService.DbInstance.CheckAndReturnConn()
	if conn == nil {
		return nil , errors.New("get db connection fail")
	}
	res := conn.Where("id = ? AND passid = ?", Id, Passid).First(&post)
	if res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}