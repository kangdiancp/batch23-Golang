create table t_policy (
	policy_number integer,
	policy_submit_date date,
	premium money,
	discount integer,
	commission money,
	client_number varchar(20) unique,
	agent_code varchar(20) unique,
	policy_status varchar(25),
	policy_due_date date,
	constraint t_policy_pk primary key (policy_number, client_number, agent_code)
)
insert into t_policy (policy_number, policy_submit_date, premium, discount, commission, client_number, agent_code, policy_status, policy_due_date) values 
(001, '2018-1-5', 18600000, 10, 930000, 'CL001', 'AG001', 'INFORCE', null),
(002, '2018-1-5', 2500000, 10, 125000, 'CL002', 'AG002', 'INFORCE', null),
(003, '2018-1-10', 2500000, 10, 125000, 'CL003', 'AG003', 'TERMINATE', null),
(004, '2018-1-25', 4000000, 10, 200000, 'CL004', 'AG004', 'PROPOSAL', null),
(005, '2018-1-25', 2625000, 10, 131250, 'CL005', 'AG005', 'PROPOSAL', null)

create table t_agent(
	agent_code varchar(20),
	agent_name varchar(50),
	agent_office varchar(30),
	basic_commission money,
	foreign key (agent_code) references t_policy(agent_code)
)
insert into t_agent (agent_code, agent_name, agent_office, basic_commission) values
('AG001', 'LIDYA', 'JAKARTA', NULL),
('AG002', 'RUDI', 'BALI', NULL),
('AG003', 'DICKY', 'BALI', NULL),
('AG004', 'DAVID', 'SURABAYA', NULL),
('AG005', 'FARIZ', 'SURABAYA', NULL)

create table t_client(
	client_number varchar(20),
	client_name varchar(50),
	birth_date date,
	foreign key (client_number) references t_policy(client_number)
)
insert into t_client (client_number, client_name, birth_date) values
('CL001', 'WAYNE ROONEY', '1956-5-11'),
('CL002', 'RYAN GIGGS', '1972-9-3'),
('CL003', 'DAVID BECKHAM', '1985-9-3'),
('CL004', 'MICHAEL OWEN', '2012-9-3')


---  a
select p.* from t_policy p 
join t_client c on p.client_number = c.client_number
where p.policy_submit_date > '2018-01-15' 

---b
SELECT * FROM t_policy a
	JOIN t_agent b
	ON a.agent_code = b.agent_code
	WHERE policy_status = 'INFORCE' AND agent_office = 'JAKARTA';

---c
SELECT a.agent_code, (b.commission/b.premium * 100) AS basic_commission
    FROM t_agent a
    JOIN t_policy b ON a.agent_code = b.agent_code;

--d
UPDATE t_policy
SET policy_due_date = DATE_TRUNC('MONTH', policy_submit_date + INTERVAL '30 days') + INTERVAL '1 MONTH' - INTERVAL '1 DAY';

--e
SELECT a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 AS MONEY) AS produksi_agent
    FROM t_agent AS a 
    JOIN t_policy as p ON p.agent_code = a.agent_code
    WHERE cast(p.premium * 0.1 AS MONEY) < cast(1000000 AS MONEY)
 
	
	
select * from t_policy 
select * from t_agent	
select * from t_client






































