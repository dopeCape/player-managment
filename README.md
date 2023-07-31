Follow the commands below to run the docker Command

Its Required to have docker installed and docker deamon running on the computer(linux) or docker service registed (windows)
following command might take some time depending on your system and network speed

Its Important that you have your lh:8000 port and lh:27017 port free for server and db to start, if they are not please free them up and try again

<h1>Linux(with make)</h1>

```
git clone https://github.com/dopeCape/player-managment
cd player-managment
sudo make docker
```

<h1>Linux(without make)</h1>

```
git clone https://github.com/dopeCape/player-managment
cd player-managment
sudo docker build -t player-managment . && sudo docker run -p 27017:27017 -p 8000:8000 player-managment
```

<h1>Windows(with make)</h1>

```
git clone https://github.com/dopeCape/player-managment
cd player-managment
make docker
```

<h1>Windows(without make)</h1>

```
git clone https://github.com/dopeCape/player-managment
cd player-managment
 docker build -t player-managment . &&  docker run -p 27017:27017 -p 8000:8000 player-managment
```

\*\* To run on windows make sure the cmd session has administrative privilege.This might not be required , try running without administrative privilege ,
if it fails try running it with administrative privileges.

<h1>API docs</h1>

1.  POST : "http://localhost:8000/players"
    will Return the ID for the player entry
    Payload ex:

```json
    {
    country:"IN",
    score:80,(should be a positive Int)
    name:"tejas"
    }


```

2.  PUT : "http://localhost:8000/players/:id"
    will return the updated Player entry
    Payload ex:

    ```json
    {
    score:90,(should be a positive Int)
    name:"tejas"
    }
    ```

3.  DELETE : "http://localhost:8000/players/:id"
    will not return any Thing

4.  GET : "http://localhost:8000/players/rank/:val"
    will return the player at vat'th rank

5.  GET : "http://localhost:8000/players/"
    will return all the players in descending order

6.  GET : "http://localhost:8000/players/random"
    will return a random player
