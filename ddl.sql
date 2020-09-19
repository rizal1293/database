CREATE DATABASE `sales` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

-- sales.Costumer definition

CREATE TABLE `Customer` (
  `CustomerID` bigint(20) NOT NULL,
  `CompanyName` varchar(50) DEFAULT NULL,
  `FirstName` varchar(30) DEFAULT NULL,
  `LastName` varchar(50) DEFAULT NULL,
  `BillingAddress` varchar(255) DEFAULT NULL,
  `City` varchar(50) DEFAULT NULL,
  `StateOrProvince` varchar(20) DEFAULT NULL,
  `ZIPCode` varchar(20) DEFAULT NULL,
  `Email` varchar(75) DEFAULT NULL,
  `CompanyWebsite` varchar(200) DEFAULT NULL,
  `PhoneNumber` varchar(30) DEFAULT NULL,
  `FaxNumber` varchar(30) DEFAULT NULL,
  `ShipAddress` varchar(255) DEFAULT NULL,
  `ShipCity` varchar(50) DEFAULT NULL,
  `ShipStateOrProvince` varchar(50) DEFAULT NULL,
  `ShipZIPCode` varchar(20) DEFAULT NULL,
  `ShipPhoneNumber` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`CostumerID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- sales.Employees definition

CREATE TABLE `Employees` (
  `EmployeeID` bigint(20) NOT NULL,
  `FisrtName` varchar(50) DEFAULT NULL,
  `LastName` varchar(50) DEFAULT NULL,
  `Title` varchar(50) DEFAULT NULL,
  `WorkPhone` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`EmployeeID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- sales.Shipping_Methods definition

CREATE TABLE `Shipping_Methods` (
  `ShippingMethodID` bigint(20) NOT NULL,
  `ShippingMethod` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`ShippingMethodID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- sales.Products definition

CREATE TABLE `Products` (
  `ProductID` bigint(20) NOT NULL,
  `ProductName` varchar(50) DEFAULT NULL,
  `UnitPrice` int(11) DEFAULT NULL,
  `InStock` char(1) DEFAULT NULL,
  PRIMARY KEY (`ProductID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- sales.Orders definition

CREATE TABLE `Orders` (
  `OrderID` bigint(20) NOT NULL,
  `CustomerID` bigint(20) DEFAULT NULL,
  `EmployeeID` bigint(20) DEFAULT NULL,
  `OrderDate` date DEFAULT NULL,
  `PurchaseOrderNumber` varchar(30) DEFAULT NULL,
  `ShipDate` date DEFAULT NULL,
  `ShippingMethodID` bigint(20) DEFAULT NULL,
  `FreightCharge` int(11) DEFAULT NULL,
  `Taxes` int(11) DEFAULT NULL,
  `PaymentReceived` char(1) DEFAULT NULL,
  `Comment` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`OrderID`),
  KEY `CustomerID` (`CustomerID`),
  KEY `EmployeeID` (`EmployeeID`),
  KEY `ShippingMethodID` (`ShippingMethodID`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`CustomerID`) REFERENCES `Customers` (`CustomerID`),
  CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`EmployeeID`) REFERENCES `Employees` (`EmployeeID`),
  CONSTRAINT `orders_ibfk_3` FOREIGN KEY (`ShippingMethodID`) REFERENCES `Shipping_Methods` (`ShippingMethodID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- sales.Order_Details definition

CREATE TABLE `Order_Details` (
  `OrderDetailID` bigint(20) NOT NULL,
  `OrderID` bigint(20) DEFAULT NULL,
  `ProductID` bigint(50) DEFAULT NULL,
  `Quantity` int(11) DEFAULT NULL,
  `UnitPrice` int(11) DEFAULT NULL,
  `Discount` int(11) DEFAULT NULL,
  PRIMARY KEY (`OrderDetailID`),
  KEY `OrderID` (`OrderID`),
  KEY `ProductID` (`ProductID`),
  CONSTRAINT `order_details_ibfk_1` FOREIGN KEY (`OrderID`) REFERENCES `Orders` (`OrderID`),
  CONSTRAINT `order_details_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `Products` (`ProductID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;