DECLARE @REQUIRED_RATE AS INT = 1600

USE SIMCOMPANIES
GO
SELECT
    [name] [Name],
    [produced_an_hour] [Produced Per Hour Per Level],
    ([bm].[wages]/[produced_an_hour]) [Worker Cost Per Unit],
    [resource].[base_salary] AS [VALUE]
FROM
    resource_main [resource]

    JOIN resource_base base
    ON [base].[id]=[resource].[resource_id]

    JOIN building_main bm
    ON bm.building_id=resource.produced_at_building_id

WHERE [name] IN ('Power', 'Silicon', 'Minerals', 'Sand', 'Chemicals', 'Processors', 'Electronic components', 'Displays')
ORDER BY [Name]
GO


-- SELECT name, wages
-- FROM building_main, building_base
-- WHERE building_base.id=building_main.building_id
-- ORDER BY name

-- SELECT 414/2440.88
