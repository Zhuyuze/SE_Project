### http://127.0.0.1/sign_up  
example: http://127.0.0.1/sign_up/username/password  
username (string) : username from the input field  
password (string) : passwrod from the input field  
return value (string) :   
username already exists &#8195;         there is a same username in database  
success    &#8195; &#8195;   &#8195; &#8195;  &#8195;     &#8195;         username and password successfully inserted to database  


### http://127.0.0.1/sign_in  
example: http://127.0.0.1/sign_in/username/password   
username (string) : username from the input field  
password (string) : passwrod from the input field  
return value (string) :   
Wrong username or password   &#8195;   there is no (username, password) pair with corresponding username  
Wrong password     &#8195;  &#8195; &#8195;      &#8195;    password doesn't match       
Success   &#8195; &#8195;  &#8195;     &#8195;   &#8195;  &#8195;         (username,password) match  


### http://127.0.0.1/order  
example: http://127.0.0.1/order/username/food/quantity  
username (string): username of the user  
food (string): food name   
quantity (int): quantity of food   
return value(string):  
food doesn't exist          &#8195; &#8195;&#8195;&#8195;   the input food doesn't exist  
food quantity not enough     &#8195;   food stock not enough  
Success     &#8195;  &#8195;   &#8195;    &#8195;   &#8195;     &#8195; &#8195;  food successfully ordered  
