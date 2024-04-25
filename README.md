<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/RickyIG/quiz-3-golang">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center"Home Cinema API</h3>

  <p align="center">
    Quiz 3 Golang
    <br />
    <a href="https://github.com/RickyIG/home-cinema"><strong>Explore the project »</strong></a>
    <br />
    <br />
    <a href="https://quiz-3-golang-production.up.railway.app/categories/1/books">View Demo</a>
    ·
    <a href="https://github.com/RickyIG/home-cinema/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href=https://github.com/RickyIG/home-cinema/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Project backend yang menggunakan Golang, Gin Framework, dan Railway deployment bertujuan untuk mengembangkan serangkaian API yang memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) terhadap dua jenis sumber daya utama: bangun datar dan kategori buku beserta buku itu sendiri. Dalam konteks ini, backend bertindak sebagai jembatan antara klien (seperti aplikasi web atau perangkat lunak lainnya) dan basis data yang berisi informasi tentang bangun datar dan daftar kategori serta buku yang berkaitan.

Pertama-tama, API yang berfokus pada bangun datar menyediakan endpoint untuk menghitung berbagai properti dan dimensi dari bangun datar seperti persegi, persegi panjang, segitiga, dan lainnya. Pengguna dapat mengirimkan permintaan dengan parameter yang sesuai untuk memperoleh hasil perhitungan yang diinginkan dalam format yang sesuai seperti JSON.

Selanjutnya, API yang terkait dengan kategori dan buku memungkinkan pengguna untuk melakukan operasi CRUD terhadap daftar kategori dan buku yang terkait dengannya. Pengguna dapat membuat, mengambil, memperbarui, dan menghapus kategori buku serta buku individu dalam kategori tertentu melalui endpoint yang disediakan. Informasi seperti judul, penulis, tahun terbit, dan kategori dapat dimasukkan dan diminta dalam permintaan dan respons API.

Proyek ini menggunakan Gin Framework untuk mempermudah pengembangan API dengan dukungan routing, middleware, dan fitur-fitur lainnya yang diperlukan untuk membangun aplikasi web yang andal dan efisien. Selain itu, deployment dilakukan menggunakan Railway, memungkinkan proyek untuk di-host dengan mudah dan dikelola dengan baik di lingkungan cloud yang dapat diskalakan sesuai kebutuhan. Dengan demikian, proyek ini menawarkan solusi yang kokoh dan skalabel untuk kebutuhan backend dalam pengembangan aplikasi berbasis web yang berkaitan dengan pengelolaan bangun datar dan katalog buku.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

