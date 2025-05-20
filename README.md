# 🏋️ SweatUp – Fitness Booking App

SweatUp adalah aplikasi berbasis web yang dirancang untuk membantu pelanggan memesan kelas kebugaran, mengelola paket membership, melakukan pembayaran, dan memantau kehadiran dengan mudah. Aplikasi ini memiliki dashboard admin lengkap serta fitur untuk pengguna umum (customer) dan instruktur.

---

## 🔧 Tech Stack

### 🚀 Frontend

- **React 18** dengan **React Router v7**
- **TailwindCSS** + **Radix UI** + **ShadCN UI**
- **React Hook Form + Zod** (validasi form)
- **Zustand** (state management)
- **TanStack React Query** (data fetching dan caching)
- **Framer Motion** (animasi)
- **Google OAuth** via `@react-oauth/google`
- **Vite** (build tool)

### 🔐 Backend

- **Golang** (Gin Framework)
- **GORM** (ORM untuk MySQL)
- **MySQL** (Database utama)
- **Redis** (Caching & token store)
- **Midtrans** (Pembayaran)
- **RabbitMQ** (Event-driven notification)
- **JWT Auth** (Access & Refresh Token)

### 🐳 Deployment

- **Docker + Docker Compose**
- **Nginx** sebagai reverse proxy

---

## 📁 Project Structure

### Frontend

```
├── pages/
│ ├── public/ → Halaman utama, signup, signin, dll
│ ├── customer/ → Halaman profil pengguna, booking, paket
│ └── admin/ → Dashboard admin, manajemen kelas & paket
├── components/ → Komponen UI reusable
├── store/ → Zustand global state
├── hooks/ → Custom hooks
└── lib/ → Konfigurasi query, axios, dll
```

### Backend

```
server/
├── internal/
│ ├── handlers/ → Controller logic
│ ├── services/ → Bisnis logic
│ ├── repositories/ → Akses ke database
│ ├── models/ → Struktur DB (GORM)
│ └── middleware/ → Middleware auth, role, dsb
├── config/ → Konfigurasi DB, env
├── routes/ → Routing API
└── main.go → Entry point aplikasi
```

---

## 🔐 Authentication & Authorization

- **OTP Login** & **Google OAuth**
- **Access Token**: untuk akses API
- **Refresh Token**: untuk perpanjangan sesi
- Middleware Role: `customer`, `admin`, `owner`

---

## 📲 API Features & Endpoints

### 👥 Auth

- `POST /auth/send-otp` → Kirim OTP
- `POST /auth/verify-otp` → Verifikasi OTP
- `POST /auth/register` / `login` / `logout`
- `GET /auth/me` → Info pengguna saat ini
- `GET /auth/google` → Login Google

### 📦 Paket

- `GET /api/packages` → List paket
- `POST /api/packages` → Tambah paket (admin)

### 🧘‍♂️ Class

- `GET /api/classes` / `active` / `:id`
- `POST /api/classes` → Tambah class (admin)
- `POST /api/classes/:id/gallery` → Upload galeri class

### 📅 Jadwal Kelas

- `GET /api/schedules` → Semua jadwal
- `GET /api/schedules/:id` → Detail
- `GET /api/schedules/status` → Status booking user
- `POST /api/schedules` → Tambah jadwal (admin)
- `GET /api/schedules/dashboard/revenue` → Statistik pendapatan (admin)

### 🧾 Booking & Attendance

- `POST /api/bookings` → Booking kelas
- `GET /api/attendances` → Daftar kehadiran
- `POST /api/attendances/:id` → Check-in
- `GET /api/attendances/:id/qr-code` → Regenerasi QR
- `POST /api/attendances/validate` → Validasi scan QR (admin)

### 🧑‍🏫 Instruktur

- `GET /api/instructors` → Daftar instruktur
- `POST /api/instructors` → Tambah instruktur (admin)

### 🗺️ Lokasi, Tipe, Level, Kategori

- CRUD lengkap untuk:
  - `locations`
  - `types`
  - `levels`
  - `categories` & `subcategories`

### 🔔 Notifikasi

- `GET /api/notifications` → Daftar notifikasi
- `PUT /api/notifications/settings` → Atur preferensi
- `POST /api/notifications/broadcast` → Kirim massal (admin)

### 💳 Pembayaran

- `POST /api/payments` → Bayar Midtrans
- `POST /api/payments/midtrans/notification` → Webhook Midtrans
- `GET /api/payments` → Data pembayaran user (admin)

### 🎫 Voucher

- `GET /api/vouchers` → Semua voucher
- `POST /api/vouchers` → Tambah voucher (admin)
- `POST /api/vouchers/apply` → Apply voucher

