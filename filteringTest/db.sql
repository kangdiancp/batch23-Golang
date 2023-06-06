CREATE TABLE T_POLICY (
	Policy_Number integer,
	Policy_Submit_Date date,
	Premium money,
	Discount integer,
	Commission money,
	Client_Number varchar(20) unique,
	Agent_Code varchar(20) unique,
	Policy_Status varchar(25) ,
	Policy_Due_Date date,
	Constraint T_POLICY_PK PRIMARY KEY (Policy_Number, Client_Number, Agent_Code)
)

	SELECT*FROM T_POLICY

CREATE TABLE T_AGENT (
	Agent_Code varchar(20),
	Agent_Name varchar(50),
	Agent_Office varchar(30),
	Basic_Commission money,
	FOREIGN KEY (Agent_Code) REFERENCES T_POLICY (Agent_Code)
)

	SELECT*FROM T_AGENT

CREATE TABLE T_CLIENT (
	Client_Number varchar(20),
	Client_Name varchar (50), 
	Birth_Date date,
	FOREIGN KEY (Client_Number) REFERENCES T_POLICY (Client_Number)
)

	SELECT*FROM T_CLIENT

INSERT INTO T_POLICY (Policy_Number, Policy_Submit_Date, Premium, Discount, Commission, Client_Number, Agent_Code, Policy_Status, Policy_Due_Date)
VALUES
(001, '2018-1-5', 18600000, 10, 930000, 'CL001', 'AG001', 'INFORCE', null),
(002, '2018-1-5', 2500000, 10, 125000, 'CL002', 'AG002', 'INFORCE', null),
(003, '2018-1-10', 2500000, 10, 125000, 'CL003', 'AG003', 'TERMINATE', null),
(004, '2018-1-25', 4000000, 10, 200000, 'CL004', 'AG004', 'PROPOSAL', null),
(005, '2018-1-25', 2625000, 10, 131250, 'CL005', 'AG005', 'PROPOSAL', null)

INSERT INTO T_AGENT (Agent_Code, Agent_Name, Agent_Office, Basic_Commission)
VALUES
('AG001', 'LIDYA', 'JAKARTA', NULL),
('AG002', 'RUDI', 'BALI', NULL),
('AG003', 'DICKY', 'BALI', NULL),
('AG004', 'DAVID', 'SURABAYA', NULL),
('AG005', 'FARIZ', 'SURABAYA', NULL)

INSERT INTO T_CLIENT (Client_Number, Client_Name, Birth_Date) 
VALUES
('CL001', 'WAYNE ROONEY', '1956-5-11'),
('CL002', 'RYAN GIGGS', '1972-9-3'),
('CL003', 'DAVID BECKHAM', '1985-9-3'),
('CL004', 'MICHAEL OWEN', '2012-9-3')


SELECT * FROM t_policy a
	LEFT JOIN t_client b
		ON a.client_number = b.client_number
			WHERE policy_submit_date > '2018-01-15' AND EXTRACT(MONTH FROM birth_date) = 9;

SELECT * FROM t_policy a
	LEFT JOIN t_agent b
		ON a.agent_code = b.agent_code
			WHERE policy_status = 'INFORCE' AND agent_office = 'JAKARTA';
	
UPDATE t_agent a
	SET basic_commission = (b.commission/b.premium) * 100 :: MONEY
		FROM t_policy b
			WHERE basic_commission IS NULL

UPDATE t_policy
    SET policy_due_date = DATE_TRUNC('MONTH', policy_submit_date + INTERVAL '30 days') + INTERVAL '1 MONTH' - INTERVAL '1 DAY';
	
	
	
	
SELECT a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 AS MONEY) AS produksi_agent
    FROM t_agent AS a 
    	LEFT JOIN t_policy as p ON p.agent_code = a.agent_code
    		WHERE cast(p.premium * 0.1 AS MONEY) < cast(1000000 AS MONEY)
    			ORDER BY produksi_agent ASC