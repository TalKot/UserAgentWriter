function reqListener () {
    console.log(this.responseText);
  }
  
window.onload = function() {
    var oReq = new XMLHttpRequest();
    oReq.addEventListener("load", reqListener);
    oReq.open("GET", "http://localhost:8000/load");
    oReq.setRequestHeader("Referer", "http://valid.referer.com");
    oReq.send();
}