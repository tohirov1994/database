package database

///////////////////////////////////// queries for SignIn ///////////////////////////////////////////////////

const GetLoginPassManager = `SELECT login, password FROM managers WHERE login = ?`
const GetLoginPassIdClient = `SELECT login, password, id FROM clients WHERE login = ?`

///////////////////////////////////// queries for Manager ///////////////////////////////////////////////////
const InsertClient = `INSERT INTO clients(name, surname, login, password) VALUES (:name, :surname, :login, :password);`
const InsertClientCard = `INSERT INTO clients_cards(pan, pin, balance, holderName, cvv, validity, client_id) VALUES (:pan, :pin, :balance, :holderName, :cvv, :validity, :clientId);`
const GetLastPAN = `SELECT ifnull(max(pan),0) FROM clients_cards;`
const InsertService = `INSERT INTO services(service) VALUES (:serviceName);`
const InsertAtm = `INSERT INTO atms(city, district, street) VALUES (:cityName, :districtName, :streetName);`

///////////////////////////////////// queries for Client ///////////////////////////////////////////////////
const GetCards = `SELECT id, pan, pin, balance, holderName, cvv, validity FROM clients_cards WHERE client_id = :idClient;`
const GetTransferCard = `SELECT count(pan) FROM clients_cards WHERE client_id = :idClient;`
const SelectCardWhoHaveManyCards  = `SELECT pan FROM clients_cards where client_id = ? AND pan = ?;`
const OutOneAmmount = `UPDATE clients_cards SET balance = balance - :amount WHERE client_id= :idClient AND balance >= :amount;`
const OutMoreOneAmmount = `UPDATE clients_cards SET balance = balance - :amount WHERE pan= :panClient AND balance >= :amount;`
const CheckPAN = `select pan from clients_cards where pan = ?;`
const InAmmount = `UPDATE clients_cards SET balance = balance + :amount WHERE pan= :PANInner;`
const CheckServiceName = `select service from services where service = ?;`
const PayService = `UPDATE services SET balance = balance + :amount WHERE service = :serviceName;`
const GetBalanceClientId = `SELECT client_id, balance FROM clients_cards WHERE client_id = ?`
const GetBalanceClientPAN = `SELECT balance FROM clients_cards WHERE pan = ?`
const GetServices = `SELECT id, service FROM services;`

///////////////////////////////////// queries for Export ///////////////////////////////////////////////////
const GetManagerData = `SELECT id, name, surname, login, password FROM managers;`
const GetClientData = `SELECT id, name, surname, login, password FROM clients;`
const GetCardsData = `SELECT id, pan, pin, balance, holderName, cvv, validity, client_id FROM clients_cards;`
const GetATMData = `SELECT id, city, district, street FROM atms;`
const GetServicesData = `SELECT id, service, ifnull((balance),0) FROM services;`