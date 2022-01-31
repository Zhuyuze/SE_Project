var XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;
var Http=new XMLHttpRequest()
var url="http://localhost:8080/get";
console.log(123)
Http.onload=function (){
    console.log(Http.responseText)
}

Http.open("get","http://localhost:8080/haha",true);
Http.send();
console.log(234)
console.log(23333)