// create two buttons for increase and decrease the text size
var textarea = document.getElementById("textarea");
var increaseSizeButton = document.getElementById("increase-size");
var decreaseSizeButton = document.getElementById("decrease-size");

// defult size is 10px
var currentFontSize = 10;

// function to increase the font size
increaseSizeButton.onclick = function () {
    currentFontSize += 2;
    if (currentFontSize >= 70) { // set maxFontSize to 70
        currentFontSize = 70
    }
    textarea.style.fontSize = currentFontSize + "px";
};

// function to decrease the font size
decreaseSizeButton.onclick = function () {
    currentFontSize -= 2;
    if (currentFontSize <= 4) { // set minFontSize to 6
        currentFontSize = 4
    }
    textarea.style.fontSize = currentFontSize + "px";
};

var shadow = document.getElementById("shadow");

shadow.onclick = function () {
    if (textarea.style.textShadow == "") {
        textarea.style.textShadow = "2px 2px #FF0000";
    } else {
        textarea.style.textShadow = "";
    }
};

function getAvargeHex(hexColor) {
    // segmnt the hex for each color
    var r = hexColor.substring(1, 3);
    var g = hexColor.substring(3, 5);
    var b = hexColor.substring(5, 7);
    // ge thte number for each color
    var rn = parseInt(r, 16);
    var gn = parseInt(g, 16);
    var bn = parseInt(b, 16);
    //calculate the avarge
    var avarge = (rn + gn + bn) / 3;
    return avarge;
}

// identify the where we will get color and align
var selectcolor = document.getElementById("color")
var selectalign = document.getElementById("align")

// function to change the color
selectcolor.addEventListener("input", function () {
    var colorV = selectcolor.value;
    textarea.style.color = colorV;
    var avarge = getAvargeHex(colorV)
    if (avarge >= 128) {
        textarea.style.backgroundColor = "#000000";
    } else {
        textarea.style.backgroundColor = "#FFFFFF";
    } 
});

// function to change the align
selectalign.addEventListener("input", function () {
    var alignV = selectalign.value;
    textarea.style.textAlign = alignV;
});

// set alarm massage to explain no font and color control to file you want to download
function alarmButton() {
    alert("The text in the file comes with standard color and font size.");
};