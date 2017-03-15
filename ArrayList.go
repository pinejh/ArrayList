package ArrayList

func New() *ArrayList {
	return &ArrayList{[]Object{}}
}

type Object interface{}

type ArrayList struct {
	array []Object
}

func (arr *ArrayList) Get(i int) Object {
	return arr.array[i]
}

func (arr *ArrayList) IndexOf(obj Object) int {
	for i, n := range arr.array {
		if n == obj {
			return i
		}
	}
	return -1
}

func (arr *ArrayList) Add(obj *Object) {
	arr.array = append(arr.array, obj)
}

func (arr *ArrayList) Remove(index int) {
	if index >= 0 && index < len(arr.array) {
		for i, _ := range arr.array {
			arr.array = append(arr.array[:i], arr.array[:i+1]...)
		}
	}
}
