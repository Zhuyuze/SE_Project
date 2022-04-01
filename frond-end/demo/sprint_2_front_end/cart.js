var url1=decodeURI(location.href);
var argsIndex=url1.split("?obj=")
var username=argsIndex[1];
const url="http://127.0.0.1:8080/get_num/"+username;
console.log(username)



const checkout = document.getElementById("checkout")

checkout.addEventListener("click",()=>{
    const url="http://127.0.0.1:8080/order_food/"+username;
    Http.onload=function(){
    console.log(Http.responseText)
    }
    Http.open("GET",url,true);
    Http.send();
    console.log(2333)
    ;})
