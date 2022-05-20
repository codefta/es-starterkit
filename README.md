# ElasticSearch Starter Kit

Project ini berisi starter kit untuk memulai pembelajaran tentang ElasticSearch di Ghazlabs. Project ini ditujukan bagi teman-teman yang memang tidak memiliki pengetahuan sama sekali tentang ElasticSearch.

Seluruh dokumen yang ada di project ini akan ditulis dalam Bahasa Indonesia. Hal ini ditujukan supaya teman-teman lebih mudah untuk memahami konsep-konsep dasar yang diberikan di project ini.

## Objektif

Starter kit ini dibuat dengan objektif sebagai berikut:

- Menunjukkan bagaimana menjalankan elasticsearch dengan menggunakan docker compose
- Menunjukkan bagaimana melakukan inisialisasi dan penambahan data awal pada elasticsearch dengan menggunakan docker compose
- Menunjukkan bagaimana menggunakan elasticsearch di Go
- Menunjukkan bagaimana menjalankan aplikasi server dengan menggunakan docker compose
- Menunjukkan bagaimana menghubungkan antara satu aplikasi dan aplikasi lainnya dengan docker compose, dalam hal ini kita akan menghubungkan aplikasi server & elasticsearch
- Menunjukkan bagaimana melakukan Go development dengan menggunakan hot reload di docker compose

Wah sangat banyak ya?

Yap, cuma memang setiap objektif ini adalah pengetahuan dasar bagi teman-teman untuk bisa berkontribusi dengan aktif di Ghazlabs. 😃

[Back to Top](#elasticsearch-starter-kit)

---

## Cara Menjalankan

Project ini dibangun diatas Docker, Docker Compose, Makefile, & Go. Jadi sebelum teman-teman memulai untuk menjalankan project ini, pastikan seluruh teknologi tersebut sudah terinstall dengan baik di komputer teman-teman.

Untuk menjalankan project ini cukup ketik perintah berikut di terminal:

```bash
> make run
```

Perintah ini akan menjalankan 2 service di komputer teman-teman:

- `ElasticSearch` => Software yang dioptimasi untuk melakukan pencarian dokumen secara umum (a.k.a Search Engine). Software ini akan dijalankan pada `http://localhost:9202`.
- `JajananAPI` => API server yang memungkinkan kita mencari, menambahkan, dan menghapus data jajanan (street food) yang ada di Indonesia. Server ini akan dijalankan pada `http://localhost:8101`.

> **Note:**
>
> Starter kit ini didesain untuk tidak menyimpan data secara permanen, jadi data-data baru yang dibuat oleh teman-teman akan terhapus ketika teman-teman menghentikan kit ini. Ketika teman-teman menjalankan kembali kit ini, data-data yang ada akan di-_reset_ seperti semula.
>
> Reset data ini dilakukan agar setiap kali starter kit ini kembali dijalankan, data-nya selalu dalam keadaan bersih (clean slate). Seringkali data yang tidak bersih bisa menimbulkan banyak masalah ketika kita melakukan development. Karena itulah biasanya untuk development kita selalu memulai dengan data yang bersih.

[Back to Top](#elasticsearch-starter-kit)

---