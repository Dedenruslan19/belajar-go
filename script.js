
// Refactor Code
    document.addEventListener('DOMContentLoaded', () => {
        const loadComponent = async (placeholderId, filePath) => {
            const placeholder = document.getElementById(placeholderId);
            if (!placeholder) {
                console.warn(`Placeholder dengan ID "${placeholderId}" tidak ditemukan.`);
                return;
            }

            try {
                const response = await fetch(filePath);
                if (!response.ok) {
                    throw new Error(`Gagal memuat ${filePath}: ${response.statusText}`);
                }
                const data = await response.text();
                
                // Setelah konten dimuat, ganti spinner dengan konten yang sebenarnya
                placeholder.innerHTML = data;

            } catch (error) {
                console.error(`Terjadi kesalahan saat memuat komponen ${filePath}:`, error);
                placeholder.innerHTML = `<p class="text-danger">Gagal memuat konten.</p>`;
            }
        };

        Promise.all([
            loadComponent('header-placeholder', '/components/header.html'),
            loadComponent('footer-placeholder', '/components/footer.html'),
        ]).then(() => {
            const mainContainer = document.getElementById('main-container');
            if (mainContainer) {
                mainContainer.style.opacity = 1;
            }
            window.scrollTo(0, 0); 
        }).catch(error => {
            console.error('Satu atau lebih komponen gagal dimuat:', error);
        });
    });

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
