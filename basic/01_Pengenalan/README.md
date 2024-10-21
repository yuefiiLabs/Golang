# Pengenalan apa itu golang?

Golang, atau Go, adalah bahasa pemrograman yang dikembangkan oleh Google dan dirilis pada tahun 2009. Bahasa ini dirancang untuk memudahkan pengembangan perangkat lunak yang efisien, cepat, dan dapat diandalkan. Beberapa fitur utama dari Go termasuk:

1. **Sintaksis Sederhana**:  
   Go memiliki sintaksis yang bersih dan mudah dipahami, sehingga memudahkan pemula dan mempercepat proses pengembangan.

2. **Kinerja Tinggi**:  
   Go adalah bahasa kompilasi, yang berarti kode yang ditulis dalam Go diubah menjadi kode mesin yang berjalan langsung di sistem operasi, menghasilkan kinerja yang sangat baik.

3. **Kebangkitan Paralelisme**:  
   Dengan fitur goroutine dan channel, Go memudahkan pengembangan aplikasi yang mampu menjalankan banyak proses secara bersamaan, cocok untuk aplikasi web dan sistem terdistribusi.

4. **Garbage Collection**:  
   Go memiliki sistem manajemen memori otomatis, yang membantu mencegah kebocoran memori dan meningkatkan efisiensi.

5. **Ekosistem yang Kuat**:  
   Go memiliki pustaka standar yang luas dan komunitas yang aktif, mendukung pengembangan berbagai jenis aplikasi, mulai dari web hingga sistem backend.

Secara keseluruhan, Go sangat cocok untuk pengembangan aplikasi skala besar dan sistem yang memerlukan performa tinggi.

## Instalasi Golang

Sebelum kamu menulis kode dengan Golang, langkah pertama yang harus kamu lakukan adalah menginstal bahasa pemrograman ini di perangkatmu. Yuk, ikuti langkah-langkah [disini](/basic/installation.md)

## Struktur pada Golang

```go
// ini adalah package yang kamu gunakan
package main

import (
   // ini package yang akan kamu gunakan
)

func main() {
  // kamu bisa menulis kode kamu disini
}
```

## Menulis Kode Program Hello World

Program Golang ini disimpan dengan extension `.go`, Silahkan kamu buat contohnya dengan nama `main.go`.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World!")
}
```

untuk lebih lengkapnya kamu bisa cek kodenya [disini](/basic/01_Pengenalan/main.go)

## Bagaimana Cara Menjalankan Kodenya?

Program Golang ini bisa dijalankan dengan 2 cara yaitu:

1. Melakukan Compile terlebih dahulu, lalu kemudian menjalankan kodenya.

   ```bash
   go build main.go
   ./main
   ```

2. Langsung menjalankan Kode nya tanpa melakukan Kompilasi.

   ```bash
   go run main.go
   ```
