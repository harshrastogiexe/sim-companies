USE SIMCOMPANIES
GO

SELECT name, sold_at_building_id, produced_at_building_id
FROM resources
SELECT *
FROM produced_froms
SELECT *
FROM buildings

DELETE FROM resources WHERE id=3

SELECT r.name, sold_at_building_id, b.name [Produced At]
FROM resources r
    JOIN buildings b ON b.id = produced_at_building_id

--
SELECT
    r.name as [Resource],
    f.name as [From],
    amount as [Amount]
FROM produced_froms pr
    JOIN resources r on r.id = pr.resource_id
    JOIN resources f on pr.produced_from_resource_id = f.id


SET IDENTITY_INSERT "resources" ON;
INSERT INTO "resources"
    ("name","image","transportation","retailable","research","exchange_tradable","realm_available","sold_at_building_id","sold_at_restaurant_id","produced_at_building_id","transport_needed","produced_an_hour","base_salary","average_retail_price","market_saturation","market_saturation_label","retail_modeling","store_base_salary","id")
VALUES
    ('Apples', 'images/resources/apples.png', 1.000000, 1, 0, 1, 1, NULL, NULL, 'P', 1.000000, 202.188442, 103.500000, NULL, NULL, NULL, NULL, NULL, 3);
SET IDENTITY_INSERT "resources" OFF
