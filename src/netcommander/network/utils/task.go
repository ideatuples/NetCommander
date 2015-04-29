package utils

type Task struct {
	
	httpMethod string
	exURL *ExURL
	
}

func NewTask(pMehtod string, pExURL *ExURL) Task{
	
	var theElement Task
	
	theElement.httpMethod = pMehtod
	theElement.exURL = pExURL

	return theElement
}

func (theElement *Task) GetMethod() string{
	
	return theElement.httpMethod
}

func (theElement *Task) GetExURL() *ExURL {
	
	return theElement.exURL
}

type TaskQueue struct {
	theList []Task
	theCurrentIndex int
	
}

func NewTaskQueue() *TaskQueue {
	
	var theQueue *TaskQueue
	theQueue = new(TaskQueue)
	
	theQueue.theList = make([]Task, 0)
	
	return theQueue
}

func (theQueue *TaskQueue) PushBack(pMethod string, pExURL *ExURL) bool {
	
	var anElement Task
	
	anElement = NewTask(pMethod, pExURL)
	
	theQueue.theList = append(theQueue.theList, anElement)
	
	return true
}

func (theQueue *TaskQueue) PullFront() Task {
	
	var theFrontTask Task
	
	theFrontTask = theQueue.theList[theQueue.theCurrentIndex]
	theQueue.theCurrentIndex++
	
	return theFrontTask
}

func (theQueue *TaskQueue) GetLength() int {
	
	return len(theQueue.theList)
}

func (theQueue *TaskQueue) GetCurrentIndex() int {
	
	return theQueue.theCurrentIndex
}

