SELECT
	Product.pkProductID,
	Product.ARKID,
	Product.LaunchDate,
	rPdctName.Value AS ProductName,
	rCodeName.Value AS CodeName 
FROM
	Product
	LEFT JOIN Resource AS rPdctName ON Product.fkProductNameID = rPdctName.fkResourceID
	LEFT JOIN Resource AS rCodeName ON Product.fkCodeNameID = rPdctName.fkResourceID 
ORDER BY
	Product.pkProductID 
	LIMIT 100