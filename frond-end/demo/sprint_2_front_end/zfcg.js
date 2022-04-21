const checkout=document.getElementById("pay")
const Http=new XMLHttpRequest();
var url1=decodeURI(location.href);
var url="http://127.0.0.1:8080/order";


checkout.addEventListener("click",()=>{
    Http.onload=function(){
        alert(" Items have been successfully ordered")
    console.log(Http.responseText)
    }
    Http.open("GET",url,true);
    Http.send();
    console.log(2333)
    ;},false)