## _215611103 - Ichsan Munadi_

## Docker Swarm
Docker Swarm merupakan teknologi pengembangan sistem terdistribusi untuk melakukan manajemen pada kelompok mesin Docker. Dengan Docker Swarm dapat menjalankan banyak kontainer sekaligus pada kelompok mesin Docker. Pada penerapan sistem terdistribusi menggunakan Docker Swarm diperlukan sebuah penyimpanan data yang persisten.

### Kesamaan
Baik Docker Swarm dan Docker-Compose memiliki kesamaan sebagai berikut:

1. Keduanya mengambil definisi berformat YAML dari tumpukan aplikasi Anda.
2. Keduanya dimaksudkan untuk menangani aplikasi multi-kontainer (layanan mikro)
3. Keduanya memiliki parameter skala yang memungkinkan Anda menjalankan beberapa wadah dari gambar yang sama yang memungkinkan layanan mikro Anda untuk menskalakan secara horizontal.
4. Keduanya dikelola oleh perusahaan yang sama, yaitu Docker, Inc.

### Perbedaan
Beberapa perbedaan antara Docker Swarm dan Docker-Compose:

1. Docker Swarm digunakan untuk menskalakan aplikasi web Anda di satu atau beberapa server. Sedangkan Docker-compose hanya akan menjalankan aplikasi web Anda pada satu host Docker.
2. Menskalakan aplikasi web Anda Docker Swarm menawarkan ketersediaan tinggi yang serius dan toleransi kesalahan. Menskalakan aplikasi web Anda menggunakan Docker-Compose pada satu host hanya berguna untuk pengujian dan pengembangan.
3. Docker Swarm dan sub-perintah terkait seperti Docker Swarm dan Docker Stack dibangun ke dalam Docker CLI itu sendiri. Mereka semua adalah bagian dari biner Docker yang Anda panggil melalui terminal Anda. Docker-Compose adalah biner mandiri.