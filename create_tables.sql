CREATE TABLE "Visitor"
(
    "Login"          varchar PRIMARY KEY,
    "Password"       varchar NOT NULL,
    "FirstName"      varchar NOT NULL,
    "LastName"       varchar NOT NULL,
    "Patronymic"     varchar NOT NULL,
    "Phone"          varchar NOT NULL,
    "Email"          varchar NOT NULL,
    "Organization"   varchar NOT NULL,
    "Note"           varchar NOT NULL,
    "Birthday"       date,
    "PassportSeries" varchar(4) NOT NULL,
    "PassportNumber" varchar(6) NOT NULL,
    "Photo"          varchar NOT NULL
);

CREATE TABLE "Branch"
(
    "Name" varchar PRIMARY KEY
);

CREATE TABLE "Employee"
(
    "ID" varchar PRIMARY KEY
);

CREATE TABLE "OrderType"
(
    "Name" varchar PRIMARY KEY
);

CREATE TABLE "OrderStatus"
(
    "Name" varchar PRIMARY KEY
);

CREATE TABLE "Order"
(
    "ID"           serial PRIMARY KEY,
    "BeginDate"    timestamp                               NOT NULL,
    "EndDate"      timestamp                               NOT NULL,
    "Purpose"      varchar                                 NOT NULL,
    "Visitor"      varchar REFERENCES "Visitor" ("Login")    NULL,
    "Branch"       varchar REFERENCES "Branch" ("Name")      NULL,
    "Employee"     varchar REFERENCES "Employee" ("ID")      NOT NULL,
    "PassportScan" varchar                                 NOT NULL,
    "Type"         varchar REFERENCES "OrderType" ("Name")   NOT NULL,
    "Status"       varchar REFERENCES "OrderStatus" ("Name") NOT NULL
);

CREATE TABLE "Account" (
    "Email" varchar NOT NULL,
    "Password" varchar NOT NULL
)