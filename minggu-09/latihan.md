## _215611103 - Ichsan Munadi_

## Clone the Lab’s GitHub Repo
1. git clone https://github.com/dockersamples/linux_tweet_app

![1](images/1.png)

## Run a single task in an Alpine Linux container
1. cek hostname dengan ketik hostname
2. docker container run alpine hostname
3. docker container ls --all

![2](images/2.png)

## Run an interactive Ubuntu container
1. docker container run --interactive --tty --rm ubuntu bash
2. ls /
3. ps aux
4. cat /etc/issue
5. exit
6. cat /etc/issue

![3](images/3.png)

## Run a background MySQL container
1. masukkan perintah ini pada CLI
```sh
docker container run \
 --detach \
 --name mydb \
 -e MYSQL_ROOT_PASSWORD=my-secret-pw \
 mysql:latest
```
2. docker container ls
3. docker container logs mydb
4.  docker container top mydb
5. docker exec -it mydb \
 mysql --user=root --password=$MYSQL_ROOT_PASSWORD --version

![4](images/4.png)

## Build a simple website image
1.  cd ~/linux_tweet_app
2. cat Dockerfile
3. export DOCKERID=your docker id
4. echo $DOCKERID
5. docker image build --tag $DOCKERID/linux_tweet_app:1.0 .
6. masukkan perintah ini pada CLI
```sh
docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```
7. docker container rm --force linux_tweet_app

![5](images/5.png)
![5](images/5a.png)

## Update the image
1. docker image build --tag $DOCKERID/linux_tweet_app:2.0 .
2.  docker image ls

![6](images/6.png)

## Test the new version
1. masukkan perintah ini pada CLI
```sh
docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:2.0
```

![7](images/7.png)
![7](images/7a.png)

## Push your images to Docker Hub
1. docker image ls -f reference="$DOCKERID/*"
2. docker login
3. docker image push $DOCKERID/linux_tweet_app:1.0
4. docker image push $DOCKERID/linux_tweet_app:2.0

![8](images/8.png)
![8](images/8a.png)