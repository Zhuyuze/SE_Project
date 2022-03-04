const btn1=document.getElementById("btn_register")
btn1.addEventListener("click",()=>{
    const input1=document.getElementById("username");
    const input2=document.getElementById("password1");
    const input3=document.getElementById("password2");
    var username=input1.value;
    var password1=input2.value;
    var password2=input3.value;
    console.log(username)
    if(username=="" || password1=="" || password2=="")
    {
        alert("Please input all username and password")
    }
    if(password1 != password2)
    {
        alert("Password doesn't match")
    }
    else
    {      
     const Http=new XMLHttpRequest();
    const url="http://127.0.0.1:8080/sign_up/"+username+"/"+password1;
    Http.onload=function(){
        console.log(Http.responseText)
    }
    
    Http.open("GET",url,true);
    Http.send();}
    ;})