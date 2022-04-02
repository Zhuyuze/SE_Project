const show=document.getElementById("hashcode")
const Http=new XMLHttpRequest();
var url1=decodeURI(location.href);
var argsIndex=url1.split("?obj=")
var username=argsIndex[1];
const url="http://127.0.0.1:8080/get_num/"+username;
Http.onload=function(){
console.log(Http.responseText)
}
Http.open("GET",url,true);
Http.send();
console.log(2333)


const cart=document.getElementById("cart")

cart.addEventListener("click",()=>{
    location.href="gouwu.html?obj="+username;
    
    ;})