### ⭐ Review

- `POST /api/reviews` → Kirim review
- `GET /api/reviews/:classId` → Review berdasarkan kelas

### 📋 Jadwal Template (Recurring)

- `GET /api/schedule-templates`
- `POST /api/schedule-templates/:id/run` → Jalankan cron job
- `POST /api/schedule-templates/:id/stop` → Hentikan cron

---

## 📌 Fitur Unggulan

- JWT + OAuth Authentication – kombinasi login konvensional dan Google OAuth

- Sistem caching otomatis menggunakan TanStack React Query untuk performa optimal

- Papan kalender interaktif untuk manajemen dan penjadwalan kelas

- Pembuatan class terjadwal dan berulang yang digenerate otomatis menggunakan cron job

- Booking kelas berbasis paket kuota pengguna

- QR Code Attendance lengkap dengan regenerasi QR dan validasi admin

- Review dan rating untuk evaluasi kualitas kelas & instruktur

- Google OAuth login langsung dari frontend

- Fitur backup database otomatis dengan cron job terjadwal

- Role-based access control (admin, customer, owner) untuk keamanan dan manajemen akses

- Dashboard admin untuk pemantauan performa, statistik pengguna, dan data revenue

---

## ⚙️ Environment

## 🔐 Konfigurasi Environment (.env)

### Frontend

```

VITE_API_SERVICES=https://www.your-backend-api-url.com
VITE_MIDTRANS_PUBLIC_KEY=your-midtrans-public-key
VITE_MIDTRANS_URL=your-midtrans-snap-url
VITE_API_KEY=your-backend-api-key
VITE_GOOGLE_CLIENT_ID=your-google-client-id

```

### Backend

```

# ==== MySQL ====

DB_HOST=your-db-host # must same with the configuration in docker-compose file
DB_PORT=your-db-port
DB_USERNAME=your-db-username
DB_NAME=your-db-name
DB_PASSWORD=your-db-password

# ==== cloudinary ====

CLOUDINARY_CLOUD_NAME=cloudinary-cloud-name
CLOUDINARY_API_KEY=cloudinary-api-key
CLOUDINARY_FOLDER_NAME=your-cloudinary-folder-name
CLOUDINARY_API_SECRET=your-cloudinary-api-secret

# ==== nodemailer ====

USER_EMAIL=your-email
USER_PASSWORD=your-mailer-password

# ====== redis ======

REDIS_ADDR=your-redis-port # default is 6379
REDIS_PASSWORD=your-redis-password # leave it blank if you dont want to use password
PORT=your-server-port
JWT_ACCESS_SECRET=your-token-access-secret
JWT_REFRESH_SECRET=your-token-refresh-secret
API_KEY=your-api-key-secret
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_CLIENT_ID=your-google-client-id
MIDTRANS_CLIENT_KEY=your-midtrans-client-key
MIDTRANS_SERVER_KEY=your-midtrans-server-key
PAYMENT_TAX_RATE=0.10 # ! adjust to your tax needs 0.10 = 10 %
NODE_ENV=development # -> change to production for deployment
TRUSTED_PROXIES=your-ip-server-address # -> for development leave it blank
TEST_MODE=true # change to false for deployment
GOOGLE_REDIRECT_URL=https://your-api-domain/api/auth/google/callback
FRONTEND_REDIRECT_URL=your-front-end-domain
COOKIE_DOMAIN=your-front-end-domain
ALLOWED_ORIGINS=your-front-end-domain

```

## ☁️ Deployment

saya menggunakan web-hosting untuk project ini, namun kamu juga bisa mendeploy ke platform Saas seperti vercel ataupun netlify :

- Hubungkan repo Git
- Tambahkan .env ke dashboard environment
- Atur build command:

```

npm run build

```

Output directory: dist

## 🤝 Kontribusi

Terbuka untuk Kontribusi bagi yang ingin mengembangkan fitur lebih jauh :

- Fork repository ini
- Buat branch: git checkout -b fitur-anda
- Commit perubahan: git commit -m 'feat: fitur baru'
- Push ke branch: git push origin fitur-anda
- Buka Pull Request

## License

**MIT License**

##👤 Developer

- name : Ahmad Fiqri oemmry
- 📁 email : fiqrioemry@gmail.com
- 🌐 Linkedin : https://www.linkedin.com/in/ahmadfiqrioemry

## 🖼️ Preview

Berikut adalah beberapa preview tampilan untuk halaman website travel planner ini. Homepage Create Trip Detail Trip My Trip List

![Preview1](./public/preview1.png)
![Preview2](./public/preview2.png)
![Preview3](./public/preview3.png)
![Preview4](./public/preview4.png)

```

```
