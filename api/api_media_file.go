package api

import (
	"demo_backend/model"
	"demo_backend/tool"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

// UploadShowBatch shows uploading multiple files page.
func UploadShowBatch(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>GF Upload Files Demo</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/api/media/upload/3" method="post">
                <input type="file" name="upload-file" />
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}

func handler_api_media_upload(r *ghttp.Request) {
	/*
		上传文件 上传路径为 api/media/upload/{projectid}
		session的accountid，header的projectid，路径为./media/projectid/accountid/filename
	*/
	//获取路径
	accountid := r.Session.MustGet("accountId").String()
	projectid := r.Get("projectid").String()
	dirpath := gfile.Join("./media", projectid, accountid)

	//保存文件
	files := r.GetUploadFiles("upload-file") // 前端file input的name是 upload-file
	names, err := files.Save(dirpath)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_media_upload")})
	} else {
		model.SaveFileMsgs(projectid, accountid, &names) //保存文件信息到库表
		r.Response.WriteJsonExit(g.Map{"status": true, "data": names})
	}
}

func handler_api_media_delete(r *ghttp.Request) {
	/*
		删除文件 路径为 api/media/delete/{projectid}
		参数为 filename
		session的accountid，header的projectid，路径为./media/projectid/accountid/filename
	*/
	//获取路径
	accountid := r.Session.MustGet("accountId").String()
	projectid := r.Get("projectid").String()
	filename := r.Get("filename").String()

	if filename == "" || projectid == "" {
		r.Response.WriteJsonExit(g.Map{
			"status": false,
			"msg":    "handler_api_media_delete: filename 或 projectid 缺失"})
	}

	dirpath := gfile.Join(".", "media", projectid, accountid, filename)
	if false == gfile.Exists(dirpath) {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "该文件可能由其他责任单位上传，您无权删除"})
	}

	if err := gfile.Remove(dirpath); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_media_delete")})
	} else {
		if model.DeleteFileMsg(projectid, accountid, filename) {
			r.Response.WriteJsonExit(g.Map{"status": true})
		} else {
			r.Response.WriteJsonExit(g.Map{"status": false, "msg": "删除失败"})
		}
	}
}

func handler_api_media_filemsgs(r *ghttp.Request) {
	// 获取一个项目的所有文件信息
	projectid := r.Get("projectid").Uint()
	if projectid == 0 {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_media_filemsgs: projectid 缺失"})
	}

	var filemsgs []model.MediaFileMsg
	result := tool.GetGormConnection().Where("projectid = ?", projectid).Find(&filemsgs)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_media_filemsgs")})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": filemsgs})
	}
}

func handler_api_media_download(r *ghttp.Request) {
	// 下载文件 /api/media/download/{projectid}/{accountid}/{filename}"
	projectid := r.Get("projectid").Uint()
	accountid := r.Get("accountid").Uint()
	filename := r.Get("filename").String()
	if projectid == 0 || accountid == 0 || filename == "" {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_media_download：参数不全"})
	}

	dirpath := gfile.Join(".", "media", gconv.String(projectid), gconv.String(accountid), filename)
	if false == gfile.Exists(dirpath) {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "该文件可能由其他责任单位上传，您无权删除"})
	}
}

func RouterGroup_Media(group *ghttp.RouterGroup) {

	group.POST("/upload/{projectid}", handler_api_media_upload)     ///api/media/upload
	group.POST("/delete/{projectid}", handler_api_media_delete)     ///api/media/delete
	group.POST("/filemsgs/{projectid}", handler_api_media_filemsgs) ///api/media/filemsgs

	group.POST("/download/{projectid}/{accountid}/{filename}", handler_api_media_download) ///api/media/download

	group.GET("/uploadshow", UploadShowBatch) //测试界面

}
