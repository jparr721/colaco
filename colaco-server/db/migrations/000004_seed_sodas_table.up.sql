BEGIN;

INSERT INTO sodas (product_name, description, cost, current_quantity, max_quantity)
VALUES ('Fizz', 'An effervescent fruity experience with hints of grape and coriander.', 1, 100, 100),
       ('Pop', 'An explosion of flavor that will knock your socks off!', 1, 100, 100),
       ('Cola', 'A basic no nonsense cola that is the perfect pick me up for any occasion.', 1, 200, 200),
       ('Mega Pop', 'Not for the faint of heart. So flavorful and so invigorating, it should probably be illegal.', 1, 50, 50);

COMMIT;