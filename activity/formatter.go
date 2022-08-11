package activity

import "time"

type ActivityFormatter struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func FormatActivity(activity Activity) ActivityFormatter {
	return ActivityFormatter{
		ID:        activity.ID,
		Email:     activity.Email,
		Title:     activity.Title,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
		DeletedAt: activity.DeletedAt,
	}
}

func FormatActivities(activities []Activity) []ActivityFormatter {
	activitiesFormatter := []ActivityFormatter{}
	for _, activity := range activities {
		activitiesFormatter = append(activitiesFormatter, FormatActivity(activity))
	}
	return activitiesFormatter
}
