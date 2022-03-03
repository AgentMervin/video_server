CREATE table sessions( 
         session_id tinytext NOT NULL,  
         TTL tinytext, 
         login_name varchar(64),
         primary key (session_id) 
         );