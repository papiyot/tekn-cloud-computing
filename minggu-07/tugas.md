## _215611103 - Ichsan Munadi_

## Install Images Wordpress
1. buat folder wordpress
2. buat file docker-compose.yml dalam folder wordpress, lalu isi script seperti dibawah ini
```sh
version: '3.1'

services:

  wordpress:
    image: wordpress
    restart: always
    ports:
      - 8080:80
    environment:
      WORDPRESS_DB_HOST: db
      WORDPRESS_DB_USER: exampleuser
      WORDPRESS_DB_PASSWORD: examplepass
      WORDPRESS_DB_NAME: exampledb
    volumes:
      - wordpress:/var/www/html

  db:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_DATABASE: exampledb
      MYSQL_USER: exampleuser
      MYSQL_PASSWORD: examplepass
      MYSQL_RANDOM_ROOT_PASSWORD: '1'
    volumes:
      - db:/var/lib/mysql

volumes:
  wordpress:
  db:

```
3. jalankan docker-compose up -d
4. buka aplikasi lewat browser http://localhost:8080

![17](images/17.png)
