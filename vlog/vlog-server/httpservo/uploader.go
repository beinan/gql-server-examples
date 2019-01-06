package httpservo

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/beinan/gql-server/logging"
)

//MaxUploadSize 是最大文件上传限制
const MaxUploadSize = 5 * 1024 * 1024 * 1024 //5GB

//MaxMemSize 是最大内存使用限制
const MaxMemSize = 20 * 1024 * 1024 //20MB

//Upload 是上传文件的处理函数
//接受两个表单字段video_file和video_id
//使用video_id作为主文件名存储视频文件
func Upload(w http.ResponseWriter, r *http.Request) {
	var logger = logging.StandardLogger(logging.DEBUG)
	logger.Debug("Receiving an upload request")
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxMemSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	formFile, _, err := r.FormFile("video_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer formFile.Close()
	fileHeader := io.LimitReader(formFile, 512) //读取文件的前512字节用于视频类型检查
	fileBytes, err := ioutil.ReadAll(fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fileType := http.DetectContentType(fileBytes)
	logger.Debug("A File received. Content-Type:", fileType, formFile)
	//vlog原型只接受mp4的视频格式
	if fileType != "video/mp4" {
		http.Error(w, "Invalid video type:"+fileType, http.StatusBadRequest)
		return
	}
	fileEndings, _ := mime.ExtensionsByType(fileType)
	logger.Debug("File extension calculated:", fileEndings[0])

	videoID := r.PostFormValue("video_id")
	logger.Debug("Video id from upload request:", videoID)
	destFilePath := filepath.Join("./", videoID+fileEndings[0])
	logger.Debug("Calculated localFilePath by video id:", destFilePath)

	//写入视频文件到本地文件夹
	destFile, err := os.Create(destFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer destFile.Close()
	formFile.Seek(0, 0) //重置源文件，从文件开始处读取
	//copy文件formFile到destFile
	if _, err := io.Copy(destFile, formFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Debug("Saved file to ", destFilePath)
	fmt.Fprintf(w, "Ok")

}
