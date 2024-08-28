// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImageDao is the data access object for table image.
type ImageDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns ImageColumns // columns contains all the column names of Table for convenient usage.
}

// ImageColumns defines and stores column names for table image.
type ImageColumns struct {
	Id          string //
	Code        string // Image Unique Code
	DeleteCode  string // Image Delete Code
	Name        string // Image File Name
	Ext         string // Image Extension
	Width       string // Image Width in Pixels
	Height      string // Image Height in Pixels
	Nsfw        string // Normal or NSFW
	UploaderIp  string // Image Uploader IP
	Fingerprint string // Image Fingerprint
	SaveName    string // Image Save Name
	Size        string // Image Size in Bits
	Views       string // Image View Counts
	CreatedAt   string // Create Time
	UpdatedAt   string // Update Time
}

// imageColumns holds the columns for table image.
var imageColumns = ImageColumns{
	Id:          "id",
	Code:        "code",
	DeleteCode:  "delete_code",
	Name:        "name",
	Ext:         "ext",
	Width:       "width",
	Height:      "height",
	Nsfw:        "nsfw",
	UploaderIp:  "uploader_ip",
	Fingerprint: "fingerprint",
	SaveName:    "save_name",
	Size:        "size",
	Views:       "views",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewImageDao creates and returns a new DAO object for table data access.
func NewImageDao() *ImageDao {
	return &ImageDao{
		group:   "default",
		table:   "image",
		columns: imageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ImageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ImageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ImageDao) Columns() ImageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ImageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ImageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ImageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
