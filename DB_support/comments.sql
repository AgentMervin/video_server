CREATE table comments( 
         id varchar(64) NOT NULL,
         video_id varchar(64),
         author_id INT UNSIGNED,  
         content TEXT, 
         time DATETIME,
         primary key (id) 
         );