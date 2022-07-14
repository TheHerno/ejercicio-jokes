# test-project-hernan

Endpoint for getting 25 random jokes from chucknorris api using concurrency to optimize the request.
http://localhost:9009/v1/client/jokes

## Development containers

In the project there is a docker-compose file used to lift the entire work environment without installing frames or tools used by the project

-   Install Docker `sudo apt-get install docker-compose`

-   Create the `.env` file based on `.env.example`.


```

-   Run docker orchestra

```bash
$ docker-compose up -d
```

-   Check service up

```bash
$ curl http://localhost:9009/ping
```

