

const btn=document.getElementById("btn");   

    btn.addEventListener("click",()=>{
        const input1=document.getElementById("username");
        const input2=document.getElementById("password");
        console.log(21312312)
        var username=input1.value;
        var password=input2.value;
        const Http=new XMLHttpRequest();
        if(username==""|| password=="")
        {
            alert("Please input username AND password")
        }
        console.log(username)
        console.log(password)
        const url="http://127.0.0.1:8080/sign_in/"+username+"/"+password;
        Http.onload=function(){
            console.log(Http.responseText)
            if(Http.responseText=="Hello Success")
            {
                location.href="my.html?obj="+username
                mypage("Roman");
            }
            else{
                alert("Wrong username or password")
            }
            
        }
        
        Http.open("GET",url,true);
        Http.send();
        ;})


