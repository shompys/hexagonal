package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/shompys/hexagonal/internal/user/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userDB struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	FirstName string        `bson:"firstName"`
	LastName  string        `bson:"lastName"`
	Email     string        `bson:"email"`
	UserName  string        `bson:"userName"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

type MongoUserRepository struct {
	db *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db: db.Collection("users")}
}

func (m *MongoUserRepository) Create(ctx context.Context, userEntity *domain.User) (*domain.User, error) {

	userResult, err := m.db.InsertOne(ctx, &userDB{
		FirstName: userEntity.FirstName(),
		LastName:  userEntity.LastName(),
		Email:     userEntity.Email(),
		UserName:  userEntity.UserName(),
		Password:  userEntity.PasswordHash(),
		CreatedAt: userEntity.CreatedAt(),
		UpdatedAt: userEntity.UpdatedAt(),
	})

	if err != nil {
		return nil, err
	}
	if oid, ok := userResult.InsertedID.(bson.ObjectID); ok {
		idVO, err := domain.NewUserID(oid.Hex())
		if err != nil {
			return nil, fmt.Errorf("error al crear el VO del ID generado: %w", err)
		}
		userEntity.SetID(idVO)
	}

	return userEntity, nil
}

func (m *MongoUserRepository) GetUserByID(ctx context.Context, id domain.UserIDVO) (*domain.User, error) {

	oid, err := getObjectID(id.Value())
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	var userDB userDB

	if err := m.db.FindOne(ctx, filter).Decode(&userDB); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return domain.RestoreUser(
		id,
		userDB.FirstName,
		userDB.LastName,
		userDB.Email,
		userDB.UserName,
		domain.RestoreUserPassword(userDB.Password),
		userDB.CreatedAt,
		userDB.UpdatedAt,
	), nil
}

func (m *MongoUserRepository) UpdateUser(ctx context.Context, id domain.UserIDVO, userEntity *domain.User) (*domain.User, error) {

	oid, err := getObjectID(id.Value())
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": oid,
	}

	userUpdate := bson.M{
		"firstName": userEntity.FirstName(),
		"lastName":  userEntity.LastName(),
		"email":     userEntity.Email(),
		"userName":  userEntity.UserName(),
		"password":  userEntity.PasswordHash(),
		// NO pongo el created_at aquí, así me aseguro de que NUNCA se pise
		"updatedAt": userEntity.UpdatedAt(),
	}

	update := bson.M{
		"$set": userUpdate,
	}

	_, errUpdate := m.db.UpdateOne(ctx, filter, update)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return userEntity, nil
}

func (m *MongoUserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}

func (m *MongoUserRepository) DeleteUser(ctx context.Context, id domain.UserIDVO) error {
	return nil
}

func getObjectID(id string) (bson.ObjectID, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.NilObjectID, fmt.Errorf("invalid id: %w", err)
	}
	return oid, nil
}
