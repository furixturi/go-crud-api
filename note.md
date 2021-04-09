This is not working
```bash
$ docker run -it -p 5432:5432 
```
```
Error: Database is uninitialized and superuser password is not specified.
       You must specify POSTGRES_PASSWORD to a non-empty value for the
       superuser. For example, "-e POSTGRES_PASSWORD=password" on "docker run".

       You may also use "POSTGRES_HOST_AUTH_METHOD=trust" to allow all
       connections without a password. This is *not* recommended.

       See PostgreSQL documentation about "trust":
       https://www.postgresql.org/docs/current/auth-trust.html
```
Use this
```bash
$ docker run -it --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=password
```