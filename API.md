### sign_up  
example: password  
username (string) : username from the input field  
password (string) : passwrod from the input field  
return value (string) :   
username already exists &#8195;         there is a same username in database  
success    &#8195; &#8195;   &#8195; &#8195;  &#8195;     &#8195;         username and password successfully inserted to database  


### sign_in  
example: username/password   
username (string) : username from the input field  
password (string) : passwrod from the input field  
return value (string) :   
Wrong username or password   &#8195;   there is no (username, password) pair with corresponding username  
Wrong password     &#8195;  &#8195; &#8195;      &#8195;    password doesn't match       
Success   &#8195; &#8195;  &#8195;     &#8195;   &#8195;  &#8195;         (username,password) match  


### order  
example: /food/quantity  
username (string): username of the user  
food (string): food name   
quantity (int): quantity of food   
return value(string):  
food doesn't exist          &#8195; &#8195;&#8195;&#8195;   the input food doesn't exist  
food quantity not enough     &#8195;   food stock not enough  
Success     &#8195;  &#8195;   &#8195;    &#8195;   &#8195;     &#8195; &#8195;  food successfully ordered  


### prev_order  
example: /username\
username (string): username of the user  \
return value(string):  
(string,int)        &#8195; &#8195;&#8195;&#8195;    (food_name,count)

### user_info 
example: /username\
username (string): username of the user  \
return value(string):  
(string,string,string)        &#8195; &#8195;&#8195;&#8195;    (username,usertype,join_date)
