USE logistics;

PREPARE i1 FROM "select id_container, to_country, create_date from container where status='1'";

PREPARE i2 FROM "select id_container, to_country, curr_weight, max_weight from container where status='1'";

PREPARE s1 FROM

PREPARE p1 FROM

PREPARE p2 FROM