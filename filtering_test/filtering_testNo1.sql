create table t_policy(
	policy_number integer primary key,
	policy_submit_date timestamp,
	premium money,
	discount smallint,
	commision money,
	client_number varchar(5),
	agent_code varchar(5),
	policy_status varchar(15),
	policy_due_date timestamp,
	foreign key(agent_code) references t_agent(agent_code),
	foreign key(client_number) references t_client(client_number)
)

insert into t_policy(policy_number, policy_submit_date, premium, discount, commision, client_number, agent_code, policy_status)
values 
(001, '2018-01-05', 18600000, 10, 930000, 'CL001', 'AG001', 'INFORCE'),
(002, '2018-01-05', 2500000, 10, 125000, 'CL002', 'AG002', 'INFORCE'),
(003, '2018-01-10', 2500000, 10, 125000, 'CL003', 'AG003', 'TERMINATE'),
(004, '2018-01-25', 4000000, 10, 200000, 'CL003', 'AG002', 'PROPOSAL'),
(005, '2018-01-25', 2625000, 10, 131250, 'CL004', 'AG002', 'PROPOSAL')

drop table t_policy
select * from t_policy
create table t_agent(
	agent_code varchar(5) primary key,
	agent_name varchar(10),
	agent_office varchar(15),
	basic_commision integer
)
insert into t_agent(agent_code, agent_name, agent_office) values
('AG001', 'LIDYA', 'JAKARTA'),
('AG002', 'RUDI', 'BALI'),
('AG003', 'DICKY', 'BALI'),
('AG004', 'DAVID', 'SURABAYA'),
('AG005', 'FARIZ', 'SURABAYA')

select * from t_agent
create table t_client(
	client_number varchar(5) primary key,
	client_name varchar(30),
	birth_date timestamp
)
insert into t_client(client_number, client_name, birth_date) values
('CL001', 'WAYNE ROONEY', '1956-05-11'),
('CL002', 'RYAN GIGGS', '1972-09-03'), 
('CL003', 'DAVID BECKHAM', '1985-09-03'),
('CL004', 'MICHAEL OWEN', '2012-09-03')

select * from t_client

-- nomor1
select p.policy_number, p.policy_submit_date, p.client_number, p.agent_code, p.policy_status from t_policy as p left join t_client as c 
on p.client_number = c.client_number 
where p.policy_submit_date > '2018-01-15' AND extract(MONTH from c.birth_date)=09
group by p.policy_number

--nomor 2
select p.policy_number, p.policy_submit_date, p.client_number, p.agent_code, p.policy_status from t_policy as p left join t_agent as a 
on p.agent_code = a.agent_code
where p.policy_status = 'INFORCE' AND a.agent_office = 'JAKARTA'
group by p.policy_number

--nomor 3
update t_agent 
set basic_commision = (p.commision/p.premium) * 100
from t_policy p
where t_agent.agent_code = p.agent_code;

-- nomor 4
update t_policy
set policy_due_date = date_trunc('day', policy_submit_date + justify_days(interval '30 days')) - interval '1 days'

-- Nomor 5
select a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 as money) as produksi_agent
from t_agent as a join t_policy as p on p.agent_code = a.agent_code
where cast(p.premium * 0.1 as money) < cast(1000000 as money)
order by produksi_agent asc