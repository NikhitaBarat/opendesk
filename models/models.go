package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	TaskTitle   string             `json:"task_title,omitempty"`
	TaskTag     string             `json:"task_tag,omitempty"`
	TimeLeft    string             `json:"time_left,omitempty"`
	DateAdded   string             `json:"date_added,omitempty"`
	Description string             `json:"description,omitempty"`
	Created     string             `json:"created,omitempty"`
	TaskStatus  bool               `json:"task_status,omitempty"`
}
