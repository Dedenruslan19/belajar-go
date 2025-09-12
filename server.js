const express = require('express');
const path = require('path');
const {
    Pool
} = require('pg');
const bcrypt = require('bcrypt');

const app = express();
const port = 8000;

// Middleware untuk memproses data dari formulir
app.use(express.json());
app.use(express.urlencoded({
    extended: true
}));

// Buat pool koneksi ke database PostgreSQL
const pool = new Pool({
    user: "postgres",
    password: "root",
    database: "milestone_project",
    port: 5432,
    host: "localhost",
});

// ========================== LOGIN ==========================
app.get('/login', async (req, res) => {
    const {
        email,
        password
    } = req.query;

    if (!email || !password) {
        return res.status(400).json({
            success: false,
            message: 'Email dan password wajib diisi.'
        });
    }

    let client;
    try {
        client = await pool.connect();
        const queryText = 'SELECT * FROM users WHERE email = $1';
        const result = await client.query(queryText, [email]);

        if (result.rows.length === 0) {
            return res.status(401).json({
                success: false,
                message: 'Email atau password salah.'
            });
        }

        const user = result.rows[0];
        const match = await bcrypt.compare(password, user.password);
        if (!match) {
            return res.status(401).json({
                success: false,
                message: 'Email atau password salah.'
            });
        }

        let redirectUrl = null;
        if (user.role === 'admin') {
            redirectUrl = '/message.html';
        }

        res.json({
            success: true,
            message: 'Login berhasil!',
            user: {
                user_id: user.user_id,
                username: user.username,
                email: user.email,
                role: user.role
            },
            redirect: redirectUrl
        });

    } catch (err) {
        console.error('Error during login:', err.stack);
        res.status(500).json({
            success: false,
            message: 'Terjadi kesalahan server saat login.'
        });
    } finally {
        if (client) client.release();
    }
});

// ========================== SIGNUP ==========================
app.post('/signup', async (req, res) => {
    const {
        username,
        email,
        password
    } = req.body;

    if (!username || !email || !password) {
        return res.status(400).json({
            success: false,
            message: 'Username, email, dan password wajib diisi.'
        });
    }

    let client;
    try {
        client = await pool.connect();

        const checkQuery = 'SELECT * FROM users WHERE username = $1 OR email = $2';
        const checkResult = await client.query(checkQuery, [username, email]);

        if (checkResult.rows.length > 0) {
            return res.status(409).json({
                success: false,
                message: 'Username atau email sudah digunakan.'
            });
        }

        const hashedPassword = await bcrypt.hash(password, 10);

        const insertQuery = `
            INSERT INTO users (username, email, password, role, created_at)
            VALUES ($1, $2, $3, 'user', NOW())
            RETURNING user_id, username, email, role, created_at
        `;
        const insertResult = await client.query(insertQuery, [username, email, hashedPassword]);

        res.status(201).json({
            success: true,
            message: 'Registrasi berhasil!',
            user: insertResult.rows[0]
        });

    } catch (err) {
        console.error('Error during signup:', err.stack);
        res.status(500).json({
            success: false,
            message: 'Terjadi kesalahan server saat registrasi.'
        });
    } finally {
        if (client) client.release();
    }
});

// ========================== SUBMIT MESSAGE ==========================
app.post('/submit-message', async (req, res) => {
    const {
        first_name,
        last_name,
        email,
        phone,
        message_subject,
        your_message
    } = req.body;
    const full_name = `${first_name} ${last_name}`.trim();

    if (!full_name || !email) {
        return res.status(400).json({
            success: false,
            message: 'Nama lengkap dan email wajib diisi.'
        });
    }

    let client;
    try {
        client = await pool.connect();
        const queryText = `
            INSERT INTO user_message (full_name, email, phone, message_subject, your_message)
            VALUES ($1, $2, $3, $4, $5)
            RETURNING message_id, created_at;
        `;
        const values = [full_name, email, phone, message_subject, your_message];

        await client.query(queryText, values);

        res.status(201).json({
            success: true,
            message: 'Pesan berhasil terkirim!'
        });

    } catch (err) {
        console.error('Error saat mengirim pesan:', err.stack);
        res.status(500).json({
            success: false,
            message: 'Terjadi kesalahan server saat mengirim pesan.'
        });
    } finally {
        if (client) client.release();
    }
});

// ========================== GET MESSAGES ==========================
app.get('/messages', async (req, res) => {
    let client;
    try {
        client = await pool.connect();
        const queryText = `
            SELECT message_id, full_name, email, phone, message_subject, your_message, created_at
            FROM user_message
            ORDER BY created_at DESC;
        `;
        const result = await client.query(queryText);

        res.status(200).json({
            success: true,
            messages: result.rows
        });

    } catch (err) {
        console.error('Error fetching messages:', err.stack);
        res.status(500).json({
            success: false,
            message: 'Terjadi kesalahan server saat mengambil data pesan.'
        });
    } finally {
        if (client) client.release();
    }
});

// Delete data message
app.delete('/messages/:id', async (req, res) => {
    const {
        id
    } = req.params;
    let client;
    try {
        client = await pool.connect();
        const queryText = `DELETE FROM user_message WHERE message_id = $1 RETURNING *;`;
        const result = await client.query(queryText, [id]);

        if (result.rowCount === 0) {
            return res.status(404).json({
                success: false,
                message: 'Pesan tidak ditemukan.'
            });
        }

        res.json({
            success: true,
            message: 'Pesan berhasil dihapus.'
        });
    } catch (err) {
        console.error('Error deleting message:', err.stack);
        res.status(500).json({
            success: false,
            message: 'Terjadi kesalahan server saat menghapus pesan.'
        });
    } finally {
        if (client) client.release();
    }
});

// ========================== STATIC FILES ==========================
app.use(express.static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'public', 'index.html'));
});

// ========================== START SERVER ==========================
app.listen(port, () => {
    console.log(`Server is running at http://localhost:${port}`);
});