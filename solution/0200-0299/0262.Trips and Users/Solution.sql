# Write your MySQL query statement below
select
    request_at as Day,
    round(avg(status != 'completed'), 2) as 'Cancellation Rate'
from
    Trips as t
    join Users as u1 on (t.client_id = u1.users_id and u1.banned = 'No')
    join Users as u2 on (t.driver_id = u2.users_id and u2.banned = 'No')
where request_at between '2013-10-01' and '2013-10-03'
group by request_at;
