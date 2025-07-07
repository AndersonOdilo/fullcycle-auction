package auction

import (
	"context"
	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)


type CreateAuctionTestSuite struct {
	suite.Suite
	ctx context.Context
	databaseConnection *mongo.Database
}

func (suite *CreateAuctionTestSuite) SetupSuite() {

	ctx := context.Background()
	databaseConnection, _ := mongodb.NewMongoDBConnection(ctx)

	suite.ctx = ctx
	suite.databaseConnection = databaseConnection;
}


func TestSuite(t *testing.T) {

	godotenv.Load("../../../../cmd/auction/.env");

	suite.Run(t, new(CreateAuctionTestSuite))
}


func (suite *CreateAuctionTestSuite) TestCreateAuction_UpdateStatusCompleted() {

	auctionEntity, _ := auction_entity.CreateAuction("Camaro", "Carro esportivo", "Camaro Amarelo", 2);

	auctionRepository := NewAuctionRepository(suite.databaseConnection)
	auctionRepository.CreateAuction(suite.ctx, auctionEntity)

	auctionEntityActive, _ := auctionRepository.FindAuctionById(suite.ctx, auctionEntity.Id);

	suite.Equal(auctionEntityActive.Status, auction_entity.Active);

	time.Sleep(getAuctionInterval() + time.Second)

	auctionEntityCompleted, _ := auctionRepository.FindAuctionById(suite.ctx, auctionEntity.Id);

	suite.Equal(auctionEntityCompleted.Status, auction_entity.Completed);


}

