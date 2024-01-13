package server

import (
	ascii "ascii-art-web/pkg/ascii-art"
	helper "ascii-art-web/pkg/helper"
	"fmt"
	"html/template"
	"net/http"
)

type HomeTmpl struct {
	AsciiArtDraw string
	Color        string
	Align        string
	Background   string
	Downloadtxt  template.HTML
	Download2    template.HTML
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	index := "templates/index.html"
	tmpl, err := template.ParseFiles(index) // make template for the main page
	if err != nil {                         // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	data := &HomeTmpl{AsciiArtDraw: "", Color: "#000000", Align: "left", Background: "#FFFFFF", Downloadtxt: template.HTML(""), Download2: template.HTML("")}
	var text string
	if r.Method == "POST" {
		text = r.PostFormValue("text")      // get the text user put in the web page
		bytes := []byte(text)               // convert the text to bytes to avoid problems form <>&
		banner := r.PostFormValue("banner") // get the banner name from the web page
		filetype := r.PostFormValue("file-type")
		data.Color = r.PostFormValue("color") // result for the color givin
		data.Align = r.PostFormValue("align") // get the selected alignment
		if !helper.CheckBannerName(banner) || !helper.CheckFileExtension(filetype) {
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			return
		}
		k, err := ascii.AsciiArt(bytes, banner)
		if err != nil { // any bad request from user will return error 400
			http.Redirect(w, r, "/400", http.StatusSeeOther)
			return
		}
		data.AsciiArtDraw = k                          //result for the text givin
		if helper.ConvertHexToInt(data.Color) >= 128 { // check if the color is dark or light to select the background color
			data.Background = "#000000"
		}
		err = helper.Createfile("txt", data.AsciiArtDraw)
		if err != nil {
			http.Redirect(w, r, "/400", http.StatusSeeOther)
			return
		}
		txtfile := "file.txt"
		data.Downloadtxt = template.HTML("<a onclick=\"alarmButton()\" class=\"download-btn\" href=\"../download?filename=" + txtfile + "\">Download " + txtfile + "</a>")
		if filetype != "" {
			err = helper.Createfile(filetype, data.AsciiArtDraw)
			if err != nil {
				http.Redirect(w, r, "/400", http.StatusSeeOther)
				return
			}
			filename := "file." + filetype
			data.Download2 = template.HTML("<a onclick=\"alarmButton()\" class=\"download-btn\" href=\"../download?filename=" + filename + "\">Download " + filename + "</a>")
		}
	}

	size := len(data.AsciiArtDraw)
	err1 := tmpl.Execute(w, data) // execute the template wiht the result
	if err1 != nil {              // any error in the template will return error 500
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(size))
	w.WriteHeader(http.StatusOK)
}
