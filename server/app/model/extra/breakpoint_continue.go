package model

import "go_base_server/server/library/global"

//@description: 文件结构体
type BreakpointContinue struct {
	global.Model
	FileMd5    string
	FileName   string
	FilePath   string
	IsFinish   bool
	ChunkTotal int
	FileChunk  []BreakpointContinueChunk `orm:"-" gorm:"foreignKey:FileID"`
}

func (b *BreakpointContinue) TableName() string {
	return "breakpoint_continue"
}

//@description: 切片结构体
type BreakpointContinueChunk struct {
	global.Model
	FileChunkPath   string
	FileID          uint
	FileChunkNumber int
}

func (b *BreakpointContinueChunk) TableName() string {
	return "breakpoint_continue_chunks"
}
