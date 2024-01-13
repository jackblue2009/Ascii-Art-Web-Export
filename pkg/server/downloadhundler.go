package server

import (
	helper "ascii-art-web/pkg/helper"
	"net/http"
)

// function to hundle any request to download the files
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename") //get the file name from url
	if !helper.CheckFileName(filename) {      // check the file name if not match to return error
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Disposition", "attachment")
	http.ServeFile(w, r, "file-request/"+filename) // serve the file to the user
}
