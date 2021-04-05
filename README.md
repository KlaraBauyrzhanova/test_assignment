# test_assignment

# About
Create a project - HTTP-service for users DB.

#Installation
Install docker-machine:
```
   curl -L https://github.com/docker/machine/releases/download/v0.13.0/docker-machine-`uname -s`-`uname -m` > /tmp/docker-machine
   # chmod +x /tmp/docker-machine
   # cp /tmp/docker-machine /usr/local/bin/docker-machine
   docker-machine --version
```
Install docker-compose:
```
  > sudo curl -L "https://github.com/docker/compose/releases/download/1.26.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  > sudo chmod +x /usr/local/bin/docker-compose
  > docker-compose version
```
  
#Usage
``` docker build -t bandlab .```
``` docker-compose up```

# Requirements
> Project has [unit-tests](https://golang.org/doc/tutorial/add-a-test) for internal functionality;
> Project well documentated using [godoc](https://blog.golang.org/godoc);

Service works as one-instance application.

# Database
Users DB scheme is presented below:
```
    CREATE TABLE `users` (
      `id` INT PRIMARY KEY,
      `data` VARCHAR,
   );
```
Field _data_ contains a JSON like
```
    {
      "first_name": "First",
      "second_name": "Last",
      "interests": "coding,golang"
    }
```
