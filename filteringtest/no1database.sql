create table T_Policy (
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
create table T_Agent(
	agent_code varchar(20),
	agent_name varchar(50),
	agent_office varchar(30),
	basic_commission money,
	foreign key (agent_code) references t_policy(agent_code)
)
create table T_Client(
	client_number varchar(20),
	client_name varchar(50),
	birth_date date,
	foreign key (client_number) references t_policy(client_number)
)

insert into T_Policy (policy_number, policy_submit_date, premium, discount, commission, client_number, agent_code, policy_status, policy_due_date) values 
(001, '2018-1-5', 18600000, 10, 930000, 'CL001', 'AG001', 'INFORCE', null),
(002, '2018-1-5', 2500000, 10, 125000, 'CL002', 'AG002', 'INFORCE', null),
(003, '2018-1-10', 2500000, 10, 125000, 'CL003', 'AG003', 'TERMINATE', null),
(004, '2018-1-25', 4000000, 10, 200000, 'CL004', 'AG004', 'PROPOSAL', null),
(005, '2018-1-25', 2625000, 10, 131250, 'CL005', 'AG005', 'PROPOSAL', null)

insert into T_Agent (agent_code, agent_name, agent_office, basic_commission) values
('AG001', 'LIDYA', 'JAKARTA', NULL),
('AG002', 'RUDI', 'BALI', NULL),
('AG003', 'DICKY', 'BALI', NULL),
('AG004', 'DAVID', 'SURABAYA', NULL),
('AG005', 'FARIZ', 'SURABAYA', NULL)

insert into T_Client (client_number, client_name, birth_date) values
('CL001', 'WAYNE ROONEY', '1956-5-11'),
('CL002', 'RYAN GIGGS', '1972-9-3'),
('CL003', 'DAVID BECKHAM', '1985-9-3'),
('CL004', 'MICHAEL OWEN', '2012-9-3')


//Nomor 1a
Select T_Policy * from T_Policy
Left join T_Client on T_Policy.client_number = T_Client.client_number
where T_Policy.policy_submit_date > '2018-01-15' and extract(month from T_Client.birth_) = 9;

//Nomor 1b
select T_Policy * from T_Policy
left join T_Agent on T_Policy.agent_code = T_Agent.agent_code
where T_Policy.policy_status = 'INFORCE' and T_Agent.agent_office = 'JAKARTA';

//Nomor 1c
select T_Agent , (T_Policy.commission / T_Policy.premium) * 100 as basic_commission
from T_Agent
left join T_Policy on T_Agent.agent_code = T_Policy.agent_code;

//Nomor 1d
update T_Policy
set policy_due_date = date_trunc('month', T_Policy.policy_submit_date) + interval '1 month' - interval '1 day' + interval '30 days';

select * from T_Policy

//Nomor 1e
select a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 as MONEY) as produksi_agent
from T_Agent as a 
left join T_Policy as p on p.agent_code = a.agent_code
where cast(p.premium * 0.1 as MONEY) < cast(1000000 AS MONEY)
order by produksi_agent asc