select distinct user_id as id, name
from orders o
         join users u on u.id = o.user_id
order by name, user_id;
