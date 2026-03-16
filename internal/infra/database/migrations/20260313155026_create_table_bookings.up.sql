CREATE TABLE bookings (
    id  SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    car_id      INT NOT NULL,
    start_rent  TIMESTAMP WITH TIME ZONE NOT NULL,
    end_rent    TIMESTAMP WITH TIME ZONE NOT NULL,
    total_cost  BIGINT NOT NULL,
    finished    BOOLEAN DEFAULT FALSE,

    CONSTRAINT fk_customer FOREIGN KEY (customer_id) 
        REFERENCES customers(id) ON DELETE CASCADE,
    CONSTRAINT fk_car FOREIGN KEY (car_id) 
        REFERENCES cars(id) ON DELETE CASCADE
);

CREATE INDEX idx_bookings_customer ON bookings(customer_id);
CREATE INDEX idx_bookings_car ON bookings(car_id);