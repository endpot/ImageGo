// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Image is the golang structure for table image.
type Image struct {
	Id          uint64      `json:"id"          description:""`
	Code        string      `json:"code"        description:"Image Unique Code"`
	DeleteCode  string      `json:"deleteCode"  description:"Image Delete Code"`
	Name        string      `json:"name"        description:"Image File Name"`
	Ext         string      `json:"ext"         description:"Image Extension"`
	Width       int         `json:"width"       description:"Image Width in Pixels"`
	Height      int         `json:"height"      description:"Image Height in Pixels"`
	Nsfw        int         `json:"nsfw"        description:"Normal or NSFW"`
	UploaderIp  string      `json:"uploaderIp"  description:"Image Uploader IP"`
	Fingerprint string      `json:"fingerprint" description:"Image Fingerprint"`
	SaveName    string      `json:"saveName"    description:"Image Save Name"`
	Size        int64       `json:"size"        description:"Image Size in Bits"`
	Views       uint64      `json:"views"       description:"Image View Counts"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"Create Time"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"Update Time"`
}
