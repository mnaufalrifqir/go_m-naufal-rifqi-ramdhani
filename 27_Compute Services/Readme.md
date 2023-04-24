# (27) Compute Services
- Langkah-langkah untuk deployment with EC2 
    1. Membuat VM di EC2 pada laman AWS dengan settingan seperti pada folder screenshoots
    2. Setelah membuat instance kita dapat melakukan koneksi ke instance tersebut melalui ssh dengan syntax: ssh -i "public_key" username@Public_IPv4_DNS
    3. Karena disini saya menggunakan docker, git, docker-compose, dan nginx, maka kita harus menginstall terlebih dahulu.
    4. Kemudian clone repository project yang akan di deploy
    5. Ubah config nginx pada nginx.conf sesuaikan dengan port yang digunakan
    6. Setelah itu kita dapat menstart docker dengan syntax: service docker start
    7. Kemudian pindahkan directory kita ke directory project yang sudah di clone, contoh: cd aws_deploy
    8. Run file docker-compose dengan syntax: docker-compose up -d
    9. Kemudian start nginx nya dengan syntax: systemctl start nginx
    10. Kita dapat langsung membuka program menggunakan public IP, contohnya public IP saya yaitu: http://3.26.178.219/