-- Membuat database slc3
CREATE DATABASE slc3;

-- Membuat table Customers
CREATE TABLE Customers (
    CustomerID INT PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    Email VARCHAR(100) UNIQUE,
    PhoneNumber VARCHAR(20)
);

-- Membuat table Revenue
CREATE TABLE Fields (
    FieldID INT PRIMARY KEY,
    FieldName VARCHAR(100) NOT NULL,
    FieldSize VARCHAR(10) NOT NULL,
    Location VARCHAR(100),
    HourlyRate DECIMAL(10, 2) NOT NULL
);

-- Membuat table Booking
CREATE TABLE Bookings (
    BookingID INT PRIMARY KEY,
    CustomerID INT NOT NULL,
    FieldID INT NOT NULL,
    BookingDate DATE NOT NULL,
    StartTime TIME NOT NULL,
    EndTime TIME NOT NULL,
    TotalAmount DECIMAL(10, 2) NOT NULL,
    
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID),
    FOREIGN KEY (FieldID) REFERENCES Fields(FieldID)
);

CREATE TABLE Payments (
    PaymentID INT PRIMARY KEY,
    BookingID INT UNIQUE NOT NULL,
    PaymentDate DATE NOT NULL,
    PaymentAmount DECIMAL(10, 2) NOT NULL,
    PaymentMethod VARCHAR(50),
    
    FOREIGN KEY (BookingID) REFERENCES Bookings(BookingID)
);


-- -- Membuat table Games
-- CREATE TABLE Games (
--     GameId SERIAL PRIMARY KEY, 
--     Title VARCHAR(255), 
--     Price DECIMAL(10,2)
-- );

-- -- Membuat table Orders
-- CREATE TABLE Orders (
--     OrderId SERIAL PRIMARY KEY,
--     CustomerID INT,
    
--     FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
-- );

-- -- Membuat table OrderDetails
-- CREATE TABLE OrderDetails (
--     OrderDetails SERIAL PRIMARY KEY,
--     OrderId INT,
--     GameId INT,
--     Quantity INT,

--     FOREIGN KEY (OrderId) REFERENCES Orders(OrderId)
--     FOREIGN KEY (GameId) REFERENCES Games(GameId)
-- )