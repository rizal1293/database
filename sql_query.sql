/* A. List of customers located in irvine city */
SELECT 
  CustomerID, 
  CompanyName, 
  FisrtName, 
  LastName, 
  BillingAddress, 
  City, 
  StateOrProvince, 
  ZIPCode, 
  Email, 
  CompanyWebsite, 
  PhoneNumber, 
  FaxNumber, 
  ShipAddress, 
  ShipCity, 
  ShipStateOrProvince, 
  ShipZIPCode, 
  ShipPhoneNumber 
FROM 
  Customers 
WHERE 
  City = 'Irvine'

/** end */

/* List of customers whose order is handled by an employee named Adam Barr */

SELECT 
  customer.CustomerID, 
  customer.CompanyName, 
  customer.FisrtName, 
  customer.LastName, 
  customer.BillingAddress, 
  customer.City, 
  customer.StateOrProvince, 
  customer.ZIPCode, 
  customer.Email, 
  customer.CompanyWebsite, 
  customer.PhoneNumber, 
  customer.FaxNumber, 
  customer.ShipAddress, 
  customer.ShipCity, 
  customer.ShipStateOrProvince, 
  customer.ShipZIPCode, 
  customer.ShipPhoneNumber 
FROM 
  Customers customer 
  INNER JOIN Orders orders ON (
    orders.CustomerID = customer.CustomerID
  ) 
  INNER JOIN Employees employees ON (
    employees.EmployeeID = orders.EmployeeID
  ) 
WHERE 
  employees.FisrtName = 'Adam' 
  and employees.LastName = 'Barr'

/** end */

/** List of products which are ordered by "Contoso, Ltd*/

SELECT 
  p.ProductID, 
  p.ProductName, 
  p.UnitPrice, 
  p.InStock 
FROM 
  Products p 
  inner join Order_Details od on (od.ProductID = p.ProductID) 
  inner join Orders o2 on (o2.OrderID = od.ProductID) 
  inner join Customers c2 on (c2.CustomerID = o2.CustomerID) 
where 
  c2.CompanyName = 'Contoso, Ltd'

/** List transaction (orders) which has "UPS Ground" as shipping method */

SELECT 
  o2.OrderID, 
  o2.CustomerID, 
  o2.EmployeeID, 
  o2.OrderDate, 
  o2.PurchaseOrderNumber, 
  o2.ShipDate, 
  o2.ShippingMethodID, 
  o2.FreightCharge, 
  o2.Taxes, 
  o2.PaymentReceived, 
  o2.Comment 
FROM 
  Orders o2 
  inner join Shipping_Methods sm on (
    sm.ShippingMethodID = o2.ShippingMethodID
  ) 
where 
  sm.ShippingMethod = "UPS Ground"

/** end */

/** List of total cost (including tax and freight charge) for every order sorted by ship date */

SELECT
  od.OrderID ,
  sum(
    (p.UnitPrice * od.Quantity) + (o.FreightCharge + o.Taxes)
  ) TotalCost
FROM 
  Order_Details od 
  inner join Products p ON p.ProductID = od.ProductID 
  inner join Orders o on o.OrderID = od.OrderID 
where 
  o.ShipDate is not NULL 
group by 
  od.OrderID
order by 
  o.ShipDate asc

/** end */