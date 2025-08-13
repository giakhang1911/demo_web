package repository

import (
	"context"
	"time"

	"github.com/giakhang1911/demo_web/path/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SurveyRepository interface {
	Create(ctx context.Context, s *models.CourseSurvey) (primitive.ObjectID, error)
	GetAll(ctx context.Context) ([]models.CourseSurvey, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.CourseSurvey, error)
}

type surveyRepository struct {
	col *mongo.Collection
}

func NewSurveyRepository(db *mongo.Database) SurveyRepository {
	return &surveyRepository{col: db.Collection("course_surveys")}
}

func (r *surveyRepository) Create(ctx context.Context, s *models.CourseSurvey) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := r.col.InsertOne(ctx, s)
	if err != nil {
		return primitive.NilObjectID, err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func (r *surveyRepository) GetAll(ctx context.Context) ([]models.CourseSurvey, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cur, err := r.col.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []models.CourseSurvey
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *surveyRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.CourseSurvey, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var s models.CourseSurvey
	if err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}
