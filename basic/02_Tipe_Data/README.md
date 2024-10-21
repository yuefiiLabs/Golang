# Tipe Data Pada Golang

Pada Bahasa Pemograman Golang, ada berbagai tipe data dasar yang bisa kamu gunakan saat pemrograman. Yuk, kita bahas tipe-tipe data utama yang ada di Go!

## 1. Tipe Data Numerik

- **Integer**: Tipe untuk bilangan bulat.
  - `int` (ukuran bervariasi, biasanya 32 atau 64 bit)
  - `int8`, `int16`, `int32`, `int64` (bilangan bulat dengan ukuran tertentu)
- **Unsigned Integer**: Tipe untuk bilangan bulat positif.

  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`

- **Floating Point**: Tipe untuk bilangan desimal.

  - `float32`, `float64`

- **Bilangan Kompleks**: Tipe untuk bilangan kompleks.
  - `complex64`, `complex128`

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/tipe_data_numerik.go)

## 2. Tipe Data Karakter dan String

- **rune**: Mewakili karakter Unicode, sama dengan `int32`.
- **string**: Tipe data untuk menyimpan teks (sekuens karakter).

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/tipe_data_string.go)

## 3. Tipe Data Boolean

- **bool**: Tipe data untuk menyimpan nilai `true` atau `false`.

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/tipe_data_boolean.go)

## 4. Tipe Data Khusus

- **Array**: Kumpulan elemen dengan tipe yang sama dan ukuran tetap.

  ```go
  var arr [5]int // Array dengan 5 elemen integer
  ```

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/array.go)

- **Slice**: Kumpulan elemen dengan tipe yang sama, tapi ukurannya bisa berubah.

  ```go
  var slice []int // Slice yang dapat tumbuh
  ```

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/slice.go)

- **Map**: Koleksi pasangan kunci-nilai.

  ```go
  var m map[string]int // Map dengan kunci string dan nilai integer
  ```

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/map.go)

- **Struct**: Tipe data yang bisa menyimpan berbagai nilai dengan tipe berbeda.

  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/struct.go)

- **Interface**: Tipe data yang mendefinisikan perilaku. Tipe yang berbeda bisa memenuhi kontrak ini.

  ```go
  type Shape interface {
      Area() float64
  }
  ```

untuk contoh kodenya bisa kamu cek [disini](/basic/02_Tipe_Data/interface.go)

Dengan memahami tipe-tipe data ini, kamu akan lebih mudah dalam memprogram menggunakan Golang! Selamat mencoba!
