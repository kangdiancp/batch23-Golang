create table T_AGENT (
	agent_code varchar primary key,
	agent_name varchar,
	agent_office varchar,
	basic_commission money
)
insert into t_agent (agent_code, agent_name, agent_office) values
	('AG001', 'LIDYA', 'JAKARTA'),
	('AG002', 'RUDI', 'BALI'),
	('AG003', 'DICKY', 'BALI'),
	('AG004', 'DAVID', 'SURABAYA'),
	('AG005', 'FARIZ', 'SURABAYA');
	
	
create table T_CLIENT(
	client_number varchar primary key,
	client_name varchar,
	birth_date timestamp
)
insert into t_client (client_number, client_name, birth_date) values
	('CL001', 'WAYNE ROONEY', '11/5/1956'),
	('CL002', 'RYAN GIGGS', '3/9/1972'),
	('CL003', 'DAVID BECKAM', '3/9/1985'),
	('CL004', 'MICHAEL OWEN', '3/9/2012');
drop table t_client cascade

create table T_POLICY (
	policy_number integer primary key,
	policy_submit_date timestamp,
	premium money,
	discount integer,
	commission money,
	client_number varchar,
	agent_code varchar,
	policy_status varchar,
	policy_due_date timestamp,
	foreign key (agent_code) references T_AGENT (agent_code),
	foreign key (client_number) references T_CLIENT (client_number)
)
insert into t_policy values
	(001, '1/5/2018', 18600000.00, 10, 930000.00, 'CL001', 'AG001', 'INFORCE'),
	(002, '1/5/2018', 2500000.00, 10, 125000.00, 'CL002', 'AG002', 'INFORCE'),
	(003, '1/10/2018', 2500000.00, 10, 125000.00, 'CL003', 'AG003', 'TERMINATE'),
	(004, '1/25/2018', 4000000.00, 10, 200000.00, 'CL003', 'AG002', 'PROPOSAL'),
	(005, '1/25/2018', 2625000.00, 10, 131250.00, 'CL004', 'AG002', 'PROPOSAL');
	
select * from t_client

-- Nomor 1 soal a tampilkan data polis yang di submit setelah tanggal 15 januari 2018 berdasarkan client yang lahir di bulan september
select * from t_policy p
join t_client as c on c.client_number = p.client_number  
where policy_submit_date > '1/15/2018'
order by birth_date


-- Nomor 1 soal b
select * from t_policy p
join t_agent a on a.agent_code = p.agent_code
where p.policy_status = 'INFORCE' and a.agent_office = 'JAKARTA'

-- Nomor 1 Soal c
select a.agent_code, a.agent_name, a.agent_office, ((p.commission / p.premium) * 100) as basic_commission 
from t_agent a
join t_policy p on p.agent_code = a.agent_code
order by a.agent_code asc

-- Nomor 1 soal d
update t_policy
set policy_due_date =  policy_submit_date + interval '30 days'

select * from t_policy

-- Nomor 1 soal e
select a.agent_code, a.agent_name, a.agent_office, p.premium - cast(premium * 0.1 as money) as produksi from t_agent a
join t_policy p on p.agent_code = a.agent_code 
where cast(premium * 0.1 as money) < cast (1000000 as money)
order by premium asc