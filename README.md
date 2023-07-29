# KAFKA
## Review
Kafka adalah salah satu message broker. Fungsinya adalah menampung semua message yang dipublish dan membagikannya 
kepada semua service yang melakukan subscribe secara berurutan (Oleh karena itu message broker juga dapat digunakan 
sebagai queue service)

## Alur Kerja
Pada dasarnya message broker terbagi menjadi 2 proses yaitu publishing dan subscribing. Oleh karena alur kerjanya akan menjadi:
1. Service publisher menambahkan message kedalam broker dengan topic tertentu (topic disini dapat digunakan untuk mengelompokkan message pada broker)
2. Service subscriber memonitor message yang ada di dalam broker dengan topic yang dibutuhkan

## Alur Pembuatan
1. --- Publisher ---
2. Koneksikan service dengan broker sebagai publisher
3. Kirim message dengan topic yang diperlukan
4. --- Subscriber ---
5. Koneksikan service dengan broker sebagai consumer
6. Lakukan monitor pada broker dengan parameter topic yang dibutuhkan
7. Ketika terjadi penambahan message lakukan sesuatu jika diperlukan
