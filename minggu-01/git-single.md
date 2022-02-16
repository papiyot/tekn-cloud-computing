## _215611103 - Ichsan Munadi_
# Latihan

## A. Instalasi GIT di Windows
Sebelum install Git di Windows kita perlu mendownload [GIT](https://code.visualstudio.com/) terlebih dahulu.

1. Setelah selesai download Git, jalankan program GIT. Akan dimunculkan lisensi. Klik **Next** untuk lanjut.
![01](images/git_install/1.png)

2. Setelah itu, pilih lokasi instalasi. Secara default akan terisi C:\Program Files\Git. Ganti lokasi jika memang anda menginginkan lokasi lain, Klik **Next** untuk lanjut.
![02](images/git_install/2.png)

3. Pilih komponen. Tidak perlu diubah-ubah, sesuai dengan default saja. Klik pada **Next**.
![03](images/git_install/3.png)

4. Mengisi shortcut untuk menu Start. Gunakan default (Git), ganti jika ingin mengganti - misalnya Git VCS.
![04](images/git_install/4.png)

5. Pilih editor yang akan digunakan bersama dengan Git. Pada pilihan ini, digunakan Visual Studio Code.
![05](images/git_install/5.png)

6. Pada saat instalasi, kita bisa melakukan setting default branch saat melakukan git init, disini saya memilih main
![06](images/git_install/6.png)

7. Pada saat instalasi, Git menyediakan akses git melalui Bash maupun command prompt. Pilih pilihan kedua supaya bisa menggunakan dari dua antarmuka tersebut. Bash adalah shell di Linux. Dengan menggunakan bash di Windows, pekerjaan di command line Windows bisa dilakukan menggunakan bash - termasuk ekskusi dari Git.
![07](images/git_install/7.png)

8. Pilih **bundled OpenSSH** untuk HTTPS. Git menggunakan SSH untuk akes ke repo GitHub atau repo-repo lain (GitLab, Assembla).
![08](images/git_install/8.png)

9. Pilih **OpenSSL** untuk HTTPS. Git menggunakan https untuk akes ke repo GitHub atau repo-repo lain (GitLab, Assembla).
![09](images/git_install/9.png)

10. Pilih pilihan pertama untuk konversi akhir baris (CR-LF).
![10](images/git_install/10.png)

11. Pilih PuTTY untuk terminal yang digunakan untuk mengakses Git Bash.
![11](images/git_install/11.png)

12. Pilih Default untuk proses git pull.
![12](images/git_install/12.png)

13. Aktifkan **Git Credential Manager** pada proses ini. 
![13](images/git_install/13.png)

14. Untuk opsi ekstra, pilih serta aktifkan **Enable File System Caching**.
![14](images/git_install/14.png)

15. Setelah itu proses instalasi akan dilakukan.
![15](images/git_install/15.png)

16. Jika selesai akan muncul dialog pemberitahuan. Klik pada Finish.
![16](images/git_install/16.png)

17. Untuk menjalankan, dari Start, ketikkan **"Git"**, akan muncul beberapa pilihan. Pilih **"Git Bash"** atau **"Git GUI"**.
![17](images/git_install/17.png)

18. Tampilan jika akan menggunakan **Git Bash**
![18](images/git_install/18.png)

19. Tampilan jika akan menggunakan **Git GUI**
![19](images/git_install/19.png)

20. Untuk mencoba dari command prompt, masuk ke command prompt, setelah itu eksekusi "git --version" untuk melihat apakah sudah terinstall atau belum. Jika sudah terinstall dengan benar, makan akan muncul hasil berikut:
![20](images/git_install/20.png)

## B. Konfigurasi GIT
Untuk melakukan konfigurasi GIT kita bisa melakukan melalui Command prompt dengan memakai perintah sebagai berikut.

1. Konfigurasi Username dan Email
Isian di bawah harus disesuaikan dengan nama serta email yang digunakan untuk mendaftar di GitHub.
```sh
$ git config --global user.name "Nama Anda di GitHub"
$ git config --global user.email email@domain.tld
```

2. untuk melihat hasil konfigurasi dengan perintah 
```sh
$ git config --list
```
untuk hasilnya seperti ini
![20](images/git_konfigurasi/1.png)

## C. Mengelola Repo Sendiri di Account Sendiri
Untuk membuat repo, gunakan langkah-langkan berikut:

1. Klik tanda + pada bagian atas setelah login, pilih **New repository**
![20](images/git/1.png)

2. Isikan nama, keterangan, serta lisensi. Jika dikehendaki, bisa membuat repo Private
![20](images/git/2.png)

## C. Mengelola Repo Sendiri di Organisasi
Untuk membuat repo, gunakan langkah-langkan berikut:

1. Masuk pada Organisasi Kalian lalu Klik tanda + pada bagian atas setelah login, pilih **New repository**
![20](images/git/3.png)

2. Isikan nama, keterangan, serta lisensi. Jika dikehendaki, bisa membuat repo Private
![20](images/git/4.png)