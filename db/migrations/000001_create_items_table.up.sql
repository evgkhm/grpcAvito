CREATE TABLE IF NOT EXISTS usr(
  id INTEGER PRIMARY KEY,
  balance DECIMAL
);

CREATE TABLE IF NOT EXISTS reservation(
  id INTEGER,
  service_id INTEGER,
  order_id INTEGER,
  cost DECIMAL,
  PRIMARY KEY (id, service_id, order_id)
);

CREATE TABLE IF NOT EXISTS revenue(
  id INTEGER,
  service_id INTEGER,
  order_id INTEGER,
  cost DECIMAL,
  curr_date DATE NOT NULL DEFAULT CURRENT_DATE
);