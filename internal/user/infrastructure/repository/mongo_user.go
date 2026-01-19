package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userDB struct {
	ID            bson.ObjectID          `bson:"_id,omitempty"`
	FirstName     string                 `bson:"firstName"`
	LastName      string                 `bson:"lastName"`
	Email         string                 `bson:"email"`
	UserName      string                 `bson:"userName"`
	Password      string                 `bson:"password"`
	CreatedAt     time.Time              `bson:"createdAt"`
	UpdatedAt     time.Time              `bson:"updatedAt"`
	Status        domain.UserStatus      `bson:"status"`
	StatusHistory []domain.StatusChanges `bson:"statusHistory"`
}

type MongoUserRepository struct {
	db *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db: db.Collection("users")}
}

func (m *MongoUserRepository) Create(ctx context.Context, userEntity *domain.User) (*domain.User, error) {

	userResult, err := m.db.InsertOne(ctx, &userDB{
		FirstName:     userEntity.FirstName(),
		LastName:      userEntity.LastName(),
		Email:         userEntity.Email(),
		UserName:      userEntity.UserName(),
		Password:      userEntity.PasswordHash(),
		CreatedAt:     userEntity.CreatedAt(),
		UpdatedAt:     userEntity.UpdatedAt(),
		Status:        userEntity.Status(),
		StatusHistory: userEntity.StatusHistory(),
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
		userDB.Status,
		userDB.StatusHistory,
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

	result, errUpdate := m.db.UpdateOne(ctx, filter, update)

	if errUpdate != nil {
		return nil, errUpdate
	}

	log.Printf("USER UPDATED: %+v", result)

	return userEntity, nil
}

func (m *MongoUserRepository) GetUsers(ctx context.Context, filters dto.Filters) ([]*domain.User, error) {

	query := bson.M{}

	if filters.Status != nil {
		query["status"] = *filters.Status
	}
	cursor, err := m.db.Find(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	defer cursor.Close(ctx)

	users := make([]*domain.User, 0)

	for cursor.Next(ctx) {
		var userDB userDB
		if err := cursor.Decode(&userDB); err != nil {
			return nil, fmt.Errorf("decoding error: %w", err)
		}
		users = append(users, domain.RestoreUser(
			domain.RestoreUserID(userDB.ID.Hex()),
			userDB.FirstName,
			userDB.LastName,
			userDB.Email,
			userDB.UserName,
			domain.RestoreUserPassword(userDB.Password),
			userDB.CreatedAt,
			userDB.UpdatedAt,
			userDB.Status,
			userDB.StatusHistory,
		))
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return users, nil
}

// ESTE NO PASA JAMAS
func (m *MongoUserRepository) DeleteUser(ctx context.Context, id domain.UserIDVO) error {
	oid, err := getObjectID(id.Value())
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": oid,
	}
	result, err := m.db.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	log.Printf("USER DELETED: %+v", result)

	return nil
}

func (m *MongoUserRepository) DeleteSoftUser(ctx context.Context, id domain.UserIDVO, userEntity *domain.User) error {

	oid, err := getObjectID(id.Value())

	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": oid,
	}

	updateDeleteSoft := bson.M{
		"status":        userEntity.Status(),
		"statusHistory": userEntity.StatusHistory(),
	}

	update := bson.M{
		"$set": updateDeleteSoft,
	}

	result, err := m.db.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	log.Printf("USER DELETED SOFT: %+v", result)

	return nil
}

func getObjectID(id string) (bson.ObjectID, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.NilObjectID, fmt.Errorf("invalid id: %w", err)
	}
	return oid, nil
}
