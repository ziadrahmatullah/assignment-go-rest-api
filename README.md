# assignment-go-rest-api

Asumsi:
1. reset expire 5 menit
2. Saat Pembuatan user, auto generate wallet dan juga attempt row
3. pada GetTransactions jika user hanya menginput salah satu antara end dan start pada endpoint, tidak akan error, tapi result tidak di filter
4. pada GetTransactions jika user hanya meninput sort pada endpoint tanpa sortBy, tidak akan error, tapi result tidak di sort
5 pada GetTransactions jika user menginput page yang melebihi total page, maka tidak akan error, hanya mengembalkan dto transaksi yang data nya null

Documentation : https://documenter.getpostman.com/view/31472691/2s9YeN1U24