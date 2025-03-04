package repository

import (
	"context"

	"github.com/ariefsn/upwork/helper"
	"github.com/ariefsn/upwork/logger"
	"github.com/ariefsn/upwork/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	db *mongo.Database
}

// GetIds implements models.UserRepository.
func (u *userRepository) GetIds(ctx context.Context) ([]string, error) {
	client := u.db.Client()
	session, err := client.StartSession()
	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	if err = session.StartTransaction(); err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	defer session.EndSession(ctx)

	d := new(models.UserData)
	coll := u.db.Collection(d.TableName())
	res := []string{}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		results, err := coll.Distinct(ctx, "_id", bson.M{})
		if err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		for _, v := range results {
			switch t := v.(type) {
			case string:
				res = append(res, t)
			}
		}

		return nil
	})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return res, nil
}

// Delete implements models.UserRepository.
func (u *userRepository) Delete(ctx context.Context, userID string) error {
	client := u.db.Client()
	session, err := client.StartSession()
	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return e
	}

	if err = session.StartTransaction(); err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return e
	}

	defer session.EndSession(ctx)

	dUser := new(models.UserProfile)
	dEarnings := new(models.EarningsData)
	collUser := u.db.Collection(dUser.TableName())
	collEarnings := u.db.Collection(dEarnings.TableName())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if _, err = collUser.DeleteOne(ctx, bson.M{"_id": userID}); err != nil {
			return err
		}

		if _, err = collEarnings.DeleteMany(ctx, bson.M{"userID": userID}); err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return e
	}

	return nil
}

// Get implements models.UserRepository.
func (u *userRepository) Get(ctx context.Context, userID string) (*models.UserProfile, error) {
	client := u.db.Client()
	session, err := client.StartSession()
	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	if err = session.StartTransaction(); err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	defer session.EndSession(ctx)

	d := new(models.UserProfile)
	coll := u.db.Collection(d.TableName())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err = coll.FindOne(ctx, bson.M{"_id": userID}).Decode(d); err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return d, nil
}

// Upsert implements models.UserRepository.
func (u *userRepository) Upsert(ctx context.Context, userData models.UserProfile) (*models.UserProfile, error) {
	client := u.db.Client()
	session, err := client.StartSession()
	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	if err = session.StartTransaction(); err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	defer session.EndSession(ctx)

	coll := u.db.Collection(userData.TableName())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if _, err = coll.UpdateOne(ctx, bson.M{"_id": userData.ID}, bson.M{"$set": userData}, &options.UpdateOptions{
			Upsert: helper.ToPtr(true),
		}); err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &userData, nil
}

func New(db *mongo.Database) models.UserRepository {
	return &userRepository{
		db: db,
	}
}
