package pupfile

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	PupFileExtName = ".pup"
	DescFileName   = "desc.json"
	ErrFileExist   = errors.New("file is exist")
)

type PupFile struct {
	desc *BookDesc
	w    *zip.Writer
	r    *zip.ReadCloser
}

func NewPupFile() *PupFile {
	return &PupFile{}
}

func (pf *PupFile) Create(zipFileName string) error {
	// if _, err := os.Stat(zipFileName); err == nil {
	// 	return ErrFileExist
	// }

	// Create a new zip archive.
	pf.desc = NewBookDesc()
	pf.desc.CreateTimestamp = time.Now().Unix()

	fzip, _ := os.Create(zipFileName)
	pf.w = zip.NewWriter(fzip)

	return nil
}

func (pf *PupFile) Open(zipFileName string) error {
	// Open a zip archive for reading.
	var err error
	pf.r, err = zip.OpenReader(zipFileName)
	if err != nil {
		return err
	}

	pf.desc = NewBookDesc()
	data := pf.getFileBytes(DescFileName)
	pf.desc.FromBytes(data)

	return nil
}

// AddEmptyPage 在末尾添加一页新书页
func (pf *PupFile) AddEmptyPage() int {
	pageDesc := NewPageDesc()
	pf.desc.Pages = append(pf.desc.Pages, pageDesc)
	return pf.GetPageCount() - 1
}

func (pf *PupFile) IsEmpty() bool {
	return pf.desc.IsEmpty()
}

func (pf *PupFile) GetPageCount() int {
	return pf.desc.GetPageCount()
}

func (pf *PupFile) SetMusicID(musicID string) {
	pf.desc.MusicId = musicID
}

func (pf *PupFile) SetEditable(editable bool) {
	pf.desc.Editable = editable
}

func (pf *PupFile) SetPageImage(index int, name string, imgData []byte) {
	pf.desc.Pages[index].Image = name
	pf.addFile(name, imgData)
}

func (pf *PupFile) SetPageVoice(index int, name string, voiceData []byte, voiceTime float32) {
	pf.desc.Pages[index].Voice = name
	pf.desc.Pages[index].VoiceTime = voiceTime
	pf.addFile(name, voiceData)
}

func (pf *PupFile) Close() error {
	var err error
	if pf.w != nil {
		pf.addDesc()
		err = pf.w.Close()
	}
	return err
}

func (pf *PupFile) addDesc() {
	pf.addFile(DescFileName, pf.desc.ToJSON())
}

func (pf *PupFile) addFile(filename string, bytes []byte) bool {
	f, err := pf.w.Create(filename)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = f.Write(bytes)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (pf *PupFile) getFile(fileName string) *zip.File {
	for _, v := range pf.r.File {
		if v.Name == fileName {
			return v
		}
	}
	return nil
}

func (pf *PupFile) getFileBytes(fileName string) []byte {
	f := pf.getFile(fileName)
	r, err := f.Open()
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil
	}

	return data
}

func (pf *PupFile) GetPageImage(index int) []byte {
	data := pf.getFileBytes(pf.desc.Pages[index].Image)
	return data
}

func (pf *PupFile) GetCoverImage() []byte {
	return pf.GetPageImage(0)
}
