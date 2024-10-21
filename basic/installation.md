# Panduan Instalasi Golang

Sebelum kamu menulis kode dengan Golang, langkah pertama yang harus kamu lakukan adalah menginstal bahasa pemrograman ini di perangkatmu. Yuk, ikuti langkah-langkah berikut ini!

## Instalasi di Windows

1. **Download Go**

   - Kunjungi [halaman unduhan resmi Go](https://golang.org/dl/).
   - Pilih installer Windows (.msi) yang terbaru dan unduh.

2. **Instalasi**

   - Jalankan file installer yang sudah kamu unduh.
   - Ikuti petunjuk yang muncul. Biasanya, Go akan terpasang di `C:\Go`.

3. **Set PATH**

   - Buka **Control Panel** > **System and Security** > **System**.
   - Klik **Advanced system settings**.
   - Klik **Environment Variables**.
   - Cari variabel `Path` di bagian sistem, lalu klik **Edit**.
   - Tambahkan path ke Go: `C:\Go\bin`.
   - Klik **OK** untuk menyimpan perubahan.

4. **Verifikasi Instalasi**

   - Buka Command Prompt dan ketik:

     ```bash
     go version
     ```

   - Jika berhasil, kamu akan melihat versi Go yang terinstal.

## Instalasi di macOS

1. **Download Go**

   - Kunjungi [halaman unduhan resmi Go](https://golang.org/dl/).
   - Pilih installer macOS (.pkg) terbaru dan unduh.

2. **Instalasi**

   - Buka file .pkg yang sudah kamu unduh, lalu ikuti petunjuk instalasinya.
   - Go akan terpasang di `/usr/local/go`.

3. **Set PATH**

   - Buka Terminal dan edit file `.bash_profile` atau `.zshrc` (tergantung shell yang kamu gunakan):

     ```bash
     nano ~/.bash_profile
     ```

     atau

     ```bash
     nano ~/.zshrc
     ```

   - Tambahkan baris ini:

     ```bash
     export PATH=$PATH:/usr/local/go/bin
     ```

   - Simpan dan tutup editor. Untuk menerapkan perubahan, jalankan:

     ```bash
     source ~/.bash_profile
     ```

     atau

     ```bash
     source ~/.zshrc
     ```

4. **Verifikasi Instalasi**

   - Ketik perintah ini di Terminal:

     ```bash
     go version
     ```

   - Kamu seharusnya melihat versi Go yang terinstal.

## Instalasi di Linux

1. **Download Go**

   - Kunjungi [halaman unduhan resmi Go](https://golang.org/dl/).
   - Salin tautan untuk versi tarball Linux (.tar.gz) terbaru dan gunakan `wget` untuk mengunduh:

     ```bash
     wget https://golang.org/dl/go1.xx.x.linux-amd64.tar.gz
     ```

     Gantilah `1.xx.x` dengan versi terbaru.

2. **Instalasi**

   - Ekstrak tarball ke `/usr/local`:

     ```bash
     sudo tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
     ```

   - Go akan terpasang di `/usr/local/go`.

3. **Set PATH**

   - Buka file `.bash_profile` atau `.bashrc`:

     ```bash
     nano ~/.bash_profile
     ```

     atau

     ```bash
     nano ~/.bashrc
     ```

   - Tambahkan baris ini:

     ```bash
     export PATH=$PATH:/usr/local/go/bin
     ```

   - Simpan dan tutup editor. Untuk menerapkan perubahan, jalankan:

     ```bash
     source ~/.bash_profile
     ```

     atau

     ```bash
     source ~/.bashrc
     ```

4. **Verifikasi Instalasi**

   - Ketik perintah berikut di terminal:

     ```bash
     go version
     ```

   - Kamu seharusnya melihat versi Go yang terinstal.

Sekarang kamu sudah berhasil menginstal Go di Windows, macOS, atau Linux. Selamat berkoding!
