-- 1. Tabel Patients
-- DDL
CREATE TABLE Patients (
    patient_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    address TEXT
);
--DML (Isi data Patients)
INSERT INTO Patients (name, date_of_birth, address) VALUES
('Deden', '1990-05-12', 'Solo'),
('Saya', '2000-08-22', 'Tangerang'),
('Anda', '1998-01-15', 'Jakarta'),
('Kamu', '2000-03-18', 'Bekasi');

-- 2. Tabel Doctors
-- DDL
CREATE TABLE Doctors (
    doctor_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    specialization VARCHAR(100) NOT NULL,
    room_number VARCHAR(10) UNIQUE 
);
-- DML (Isi data Doctors)
INSERT INTO Doctors (name, specialization, room_number) VALUES
('Dr. Anda', 'Cardiologist', 'A-101'),
('Dr. Andes', 'Dermatologist', 'A-102'),
('Dr. Budi', 'Orthopedic', 'B-101'),
('Dr. Nina', 'Pediatrician', 'B-102');

-- 3. Tabel Treatments
-- DDL
CREATE TABLE Treatments (
    treatment_id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);
-- DML (Isi data Treatments)
INSERT INTO Treatments (name, description) VALUES
('ECG', 'Heart Rate Test'),
('Skin Therapy', 'Skin Treatments'),
('Physiotherapy', 'Treatments for body health'),
('Vaccination', 'Immunization');

-- 4. Tabel Appointments
-- DDL
CREATE TABLE Appointments (
    appointment_id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL,
    doctor_id INT NOT NULL,
    appointment_date TIMESTAMP WITH TIME ZONE NOT NULL,

    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES Doctors(doctor_id),

    -- Menambah constraint agar Appointment yang memiliki id Patients, id Doctors dan date Appointments tidak dapat diduplikat
    CONSTRAINT unique_patient_doctor_appointment
    UNIQUE (patient_id, doctor_id, appointment_date)
);
-- DML (Isi data Appointments antara id Patients dengan id Doctors)
INSERT INTO Appointments (patient_id, doctor_id, appointment_date) VALUES
(1, 1, '2025-10-05 10:00:00+07'),
(2, 1, '2025-10-05 11:00:00+07'),
(3, 2, '2025-10-06 09:00:00+07'),
(4, 4, '2025-10-06 13:00:00+07');

-- 5. Tabel Billing (Tagihan)
CREATE TABLE Billing (
    billing_id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL,
    treatment_id INT NOT NULL,
    amount_billed NUMERIC(10, 2) NOT NULL CHECK (amount_billed >= 0),
    billing_date DATE NOT NULL DEFAULT CURRENT_DATE,

    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (treatment_id) REFERENCES Treatments(treatment_id),

    CONSTRAINT unique_patient_treatment_date
    UNIQUE (patient_id, treatment_id, billing_date)
);

-- DML (Menentukan amount_billed berdasarkan jenis treatment yang diambil dari table Treatments (diurutkan by ID))
INSERT INTO Billing (patient_id, treatment_id, amount_billed, billing_date)
SELECT 
    a.patient_id,
    at.treatment_id,
    CASE at.treatment_id
        WHEN 1 THEN 250000.00  
        WHEN 2 THEN 400000.00  
        WHEN 3 THEN 500000.00  
        WHEN 4 THEN 150000.00  
        ELSE 100000.00
    END AS amount_billed,
    a.appointment_date::date AS billing_date
FROM Appointments a
JOIN AppointmentTreatments at ON a.appointment_id = at.appointment_id;

-- 6. Tabel Relasi: Treatment
CREATE TABLE AppointmentTreatments (
    appointment_treatment_id SERIAL PRIMARY KEY,
    appointment_id INT NOT NULL,
    treatment_id INT NOT NULL,
    
    FOREIGN KEY (appointment_id) REFERENCES Appointments(appointment_id),
    FOREIGN KEY (treatment_id) REFERENCES Treatments(treatment_id),

    UNIQUE (appointment_id, treatment_id) 
);
-- DML
INSERT INTO AppointmentTreatments (appointment_id, treatment_id) VALUES
(1, 1),  (2, 1), (3, 2), (4, 4);

-- a. Retrieve the names of patients who have appointments with a particular doctor.
SELECT p.name AS patient_name, d.name AS doctor_name
FROM Appointments a
JOIN Patients p ON a.patient_id = p.patient_id
JOIN Doctors d ON a.doctor_id = d.doctor_id
WHERE d.name = 'Dr. Anda';

