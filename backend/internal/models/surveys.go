package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Working struct {
	CompanyName            string `json:"company_name" bson:"company_name"`
	IndustrySector         string `json:"industry_sector" bson:"industry_sector"`
	Location               string `json:"location" bson:"location"`
	Position               string `json:"position" bson:"position"`
	Status                 string `json:"status" bson:"status"`
	GrossSalary            int    `json:"gross_salary" bson:"gross_salary"`
	WorkStudyConnectedness string `json:"work_study_connectedness" bson:"work_study_connectedness"`
	FirstJob               bool   `json:"first_job" bson:"first_job"`
	FirstDayWork           string `json:"first_day_work" bson:"first_day_work"`
	IsAverageMinimumWage   bool   `json:"is_average_minimum_wage" bson:"is_average_minimum_wage"`
}

type Enterpreneurship struct {
	Product          string `json:"product" bson:"product"`
	NetIncomeAverage int    `json:"net_income_average" bson:"net_income_average"`
	Unit             string `json:"unit" bson:"unit"`
	StartDate        string `json:"start_date" bson:"start_date"`
}

type FurtherStudy struct {
	StudyProgram string `json:"study_program" bson:"study_program"`
	University   string `json:"university" bson:"university"`
	Level        string `json:"level" bson:"level"`
	StartDate    string `json:"start_date" bson:"start_date"`
}

type NotWorking struct {
	AlreadyWorkingOrEnterpreneurship bool   `json:"already_working_or_enterpreneurship" bson:"already_working_or_enterpreneurship"`
	Reason                           string `json:"reason" bson:"reason"`
	StartWorkingOrEnterpreneurship   string `json:"start_working_or_enterpreneurship" bson:"start_working_or_enterpreneurship"`
	EndWorkingOrEnterpreneurship     string `json:"end_working_or_enterpreneurship" bson:"end_working_or_enterpreneurship"`
}

type Survey struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId          string             `json:"user_id" bson:"user_id"`
	CurrentActivity string             `json:"current_activity" bson:"current_activity"`
	Details         interface{}        `json:"details" bson:"details"`
	Satisfaction    string             `json:"satisfaction" bson:"satisfaction"`
	Suggestions     string             `json:"suggestions" bson:"suggestions"`
}
