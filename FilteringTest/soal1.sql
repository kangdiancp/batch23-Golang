create dataase soal1
CREATE TABLE T_AGENT(
    AGENT_CODE VARCHAR(10) PRIMARY KEY,
    AGENT_NAME VARCHAR(20),
    AGENT_OFFICE VARCHAR(20),
    BASIC_COMMISION INTEGER,

)

insert into t_agent(agent_code, agent_name, agent_office) values
('AG001', 'LIDYA', 'JAKARTA'),
('AG002', 'RUDI', 'BALI'),
('AG003', 'DICKY', 'BALI'),
('AG004', 'DAVID', 'SURABAYA'),
('AG005', 'FARIZ', 'SURABAYA')

CREATE TABLE T_CLIENT(
    CLIENT_NUMBER VARCHAR(10) PRIMARY KEY,
    CLIENT_NAME VARCHAR(25),
    BIRTH_DATE DATE,
)
insert into t_client(client_number, client_name, birth_date) values
('CL001', 'WAYNE ROONEY', '1956-05-11'),
('CL002', 'RYAN GIGGS', '1972-09-03'), 
('CL003', 'DAVID BECKHAM', '1985-09-03'),
('CL004', 'MICHAEL OWEN', '2012-09-03')

create table T_POLICY (
    POLICY_NUMBER serial PRIMARY KEY,
    POLICY_SUBMIT_DATE DATE,
    PREMIUM MONEY,
    DISCOUNT INTEGER,
    COMMISION MONEY,
    CLIENT_NUMBER VARCHAR(10),
    AGENT_CODE VARCHAR(10),
    POLICY_STATUS VARCHAR(20),
    POLICY_DUE_DATE DATE,
    FOREIGN KEY (AGENT_CODE) REFERENCES T_AGENT(AGENT_CODE),
    FOREIGN KEY (CLIENT_NUMBER) REFERENCES T_CLIENT(CLIENT_NUMBER),


)

insert into t_policy(policy_number, policy_submit_date, premium, discount, commision, client_number, agent_code, policy_status)
values 
(001, '2018-01-05', 18600000, 10, 930000, 'CL001', 'AG001', 'INFORCE'),
(002, '2018-01-05', 2500000, 10, 125000, 'CL002', 'AG002', 'INFORCE'),
(003, '2018-01-10', 2500000, 10, 125000, 'CL003', 'AG003', 'TERMINATE'),
(004, '2018-01-25', 4000000, 10, 200000, 'CL003', 'AG002', 'PROPOSAL'),
(005, '2018-01-25', 2625000, 10, 131250, 'CL004', 'AG002', 'PROPOSAL')


-- nomor a
select p.policy_number, p.policy_submit_date,p.premium, p.discount,p.commision, p.client_number, p.agent_code, p.policy_status,p.policy_due_date 
from t_policy as p 
left join t_client as c 
on p.client_number = c.client_number 
where p.policy_submit_date > '2018-01-15' AND extract(MONTH from c.birth_date)=9
group by p.policy_number

--nomor b
select p.policy_number, p.policy_submit_date,p.premium, p.discount,p.commision, p.client_number, p.agent_code, p.policy_status,p.policy_due_date
from t_policy as p 
left join t_agent as a on p.agent_code = a.agent_code
where p.policy_status = 'INFORCE' AND a.agent_office = 'JAKARTA'
group by p.policy_number

--nomor c
update t_agent 
set basic_commision = (p.commision/p.premium) * 100
from t_policy as p
where t_agent.agent_code = p.agent_code;

-- nomor d
update t_policy
set policy_due_date = date_trunc('month', policy_submit_date + interval '30 days') +interval '1 month' - interval '1 days'



-- nomor e
select a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 as money) as produksi_agent
from t_agent as a 
join t_policy as p on p.agent_code = a.agent_code
where cast(p.premium * 0.1 as money) < cast(1000000 as money)
order by produksi_agent asc