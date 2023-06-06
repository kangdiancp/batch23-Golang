CREATE TABLE T_POLICY (
	POLICY_NUMBER VARCHAR(10),
	POLICY_SUBMIT_DATE DATE,
	PREMIUM NUMERIC(18, 2),
	DISCOUNT INTEGER,
	COMMISSION NUMERIC(18, 2),
	CLIENT_NUMBER VARCHAR(10) UNIQUE,
	AGENT_CODE VARCHAR(10) UNIQUE,
	POLICY_STATUS VARCHAR(10),
	POLICY_DUE_DATE DATE,
	CONSTRAINT T_POLICY_PK PRIMARY KEY (POLICY_NUMBER, CLIENT_NUMBER, AGENT_CODE)
)

SELECT * FROM T_POLICY
DROP TABLE T_POLICY cascade

CREATE TABLE T_AGENT (
	AGENT_CODE VARCHAR(10),
	AGENT_NAME VARCHAR(50),
	AGENT_OFFICE VARCHAR(50),
	BASIC_COMMISSION NUMERIC(18, 2),
	FOREIGN KEY (AGENT_CODE) REFERENCES T_POLICY (AGENT_CODE)
)
SELECT * FROM T_AGENT
DROP TABLE T_AGENT

CREATE TABLE T_CLIENT (
	CLIENT_NUMBER VARCHAR(10),
	CLIENT_NAME VARCHAR(50),
	BIRTH_DATE DATE,
	FOREIGN KEY (CLIENT_NUMBER) REFERENCES T_POLICY (CLIENT_NUMBER)
)
SELECT * FROM T_CLIENT
DROP TABLE T_CLIENT

-- Menambahkan data ke dalam tabel T_POLICY
INSERT INTO T_POLICY (POLICY_NUMBER, POLICY_SUBMIT_DATE, PREMIUM, DISCOUNT, COMMISSION, CLIENT_NUMBER, AGENT_CODE, POLICY_STATUS, POLICY_DUE_DATE)
VALUES
    ('001', '2018-01-05', 18600000.00, 10, 930000.00, 'CL001', 'AG001', 'INFORCE', null),
    ('002', '2018-01-05', 2500000.00, 10, 125000.00, 'CL002', 'AG002', 'INFORCE', null),
    ('003', '2018-01-10', 2500000.00, 10, 125000.00, 'CL003', 'AG003', 'TERMINATE', null),
    ('004', '2018-01-25', 4000000.00, 10, 200000.00, 'CL004', 'AG004', 'PROPOSAL', null),
    ('005', '2018-01-25', 2625000.00, 10, 131250.00, 'CL005', 'AG005', 'PROPOSAL', null);
	
-- Menambahkan data ke dalam tabel T_AGENT
INSERT INTO T_AGENT (AGENT_CODE, AGENT_NAME, AGENT_OFFICE, BASIC_COMMISSION)
VALUES
    ('AG001', 'LIDYA', 'JAKARTA', null),
    ('AG002', 'RUDI', 'BALI', null),
    ('AG003', 'DICKY', 'BALI', null),
    ('AG004', 'DAVID', 'SURABAYA', null),
    ('AG005', 'FARIZ', 'SURABAYA', null);
	
-- Menambahkan data ke dalam tabel T_CLIENT
INSERT INTO T_CLIENT (CLIENT_NUMBER, CLIENT_NAME, BIRTH_DATE)
VALUES
    ('CL001', 'WAYNE ROONEY', '1956-11-05'),
    ('CL002', 'RYAN GIGGS', '1972-09-03'),
    ('CL003', 'DAVID BECKHAM', '1985-09-03'),
    ('CL004', 'MICHAEL OWEN', '2012-09-03');
	
	
-- a. Menampilkan data polis yang disubmit setelah tanggal 15 JANUARI 2018 berdasarkan client yang lahir di bulan SEPTEMBER.
SELECT * FROM T_POLICY AS P
LEFT JOIN T_CLIENT AS C 
ON P.CLIENT_NUMBER = C.CLIENT_NUMBER
WHERE P.POLICY_SUBMIT_DATE > '2018-01-15' AND EXTRACT(MONTH FROM C.BIRTH_DATE) = 9;

-- b. Menampilkan data polis dengan status INFORCE, yang agentnya berkantor di "JAKARTA"
SELECT * FROM T_POLICY AS P
LEFT JOIN T_AGENT AS A
ON P.AGENT_CODE = A.AGENT_CODE
WHERE P.POLICY_STATUS = 'INFORCE' AND A.AGENT_OFFICE = 'JAKARTA';


-- c. Menghitung Kolom Basic Commission pada Table T_AGENT dengan rumus: nilai (COMMISSION/ PREMIUM) 100 dari table T_POLICY
SELECT *, (P.COMMISSION / P.PREMIUM) * 100 AS BASIC_COMMISSION
FROM T_AGENT AS A
LEFT JOIN T_POLICY AS P 
ON A.AGENT_CODE = P.AGENT_CODE;

-- d. Mengisi kolom POLICY_DUE _DATE dari table T_POLICY, dengan kondisi tanggal akhir bulan dari 30 Hari setelah tanggal POLICY_SUBMIT_DATE
UPDATE T_POLICY
SET POLICY_DUE_DATE = DATE_TRUNC('MONTH', POLICY_SUBMIT_DATE) + INTERVAL '1 MONTH' - INTERVAL '1 DAY' + INTERVAL '30 DAYS';

-- e. Menampilkan semua data produksi agent yang nilai premiumnya setelah dipotong DISCOUNT dibawah SATU JUTA (diurutkan dari nilai terkecil).
SELECT *
FROM T_POLICY
WHERE (CAST(PREMIUM AS NUMERIC) / CAST(DISCOUNT AS NUMERIC)) < 1000000
ORDER BY (CAST(PREMIUM AS NUMERIC) / CAST(DISCOUNT AS NUMERIC));