Dibangun dengan:
[![Go][Golang]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Go 1.22.1
* Postman

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo (if needed)
   ```sh
   git clone https://github.com/RickyIG/home-cinema.git
   ```
2. Buka Postman
3. Lakukan pemanggilan Rest API via Postman

Berikut link yang telah di-deploy:
(https://home-cinema-production.up.railway.app/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

### User Roles
- **Role ID 1:** Customer
- **Role ID 2:** Admin
- **Role ID 3:** Fintech


### User

#### POST Register User
- Endpoint: `https://home-cinema-production.up.railway.app/register/user`
- Notes:
  - Panjang password harus minimal panjangnya 8 character, harus memuat setidaknya 1 special character, 1 lowercase, 1 uppercase, dan 1 angka
  - Email harus valid
- Contoh Input Raw Body (JSON):
  ```json
  {
    "username": "Test12345#",
    "password": "Test12345#",
    "first_name": "Test",
    "last_name" : "Test",
    "email": "test@gmail.com",
    "phone_number": "082839128762"
  }


#### POST Register User
- Endpoint: `https://home-cinema-production.up.railway.app/register/user`
- Contoh Input Raw Body (JSON):
  ```json
  {
      "username": "rickyindrag",
      "password": "Kocak@12345"
  }

#### POST Register User
- Endpoint: `https://home-cinema-production.up.railway.app/register/user`
- Contoh Input Raw Body (JSON):
  ```json
  {
      "username": "Fintech@12345",
      "password": "Fintech@12345"
  }

#### POST Login
- Endpoint: `https://home-cinema-production.up.railway.app/login`
- Contoh Input Raw Body (JSON):
  ```json
  {
      "username": "rickyindrag",
      "password": "Kocak@12345"
  }


https://quiz-3-golang-production.up.railway.app/bangun-datar/segitiga-sama-sisi?hitung=luas&alas=7&tinggi=4

https://quiz-3-golang-production.up.railway.app/bangun-datar/segitiga-sama-sisi?hitung=keliling&alas=7&tinggi=4


https://quiz-3-golang-production.up.railway.app/bangun-datar/persegi?hitung=luas&sisi=7

https://quiz-3-golang-production.up.railway.app/bangun-datar/persegi?hitung=keliling&sisi=7


https://quiz-3-golang-production.up.railway.app/bangun-datar/persegi-panjang?hitung=luas&panjang=7&lebar=4

https://quiz-3-golang-production.up.railway.app/bangun-datar/persegi-panjang?hitung=keliling&panjang=7&lebar=4


https://quiz-3-golang-production.up.railway.app/bangun-datar/lingkaran?hitung=luas&jariJari=7

https://quiz-3-golang-production.up.railway.app/bangun-datar/persegi-panjang?hitung=keliling&panjang=7&lebar=4


Perlu diingat bahwa hitung itu nilainya antara luas / keliling.

Perhatikan juga parameter-parameternya.


https://quiz-3-golang-production.up.railway.app/categories -> GET -> Mendapatkan/menampilkan semua categories

https://quiz-3-golang-production.up.railway.app/categories -> POST -> Menambah categories

https://quiz-3-golang-production.up.railway.app/categories/:id -> PUT -> Mengedit categories

https://quiz-3-golang-production.up.railway.app/categories/:id -> DELETE -> Menghapus categories


https://quiz-3-golang-production.up.railway.app/books -> GET -> Mendapatkan/menampilkan semua buku

https://quiz-3-golang-production.up.railway.app/books -> POST -> Menambah buku

https://quiz-3-golang-production.up.railway.app/books/:id -> PUT -> Mengedit buku

https://quiz-3-golang-production.up.railway.app/books/:id -> DELETE -> Menghapus buku



https://quiz-3-golang-production.up.railway.app/categories/:id/books -> GET -> Melihat semua buku berdasarkan category ID

Perlu diingat bahwa id dalam tabel categories dan category_id dalam tabel books itu berhubungan (FOREIGN KEY).

Untuk POST, PUT, DELETE, perlu auth 

username: admin

password: password


username: editor

password: secret


Contoh raw body untuk insert category
```
{
     "name": "Sastra Indonesia"
}
```
Contoh raw body untuk insert buku

```
{
    "title": "Laskar Pelangi",
    "description": "Novel inspiratif tentang persahabatan dan semangat meraih mimpi di tengah keterbatasan.",
    "image_url": "https://upload.wikimedia.org/wikipedia/id/1/17/Laskar_Pelangi_film.jpg",
    "release_year": 2008,
    "price": "Rp 50.000",
    "total_page": 384,
    "category_id": 1
}
```


Note :

Keadaan saat ini:

- buku dengan id 1 baru saja diedit
- buku dengan id 2 baru saja dihapus
- unauthorized user cuma bisa get



<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Setup project
- [x] Pembuatan API
  - [x] User
    - [x] Auth
      - [x] Register, Login (+ JWT, Role-Based)
    - [x] See Profile
    - [x] Put Balance oleh Fintech
  - [x] CRUD
    - [x] Film
    - [x] Studio
    - [x] Jadwal
  - [x] Special Case di luar CRUD
    - [x] Kursi
      - [x] Tersedia/Tidak Tersedia
      - [x] Auto-Add dari Jadwal 
  - [x] Pesan Ticket (Melakukan Transaksi)
  - [x] Transaksi
    - [x] GetUserTransactionHistory
    - [x] GetUserTransactionHistoryByID
    - [x] GetUserTransactionHistoryDetail
  - [x] Ticket
    - [x] GetUserTicketHistory
    - [x] GetUserTicketHistoryByTransactionID
    - [x] GetUserTicketHistoryByIDs (by transaction ID + ticket ID)
- [x] Deploy to Railway

See the [open issues](https://github.com/RickyIG/quiz-3-golang/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Apache 2.0 License.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Ricky Indra Gunawan - Instagram: [@rickyindrag](https://instagram.com/rickyindrag) - rickyindra53@gmail.com

Project Link: [https://github.com/RickyIG/home-cinema](https://github.com/RickyIG/home-cinema)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/rickyindrag
[product-screenshot]: images/screenshot.png
[Golang]: https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png
[Go-url]: https://go.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
