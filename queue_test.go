package List

import (
	"DbxDev/Logger"
	"strconv"
	"testing"
)

func TestNode(t *testing.T) {
	libelle := "ValueTest"
	n := NewStringNode(libelle)
	if n.Value() == libelle {
		Logger.Debugf("New node successfully created : %v", n.Value())
	} else {
		t.Errorf("Failed to create a new node.")
	}
}

func TestQueueEnQueue(t *testing.T) {
	size := 4
	q := NewGenericQueue()
	Logger.Debugf("New queue built")
	for i := 0; i < size; i++ {
		Logger.Debugf("Enqueuing node #%v", i)
		q.EnQueue(NewStringNode("Node-" + strconv.Itoa(i)))
	}
	if q.Size() != size {
		t.Error("Expecting queue of size %v. Found %v", size, q.Size())
	}
	count := q.Size()
	for q.Size() > 0 {
		count--
		Logger.Debugf("Dequeuing : %v", q.DeQueue())
		if count == size/2 {
			q.EnQueue(NewStringNode("One more"))
		}
	}
}

func TestStack(t *testing.T) {
	size := 4
	s := NewGenericStack()
	Logger.Debugf("New stack built")
	for i := 0; i < size; i++ {
		Logger.Debugf("Put node #%v", i)
		s.Put(NewStringNode("Node-" + strconv.Itoa(i)))
	}
	if s.Size() != size {
		t.Error("Expecting stack of size %v. Found %v", size, s.Size())
	}
	count := s.Size()
	for s.Size() > 0 {
		count--
		Logger.Debugf("Pop : %v", s.Pop())
		if count == size/2 {
			s.Put(NewStringNode("One more"))
		}
	}
}
