# ASCII-ART-WEB

A web server for outputting the text in the graphical representation of ASCII and downloading a file with it.

## AUTHORS

* Hussain
* Abdulrahman

## USAGE

### Local Run:
Normal run in local host we use:
```
go run main.go
```

### Docker Run:
For run the web application in container using docker will be as follows:
1. Create an image
```
docker image build -f Dockerfile -t <name_of_the_image> .
```
example:
```
docker image build -f Dockerfile -t asciiart .
```
2. Run the image as a container
```
docker container run -p <port_you_what_to_run> --detach --name <name_of_the_container> <name_of_the_image>
```
example:
```
docker container run -p 8080:8080 --detach --name webcontainer asciiart
```
3. Docker exec to check application contains
```
docker exec -it <container_name> /bin/sh
```
example:
```
docker exec -it webcontainer /bin/sh

/app # ls -l
```

This command shall start the web server upon execution on the specified port 80. You can afterwards access the web server at 'http://localhost:8080'.

To generate ASCII art, enter the text you want to convert in the text field, choose the type of ASCII art, and then click 'Generate' button. The ASCII art will be displayed in the ASCII art text area field. You can also download the ASCII art as a txt file by clicking the 'Download' button.

## OPTIONS

You can specify the following options when generating ASCII art:

### Banner Type
* 'standard': The standard default ascii art type.
* 'shadow': The shadowed/ darkened ascii art type.
* 'thinkertoy': The robotic ascii art type.
* '3d': The addtional ascii art we add.

### Text Alignment
* 'Left': Align text to left of text area.
* 'Center': Align text to center of text area.
* 'Right': Align text to right of text area.

### Other options
* Select color of text
* Increase/Decrease the font size
* Put shadow on the text

## ALGORITHMS USED

We used *SEARCHING ALGORITHMS* to search for the ascii art characters need to be printed in the website upon the user requeste on the POST method.

The following functions are implemented in the backend server side to process Ascii Art graphical text through POST method:

* 'GetAscii': It runs CheckString() and returns a string of lines specified per character ascii code from banner text file.
* 'CheckString': It returns false if a character of string parameter is less than 32 or greater than 126, otherwise returns true.
* 'StoreAsciiArt': It  returns a string of a generated ascii art word of string passed as parameter.

## LICENSE

This project is licensed under the Reboot License.