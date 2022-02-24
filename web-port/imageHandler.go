package main

import (
	"errors"
	"net/http"
)

func imageUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	println("imageUploadFileHandler ....")
	switch r.Method {
	//POST takes the uploaded file(s) and saves it to disk.
	case http.MethodPost:
		//parse the multipart form in the request
		err := r.ParseMultipartForm(1024)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			println("error: ", err.Error())
		}
		//get a ref to the parsed multipart form
		m := r.MultipartForm
		files := m.File["file"]

		for i := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				println("error: ", err.Error())
			}
			// base64Image = "data:image/png;base64, " + data

			// []byte
			data := make([]byte, files[i].Size)
			// count, err :=
			file.Read(data)
			// if err != nil {
			// 	println("error: ", err.Error())
			// }
			//base64压缩
			// sourcestring := base64.StdEncoding.EncodeToString(data)
			//写入临时文件
			// ioutil.WriteFile("a.png.txt", data, 0667)
			// grpc []byte 传给  images
			daprGrpcClientSend("image-api-go", "/api/image", data, w)
			// daprHttpClientSend("image-api-rs", "/api/image", data, w)

			// httpClientSend(data, w)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		println("error: ", errors.New("Only POST method is supported"))
	}

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%v", "txs")
}
