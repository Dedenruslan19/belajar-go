document.addEventListener('DOMContentLoaded', () => {
    // ---------------- Load Header & Footer ----------------
    const loadComponent = async (placeholderId, filePath) => {
        const placeholder = document.getElementById(placeholderId);
        if (!placeholder) return;
        try {
            const response = await fetch(filePath);
            if (!response.ok) throw new Error(`Gagal memuat ${filePath}: ${response.statusText}`);
            const data = await response.text();
            placeholder.innerHTML = data;
        } catch (error) {
            console.error(error);
            placeholder.innerHTML = `<p class="text-danger">Gagal memuat konten.</p>`;
        }
    };

    Promise.all([
        loadComponent('header-placeholder', 'components/header.html'),
        loadComponent('footer-placeholder', 'components/footer.html')
    ]).then(() => {
        const mainContainer = document.getElementById('main-container');
        if (mainContainer) mainContainer.style.opacity = 1;

        // ---------------- Elemen Login ----------------
        const loginWrapper = document.querySelector('.login-link');
        const loginLinkText = loginWrapper ? loginWrapper.querySelector('span') : null;
        const loginModal = document.getElementById('myModal');
        const loginForm = document.getElementById('loginForm');
        const loginMessage = document.getElementById('login-message');
        const loginEmail = document.getElementById('loginEmail');
        const loginPassword = document.getElementById('loginPassword');
        const rememberMe = document.getElementById('loginRememberMe');

        // ---------------- Elemen Signup ----------------
        const signupModal = document.getElementById('signupModal');
        const signupForm = document.getElementById('signupForm');
        const signupUsername = document.getElementById('signupUsername');
        const signupEmail = document.getElementById('signupEmail');
        const signupPassword = document.getElementById('signupPassword');
        const confirmPassword = document.getElementById('confirmPassword');
        const signupTerms = document.getElementById('signupTerms');

        // ---------------- Toggle Password ----------------
        function setupPasswordToggle(modal) {
            if (!modal) return;
            const toggles = modal.querySelectorAll('.password-toggle-icon');
            toggles.forEach(toggle => {
                toggle.addEventListener('click', () => {
                    const input = toggle.closest('.form-floating').querySelector('input');
                    if (!input) return;
                    if (input.type === 'password') {
                        input.type = 'text';
                        toggle.querySelector('i').classList.replace('fa-eye', 'fa-eye-slash');
                    } else {
                        input.type = 'password';
                        toggle.querySelector('i').classList.replace('fa-eye-slash', 'fa-eye');
                    }
                });
            });
        }
        setupPasswordToggle(loginModal);
        setupPasswordToggle(signupModal);

        // ---------------- Join Now Link ----------------
        const joinNowLink = document.getElementById('joinNowLink');
        if (joinNowLink) {
            joinNowLink.addEventListener('click', (e) => {
                e.preventDefault();
                const loginInstance = bootstrap.Modal.getInstance(loginModal);
                if (loginInstance) loginInstance.hide();
                new bootstrap.Modal(signupModal).show();
            });
        }

        if (!loginWrapper) return;

        // ---------------- Set Status LOGIN / LOGOUT ----------------
        if (localStorage.getItem('isLoggedIn') === 'true') {
            loginLinkText.textContent = 'LOGOUT';

            // Jika login sudah ada, update tombol Connect
            const role = localStorage.getItem('role');
            const connectBtn = document.querySelector('.connect-btn');
            if (connectBtn && role === 'admin') {
                connectBtn.href = '/message.html';
            }
        }

        // ---------------- Klik LOGIN / LOGOUT ----------------
        loginWrapper.addEventListener('click', (e) => {
            e.preventDefault();
            if (loginLinkText.textContent === 'LOGOUT') {
                localStorage.removeItem('isLoggedIn');
                localStorage.removeItem('role');
                loginLinkText.textContent = 'LOGIN';
                loginForm.reset();
                loginMessage.textContent = '';
                loginMessage.classList.add('d-none');

                // Reset tombol Connect ke default
                const connectBtn = document.querySelector('.connect-btn');
                if (connectBtn) connectBtn.href = '/connect.html';

                return false;
            } else {
                const modal = new bootstrap.Modal(loginModal);
                modal.show();

                const savedEmail = localStorage.getItem('savedEmail');
                const savedPassword = localStorage.getItem('savedPassword');
                if (savedEmail) loginEmail.value = savedEmail;
                if (savedPassword) loginPassword.value = savedPassword;
                rememberMe.checked = !!savedEmail && !!savedPassword;
            }
        });

        // ---------------- Login Form Submit ----------------
        loginForm?.addEventListener('submit', async (e) => {
            e.preventDefault();
            loginMessage.textContent = '';
            loginMessage.classList.add('d-none');

            const email = loginEmail.value.trim();
            const password = loginPassword.value.trim();

            try {
                const response = await fetch(`/login?email=${encodeURIComponent(email)}&password=${encodeURIComponent(password)}`);
                const data = await response.json();

                if (response.ok && data.success) {
                    // Simpan info login
                    if (rememberMe.checked) {
                        localStorage.setItem('savedEmail', email);
                        localStorage.setItem('savedPassword', password);
                    } else {
                        localStorage.removeItem('savedEmail');
                        localStorage.removeItem('savedPassword');
                    }
                    localStorage.setItem('isLoggedIn', 'true');
                    localStorage.setItem('role', data.user.role);
                    loginLinkText.textContent = 'LOGOUT';

                    // Tutup modal login
                    const modalInstance = bootstrap.Modal.getInstance(loginModal);
                    if (modalInstance) modalInstance.hide();

                    // Update tombol Connect jika admin
                    const connectBtn = document.querySelector('.connect-btn');
                    if (connectBtn && data.user.role === 'admin') {
                        connectBtn.href = '/message.html';
                    }

                } else {
                    loginMessage.textContent = data.message || 'Email atau password salah!';
                    loginMessage.classList.remove('d-none');
                    loginMessage.classList.add('alert-danger');
                }
            } catch (error) {
                loginMessage.textContent = 'Terjadi kesalahan. Silakan coba lagi.';
                loginMessage.classList.remove('d-none');
                loginMessage.classList.add('alert-danger');
                console.error(error);
            }
        });

        // ---------------- Signup Form Submit ----------------
        signupForm?.addEventListener('submit', async (e) => {
            e.preventDefault();
            if (signupPassword.value !== confirmPassword.value) {
                alert('Password dan Confirm Password tidak sama!');
                return;
            }
            if (!signupTerms.checked) {
                alert('Silakan setujui Terms & Conditions!');
                return;
            }

            const data = {
                username: signupUsername.value.trim(),
                email: signupEmail.value.trim(),
                password: signupPassword.value.trim()
            };

            try {
                const response = await fetch('/signup', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data)
                });

                if (response.ok) {
                    alert('Registrasi berhasil! Silakan login.');
                    signupForm.reset();
                    const signupInstance = bootstrap.Modal.getInstance(signupModal);
                    if (signupInstance) signupInstance.hide();
                    new bootstrap.Modal(loginModal).show();
                } else {
                    const resData = await response.json();
                    alert(resData.message || 'Terjadi kesalahan saat registrasi.');
                }
            } catch (error) {
                alert('Terjadi kesalahan. Silakan coba lagi.');
                console.error(error);
            }
        });
    });

    // ---------------- Signup Form Submit ----------------
    document.getElementById("contactForm").addEventListener("submit", async function(event) {
        event.preventDefault();

        const formData = {
            first_name: document.getElementById("firstName").value,
            last_name: document.getElementById("lastName").value,
            email: document.getElementById("email").value,
            phone: document.getElementById("phone").value,
            message_subject: document.getElementById("subject").value,
            your_message: document.getElementById("message").value
        };

        try {
            const response = await fetch("http://localhost:8000/submit-message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(formData)
            });

            const result = await response.json();
            alert(result.message);

        } catch (error) {
            console.error("Error:", error);
            alert("Terjadi kesalahan saat mengirim pesan.");
        }
    });

    // ---------------- Pagination Produk ----------------
    const productRow = document.querySelector('.row.g-0');
    const prevButton = document.querySelector('.page-item:first-child .page-link');
    const nextButton = document.querySelector('.page-item:last-child .page-link');

    if (productRow && prevButton && nextButton) {
        function shiftProductsLeft() {
            const firstProduct = productRow.firstElementChild;
            productRow.appendChild(firstProduct);
        }

        function shiftProductsRight() {
            const lastProduct = productRow.lastElementChild;
            productRow.prepend(lastProduct);
        }
        prevButton.addEventListener('click', (e) => {
            e.preventDefault();
            shiftProductsRight();
        });
        nextButton.addEventListener('click', (e) => {
            e.preventDefault();
            shiftProductsLeft();
        });
    }
});