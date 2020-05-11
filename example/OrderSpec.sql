SELECT
	Product.pkProductID,
	ProductOrderSpec.pkProductOrderSpecID,
	rOrderSpecName.Value AS OrderSpecName,
	rRcp.Value AS Rcp,
	rVtx.Value AS Vtx,
	rEccn.Value AS Eccn,
	rStep.Value AS Step,
	rCcats.Value AS Ccats,
	rUsHts.Value AS UsHts,
	rSocket.Value AS Socket,
	rStepTdp.Value AS StepTdp,
	rSpecCode.Value AS ASSpecCode,
	rOrderingCode.Value AS OrderingCode
FROM
	Product
	LEFT JOIN ProductOrderSpec ON Product.pkProductID = ProductOrderSpec.fkProductID
	LEFT JOIN Resource AS rOrderSpecName ON ProductOrderSpec.fkNameID = rOrderSpecName.fkResourceID
	LEFT JOIN Resource AS rRcp ON ProductOrderSpec.fkRcpResourceID = rRcp.fkResourceID
	LEFT JOIN Resource AS rVtx ON ProductOrderSpec.fkVtxResourceID = rVtx.fkResourceID
	LEFT JOIN Resource AS rEccn ON ProductOrderSpec.fkEccnResourceID = rEccn.fkResourceID
	LEFT JOIN Resource AS rStep ON ProductOrderSpec.fkStepResourceID = rStep.fkResourceID
	LEFT JOIN Resource AS rCcats ON ProductOrderSpec.fkCcatsResourceID = rCcats.fkResourceID
	LEFT JOIN Resource AS rUsHts ON ProductOrderSpec.fkUsHtsResourceID = rUsHts.fkResourceID
	LEFT JOIN Resource AS rSocket ON ProductOrderSpec.fkSocketResourceID = rSocket.fkResourceID
	LEFT JOIN Resource AS rStepTdp ON ProductOrderSpec.fkStepTdpResourceID = rStepTdp.fkResourceID
	LEFT JOIN Resource AS rSpecCode ON ProductOrderSpec.fkSpecCodeResourceID = rSpecCode.fkResourceID
	LEFT JOIN Resource AS rOrderingCode ON ProductOrderSpec.fkOrderingCodeResourceID = rOrderingCode.fkResourceID 
WHERE
	Product.ARKID = 97144 
ORDER BY
	ProductOrderSpec.pkProductOrderSpecID