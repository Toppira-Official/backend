package constants

type ReminderPriority string //	@name	ReminderPriority

const (
	Low      ReminderPriority = "low"
	Medium   ReminderPriority = "medium"
	High     ReminderPriority = "high"
	Critical ReminderPriority = "critical"
)

var ReminderPriorities = [...]ReminderPriority{Low, Medium, High, Critical}
