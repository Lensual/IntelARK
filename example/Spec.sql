SELECT
	Product.pkProductID,
	Specification.pkSpecificationID,
	rSpecName.Value AS SpecName,
	rSpecVal.Value AS SpecValue 
FROM
	Product
	LEFT JOIN ProductSpecificationLink ON Product.pkProductID = ProductSpecificationLink.fkProductID
	LEFT JOIN Specification ON ProductSpecificationLink.fkSpecificationID = Specification.pkSpecificationID
	LEFT JOIN Resource AS rSpecName ON Specification.fkNameID = rSpecName.fkResourceID
	LEFT JOIN Resource AS rSpecVal ON Specification.fkValueID = rSpecVal.fkResourceID 
WHERE
	Product.ARKID = 97144 
ORDER BY
	Specification.pkSpecificationID