package pupfile

import (
	"encoding/json"
)

type BookDesc struct {
	// 书本格式版本号
	Version string `json:"version"`

	// 书本的背景音乐ID
	MusicId string `json:"musicId"`

	// 是否可编辑的
	Editable bool `json:"editable"`

	// 书本名字(暂时没用)
	Name string `json:"name"`

	// 创建书本的时间戳
	CreateTimestamp int64 `json:"createTimestamp"`

	Pages []*PageDesc `json:"pages"`
}

func NewBookDesc() *BookDesc {
	return &BookDesc{
		Version:  "1.0",
		Editable: true,
	}
}

func (desc *BookDesc) GetPageCount() int {
	return len(desc.Pages)
}

func (desc *BookDesc) IsEmpty() bool {
	return desc.GetPageCount() <= 0
}

func (desc *BookDesc) ToJSON() []byte {
	bytes, _ := json.Marshal(desc)
	return bytes
}

func (desc *BookDesc) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, desc)
}

type PageDesc struct {
	// zip中的文件名
	Image string `json:"image"`

	// zip中的文件名
	Voice string `json:"voice"`

	// 音效时长,主要用于列表中的显示
	VoiceTime float32 `json:"voiceTime"`
}

func NewPageDesc() *PageDesc {
	return &PageDesc{}
}
