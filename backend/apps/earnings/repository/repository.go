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

type earningsRepository struct {
	db *mongo.Database
}

// GetYearsByUserID implements models.EarningsRepository.
func (u *earningsRepository) GetYearsByUserID(ctx context.Context, userID string) ([]int, error) {
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

	d := new(models.EarningsData)
	coll := u.db.Collection(d.TableName())
	res := []int{}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		results, err := coll.Distinct(ctx, "year", bson.M{"userID": userID})
		if err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		for _, v := range results {
			switch t := v.(type) {
			case int32:
				res = append(res, int(t))
			case int:
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

// EarningUsersYearly implements models.EarningsRepository.
func (u *earningsRepository) EarningUsersYearly(ctx context.Context, year int) ([]*models.EarningsUserPerYear, error) {
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

	earning := new(models.EarningsData)
	d := new(models.UserData)
	coll := u.db.Collection(d.TableName())

	lookup := helper.MongoLookup(helper.MongoLookupOptions{
		From:         earning.TableName(),
		LocalField:   "_id",
		ForeignField: "userID",
		As:           "earnings",
	})

	unwind := helper.MongoUnwind(helper.MongoUnwindOptions{
		Path:          "$earnings",
		PreserveEmpty: true,
	})

	match := helper.MongoMatch(helper.MongoFilter(helper.FoEq, "earnings.year", year))

	group := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"year":   "$year",
				"userID": "$_id",
			},
			"userID": bson.M{
				"$first": "$_id",
			},
			"fullName": bson.M{
				"$first": "$fullName",
			},
			"title": bson.M{
				"$first": "$title",
			},
			"city": bson.M{
				"$first": "$city",
			},
			"country": bson.M{
				"$first": "$country",
			},
			"amount": bson.M{
				"$sum": "$earnings.amount",
			},
			"fee": bson.M{
				"$sum": "$earnings.fee",
			},
		},
	}

	project := bson.M{
		"$project": bson.M{
			"user": bson.M{
				"_id":      "$userID",
				"fullName": "$fullName",
				"city":     "$city",
				"country":  "$country",
				"title":    "$title",
			},
			"amount": "$amount",
			"fee":    "$fee",
		},
	}

	sort := bson.M{
		"$sort": helper.MongoSorting(helper.MongoSort{
			SortField: "amount",
			SortBy:    helper.SortByDesc,
		}),
	}

	pipe := []bson.M{
		lookup, unwind, match, group, project, sort,
	}

	result := []*models.EarningsUserPerYear{}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		cursor, err := coll.Aggregate(ctx, pipe)
		if err != nil {
			return err
		}

		if err = cursor.All(ctx, &result); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return result, nil
}

// GetByUserIDPerYear implements models.EarningsRepository.
func (u *earningsRepository) GetByUserIDPerYear(ctx context.Context, userID string, year int) ([]*models.EarningsDataMonthly, error) {
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

	d := new(models.EarningsDataMonthly)
	coll := u.db.Collection(d.TableName())

	filters := models.M{}

	filters["userID"] = helper.MongoFilter(helper.FoEq, "userID", userID)["userID"]
	filters["year"] = helper.MongoFilter(helper.FoEq, "year", year)["year"]

	pipe := helper.MongoPipe(helper.MongoAggregate{
		Match: bson.M(filters),
		Sort: []helper.MongoSort{
			{
				SortField: "year",
				SortBy:    helper.SortByDesc,
			},
			{
				SortField: "month",
				SortBy:    helper.SortByDesc,
			},
		},
	})

	group := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"year":  "$year",
				"month": "$month",
			},
			"userID": bson.M{
				"$first": "$userID",
			},
			"month": bson.M{
				"$first": "$month",
			},
			"year": bson.M{
				"$first": "$year",
			},
			"totalAmount": bson.M{
				"$sum": "$amount",
			},
			"totalFee": bson.M{
				"$sum": "$fee",
			},
			"items": bson.M{
				"$push": bson.M{
					"type":   "$type",
					"amount": "$amount",
					"fee":    "$fee",
				},
			},
		},
	}

	sort := bson.M{
		"$sort": helper.MongoSorting(helper.MongoSort{
			SortField: "month",
			SortBy:    helper.SortByDesc,
		}, helper.MongoSort{
			SortField: "year",
			SortBy:    helper.SortByDesc,
		}),
	}

	pipe = append(pipe, group, sort)

	result := []*models.EarningsDataMonthly{}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		cursor, err := coll.Aggregate(ctx, pipe)
		if err != nil {
			return err
		}

		if err = cursor.All(ctx, &result); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return result, nil
}

// DeleteByUserID implements models.EarningsRepository.
func (u *earningsRepository) DeleteByUserID(ctx context.Context, userID string) (*int64, error) {
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
	count := int64(0)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		deleteResult, err := coll.DeleteMany(ctx, bson.M{"_id": userID})
		if err != nil {
			return err
		}

		count = deleteResult.DeletedCount

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

	return &count, nil
}

// Upsert implements models.EarningsRepository.
func (u *earningsRepository) Upsert(ctx context.Context, earningData models.EarningsData) (*models.EarningsData, error) {
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

	coll := u.db.Collection(earningData.TableName())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if _, err = coll.UpdateOne(ctx, bson.M{"_id": earningData.ID}, bson.M{"$set": earningData}, &options.UpdateOptions{
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

	return &earningData, nil
}

func New(db *mongo.Database) models.EarningsRepository {
	return &earningsRepository{
		db: db,
	}
}
