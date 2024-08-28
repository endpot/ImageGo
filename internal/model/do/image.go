// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Image is the golang structure of table image for DAO operations like Where/Data.
type Image struct {
	g.Meta      `orm:"table:image, do:true"`
	Id          interface{} //
	Code        interface{} // Image Unique Code
	DeleteCode  interface{} // Image Delete Code
	Name        interface{} // Image File Name
	Ext         interface{} // Image Extension
	Width       interface{} // Image Width in Pixels
	Height      interface{} // Image Height in Pixels
	Nsfw        interface{} // Normal or NSFW
	UploaderIp  interface{} // Image Uploader IP
	Fingerprint interface{} // Image Fingerprint
	SaveName    interface{} // Image Save Name
	Size        interface{} // Image Size in Bits
	Views       interface{} // Image View Counts
	CreatedAt   *gtime.Time // Create Time
	UpdatedAt   *gtime.Time // Update Time
}
