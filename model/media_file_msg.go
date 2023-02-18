package model

import (
	"demo_backend/tool"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

/*
媒体文件信息表
*/
type MediaFileMsg struct {
	Projectid uint      `gorm:"index:index1;"`
	Accountid uint      `gorm:"index:index2;"`
	Filename  string    `gorm:"index:index3;size:50"`
	CreatedAt time.Time `json:"time"`
}

func SaveFileMsgs(projectid, accountid string, filenames *[]string) {
	// 多个文件信息保存到库表
	for _, fname := range *filenames {
		result := tool.GetGormConnection().Create(&MediaFileMsg{
			Projectid: gconv.Uint(projectid), Accountid: gconv.Uint(accountid), Filename: fname,
		})
		if result.Error != nil {
			panic(gerror.Wrap(result.Error, "SaveFileMsgs"))
		}
	}
}

func DeleteFileMsg(projectid, accountid, filename string) {
	// 单个文件信息删除
	result := tool.GetGormConnection().Where("projectid = ? and accountid = ? and filename = ?",
		gconv.Uint(projectid), gconv.Uint(accountid), filename).Delete(&MediaFileMsg{})
	if result.Error != nil {
		panic(gerror.Wrap(result.Error, "deleteFileMsg"))
	}
}

func DeleteFileMsgByProjectid(proid uint) {
	// 删除整个项目的文件信息，并删除其文件。删除项目时使用
	tool.GetGormConnection().Where("projectid = ?", proid).Delete(&MediaFileMsg{})
	dirpath := gfile.Join("./media", gconv.String(proid))
	gfile.Remove(dirpath)
}

func DeleteFileMsgByAccountid(acid uint) {
	// 删除个人账号的文件信息，并删除其文件。删除账号时使用

}

func init() {
	tool.GetGormConnection().AutoMigrate(&MediaFileMsg{})
}
