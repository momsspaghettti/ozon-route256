select ff.id                                  as fulfilment_id,
       ff.name                                as fulfilment_name,
       p.id                                   as package_id,
       p.name                                 as package_name,
       t.arrival_time                         as time_in,
       t_in.departure_time                    as time_out,
       (t_in.departure_time - t.arrival_time) as storage_time
from transportations t
         inner join fulfilments ff on t.destination_id = ff.id
         inner join packages p on t.package_id = p.id
         inner join transportations t_in on t.destination_id = t_in.source_id and
                                            t.package_id = t_in.package_id
order by storage_time desc, time_in, fulfilment_id, package_id;
