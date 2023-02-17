package api

import (
	"demo_backend/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
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
	dirpath := gfile.Join("./media", projectid, accountid, filename)

	if err := gfile.Remove(dirpath); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_media_delfile")})
	} else {
		model.DeleteFileMsg(projectid, accountid, filename)
		r.Response.WriteJsonExit(g.Map{"status": true})
	}

}

func RouterGroup_Media(group *ghttp.RouterGroup) {

	group.POST("/upload/{projectid}", handler_api_media_upload) ///api/media/upload
	group.POST("/delete/{projectid}", handler_api_media_delete) ///api/media/delete

	// group.GET("/uploadshow", UploadShowBatch)

}
