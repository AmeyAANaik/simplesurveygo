package dao

import (
	"fmt"
//        "reflect"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Question struct {
	QuestionString string   `json:"questionString" bson:"questionString"`
	Options        []string `json:"options" bson:"options"`
}

type Answer struct {
	Question Question `json:"question" bson:"question"`
	Answer   string   `json:"answer" bson:"answer"`
}

type Survey struct {
	ID   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	SurveyName  string     `json:"surveyName" bson:"surveyName"`
	Heading     string     `json:"heading" bson:"heading"`
	Description string     `json:"description" bson:"description"`
	Questions   []Question `json:"questions" bson:"questions"`
	Status      bool       `json:"status" bson:"status"`
	Time        time.Time  `json:"Time" bson:"Time"`
}

type SurveyResponse struct {
	UserName string   `json:"userName" bson:"userName"`
	Survey   Survey   `json:"survey" bson:"survey"`
	Answers  []Answer `json:"answers" bson:"answers"`
}

func GetActiveSurveys() interface{} {
	session := MgoSession.Clone()
	defer session.Close()

	var response []interface{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{"status": true})
	err := query.All(&response)

	if err != nil {
		return nil
	} else {
		return response
	}
}

func GetSurveysForUser(userName string) interface{} {
	session := MgoSession.Clone()
	defer session.Close()

	var response []interface{}
	clctn := session.DB("simplesurveys").C("survey_response")
	query := clctn.Find(bson.M{"userName": userName})
	err := query.All(&response)

	if err != nil {
		return nil
	} else {
		return response
	}
}

func GetSurveyByName(surveyName string) interface{} {
	fmt.Println("GetSurveyByName:" + surveyName)
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{"surveyname": surveyName})
	err := query.One(&response)

	if err != nil {
		return nil
	} else {
		return response
	}
}

func InsertUserResponse(userResponse SurveyResponse) {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("survey_response")
	clctn.Insert(userResponse)
}

func DeactiveAllSurveyByDate(){
	session := MgoSession.Clone()
	defer session.Close()

	response :=[]Survey{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{})
	err := query.All(&response)

        fmt.Println("in Deactive")
        fmt.Println(len(response))
	if err != nil{
            return 
	}
	currTime := time.Now()
	for _,dat := range response{
               if( currTime.After(dat.Time)){
	        colQuerier := bson.M{"_id": dat.ID}
                change := bson.M{"$set": bson.M{"status" :false }}
		err1 := clctn.Update(colQuerier, change)
	        if err1 != nil {
		    panic(err1)
	         } 
	       }
		
	}
	time.Sleep(10*time.Second)
	}
