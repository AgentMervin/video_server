
# Share+: video-sharing platform

A streaming video website for internal use. User can watch/upload/download videos from the website. Comments are also permitted if the owner(s) of the video agrees.

## Features
- Very simple to install and use;
- Pure Golang, high performance, and cross-platform;
- Supports commonly used transmission protocols, file formats, and encoding formats;



## Architecture
![image](./asset/Design.jpeg)

## API Design
Using HTTP Protocol to fulfill the operation on resource.
Three types of APIs:
- USER API: return states of each user.
- RESOURCE API: returns the states of video(s).
- COMMENT API: returns all the comments under one specific video/

#### Streaming Server Design
Prerequisite:
[bitbucket](https://godoc.org/github.com/DavidCai1993/token-bucket)

- UDP protocol to implement file uploading
- Token Bucket to control rate limit

#### Scheduler Design
[**Channels are the pipes that connect concurrent goroutines.**](https://tour.golang.org/concurrency/2)

- Asynchronous Delete
- Producer-Consumer Model
- Timer: Run and Stop


## Future goals

Cloud native Optimization

