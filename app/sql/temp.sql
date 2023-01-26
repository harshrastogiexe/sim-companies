use SIMCOMPANIES
GO

SELECT * FROM [Resource]


CREATE TABLE [Building] (
    id NVARCHAR(10) NOT NULL PRIMARY KEY,
    [name] NVARCHAR(100) NOT NULL,
    [image] NVARCHAR(100) NOT NULL,
    [category] NVARCHAR(100) NOT NULL,
    [cost] INT NOT NULL,
    [robots_needed] FLOAT NOT NULL,
    [realm_available] BIT NOT NULL
)

CREATE TABLE [BuildingImages] (
    [id] NVARCHAR(10) NOT NULL ,
    [level] INT NOT NULL,
    [path] NVARCHAR(100) NOT NULL,
    FOREIGN KEY (id) REFERENCES Building(id),
    CONSTRAINT UNI_REF UNIQUE(id, [level])
)

INSERT INTO [BuildingImages]
VALUES
('100', 1, 'path')

SELECT * FROM [BuildingImages]
SELECT * FROM [Building]
-- DELETE FROM Building

SELECT img.* FROM BuildingImages [img], Building[b]
WHERE img.[id] = b.[id] AND b.[id]=5

DROP TABLE [Building]

EXEC sp_rename [Resourse.RealmAvailable], [realmAvailable], [COLUMN]
