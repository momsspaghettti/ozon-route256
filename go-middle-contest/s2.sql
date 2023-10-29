select ff.id                                      AS fulfilment_id,
       ff.name                                    AS fulfilment_name,
       p.id                                       AS package_id,
       p.name                                     AS package_name,
       t_in.arrival_time                          AS time_in,
       t_out.departure_time                       AS time_out,
       (t_out.departure_time - t_in.arrival_time) AS storage_time
from fulfilments ff
         JOIN transportations t_in ON ff.id = t_in.destination_id
         JOIN transportations t_out ON ff.id = t_out.source_id
         JOIN packages p ON t_in.package_id = p.id
WHERE t_in.package_id = t_out.package_id
  AND t_in.arrival_time <= t_out.departure_time
  AND NOT EXISTS(
        SELECT 1
        FROM transportations t1
        WHERE t1.package_id = t_in.package_id
          AND t1.source_id = ff.id
          AND t1.departure_time > t_in.arrival_time
          AND t1.departure_time < t_out.departure_time
    )
  AND NOT EXISTS(
        SELECT 1
        FROM transportations t2
        WHERE t2.package_id = t_in.package_id
          AND t2.destination_id = ff.id
          AND t2.arrival_time > t_in.arrival_time
          AND t2.arrival_time < t_out.departure_time
    )
order BY storage_time DESC, time_in, fulfilment_id, package_id;
