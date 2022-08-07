package controller

import (
	v1 "config-deliver/api/v1"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	GetConfig = cGetConfig{}
)

type cGetConfig struct{}

func (c *cGetConfig) GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error) {
	g.Log().Debugf(ctx, "received request: %+v", req)
	// g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	// var reply string = "Hello World!"

	homedir, err := g.Cfg().Get(ctx, "dirmap."+req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "get config error: %+v", err)
		return nil, err
	}
	g.Log().Debugf(ctx, "查询到映射目录homedir: %s,", homedir)
	// 判断homedir是存为空
	if len(homedir.String()) == 0 {
		g.Log().Debugf(ctx, "homedir is empty")
		g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{"code": 1, "message": "homedir is empty"})
		return nil, nil
	}
	// 判断homedir是否存在
	if !gfile.Exists(homedir.String()) {
		g.Log().Errorf(ctx, "homedir not exists: %s", homedir)
		g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{"code": 1, "message": "homedir not exists"})
		return nil, nil
	}
	// 获取索引目录下所有文件名
	files, err := gfile.ScanDirFile(homedir.String(), "*", false) // false表示不遍历子目录
	if err != nil {
		g.Log().Errorf(ctx, "get file error: %+v", err)
		return nil, err
	}
	// g.Log().Debugf(ctx, "files: %+v", files)
	// 遍历文件名,当文件名为req.Filename时,返回文件内容
	// var pureFiles []string

	var pureFiles []string
	for _, file := range files {
		// reply += file + ","
		// g.Log().Debugf(ctx, "file: %s", file)
		basename := gfile.Basename(file)
		// g.Log().Debugf(ctx, "basename: %s", basename)
		if basename == req.Filename {
			g.Log().Debugf(ctx, "find file: %s", file)
			if req.Dl == "true" {
				g.RequestFromCtx(ctx).Response.ServeFile(file)
				return nil, nil

			} else {
				md5, err := gmd5.EncryptFile(file)
				if err != nil {
					g.Log().Errorf(ctx, "get md5 error: %+v", err)
					return nil, err
				}
				g.RequestFromCtx(ctx).Response.WriteJsonExit(v1.GetConfigRes{
					Code:    0,
					Message: "success",
					Data:    md5,
				})
			}
		}
		pureFiles = append(pureFiles, basename)
	}
	// 判断filename是否存在
	if len(req.Filename) == 0 {
		g.Log().Debugf(ctx, "filename is empty")
		g.Log().Debugf(ctx, "返回文件列表pureFiles: %+v", pureFiles)
		g.RequestFromCtx(ctx).Response.WriteJsonExit(v1.GetConfigRes{
			Code:    0,
			Message: "success",
			Data:    pureFiles,
		})
		return nil, nil
	} else {
		g.Log().Debugf(ctx, "filename not exists: %s", req.Filename)
		g.RequestFromCtx(ctx).Response.WriteJsonExit(v1.GetConfigRes{
			Code:    1,
			Message: "file not exists:" + req.Filename,
			Data:    nil,
		})
		return nil, nil
	}
}
