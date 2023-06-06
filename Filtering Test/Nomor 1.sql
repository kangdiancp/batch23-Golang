-- A
select p.* from policy p join client c on p.client_number = c.client_number
where p.policy_submit_date > '2018-01-15' and EXTRACT (MONTH FROM c.birth_date) = 9

-- B
select p.* from policy p join agent a on p.agent_code = a.agent_code
where p.policy_status = 'INFORCE' and a.agent_office ='JAKARTA'

-- C
update agent a set basic_commision = (p.commision / p.premium) * 100::money  from policy p
where a.agent_code = p.agent_code
select * from agent

-- D
update policy set policy_due_date = DATE_TRUNC('MONTH', policy_submit_date) + INTERVAL '1 MONTH' - INTERVAL '1 day'
where policy_due_date is NULL
select * from policy

-- E
select * from policy
where cast(premium as numeric) / cast(discount as numeric) < 1000000
order by cast(premium as numeric) / cast(discount as numeric) asc




