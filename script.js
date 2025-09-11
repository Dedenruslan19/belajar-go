// Untuk pagination di page products
    // Ambil elemen-elemen yang dibutuhkan
    const productRow = document.querySelector('.row.g-0');
    const prevButton = document.querySelector('.page-item:first-child .page-link');
    const nextButton = document.querySelector('.page-item:last-child .page-link');

    // Fungsi untuk menggeser produk ke kiri
    function shiftProductsLeft() {
        // Ambil elemen produk pertama
        const firstProduct = productRow.firstElementChild;
        // Pindahkan elemen produk pertama ke akhir barisan
        productRow.appendChild(firstProduct);
    }

    // Fungsi untuk menggeser produk ke kanan
    function shiftProductsRight() {
        // Ambil elemen produk terakhir
        const lastProduct = productRow.lastElementChild;
        // Pindahkan elemen produk terakhir ke awal barisan
        productRow.prepend(lastProduct);
    }

    // Tambahkan 'event listener' untuk setiap tombol
    prevButton.addEventListener('click', function(event) {
        event.preventDefault(); // Mencegah link pindah halaman
        shiftProductsRight();
    });

    nextButton.addEventListener('click', function(event) {
        event.preventDefault(); // Mencegah link pindah halaman
        shiftProductsLeft();
    });