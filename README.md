# assignment-go-rest-api

Asumsi:
1. reset expire 5 menit
2. Saat Pembuatan user, auto generate wallet dan juga attempt row
3. pada GetTransactions jika user hanya menginput salah satu antara end dan start pada endpoint, maka error invalid filter format
4. pada GetTransactions jika user hanya meninput sort pada endpoint tanpa sortBy, maka error invalid sort format
5 pada GetTransactions jika user menginput page yang melebihi total page, maka error page not found

Documentation : https://documenter.getpostman.com/view/31472691/2s9YeN1U24