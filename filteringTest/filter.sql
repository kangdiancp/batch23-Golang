create table t_agent (
	agent_code varchar(10) primary key,
	agent_name varchar(30),
	agent_office varchar(20),
	basic_commission integer
)

create table t_client (
	client_number varchar primary key,
	client_name varchar(30),
	birth_date date
)

create table t_policy (
	policy_number serial primary key,
	policy_submit_date date,
	premium money,
	discount integer,
	commission money,
	client_number varchar(10),
	agent_code varchar(10),
	policy_status char(50),
	policy_due_date date,
	foreign key (client_number) references t_client(client_number),
	foreign key (agent_code) references t_agent(agent_code)
)

Select * from t_agent
select * from t_client
select * from t_policy

insert into t_agent(agent_code,agent_name,agent_office,basic_commission) values
('AG001','LIDYA','JAKARTA',null),
('AG002','RUDI','BALI',null),
('AG003','DICKY','BALI',null),
('AG004','DAVID','SURABAYA',null),
('AG005','FARIZ','SURABAYA',null)

insert into t_client(client_number,client_name,birth_date) values
('CL001','WAYNE ROONEY','11/5/1959'),
('CL002','RYAN GIGGS','3/9/1972'),
('CL003','DAVID BECKHAM','3/9/1985'),
('CL004','MICHAEL OWEN','3/9/2012')

insert into t_policy(policy_number,policy_submit_date,premium,discount,commission,client_number,agent_code,policy_status,policy_due_date) values
('001','5/1/2018','18600000','10','930000','CL001','AG001','INFORCE',null),
('002','5/1/2018','2500000','10','125000','CL002','AG002','INFORCE',null),
('003','10/1/2018','2500000','10','125000','CL003','AG003','TERMINATE',null),
('004','25/1/2018','4000000','10','200000','CL003','AG002','PROPOSAL',null),
('005','25/1/2018','2625000','10','131250','CL004','AG002','PROPOSAL',null)

//nomor 1a
select * from t_policy p join t_client c on p.client_number = c.client_number 
where policy_submit_date > '15/1/2018' and extract (month from birth_date)=9

//nomor 1b
select p.policy_number, p.policy_submit_date,p.premium, p.discount,p.commission, 
p.client_number, p.agent_code, p.policy_status,p.policy_due_date
from t_policy p join t_agent a on p.agent_code=a.agent_code 
where p.policy_status='INFORCE' and a.agent_office='JAKARTA' group by p.policy_number

//nomor 1c
update t_agent a set basic_commission=(p.commission/p.premium)*100 from t_policy p
where p.agent_code=a.agent_code

//nomor 1d
update t_policy set policy_due_date = date_trunc('month', policy_submit_date + interval '30 days') 
+interval '1 month' - interval '1 days'

//nomor 1e
select a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 as money) produksi_agent
from t_agent a join t_policy p on p.agent_code = a.agent_code
where cast(p.premium * 0.1 as money) < cast(1000000 as money)
order by produksi_agent asc