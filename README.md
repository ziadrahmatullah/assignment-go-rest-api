# assignment-go-rest-api

Asumsi:
- reset expire 1 menit
- Saat Pembuatan user, auto generate wallet dan juga attempt row
- pada GetTransactions jika user hanya menginput salah satu antara end dan start pada endpoint, maka error invalid filter format
- pada GetTransactions jika user hanya meninput sort pada endpoint tanpa sortBy, maka error invalid sort format
- pada GetTransactions jika user menginput page yang melebihi total page, maka error page not found
- sort pada GetTransaction, date: created_at, amount: amount, to : receiver (wallet number)
- search menggunakan case insensitive

Seeding: /sql

Unit Test Coverage:
- Handler : 100%
- Usecase : 100%
- Util : 100%

Documentation : https://documenter.getpostman.com/view/31472691/2s9YeN1U24