-- Menggunakan CTE
WITH DoctorAppointments AS (
    SELECT 
        a.appointment_id,
        p.name AS patient_name,
        d.name AS doctor_name,
        a.appointment_date
    FROM Appointments a
    JOIN Patients p ON a.patient_id = p.patient_id
    JOIN Doctors d ON a.doctor_id = d.doctor_id
)
SELECT 
    patient_name,
    doctor_name
FROM DoctorAppointments
WHERE doctor_name = 'Dr. Anda';

--Output:
 patient_name | doctor_name
--------------+-------------
 Deden        | Dr. Anda
 Saya         | Dr. Anda
(2 rows)


-- b. Calculate the total billing amount for a particular patient.
SELECT 
    p.patient_id,
    p.name AS patient_name,
    SUM(b.amount_billed) AS total_billing
FROM Billing b
JOIN Patients p ON b.patient_id = p.patient_id
WHERE p.name = 'Deden'
GROUP BY p.patient_id, p.name;

-- Menggunakan CTE
WITH PatientBilling AS (
    SELECT 
        b.patient_id,
        p.name AS patient_name,
        b.amount_billed
    FROM Billing b
    JOIN Patients p ON b.patient_id = p.patient_id
)
SELECT 
    patient_id,
    patient_name,
    SUM(amount_billed) AS total_billing
FROM PatientBilling
WHERE patient_name = 'Deden'
GROUP BY patient_id, patient_name;

-- Membuat fungsi untuk view
CREATE VIEW TotalPatientBilling AS
SELECT 
    p.patient_id,
    p.name AS patient_name,
    SUM(b.amount_billed) AS total_billing
FROM Billing b
JOIN Patients p ON b.patient_id = p.patient_id
GROUP BY p.patient_id, p.name;

-- Memanggil data dengan fungsi View
SELECT * FROM TotalPatientBilling WHERE patient_name = 'Deden';

-- Output:
 patient_id | patient_name | total_billing
------------+--------------+---------------
          1 | Deden        |     250000.00
(1 row)

-- c. List the doctors who currently have no appointments scheduled.
SELECT d.name AS doctor_name
FROM Doctors d
LEFT JOIN Appointments a ON d.doctor_id = a.doctor_id
WHERE a.doctor_id IS NULL;

-- Menggunakan CTE
WITH DoctorAppointments AS (
    SELECT 
        d.doctor_id,
        d.name AS doctor_name,
        a.appointment_id
    FROM Doctors d
    LEFT JOIN Appointments a ON d.doctor_id = a.doctor_id
)
SELECT 
    doctor_name
FROM DoctorAppointments
WHERE appointment_id IS NULL;

-- Membuat fungsi View
CREATE VIEW DoctorsWithoutAppointments AS
SELECT 
    d.doctor_id,
    d.name AS doctor_name
FROM Doctors d
LEFT JOIN Appointments a ON d.doctor_id = a.doctor_id
WHERE a.doctor_id IS NULL;

-- Memanggil data dengan fungsi View
SELECT * FROM DoctorsWithoutAppointments;

-- Output
 doctor_name
-------------
 Dr. Budi
(1 row)

-- d. Find the treatments prescribed most frequently.
SELECT t.name AS treatment_name, COUNT(at.treatment_id) AS frequency
FROM AppointmentTreatments at
JOIN Treatments t ON at.treatment_id = t.treatment_id
GROUP BY t.name
ORDER BY frequency DESC
LIMIT 1;

-- Menggunakan CTE
WITH TreatmentFrequency AS (
    SELECT 
        t.name AS treatment_name,
        COUNT(at.treatment_id) AS frequency
    FROM AppointmentTreatments at
    JOIN Treatments t ON at.treatment_id = t.treatment_id
    GROUP BY t.name
)
SELECT 
    treatment_name, 
    frequency
FROM TreatmentFrequency
ORDER BY frequency DESC
LIMIT 1;

-- Membuat fungsi view
CREATE VIEW TreatmentUsageFrequency AS
SELECT 
    t.name AS treatment_name,
    COUNT(at.treatment_id) AS frequency
FROM AppointmentTreatments at
JOIN Treatments t ON at.treatment_id = t.treatment_id
GROUP BY t.name;

-- Memanggil data dengan view
SELECT * FROM TreatmentUsageFrequency ORDER BY frequency DESC LIMIT 1;

 treatment_name | frequency
----------------+-----------
 ECG            |         2
(1 row)