//var XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;
var Http=new XMLHttpRequest()
var url="http://localhost:8080/get";

// login
const elem=document.getElementById("container");
const btn1=document.getElementById("btn-login");
const btn2=document.getElementById("bnt-register");
const username=document.getElementById("username");
const password=document.getElementById("password");
btn1.addEventListener("click",()=>{
    var Http=new XMLHttpRequest()
    Http.onload=function (){
        console.log(Http.responseText)
    }
    var un=username.value;
    var ps=password.value;
    Http.open("GET","http://localhost:8080/login/"+un+"/"+ps,true);
    Http.send(un);
})



btn2.addEventListener("click",()=>{
    var Http=new XMLHttpRequest()
    Http.onload=function (){
        console.log(Http.responseText)
    }
    var un=username.value;
    var ps=username.value;
    Http.open("GET","http://localhost:8080/register/"+un+"/"+ps,true);
    Http.send()
})
