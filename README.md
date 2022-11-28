# Pertanyaan
Tulis dan lampirkan kode untuk permasalahan berikut. Anda boleh menggunakan bahasa apapun yang Anda familiar.

Tulis fungsi sebagai berikut:

Input: Sebuah bilangan n

Output: Array berisi seluruh angka dari 1..n yang merupakan bilangan prima, dan juga merupakan palindrom (Palindromic Prime)

Contoh:

> palindromic_prime(200)
[2, 3, 5, 7, 11, 101, 131, 151, 181, 191]

Berikan juga penjelasan tertulis mengapa kode Anda memecahkan masalah di atas dengan benar.

# Jawaban Kode Online
Buka link berikut:
[Disini](https://go.dev/play/p/SDPXzpOhCOX)

# Penjelasan
Diketahui Angka Palindrom adalah angka yang dibaca sama dari belakang ke depan.

Bilangan Prima adalah bilangan yang punya dua faktor, yaitu 1 dan bilangan itu sendiri.

Palindrom Prime sendiri adalah angka bilangan prima yang terdapat palindrom

Secara sederhana, saya membaginya menjadi 2 logic:
1. Hasilkan bilangan prima ke n
2. Cari angka palindrom

Cara tersebut mungkin tidak cepat, tetapi efektif untuk menghasilkan jawaban yang benar.