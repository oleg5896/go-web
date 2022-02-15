package handler

import (
	"fmt"
	"io"
	"net/http"

	// "net/http"
	// "path/filepath"

	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	goweb "github.com/oleg5896/go-web"
)

func (h *Handler) getItem(c *gin.Context) {
	Display(*c, "upload", nil)
}

type Progress struct {
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		logrus.Println("DONE!")
		return
	}
	status := float32(pr.BytesRead) / float32(pr.TotalSize) * 100
	logrus.Println(status)
}

func (h *Handler) addItem(c *gin.Context) {
	var input goweb.File
	var ids []int
	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, fileHeader := range files {

		// Open the file
		file, err := fileHeader.Open()
		input.Path = fileHeader.Filename
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		defer file.Close()
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%s", input.Path))
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		defer f.Close()

		pr := &Progress{
			TotalSize: fileHeader.Size,
		}

		_, err = io.Copy(f, io.TeeReader(file, pr))
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		id, err := h.services.AddItem.AddFile(input)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			ids = append(ids, id)
		}
	}
	

	c.JSON(http.StatusOK, map[string]interface{}{
		"ids": ids,
	})

	// err := c.Request.ParseMultipartForm(200000) // grab the multipart form
	// if err != nil {
	// 	fmt.Fprintln(c.Writer, err)
	// 	return
	// }

	// formdata := c.Request.MultipartForm // ok, no problem so far, read the Form data

	// //get the *fileheaders
	// files := formdata.File["myFiles"] // grab the filenames

	// for i, _ := range files { // loop through the files one by one
	// 	input.Path = viper.GetString("upload_dir") + files[i].Filename
	// 	file, err := files[i].Open()
	// 	if err != nil {
	// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	defer file.Close()

	// 	out, err := os.Create(input.Path)
	// 	if err != nil {
	// 		NewErrorResponse(c, http.StatusInternalServerError, "Unable to create the file for writing. Check your write access privilege: "+err.Error())
	// 		return
	// 	}
	// 	defer out.Close()

	// 	_, err = io.Copy(out, file)

	// 	if err != nil {
	// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}

	// 	id, err := h.services.AddItem.AddFile(input)
	// 	if err != nil {
	// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, map[string]interface{}{
	// 		"id": id,
	// 	})
	// }

	// var input goweb.File
	// c.Request.ParseMultipartForm(1000)

	// mr, err := c.Request.MultipartReader()
	// if err != nil {
	// 	return
	// }
	// length := c.Request.ContentLength
	// for {

	// 	part, err := mr.NextPart()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	var read int64
	// 	var p float32
	// 	dst, err := os.OpenFile("dstfile", os.O_WRONLY|os.O_CREATE, 0644)
	// 	if err != nil {
	// 		return
	// 	}
	// 	for {
	// 		buffer := make([]byte, 100000)
	// 		cBytes, err := part.Read(buffer)
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		read = read + int64(cBytes)
	// 		//fmt.Printf("read: %v \n",read )
	// 		p = float32(read) / float32(length) * 100
	// 		fmt.Printf("progress: %v \n", p)
	// 		dst.Write(buffer[0:cBytes])
	// 	}
	// }
	// input.Path = viper.GetString("upload_dir")

	// var input goweb.File
	// file, err := c.FormFile("myFile")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// input.Path = viper.GetString("upload_dir") +file.Filename
	// err = c.SaveUploadedFile(file, input.Path)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// id, err := h.services.AddItem.AddFile(input)
	// if err != nil {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id": id,
	// })
}
