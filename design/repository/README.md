# Repository Pattern.

[![GoReference](https://pkg.go.dev/badge/golang.org/x/tools)](https://pkg.go.dev/golang.org/x/tools)

# Contents
- [Repository Pattern.](#repositoryPattern)
  - [Introduction.](#introduction)
  - [How looks the Repository Pattern?](#howLooks)
  - [Beneficies](#beneficies)
    - [Single Responsibility Principle](#srp)
    - [Clean Architectures](#cleanArchitectures)
    - [Unit Testing and Repositories](#testing)
  - [API Practical Example](#api)
    - [Install](#install)
    - [Usage](#usage)
      - [Consume API](#consume)
      - [Run Tests in local](#localTesting)
    - [Code Review](#codeReview)
      - [type BitcoinRepository.](#BitcoinRepository)
      - [type BitcoinSrv.](#BitcoinSrv)
      - [type MockBitcoinRepository.](#MockBitcoinRepository)
      - [type RedisBitcoinRepository.](#RedisBitcoinRepository)
      - [type VendorBitcoinRepository.](#VendorBitcoinRepository)
      - [type BitcoinRepositoryFactory.](#BitcoinRepositoryFactory)

# Repository Pattern. <a name="repositoryPattern"></a>
## Introduction. <a name="introduction"></a>
Any business application, no matter how small, probably needs a data persistence system.
As a result of this demand regarding data management, today we find a market full of different tools and technologies
dedicated to persistence: SQL Databases, non-SQL databases, cloud services, spreadsheets, etc.

On the other hand, one of the biggest challenges that developers face when designing scalable applications is to
achieve a **low coupling between the different modules of the system**, in order to deliver a solution that provides
**easy testability and deployment**.

The **_Repository pattern_** comes to solve a recurring problem for us:
The coupling to persistence systems.
**_Your primary concern is to decouple data access and writing of the logic that is part of the domain of the Application_**,
speaking in terms of DDD.
The goal is to build a domain totally independent of the different technologies used for data persistence.

The pattern proposes that domain objects interact only with an abstraction of the data layer.
In this way, our domain does not know any details regarding the implementation of the infrastructure, which allows a
clear **separation of responsibilities**, allowing domain objects to have as their only task to implement business logic
and derive the responsibility of data persistence towards the specific implementation or implementations that exist in
external layers of the application, in other words, infrastructure layers.

Data layer or infrastructure is understood as any tool read and/or write data, be it a database, a cache,
external service, files or, for example, any array in memory.

This solution was introduced first time in the **_book Domain Driven Design_**, written by **_Eric Evans_** in year 2004.

## How looks the Repository Pattern? <a name="howLooks"></a>
Basically, **it is an abstraction**, that is, an interface with its different implementations.

![repositoy_pattern_uml](img/repository_pattern.drawio.png.png)
An important point is that **the interface Repository is part of our Domain** (yellow color).
While they are the implementations of the interface who are in the infrastructure layer (green color).
This makes sense, it’s the domain that define the contract that must be fulfilled, since it’s the domain that knows
what operations are necessary to perform on the data in order to complete the uses cases.

As we see, the proposal is that both **Entities, Aggregates and Domain services interact with data through the Repository
interface**.
![DDD_diagram_repository](img/DDD_Diagram.png)

## Beneficies <a name="beneficies"></a>
### Single Responsibility Principle <a name="srp"></a>
In principle, the use of Repository encourages the **S of SOLID**: **Single Responsibility Principle**.
That is, each thing has only one reason for change.
A use case coupled to the infrastructure means that our domain have to deal with their own business rules and, as if
this wasn’t enough, which also bears the responsibility of correctly implementing the persistence system.

Also keep in mind that in an infrastructure-coupled scenario, if change the way the data is persisted, it would
directly affect our Domain, having to modify the use cases, adapting them to the new infrastructure.
This is detrimental to system design. Since, following DDD approach, **domain modifications should be as a result of
changes in business rules**, not due to changes in dependencies technological.

### Clean Architectures <a name="cleanArchitectures"></a>
On the other hand, it also encourages the main rule of **Clean
Architectures**: Whose main premise is the **separation of the domain layer
of the infrastructure layer**.
Ensuring that the domain does not know infrastructure details. Of
Inversely, it’s the infrastructure that knows the details of the domain,
such that, it is the outer layer that is coupled to the needs of the layers
more interiors.
In **SOLID** terms, **this refers to Dependency Inversion Principle** since the Application services just depends on
abstraction.

The repository interface belongs to the _Interface
Adapters layer_. Sometimes they ara called Gateways.
While the different implementations of Repository Interface belongs to the _Frameworks & Drivers layer_.
![Clean_architecture_img](img/CleanArchitecture_img.jpg)

### Unit Testing and Repositories <a name="testing"></a>
Lastly, there is a key point that, in my opinion, drives the use of repositories in projects of any size:
**Unit testing**.

Let us imagine that we have an Application service that registers new users, it could be called UserRegister for example.
Now, this service has different validations that imply perform database queries, such as checking that the name
username or email do not exist.
We could run unit tests of the service using a database actual, either production copy or a specific base for testing.
Anyway, we should take care of updating the records of the tables before executing each unit test, so that each use
case find the necessary values in the persistence system.

On the other hand, we should also deal with the latency involved in wake up
a real database every time we run unit tests, which directly impacts our CI/CD cycle.

Analyzing this scenario carefully, every time we add a new use case for unit tests, their execution time could grow
exponentially, since each use case possibly needs one or more CRUD operations on the database.

This growth in execution times causes the introduction of a new feature or a simple refactor becomes tedious and a waste
of resources.
In large-scale projects, unit tests are huge, and it’s not optimal to build, clean and update the database to be able to
execute them.
In any case, **the specific implementation of each database could be tested in the infrastructure layer**
(integration/black box tests) with a reduced number of use cases.
This scenario gives us the opportunity to use a **Mock Repository** to use in unit tests, through an in-memory database,
redis, an in-memory array or whatever that occurs to us and is useful for our context.

**That is precisely the power of the Repository, the changeability it offers**.
Today we could be using a persistence system A, tomorrow one B, but in turn using one C on the QA server and one D on
the PRD server.
In this way, persistence systems are placed in the place of detail of implementation, being able to use the database
that best suits the needs of the project in each phase of the project.

## API Practical Example <a name="api"></a>
This example consists of an HTTP API that returns the currency quote Bitcoin relative to other non-digital currencies.
To deliver values to the Client, it uses two Gateways:
- Redis implementation in memory.
- Vendor implementation: Public API --> [https://api.coindesk.com/v1/bpi/currentprice.json](https://api.coindesk.com/v1/bpi/currentprice.json)

The peculiarity of the API is that depending on the time in which to perform the Request, it’s decided at runtime which
Gateway to use.
The idea is to make several Requests so that, as a Client, you can visualize how the data is obtained from different
data sources depending on the time.

### Install <a name="install"></a>
- Dependencies: Docker & Docker-compose.
```shell
git clone https://github.com/rcrespodev/Blogs
```
```shell
cd Blogs/design/repository
```

Build image:
```shell
sudo docker build -t repository_pattern:latest .
```

The next command
```shell
docker image ls -f reference=repository_pattern:latest
```

should be showed an output like this:
```shell
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
repository_pattern   latest    09a179937fae   7 minutes ago   522MB
```

### Usage <a name="usage"></a>
Run api in container:
```shell
docker compose up -d
```

#### Consume API <a name="consume"></a>
Request:
```shell
curl 0.0.0.0:8080/bitcoin-price | json_pp
```

Response:
```shell
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 01 Jun 2022 02:46:28 GMT
Content-Length: 277

{
   "bitcoin_price" : {
      "crypto_name" : "Bitcoin",
      "currencies" : [
         {
            "code" : "USD",
            "description" : "United States Dollar",
            "rate" : 30441.5389
         },
         {
            "code" : "EUR",
            "description" : "Euro",
            "rate" : 28306.8565
         },
         {
            "code" : "GBP",
            "description" : "British Pound Sterling",
            "rate" : 24084.1584
         }
      ],
      "updated_at" : "2022-05-30T14:59:45.663265426Z"
   },
   "error" : "",
   "implementation_name" : "Vendor implementation( <https://api.coindesk.com/v1/bpi/currentprice.json> )"
}
```
Where implementation_name return the name of gateway used the API in runtime. The possible values are:
- "implementation_name" : "Vendor implementation( [https://api.coindesk.com/v1/bpi/currentprice.json](https://api.coindesk.com/v1/bpi/currentprice.json) )"
- "implementation_name" : "Redis"

#### Run Tests in local <a name="localTesting"></a>
Run tests in local:
```shell
make tests
```

### Code Review <a name="codeReview"></a>
#### type BitcoinRepository. <a name="BitcoinRepository"></a>
BitcoinRepository is the interface that define the contract in own domain layer.
```go
type BitcoinRepository interface {
	BitcoinPrice() (error, *BitcoinPrice)
	ImplementationName() (error, string)
}
```

#### type BitcoinSrv. <a name="BitcoinSrv"></a>
The BitcoinSrv domain Services knows only the interface BitcoinRepository.
```go
type BitcoinSrv struct {
	bitcoinRepository BitcoinRepository
}

func NewBitcoinSrv(bitcoinRepository BitcoinRepository) *BitcoinSrv {
	return &BitcoinSrv{bitcoinRepository: bitcoinRepository}
}

func (b BitcoinSrv) GetBitcoinPrice() *BitcoinResponse {
	err, bitcoinPrice := b.bitcoinRepository.BitcoinPrice()
	err, implementationName := b.bitcoinRepository.ImplementationName()

	// more code
	
	return &BitcoinResponse{
		BitcoinPrice: &BitcoinPriceResponse{
			UpdatedAt:  bitcoinPrice.updatedAt,
			CryptoName: bitcoinPrice.cryptoName,
			Currencies: bitcoinPrice.currencies,
		},
		ImplementationName: implementationName,
	}
}
```

#### type MockBitcoinRepository. <a name="MockBitcoinRepository"></a>
In the other hand, was found the concrete implementations of Repository.
MockBitcoinRepository is an implementation with only testing purposes. He has an only in memory register.
```go
type MockBitcoinRepository struct {
	data           *domain.BitcoinPrice
	implementation string
}

func New() domain.BitcoinRepository {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	return MockBitcoinRepository{
		data: domain.NewBitcoinPrice(t, []domain.Currency{
			{
				Code:        "USD",
				Rate:        29055.3222,
				Description: "United States Dollar",
			},
		}),
		implementation: "Mock_Repository",
	}
}

func (m MockBitcoinRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	return nil, m.data
}

func (m MockBitcoinRepository) ImplementationName() (error, string) {
	return nil, m.implementation
}
```

#### type RedisBitcoinRepository. <a name="RedisBitcoinRepository"></a>
The concrete implementation in redis db.
```go
type RedisBitcoinRepository struct {
	redisCliente *redis.Client
	ctx          context.Context
	time         time.Time
}

func New(host string, port int, db int) (*RedisBitcoinRepository, error) {
    // more code
	return repository, nil
}

func (r *RedisBitcoinRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	result, err := r.redisCliente.Do(r.ctx, "get", r.time.String()).Result()
	if err != nil {
		return err, nil
	}

	strResult, ok := result.(string)
	if !ok {
		return fmt.Errorf("internal Server Error"), nil
	}

	var redisBitcoinPrice RedisBitcoinPrice
	err = json.Unmarshal([]byte(strResult), &redisBitcoinPrice)
	if err != nil {
		return fmt.Errorf("internal Server Error"), nil
	}

	return nil, domain.NewBitcoinPrice(redisBitcoinPrice.UpdatedAt, redisBitcoinPrice.Currencies)
}

func (r RedisBitcoinRepository) ImplementationName() (error, string) {
	return nil, "Redis"
}
```

#### type VendorBitcoinRepository. <a name="VendorBitcoinRepository"></a>
The Vendor implementation. He obtains the data from https://api.coindesk.com/v1/bpi/currentprice.json api.
```go
type VendorBitcoinRepository struct {
	endpoint string
}

func NewVendorRepository(endpoint string) *VendorBitcoinRepository {
	return &VendorBitcoinRepository{
		endpoint: endpoint,
	}
}

func (v VendorBitcoinRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	resp, err := http.Get(v.endpoint)
	if err != nil {
		return fmt.Errorf("internal server error"), nil
	}

	defer resp.Body.Close()

	var VendorBitcoinPrice VendorBitcoinPrice
	err = json.NewDecoder(resp.Body).Decode(&VendorBitcoinPrice)
	if err != nil {
		if err != nil {
			return fmt.Errorf("internal server error"), nil
		}
	}

	bitcoinPrice, err := v.newBitcoinPrice(VendorBitcoinPrice)
	if err != nil {
		return err, nil
	}

	return nil, bitcoinPrice
}

func (v VendorBitcoinRepository) newBitcoinPrice(price VendorBitcoinPrice) (*domain.BitcoinPrice, error) {

	bitcoinPrice := domain.NewBitcoinPrice(time.Now(), []domain.Currency{
		{
			Code:        price.Bpi.USD.Code,
			Rate:        price.Bpi.USD.RateFloat,
			Description: price.Bpi.USD.Description,
		},
		{
			Code:        price.Bpi.EUR.Code,
			Rate:        price.Bpi.EUR.RateFloat,
			Description: price.Bpi.EUR.Description,
		},
		{
			Code:        price.Bpi.GBP.Code,
			Rate:        price.Bpi.GBP.RateFloat,
			Description: price.Bpi.GBP.Description,
		},
	})

	return bitcoinPrice, nil
}

func (v VendorBitcoinRepository) ImplementationName() (error, string) {
	return nil, fmt.Sprintf("Vendor implementation( %v )", v.endpoint)
}
```

#### type BitcoinRepositoryFactory. <a name="BitcoinRepositoryFactory"></a>
Finally, we find a BitcoinRepositoryFactory whose responsibility is instance all needed repositories to wake up the API
and return one of them in function of server actual hour.
```go
type BitcoinRepositoryFactory struct {
	test                                              bool
	mockRepository, redisRepository, vendorRepository domain.BitcoinRepository
}

func NewBitcoinRepositoryFactory(test bool) (error, *BitcoinRepositoryFactory) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return err, nil
	}

	redisRepository, err := redisBitcoinRepository.New(redisHost, redisPort, 0)
	if err != nil {
		return err, nil
	}

	vendorRepository := vendorBitcoinRepository.NewVendorRepository(os.Getenv("VENDOR_ENDPOINT"))

	return nil, &BitcoinRepositoryFactory{
		test:             test,
		mockRepository:   mockBitcoinRepository.New(),
		redisRepository:  redisRepository,
		vendorRepository: vendorRepository,
	}
}

func (b BitcoinRepositoryFactory) Repository() domain.BitcoinRepository {
	if b.test {
		return b.mockRepository
	}

	t := time.Now()
	switch t.Unix() % 2 {
	case 0:
		return b.redisRepository
	default:
		return b.vendorRepository
	}
}
```