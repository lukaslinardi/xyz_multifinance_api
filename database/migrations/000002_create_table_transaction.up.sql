CREATE TABLE `transaction_record` 
(
    `id` SERIAL PRIMARY KEY,
    `nik_user` VARCHAR(255) NOT NULL, 
    `contract_number`INT NOT NULL,
    `otr` INT NOT NULL,
    `admin_fee` INT NOT NULL,
    `total_instalment` INT NOT NULL,
    `total_interest` INT NOT NULL,
    `asset_name` INT NOT NULL,
    `picture` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP
);
