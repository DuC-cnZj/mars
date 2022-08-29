package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/dustin/go-humanize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"

	"github.com/duc-cnzj/mars-client/v4/file"
	"github.com/duc-cnzj/mars-client/v4/types"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/event/events"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/scopes"
)

func init() {
	RegisterServer(func(s grpc.ServiceRegistrar, app contracts.ApplicationInterface) {
		file.RegisterFileServer(s, new(File))
	})
	RegisterEndpoint(file.RegisterFileHandlerFromEndpoint)
}

type File struct {
	file.UnimplementedFileServer
}

func (m *File) List(ctx context.Context, request *file.ListRequest) (*file.ListResponse, error) {
	var (
		page     = int(request.Page)
		pageSize = int(request.PageSize)
		files    []models.File
		count    int64
	)
	ss := []func(*gorm.DB) *gorm.DB{
		func(db *gorm.DB) *gorm.DB {
			if !request.GetWithoutDeleted() {
				db = db.Unscoped()
			}
			return db
		},
	}
	if err := app.DB().Scopes(append(ss, scopes.Paginate(&page, &pageSize))...).Order("`id` DESC").Find(&files).Error; err != nil {
		return nil, err
	}
	app.DB().Model(&models.File{}).Scopes(ss...).Count(&count)

	var res = make([]*types.FileModel, 0, len(files))
	for _, ff := range files {
		res = append(res, ff.ProtoTransform())
	}

	return &file.ListResponse{
		Page:     int64(page),
		PageSize: int64(pageSize),
		Items:    res,
		Count:    count,
	}, nil
}

func (m *File) DiskInfo(ctx context.Context, request *file.DiskInfoRequest) (*file.DiskInfoResponse, error) {
	size, err := app.Uploader().DirSize()
	if err != nil {
		return nil, err
	}
	return &file.DiskInfoResponse{
		Usage:         size,
		HumanizeUsage: humanize.Bytes(uint64(size)),
	}, nil
}

type listFiles []*types.FileModel

type item struct {
	Name string `yaml:"name"`
	Size string `yaml:"size"`
}

func (l listFiles) PrettyYaml() string {
	var items = make([]item, 0, len(l))
	for _, f := range l {
		items = append(items, item{
			Name: f.Path,
			Size: f.HumanizeSize,
		})
	}
	marshal, _ := yaml.Marshal(items)
	return string(marshal)
}

// DeleteUndocumentedFiles
// TODO: rename method
func (m *File) DeleteUndocumentedFiles(ctx context.Context, _ *file.DeleteUndocumentedFilesRequest) (*file.DeleteUndocumentedFilesResponse, error) {
	var (
		files []models.File

		clearList     = make(listFiles, 0)
		uploader      = app.Uploader()
		localUploader = app.LocalUploader()

		cleanFunc = func(up contracts.Uploader, db *gorm.DB, fileID int, filePath string) bool {
			if !up.Exists(filePath) {
				tx := db.Delete(&models.File{ID: fileID})
				if tx.Error != nil {
					mlog.Error(tx.Error)
				}
				return true
			}
			return false
		}
	)

	app.DB().FindInBatches(&files, 100, func(tx *gorm.DB, batch int) error {
		for _, f := range files {
			var deleted bool
			switch f.UploadType {
			case uploader.Type():
				deleted = cleanFunc(uploader, tx, f.ID, f.Path)
			case localUploader.Type():
				deleted = cleanFunc(localUploader, tx, f.ID, f.Path)
			}
			if deleted {
				clearList = append(clearList, f.ProtoTransform())
			}
		}
		return nil
	})

	localUploader.RemoveEmptyDir()
	events.AuditLog(MustGetUser(ctx).Name, types.EventActionType_Delete, "删除未被记录的文件", clearList, nil)

	return &file.DeleteUndocumentedFilesResponse{Items: clearList}, nil
}

func (*File) Delete(ctx context.Context, request *file.DeleteRequest) (*file.DeleteResponse, error) {
	var f = &models.File{ID: int(request.Id)}
	if err := app.DB().First(&f).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	f.DeleteFile()
	AuditLog(
		MustGetUser(ctx).Name,
		types.EventActionType_Delete,
		fmt.Sprintf("删除文件: '%s', 该文件由 %s 上传, 大小是 %s", f.Path, f.Username, humanize.Bytes(f.Size)))

	return &file.DeleteResponse{}, nil
}

func (*File) MaxUploadSize(ctx context.Context, request *file.MaxUploadSizeRequest) (*file.MaxUploadSizeResponse, error) {
	return &file.MaxUploadSizeResponse{
		HumanizeSize: humanize.Bytes(app.Config().MaxUploadSize()),
		Bytes:        app.Config().MaxUploadSize(),
	}, nil
}

func (m *File) Authorize(ctx context.Context, fullMethodName string) (context.Context, error) {
	if fullMethodName == "MaxUploadSize" {
		return ctx, nil
	}

	if !MustGetUser(ctx).IsAdmin() {
		return nil, status.Error(codes.PermissionDenied, ErrorPermissionDenied.Error())
	}

	return ctx, nil
}
