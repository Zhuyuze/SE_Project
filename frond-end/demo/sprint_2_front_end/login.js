
      
      
const btn=document.getElementById("btn");
const input1=document.getElementById("username");
const input2=document.getElementById("password");
console.log(2323)
btn.addEventListener("click",()=>{
console.log(21312312)
var username=input1.value;
var password=input2.value;
const Http=new XMLHttpRequest();
const url="http://127.0.0.1:8080/sign_in/"+username+"/"+password;
Http.onload=function(){
    console.log(Http.responseText)
}

Http.open("GET",url,true);
Http.send();

;})