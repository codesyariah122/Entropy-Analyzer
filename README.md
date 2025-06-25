# EntropyCheck - Aplikasi Hitung Entropi Teks, CSV, dan Gambar

EntropyCheck adalah aplikasi web untuk menghitung entropi — ukuran variasi data — pada teks, file CSV, dan gambar. Cocok untuk analisis keunikan konten, deteksi plagiarisme, dan visualisasi distribusi karakter atau pixel.

---

## Fitur

- Hitung entropi dari input teks, file CSV, atau gambar.
- Visualisasi frekuensi karakter/pixel dalam grafik.
- Unduh laporan hasil analisis dalam format PDF.
- Antarmuka simpel dan responsif.

---

## Cara Install

### Backend (Go API)

1. Install [Go](https://go.dev/doc/install) versi minimal 1.18.
2. Clone repo backend, masuk foldernya, lalu jalankan:
   ```bash
   go mod tidy
   go run main.go
   ```  
3. API berjalan di http://localhost:8080

## Frontend (Nuxt 3)

1. Pastikan sudah install Node.js dan package manager (pnpm, npm, atau yarn).

3. Clone repo frontend, masuk foldernya, lalu jalankan:
```
pnpm install
pnpm dev
``` 
3. Buka browser ke http://localhost:3000  


## API Endpoint  
| Endpoint         | Method | Keterangan                        |
| ---------------- | ------ | --------------------------------- |
| `/entropi/text`  | POST   | Hitung entropi dari teks          |
| `/entropi/csv`   | POST   | Hitung entropi dari file CSV      |
| `/entropi/image` | POST   | Hitung entropi dari file gambar   |
| `/entropi/pdf`   | GET    | Unduh laporan hasil entropi (PDF) |

> Manfaat

    - Deteksi variasi dan keunikan teks untuk mencegah plagiarisme.

    - Analisis kompleksitas gambar lewat grayscale.

    - Mempermudah riset dalam bidang kriptografi dan analisis data.

    - Visualisasi data distribusi yang intuitif.

### Lisensi

**MIT License © 2025**

### Kontak

```
    Butuh bantuan atau ingin kontribusi? Hubungi:
    Email: pujiermanto@gmail.com
    GitHub: https://github.com/codesyariah122/Entropy-Analyzer
```