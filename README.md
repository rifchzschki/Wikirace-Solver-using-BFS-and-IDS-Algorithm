# Wikirace-Solver-using-BFS-and-IDS-Algorithm

Membuat program wikirace solver menggunakan algoritma IDS dan BFS, dengan interface berupa website

## What is BFS and IDS

Breadth-First Search (BFS) dan Iterative Deepening Search (IDS) adalah dua pendekatan yang berbeda dalam pencarian solusi dalam struktur data graf. BFS memulai pencarian dari simpul awal dan secara berurutan mengeksplorasi semua simpul tetangga pada kedalaman yang sama sebelum melanjutkan ke simpul-simpul yang lebih dalam. Dalam BFS, simpul-simpul disimpan dalam antrian dan dikunjungi dalam urutan yang sesuai. Kelebihan BFS adalah kemampuannya untuk menemukan jalur terpendek dari simpul awal ke simpul tujuan dalam graf tak berbobot, namun memerlukan banyak memori karena perlu menyimpan banyak simpul dalam antrian.

Sementara itu, IDS adalah variasi dari algoritma pencarian kedalaman terbatas (DFS) yang memperkenalkan strategi pencarian berulang dengan peningkatan kedalaman. IDS memulai dengan pencarian pada kedalaman terbatas, kemudian secara bertahap meningkatkan batasan kedalaman setiap kali pencarian tidak berhasil pada kedalaman sebelumnya. Hal ini memungkinkan IDS untuk mempertahankan keuntungan dari kedalaman pertama kali sambil mencoba membatasi kompleksitas waktu yang tinggi yang biasa terjadi pada pencarian kedalaman terdalam. IDS cocok untuk digunakan dalam situasi di mana kedalaman pencarian tidak diketahui dan terdapat batasan sumber daya seperti memori. Baik BFS maupun IDS memiliki peran penting dalam berbagai aplikasi seperti pencarian jalur terpendek, kecerdasan buatan, dan permainan strategi. Pemilihan antara keduanya tergantung pada sifat dari struktur data graf, sifat solusi yang dicari, dan batasan sumber daya yang ada.

## Prerequisite

1. [Install Golang](https://go.dev/doc/install)
2. [Install Docker](https://www.docker.com/products/docker-desktop/)

## Cara Mengkompilasi

1. Buka aplikasi Docker 


2. Masuk ke dalam direktori Backend

   ```bash
   cd src/Backend
   ```

3. Lakukan build pada direktori tersebut

   ```bash
   docker build -t stima .
   ```

4. Jalankan program

   ```bash
   docker run -p 8080:8080 stima
   ```

5. Kemudian buka browser dan buka halaman

   ```bash
   localhost:8080
   ```

## Anggota

| NIM      | NAMA                               |
| -------- | ---------------------------------- |
| 10023608 | Tazkirah Amaliah                   |
| 13522001 | Mohammad Nugraha Eka Prawira       |
| 13522120 | Muhamad Rifki Virziadeili Harisman |
