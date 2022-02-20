package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AdiF1/solidity/bookings/ether"
	"github.com/AdiF1/solidity/bookings/helpers"
	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/AdiF1/solidity/bookings/internal/driver"
	"github.com/AdiF1/solidity/bookings/internal/handlers"
	"github.com/AdiF1/solidity/bookings/internal/models"
	"github.com/AdiF1/solidity/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	//"golang.org/x/crypto/bcrypt"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {

	/* password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	log.Println(string(hashedPassword)) */

	//ether.EtherNet()
	//ether.EtherAddress()
	ether.EtherAccountBalances()
	//ether.GenerateWallet()
	//ether.CreateKs()
	//ether.Address_check()
	//ether.GetBlockInfo()
	//ether.GetTransactions()
	//ether.TransferEther()
	//ether.TransferTokens()
	//ether.BlockSubscribe()
	//ether.TransactionRawCreate()
	//ether.TransactionRawSend()

	//ether.ContractDeploy()
	//ether.ContractLoad()
	//ether.ContractQuery()
	//ether.ContractWrite()
	//ether.ContractBytecode()

//ether.BaseCampaignDeploy()

	//ether.BaseCampaignCreator()
	//ether.BaseCampaignMinimumContribution()

//ether.BaseCampaignContribute()
	
	//ether.BaseCampaignBalance()

//ether.BaseCampaignSetRequest()
//ether.BaseCampaignGetRequest()

//ether.BaseCampaignApproveRequest()
//ether.BaseCampaignFinalizeRequest()
	//ether.SignatureVerify()
	//ether.CheckIt()
	
	//token := ether.ContractReadERC20()
	//fmt.Println(token.Symbol)

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	// Mail channel 
	defer close(app.MailChan)
	fmt.Println("Starting mail listener...")
	listenForMail()
	fmt.Println(fmt.Printf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}


func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	// read flags
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbHost := flag.String("db_host", "localhost", "DB host")
	dbName := flag.String("db_name", "", "DB name")
	dbUser := flag.String("db_user", "", "DB user")
	dbPass := flag.String("db_pwd", "", "DB password")
	dbPort := flag.String("db_port", "5432", "DB port")
	dbSSL := flag.String("db_ssl", "disable", "DB SSL settings (disable, prefer, require)")

	flag.Parse()

	if *dbName == "" || *dbUser == "" {
		fmt.Println("Missing required flags")
		os.Exit(1)
	}



	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", 
									*dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	log.Println("\n\nConnecting to database...")
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

/* 
Port conflict issue:

Sessions: https://github.com/alexedwards/scs

ms@MBP14inch2021 bookings % lsof -i :8080              
COMMAND    PID USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
bookings 64003   ms    9u  IPv6 0xcfa7a0d7fffbbf73      0t0  TCP *:http-alt (LISTEN)
ms@MBP14inch2021 bookings % kill -9 64003

Test all from root folder:
-	go test -v ./...

In package: 
-	go test -v
-	go test -cover
-	go test -coverprofile=coverage.out && go tool cover -html=coverage.out

Commit to GitHub:
-	git push -u origin main

Run all non-test files with z sript:
-	./run.sh

Create migrations:
-	soda reset
-	soda generate fizz(sql) (file name)
-	soda migrate 

Auto fill struct:
-	cmd shift p -> Go: Fill struct

https://cloud.linode.com/linodes
https://caddyserver.com/docs/install

root@139.177.199.26
Wo
adf@139.177.199.26
Wa

:wq<Return>

sudo service postgresql start

cd bookings
git pull

adf@localhost:/etc/postgresql/13/main$ ps ax | grep postgr

go build -o bookings cmd/web/*.go

adf@localhost:/var/www/book$ ./update.sh

adf@localhost:~$ ps ax | grep supervisorctl

Caddyfile:

{
        email   adifletcher@mac.com
}

(static) {
        file
        path *.ico *.css *.js *.gif *.jpg *.jpeg *.png *.svg *.woff *.json
        }

        header @static Cache-Control max-age=51840000
}

(security) {
        header {
                # enable HSTS
                Strict-Transport-Security max-age=315360000
                # disbale cleints from sniffing header types
                X-Content-Type-Options nosniff
                # keep referrer data off of HTTP connections
                Referrer-Policy no-referrer-when-downgrade
        }
}

import conf.d/*.conf

sudo tail -f /var/log/syslog

/ solc --optimize --abi ./contracts/Store.sol -o build --overwrite
// solc --optimize --bin ./contracts/Store.sol -o build 
// abigen --abi=./build/MyStore.abi --bin=./build/MyStore.bin --pkg=Store --out=./contracts/Store.go

// solc-select use 0.6.12 

/* 	smallCHANGE 

	user-driven fund raisers maybe eventually
	accept fiat for local causes, with 15% ADA/ETH to stake for global initiatives

 */




