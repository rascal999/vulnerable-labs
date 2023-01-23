## Vulnerable Lab

## Run with docker

```
$ docker run --rm -p 80:80 --name php_rce -v "$PWD":/var/www/html php:5.5-apache
```

## Expected result

After the application starts, navigate to `http://localhost/?pattern=/test/&replace=TEST&subject=This+is+a+test` in your web browser.
