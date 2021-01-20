##Local Database Connection Test
```
go test -v

=== RUN   TestUserWorkFlow
=== RUN   TestUserWorkFlow/Add
=== RUN   TestUserWorkFlow/Get
=== RUN   TestUserWorkFlow/Del
=== RUN   TestUserWorkFlow/Reget
--- PASS: TestUserWorkFlow (0.00s)
    --- PASS: TestUserWorkFlow/Add (0.00s)
    --- PASS: TestUserWorkFlow/Get (0.00s)
    --- PASS: TestUserWorkFlow/Del (0.00s)
    --- PASS: TestUserWorkFlow/Reget (0.00s)
=== RUN   TestVideoWorkFlow
=== RUN   TestVideoWorkFlow/PrepareUser
=== RUN   TestVideoWorkFlow/AddVideo
=== RUN   TestVideoWorkFlow/GetVideo
=== RUN   TestVideoWorkFlow/DelVideo
=== RUN   TestVideoWorkFlow/RegetVideo
--- PASS: TestVideoWorkFlow (0.03s)
    --- PASS: TestVideoWorkFlow/PrepareUser (0.00s)
    --- PASS: TestVideoWorkFlow/AddVideo (0.00s)
    --- PASS: TestVideoWorkFlow/GetVideo (0.00s)
    --- PASS: TestVideoWorkFlow/DelVideo (0.00s)
    --- PASS: TestVideoWorkFlow/RegetVideo (0.00s)
=== RUN   TestComments
=== RUN   TestComments/AddUser
=== RUN   TestComments/AddCommnets
=== RUN   TestComments/ListComments
--- PASS: TestComments (0.04s)
    --- PASS: TestComments/AddUser (0.00s)
    --- PASS: TestComments/AddCommnets (0.00s)
    --- PASS: TestComments/ListComments (0.00s)
PASS
ok      video_server/api/dbops  0.312s

```

