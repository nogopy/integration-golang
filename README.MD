# 1. make test.init
make test.init starts 2 container: http server and mysql server.
# 2. make test.start
make test.start runs all integration tests inside 2 above containers.
# 3. make test.end
make test.end stops, removes 2 above containers and their images
# 4. enjoy 
All you need to do is a docker.\
Everything is inside docker, it does not ffect your local machine.\
Don't worry