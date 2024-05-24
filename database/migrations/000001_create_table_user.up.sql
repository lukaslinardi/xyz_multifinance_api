CREATE TABLE `users` 
(
    `id` SERIAL PRIMARY KEY,
    `nik` VARCHAR(255) NOT NULL,
    `fullname` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `legal_name` VARCHAR(255) NOT NULL,
    `birth_place` VARCHAR(255) NOT NULL,
    `birth_date` DATE NOT NULL,
    `salary` INT NOT NULL,
    `ktp_picture` VARCHAR(255) NOT NULL,
    `picture` VARCHAR(255) NOT NULL
);
