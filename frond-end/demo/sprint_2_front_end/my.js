const show=document.getElementById("hashcode")
const Http=new XMLHttpRequest();
var url1=decodeURI(location.href);
var argsIndex=url1.split("?obj=")
var username=argsIndex[1];
var url="http://127.0.0.1:8080/get_num";
var num1=document.getElementById("num1")
var num2=document.getElementById("num2")
const Http1=new XMLHttpRequest();
Http1.onload=function(){
console.log(Http1.responseText)
console.log(1111)
num1.innerHTML=Http1.responseText/2;
num2.innerHTML=Http1.responseText/2;
}
Http1.open("GET",url,true);
Http1.send();

console.log(1111)


console.log(2333)
url="http://127.0.0.1:8080/get_info";
var username=document.getElementById("username");
Http.onload=function(){
console.log(Http.responseText)
username.innerHTML=Http.responseText;
}


Http.open("GET",url,true);
Http.send();
const cart=document.getElementById("cart")

cart.addEventListener("click",()=>{
    location.href="gouwu.html?obj="+username;
    ;})